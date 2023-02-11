package test

// @path: video/service/test/unlike_video_test.go
// @description: UnLike service test
// @author: Chongzhi <dczdcz2001@aliyun.com>
import (
	"context"
	"douyin/code_gen/kitex_gen/videoproto"
	"douyin/video/service"
	"testing"
)

func TestUnLikeVideo(t *testing.T) {
	testInit()
	req := &videoproto.UnLikeVideoReq{
		UserId:  int64(1),
		VideoId: int64(8),
	}
	err := service.NewUnLikeVideoService(context.Background()).UnLikeVideo(req)
	if err != nil {
		panic(err)
	}
}
