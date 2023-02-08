package redis

import (
	"douyin/common/conf"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var redisPool *redis.Pool

var expireTimeUtil ExpireTimeUtil

func Init() {
	redisPool = &redis.Pool{
		MaxIdle:   conf.Redis.MaxIdle,
		MaxActive: conf.Redis.MaxActive,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", conf.Redis.Address)
		},
	}
	expireTimeUtil = ExpireTimeUtil{
		expireTime:     conf.Redis.ExpireTime,
		maxRandAddTime: conf.Redis.MaxRandAddTime,
	}
	fmt.Println(expireTimeUtil.maxRandAddTime)
}
