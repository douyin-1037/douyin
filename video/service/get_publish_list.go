package service

// @path: video/service/get_publish_list.go
// @description: GetVideoByUserId service of video
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
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

type MGetVideoByUserIdService struct {
	ctx context.Context
}

func NewMGetVideoByUserIdService(ctx context.Context) *MGetVideoByUserIdService {
	return &MGetVideoByUserIdService{ctx: ctx}
}

// MGetVideo 通过UserID从DAO层获取视频基本信息，并查出当前用户是否点赞，组装后返回
func (s *MGetVideoByUserIdService) MGetVideo(req *videoproto.GetVideoListByUserIdReq) ([]*videoproto.VideoInfo, error) {
	span := Tracer.StartSpan("get_publish_list")
	defer span.Finish()
	s.ctx = opentracing.ContextWithSpan(s.ctx, span)
	userId := req.UserId
	// 只能得到视频id,uid,title，play_url,cover_url,created_time
	var videoModels []*model.Video
	videoModels, err := redis.GetPublishList(userId)
	if err != nil {
		if errors.Is(err, goredis.ErrNil) == false {
			return nil, err
		}
		// 缓存未命中，去数据库查询，然后缓存到redis
		videoModels, err = dal.MGetVideoByUserID(s.ctx, userId)
		if err != nil {
			return nil, err
		}
		if err := redis.AddPublishList(videoModels, userId); err != nil {
			klog.Error(err)
		}
	}
	videos := pack.Videos(videoModels) // 做类型转换：视频id、base_info、点赞数、评论数已经得到，还需要判断是否点赞
	// 把视频的其他信息进行绑定
	appUserID := req.AppUserId
	isLikeKeyExist, err := redis.IsLikeKeyExist(appUserID)
	if err != nil {
		klog.Error(err)
	}
	if isLikeKeyExist == false {
		// 如果redis没有appUserID的记录，则去mysql查询一次点赞列表进行缓存
		likeList, err := dal.MGetLikeList(s.ctx, appUserID)
		if err != nil {
			return nil, err
		}
		if err := redis.AddLikeList(appUserID, likeList); err != nil {
			klog.Error(err)
		}
	}
	for i := 0; i < len(videos); i++ {
		isFavorite, err := redis.GetIsLikeById(appUserID, videos[i].VideoId)
		if err != nil {
			isFavorite, _ = dal.IsFavorite(s.ctx, videos[i].VideoId, appUserID)
		}
		videos[i].IsFavorite = isFavorite
	}
	return videos, nil
}
