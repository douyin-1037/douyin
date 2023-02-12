package test

// @path: video/service/test/publish_video_test.go
// @description: CreateVideo service test
// @author: Chongzhi <dczdcz2001@aliyun.com>
import (
	"context"
	"douyin/code_gen/kitex_gen/videoproto"
	"douyin/video/service"
	"testing"
)

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
	err := service.NewCreateVideoService(context.Background()).CreateVideo(req)
	if err != nil {
		panic(err)
	}
}
