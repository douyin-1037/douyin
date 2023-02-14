package service

// @path: video/service/like_video.go
// @description: LikeVideo service of video
// @author: Chongzhi <dczdcz2001@aliyun.com>
import (
	"context"
	"douyin/code_gen/kitex_gen/videoproto"
	"douyin/video/infra/dal"
	"douyin/video/infra/redis"
)

type LikeVideoService struct {
	ctx context.Context
}

func NewLikeVideoService(ctx context.Context) *LikeVideoService {
	return &LikeVideoService{ctx: ctx}
}

func (s *LikeVideoService) LikeVideo(req *videoproto.LikeVideoReq) error {
	userID := req.UserId
	videoID := req.VideoId
	isLikeKeyExist, err := redis.IsLikeKeyExist(userID)
	if err != nil {
		return err
	}
	if isLikeKeyExist == true {
		// 如果redis有这个userID的记录，则需要在redis中再加入这条新的点赞的操作，确保和mysql一致
		isLikeById, err := redis.GetIsLikeById(userID, videoID)
		if err != nil {
			return err
		}
		if isLikeById == true {
			return nil
		}
		if err := redis.AddLike(userID, videoID); err != nil {
			return err
		}
	} else {
		// 如果redis没有这个userID的记录，则去mysql查询一次点赞列表进行缓存
		likeList, err := dal.MGetLikeList(s.ctx, userID)
		if err != nil {
			return err
		}
		if err := redis.AddLikeList(userID, likeList); err != nil {
			return err
		}
		isLikeById, err := redis.GetIsLikeById(userID, videoID)
		if err != nil {
			return err
		}
		if isLikeById == true {
			return nil
		}
		if err := redis.AddLike(userID, videoID); err != nil {
			return err
		}
	}
	if err := dal.LikeVideo(s.ctx, userID, videoID); err != nil {
		return err
	}
	return nil
}
