package test

import (
	"context"
	"douyin/code_gen/kitex_gen/videoproto"
	"douyin/video/service"
	"fmt"
	"testing"
)

func TestMGetLikeVideo(t *testing.T) {
	testInit()
	req := &videoproto.GetLikeVideoListReq{
		AppUserId: int64(19),
	}
	videoInfos, err := service.NewMGetLikeVideoService(context.Background()).MGetLikeVideo(req)
	fmt.Printf("%v\n", videoInfos)
	if err != nil {
		panic(err)
	}
}
