package service

// @path: video/service/unlike_video.go
// @description: UnLike service of video
// @author: Chongzhi <dczdcz2001@aliyun.com>
import (
	"context"
	"douyin/code_gen/kitex_gen/videoproto"
	"douyin/video/infra/dal"
	"douyin/video/infra/redis"
)

type UnLikeVideoService struct {
	ctx context.Context
}

func NewUnLikeVideoService(ctx context.Context) *UnLikeVideoService {
	return &UnLikeVideoService{ctx: ctx}
}

func (s *UnLikeVideoService) UnLikeVideo(req *videoproto.UnLikeVideoReq) error {
	userID := req.UserId
	videoID := req.VideoId
	isLikeKeyExist, err := redis.IsLikeKeyExist(userID)
	if err != nil {
		return err
	}
	if isLikeKeyExist == true {
		// 如果redis有这个userID的记录，则需要在redis中删去这条like记录，确保和mysql一致
		isLikeById, err := redis.GetIsLikeById(userID, videoID)
		if err != nil {
			return err
		}
		if isLikeById == false {
			return nil
		}
		if err := redis.DeleteLike(userID, videoID); err != nil {
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
		if isLikeById == false {
			return nil
		}
		if err := redis.DeleteLike(userID, videoID); err != nil {
			return err
		}
	}
	if err := dal.UnLikeVideo(s.ctx, userID, videoID); err != nil {
		return err
	}
	return nil
}
