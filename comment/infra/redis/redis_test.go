package redis

import (
	"douyin/common/conf"
	"douyin/common/util"
	"github.com/gomodule/redigo/redis"
	"testing"
)

func testInit() {
	conf.InitConfig()
	expireTimeUtil = util.ExpireTimeUtil{
		ExpireTime:     conf.Redis.ExpireTime,
		MaxRandAddTime: conf.Redis.MaxRandAddTime,
	}
	redisPool = &redis.Pool{
		MaxIdle:   conf.Redis.MaxIdle,
		MaxActive: conf.Redis.MaxActive,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", conf.Redis.Address)
		},
	}
}

func TestGetCommentList(t *testing.T) {
	testInit()
	GetCommentList(14)
}
