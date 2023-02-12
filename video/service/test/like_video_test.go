package test

// @path: video/service/test/like_video_test.go
// @description: LikeVideo service test
// @author: Chongzhi <dczdcz2001@aliyun.com>
import (
	"context"
	"douyin/code_gen/kitex_gen/videoproto"
	"douyin/video/service"
	"testing"
)

func TestLikeVideo(t *testing.T) {
	testInit()
	req := &videoproto.LikeVideoReq{
		UserId:  int64(1),
		VideoId: int64(8),
	}
	err := service.NewLikeVideoService(context.Background()).LikeVideo(req)
	if err != nil {
		panic(err)
	}
}
