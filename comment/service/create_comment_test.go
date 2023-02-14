package service

import (
	"context"
	"douyin/code_gen/kitex_gen/commentproto"
	"douyin/comment/infra/dal"
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
}

func TestCreateCommentService(t *testing.T) {
	initCreateCommentTest()
	ctx := context.Background()
	req := &commentproto.CreateCommentReq{
		UserId:  22,
		VideoId: 7,
		Content: "redis测试完毕",
	}
	comments, err := NewCreateCommentService(ctx).CreateComment(req.UserId, req.VideoId, req.Content)
	if err != nil {
		klog.Error(err.Error())
	}
	fmt.Println(comments)
}
