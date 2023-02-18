package service

// @path: video/service/like_video.go
// @description: LikeVideo service of video
// @author: Chongzhi <dczdcz2001@aliyun.com>
import (
	"context"
	"douyin/code_gen/kitex_gen/videoproto"
	"douyin/video/infra/dal"
	"douyin/video/infra/pulsar"
	"douyin/video/infra/redis"
	"github.com/cloudwego/kitex/pkg/klog"
)

type LikeVideoService struct {
	ctx context.Context
}

func NewLikeVideoService(ctx context.Context) *LikeVideoService {
	return &LikeVideoService{ctx: ctx}
}

func (s *LikeVideoService) LikeVideo(req *videoproto.LikeVideoReq) error {
	userId := req.UserId
	videoID := req.VideoId
	isLikeKeyExist, err := redis.IsLikeKeyExist(userId)
	if err != nil {
		klog.Error(err)
	}
	if isLikeKeyExist == true {
		// 如果redis有这个userId的记录，则需要在redis中再加入这条新的点赞的操作，确保和mysql一致
		isLikeById, err := redis.GetIsLikeById(userId, videoID)
		if err != nil {
			klog.Error(err)
		}
		if isLikeById == true {
			return nil
		}
		if err := redis.AddLike(userId, videoID); err != nil {
			klog.Error(err)
		}
	} else {
		// 如果redis没有这个userId的记录，则去mysql查询一次点赞列表进行缓存
		likeList, err := dal.MGetLikeList(s.ctx, userId)
		if err != nil {
			klog.Error(err)
		}
		if err := redis.AddLikeList(userId, likeList); err != nil {
			klog.Error(err)
			return err
		}
		isLikeById, err := redis.GetIsLikeById(userId, videoID)
		if err != nil {
			klog.Error(err)
		}
		if isLikeById == true {
			return nil
		}
		if err := redis.AddLike(userId, videoID); err != nil {
			klog.Error(err)
		}
	}
	if err := pulsar.LikeVideoProduce(s.ctx, userId, videoID); err != nil {
		return err
	}
	return nil
}
