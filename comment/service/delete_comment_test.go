package service

import (
	"context"
	"douyin/code_gen/kitex_gen/commentproto"
	"douyin/comment/infra/dal"
	"douyin/comment/infra/redis"
	"douyin/common/conf"
	"github.com/cloudwego/kitex/pkg/klog"
	"testing"
)

func initDeleteCommentTest() {
	conf.InitConfig()
	dal.Init()
	redis.Init()
}

func TestDeleteCommentService(t *testing.T) {
	initDeleteCommentTest()
	ctx := context.Background()
	req := &commentproto.DeleteCommentReq{
		CommentId: 447637918983389184,
		VideoId:   7,
	}
	err := NewDeleteCommentService(ctx).DeleteComment(req.CommentId, req.VideoId)
	if err != nil {
		klog.Error(err.Error())
	}
	//fmt.Println(comments)
}
