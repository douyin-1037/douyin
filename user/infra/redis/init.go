package redis

import (
	"douyin/common/conf"
	"douyin/common/util"
	"github.com/gomodule/redigo/redis"
)

var redisPool *redis.Pool

var expireTimeUtil util.ExpireTimeUtil
var bloomKeyOpen bool

func Init() {
	redisPool = &redis.Pool{
		MaxIdle:   conf.Redis.MaxIdle,
		MaxActive: conf.Redis.MaxActive,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", conf.Redis.Address)
		},
	}
	expireTimeUtil = util.ExpireTimeUtil{
		ExpireTime:     conf.Redis.ExpireTime,
		MaxRandAddTime: conf.Redis.MaxRandAddTime,
	}

	bloomKeyOpen = false

}
