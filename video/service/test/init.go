package test

import (
	"douyin/video/infra/dal"
	"douyin/video/infra/redis"
)

func testInit() {
	dal.Init()
	redis.Init()
}
