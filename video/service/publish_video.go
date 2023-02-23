package service

// @path: video/service/publish_video.go
// @description: CreateVideo service of video
// @author: Chongzhi <dczdcz2001@aliyun.com>
import (
	"context"
	"douyin/code_gen/kitex_gen/videoproto"
	"douyin/video/infra/dal"
	"douyin/video/infra/redis"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/opentracing/opentracing-go"
)

type CreateVideoService struct {
	ctx context.Context
}

func NewCreateVideoService(ctx context.Context) *CreateVideoService {
	return &CreateVideoService{ctx: ctx}
}

func (s *CreateVideoService) CreateVideo(req *videoproto.CreateVideoReq) error {
	span := Tracer.StartSpan("publish")
	defer span.Finish()
	s.ctx = opentracing.ContextWithSpan(s.ctx, span)
	// 投稿的时候，直接删除缓存里面的pulish:id
	// 因为redis和数据库不一致，方便起见直接删除，下次获取投稿列表的时候再缓存
	// 因为投稿是个低频操作，所以可以这样处理；点赞则不同
	if err := redis.DelPublishList(req.VideoBaseInfo.UserId); err != nil {
		klog.Error(err)
	}
	// 如果添加失败，返回error
	return dal.CreateVideo(s.ctx, req.VideoBaseInfo.UserId, req.VideoBaseInfo.Title, req.VideoBaseInfo.PlayUrl, req.VideoBaseInfo.CoverUrl)
}
