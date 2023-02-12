package test

// @path: video/service/test/get_publish_list_test.go
// @description: GetVideoByUserId service test
// @author: Chongzhi <dczdcz2001@aliyun.com>
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
