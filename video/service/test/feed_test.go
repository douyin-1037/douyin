package test

// @path: video/service/test/feed_test.go
// @description: GetVideoByTime service test
// @author: Chongzhi <dczdcz2001@aliyun.com>
import (
	"context"
	"douyin/code_gen/kitex_gen/videoproto"
	"douyin/video/service"
	"fmt"
	"testing"
	"time"
)

func TestMGetVideoByTime(t *testing.T) {
	testInit()
	req := &videoproto.GetVideoListByTimeReq{
		AppUserId:  int64(1),
		LatestTime: time.Now().Unix(),
		Count:      5,
	}
	videos, nextTime, err := service.NewMGetVideoByTimeService(context.Background()).MGetVideoByTime(req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", videos)
	fmt.Println("nextTime: ", nextTime)
}
