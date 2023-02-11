package test

import (
	"context"
	"douyin/code_gen/kitex_gen/videoproto"
	"douyin/video/service"
	"fmt"
	"testing"
)

func TestMGetVideo(t *testing.T) {
	testInit()
	req := &videoproto.GetVideoListByUserIdReq{
		UserId: int64(5),
	}
	videoInfos, err := service.NewMGetVideoByUserIdService(context.Background()).MGetVideo(req)
	fmt.Printf("%v\n", videoInfos)
	if err != nil {
		panic(err)
	}
}
