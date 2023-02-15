package redis

import (
	"douyin/common/conf"
	"douyin/common/constant"
	"douyin/common/util"
	"fmt"
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

func TestAddRelation(t *testing.T) {
	testInit()
	//ctrl := gomock.NewController(t)

	err := AddRelation(1, 3)
	if err != nil {
		fmt.Println(err)
	}
}

func TestDeleteRelation(t *testing.T) {
	testInit()
	err := DeleteRelation(1, 3)
	if err != nil {
		fmt.Println(err)
	}
}

func TestAddFollowList(t *testing.T) {
	testInit()
	followList := []int64{2, 3, 4, 5, 6}
	err := AddFollowList(1, followList)
	if err != nil {
		fmt.Println(err)
	}

}

func TestGetFollowList(t *testing.T) {
	testInit()
	result, err := GetFollowList(38)
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

func TestIsFollowKeyExist(t *testing.T) {
	testInit()
	result, err := IsFollowKeyExist(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func TestAddBloomKey(t *testing.T) {
	testInit()
	err := AddBloomKey(constant.UserInfoRedisPrefix, 1)
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsKeyExistByBloom(t *testing.T) {
	testInit()
	result, err := IsKeyExistByBloom(constant.UserInfoRedisPrefix, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
