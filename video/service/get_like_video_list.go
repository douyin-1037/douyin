package service

// @path: video/service/get_like_video_list.go
// @description: GetLikeVideo service of video
// @author: Chongzhi <dczdcz2001@aliyun.com>
import (
	"context"
	"douyin/code_gen/kitex_gen/videoproto"
	"douyin/video/infra/dal"
	"douyin/video/infra/dal/model"
	"douyin/video/infra/redis"
	"douyin/video/pack"
	"github.com/cloudwego/kitex/pkg/klog"
	goredis "github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
)

type MGetLikeVideoService struct {
	ctx context.Context
}

func NewMGetLikeVideoService(ctx context.Context) *MGetLikeVideoService {
	return &MGetLikeVideoService{ctx: ctx}
}

// MGetLikeVideo 通过用户ID从DAO层获取喜欢视频的基本信息，并查出当前用户是否点赞，组装后返回
func (s *MGetLikeVideoService) MGetLikeVideo(req *videoproto.GetLikeVideoListReq) ([]*videoproto.VideoInfo, error) {
	appUserId := req.AppUserId
	userId := req.UserId
	var likeList []int64
	isLikeKeyExist, err := redis.IsLikeKeyExist(userId)
	if err != nil {
		klog.Error(err)
	}
	if isLikeKeyExist == true {
		likeList, err = redis.GetLikeList(userId)
		if err != nil {
			return nil, err
		}
	} else {
		// 缓存未命中，去数据库查询，然后缓存到redis
		likeList, err = dal.MGetLikeList(s.ctx, userId)
		if err != nil {
			return nil, err
		}
		if err := redis.AddLikeList(userId, likeList); err != nil {
			klog.Error(err)
		}
	}
	var videoModels = make([]*model.Video, len(likeList))
	for i, videoID := range likeList {
		videoModels[i], err = redis.GetVideoInfo(videoID)
		if err != nil {
			if errors.Is(err, goredis.ErrNil) == false {
				return nil, err
			}
			// 缓存未命中，去数据库查询，然后缓存到redis
			videoModels[i], err = dal.MGetVideoInfo(s.ctx, videoID)
			if err != nil {
				return nil, err
			}
			if err := redis.AddVideoInfo(*videoModels[i]); err != nil {
				klog.Error(err)
			}
		}
	}
	videos := pack.Videos(videoModels) // 做类型转换：视频id、base_info、点赞数、评论数已经得到，还需要判断是否点赞
	// 把视频的其他信息进行绑定
	isLikeKeyExist, err = redis.IsLikeKeyExist(appUserId)
	if err != nil {
		klog.Error(err)
	}
	if isLikeKeyExist == false {
		likeList, err = dal.MGetLikeList(s.ctx, userId)
		if err != nil {
			return nil, err
		}
		if err := redis.AddLikeList(userId, likeList); err != nil {
			klog.Error(err)
		}
	}
	for i := 0; i < len(videos); i++ {
		isLike, err := redis.GetIsLikeById(appUserId, videos[i].VideoId)
		if err != nil {
			isLike, _ = dal.IsFavorite(s.ctx, videos[i].VideoId, appUserId)
		}
		videos[i].IsFavorite = isLike // 返回的视频都是已经点赞的
	}
	return videos, nil
}
