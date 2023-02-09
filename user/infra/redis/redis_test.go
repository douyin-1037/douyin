package redis

import (
	"douyin/common/conf"
	"douyin/user/infra/redis/model"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"testing"
)

func testInit() {
	conf.InitConfig()
	expireTimeUtil = ExpireTimeUtil{
		expireTime:     conf.Redis.ExpireTime,
		maxRandAddTime: conf.Redis.MaxRandAddTime,
	}
	redisPool = &redis.Pool{
		MaxIdle:   conf.Redis.MaxIdle,
		MaxActive: conf.Redis.MaxActive,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", conf.Redis.Address)
		},
	}
}

func TestAddRelation(t *testing.T) {
	testInit()
	err := AddRelation(2, 3)
	if err != nil {
		fmt.Println(err)
	}
}

func TestDeleteRelation(t *testing.T) {
	testInit()
	err := DeleteRelation(2, 3)
	if err != nil {
		fmt.Println(err)
	}
}

func TestGetFollowList(t *testing.T) {
	testInit()
	result, err := GetFollowList(1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)

}

func TestGetFanList(t *testing.T) {
	testInit()
	result, err := GetFanList(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func TestGetFriendList(t *testing.T) {
	testInit()
	result, err := GetFriendList(2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func TestGetIsFollowById(t *testing.T) {
	testInit()
	result, err := GetIsFollowById(2, 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func TestAddUserInfo(t *testing.T) {
	testInit()
	userinfo := model.UserRedis{
		UserId:   1,
		UserName: "yuirito",
	}
	err := AddUserInfo(userinfo)
	if err != nil {
		fmt.Println(err)
	}
}

func TestGetUserInfo(t *testing.T) {
	testInit()

	userinfo, err := GetUserInfo(4)
	if err != nil {
		fmt.Println("err:")
		fmt.Println(err)
	}
	fmt.Println(userinfo)
}
