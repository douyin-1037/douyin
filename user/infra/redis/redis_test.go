package redis

import (
	"context"
	"douyin/common/conf"
	"douyin/common/constant"
	"douyin/common/util"
	"douyin/user/infra/redis/model"
	"fmt"
	"github.com/gomodule/redigo/redis"
	. "github.com/smartystreets/goconvey/convey"
	"strconv"
	"testing"
	"time"
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
	bloomOpen = true
}

func TestAddRelation(t *testing.T) {
	testInit()
	Convey("TestAddRelation", t, func() {
		So(AddRelation(1, 3), ShouldBeNil)
	})
}

func TestDeleteRelation(t *testing.T) {
	testInit()
	Convey("TestDeleteRelation", t, func() {
		So(DeleteRelation(1, 3), ShouldBeNil)
	})
}

func TestAddFollowList(t *testing.T) {
	testInit()
	Convey("TestAddFollowList", t, func() {
		followList := []int64{2, 3, 4, 5, 6}
		So(AddFollowList(1, followList), ShouldBeNil)
	})
}

func TestGetFollowList(t *testing.T) {
	testInit()

	Convey("TestGetFollowList", t, func() {
		err := deleteTestKey(1, constant.FollowRedisPrefix)
		So(err, ShouldBeNil)
		followList := []int64{2, 3, 4, 5, 6}
		So(AddFollowList(1, followList), ShouldBeNil)
		result, err := GetFollowList(1)
		So(err, ShouldBeNil)
		So(result, ShouldResemble, followList)
		err = deleteTestKey(1, constant.FollowRedisPrefix)
		So(err, ShouldBeNil)
		result, err = GetFollowList(1)
		So(err, ShouldBeNil)
		So(result, ShouldBeEmpty)
	})
}

func TestGetFanList(t *testing.T) {
	testInit()
	Convey("TestGetFanList", t, func() {
		err := deleteTestKey(1, constant.FanRedisPrefix)
		So(err, ShouldBeNil)
		fanList := []int64{2, 3, 4, 5, 6}
		So(AddFanList(1, fanList), ShouldBeNil)
		result, err := GetFanList(1)
		So(err, ShouldBeNil)
		So(result, ShouldResemble, fanList)
		err = deleteTestKey(1, constant.FanRedisPrefix)
		So(err, ShouldBeNil)
		result, err = GetFanList(1)
		So(err, ShouldBeNil)
		So(result, ShouldBeEmpty)
	})
}

func TestGetFriendList(t *testing.T) {
	testInit()
	Convey("TestGetFriendList", t, func() {
		err := deleteTestKey(1, constant.FollowRedisPrefix)
		So(err, ShouldBeNil)
		err = deleteTestKey(1, constant.FanRedisPrefix)
		So(err, ShouldBeNil)
		followList := []int64{3, 4, 5, 6}
		So(AddFollowList(1, followList), ShouldBeNil)
		fanList := []int64{2, 3, 4, 5}
		So(AddFanList(1, fanList), ShouldBeNil)
		result, err := GetFriendList(1)
		So(err, ShouldBeNil)
		So(result, ShouldResemble, []int64{5, 4, 3})
		err = deleteTestKey(1, constant.FollowRedisPrefix)
		So(err, ShouldBeNil)
		err = deleteTestKey(1, constant.FanRedisPrefix)
		So(err, ShouldBeNil)
		result, err = GetFriendList(1)
		So(err, ShouldBeNil)
		So(result, ShouldBeEmpty)
	})
}

func TestGetIsFollowById(t *testing.T) {
	testInit()
	Convey("TestGetIsFollowById", t, func() {
		So(AddRelation(1, 3), ShouldBeNil)
		result, err := GetIsFollowById(1, 3)
		So(err, ShouldBeNil)
		So(result, ShouldBeTrue)
		So(DeleteRelation(1, 3), ShouldBeNil)
		result, err = GetIsFollowById(1, 3)
		So(err, ShouldBeNil)
		So(result, ShouldBeFalse)
	})
}

func TestAddUserInfo(t *testing.T) {
	testInit()
	Convey("TestAddUserInfo", t, func() {
		So(AddUserInfo(model.UserInfoRedis{
			UserId:   1,
			UserName: "test_name",
		}, model.UserCntRedis{
			FollowCnt:   1,
			FanCnt:      2,
			WorkCnt:     3,
			FavoriteCnt: 4,
		}), ShouldBeNil)
	})
}

func TestGetUserInfo(t *testing.T) {
	testInit()
	Convey("TestGetUserInfo", t, func() {
		So(AddUserInfo(model.UserInfoRedis{
			UserId:   1,
			UserName: "test_name",
		}, model.UserCntRedis{
			FollowCnt:   1,
			FanCnt:      2,
			WorkCnt:     3,
			FavoriteCnt: 4,
		}), ShouldBeNil)
		result, err := GetUserInfo(1)
		So(err, ShouldBeNil)
		So(result, ShouldResemble, &model.UserRedis{
			UserId:      1,
			UserName:    "test_name",
			FollowCnt:   1,
			FanCnt:      2,
			WorkCnt:     3,
			FavoriteCnt: 4,
		})
		err = deleteTestKey(1, constant.UserInfoRedisPrefix)
		So(err, ShouldBeNil)
		result, err = GetUserInfo(1)
		So(err, ShouldNotBeNil)
	})
}

func TestIsFollowKeyExist(t *testing.T) {
	testInit()
	Convey("TestIsFollowKeyExist", t, func() {
		So(deleteTestKey(1, constant.FollowRedisPrefix), ShouldBeNil)
		result, err := IsFollowKeyExist(1)
		So(err, ShouldBeNil)
		So(result, ShouldBeFalse)
		So(AddRelation(1, 3), ShouldBeNil)
		result, err = IsFollowKeyExist(1)
		So(err, ShouldBeNil)
		So(result, ShouldBeTrue)
		So(deleteTestKey(1, constant.FollowRedisPrefix), ShouldBeNil)
	})
}

func TestAddBloomKey(t *testing.T) {
	testInit()
	Convey("TestAddBloomKey", t, func() {
		So(AddBloomKey(constant.UserInfoRedisPrefix, 1), ShouldBeNil)
	})
}

func TestIsKeyExistByBloom(t *testing.T) {
	testInit()
	Convey("TestIsKeyExistByBloom", t, func() {
		So(AddBloomKey(constant.UserInfoRedisPrefix, 1), ShouldBeNil)
		result, err := IsKeyExistByBloom(constant.UserInfoRedisPrefix, 1)
		So(err, ShouldBeNil)
		So(result, ShouldBeTrue)
	})
}

func deleteTestKey(keyId int64, prefix string) error {
	redisConn := redisPool.Get()
	defer redisConn.Close()
	key := prefix + strconv.FormatInt(keyId, 10)
	_, err := redisConn.Do("del", key)
	return err
}

func TestDistributedLock_TryLock(t *testing.T) {
	testInit()
	redisConn := redisPool.Get()
	defer redisConn.Close()
	lock := DistributedLock{
		TTL:         60,
		Key:         "testkey",
		RandomValue: 10,
	}
	fmt.Println(lock.TryLock(redisConn))
}

func TestDistributedLock_Unlock(t *testing.T) {
	testInit()
	lock := DistributedLock{
		TTL:             60,
		Key:             "testkey",
		RandomValue:     100,
		TryLockInterval: time.Duration(100),
		watchDog:        make(chan bool),
	}
	fmt.Println(lock.Unlock())
}

func TestDistributedLock_Lock(t *testing.T) {
	testInit()
	lock := DistributedLock{
		TTL:             60,
		Key:             "testkey",
		RandomValue:     100,
		TryLockInterval: time.Duration(100),
		watchDog:        make(chan bool),
	}
	ctx := context.Background()
	fmt.Println(lock.Lock(ctx))
}
