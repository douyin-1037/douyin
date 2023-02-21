package service

import (
	"context"
	"douyin/code_gen/kitex_gen/commentproto"
	"douyin/comment/infra/dal"
	"douyin/comment/infra/pulsar"
	"douyin/comment/infra/redis"
	"douyin/common/conf"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"testing"
)

func initCreateCommentTest() {
	conf.InitConfig()
	dal.Init()
	redis.Init()
	pulsar.Init()
}

func TestCreateCommentService(t *testing.T) {
	initCreateCommentTest()
	ctx := context.Background()
	req := &commentproto.CreateCommentReq{
		UserId:  15,
		VideoId: 14,
		Content: "test comment",
	}
	comments, err := NewCreateCommentService(ctx).CreateComment(req.UserId, req.VideoId, req.Content)
	if err != nil {
		klog.Error(err.Error())
	}
	fmt.Println(comments)
}
