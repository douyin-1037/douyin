package test

// @path: video/service/test/init.go
// @description: init test
// @author: Chongzhi <dczdcz2001@aliyun.com>
import (
	"douyin/video/infra/dal"
	"douyin/video/infra/redis"
)

func testInit() {
	dal.Init()
	redis.Init()
}
