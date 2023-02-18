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

func initGetCommentListTest() {
	conf.InitConfig()
	dal.Init()
	redis.Init()
}

func TestGetCommentListService(t *testing.T) {
	initGetCommentListTest()
	ctx := context.Background()
	req := &commentproto.GetCommentListReq{VideoId: 3}
	comments, err := NewGetCommentListService(ctx).GetCommentList(req.VideoId)
	if err != nil {
		klog.Error(err.Error())
	}
	fmt.Println(comments)
}
