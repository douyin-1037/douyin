package service

import (
	"context"
	"douyin/code_gen/kitex_gen/videoproto"
	"douyin/video/infra/dal"
	"douyin/video/infra/redis"
	"testing"
)

func testInit() {
	dal.Init()
	redis.Init()
}

func TestCreateVideo(t *testing.T) {
	testInit()
	req := &videoproto.CreateVideoReq{
		VideoBaseInfo: &videoproto.VideoBaseInfo{
			UserId:   int64(5),
			PlayUrl:  "testPlayUrl",
			CoverUrl: "testCoverUrl",
			Title:    "testTitle",
		},
	}
	err := NewCreateVideoService(context.Background()).CreateVideo(req)
	if err != nil {
		panic(err)
	}
}
