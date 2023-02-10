package redis

import (
	"douyin/user/infra/redis/model"
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"strconv"
)

func GetFollowList(userId int64) ([]int64, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := "follow:" + strconv.FormatInt(userId, 10)
	result, err := redis.Int64s(redisConn.Do("zrevrange", key, 0, -1))
	if err != nil {
		return nil, err
	}
	return result, nil

}

func GetFanList(userId int64) ([]int64, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := "fan:" + strconv.FormatInt(userId, 10)
	result, err := redis.Int64s(redisConn.Do("zrevrange", key, 0, -1))
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetFriendList(userid int64) ([]int64, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()
	followkey := "follow:" + strconv.FormatInt(userid, 10)
	fankey := "fan:" + strconv.FormatInt(userid, 10)
	result, err := redis.Int64s(redisConn.Do("zinter", 2, followkey, fankey))
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetIsFollowById(userid int64, followid int64) (bool, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := "follow:" + strconv.FormatInt(userid, 10)
	followkey := strconv.FormatInt(followid, 10)
	result, err := redisConn.Do("zscore", key, followkey)
	if err != nil {
		return false, err
	}
	if result == nil {
		return false, nil
	}
	return true, nil
}

func GetUserInfo(userid int64) (*model.UserRedis, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := "userinfo:" + strconv.FormatInt(userid, 10)
	result, err := redis.Bytes(redisConn.Do("get", key))
	if err != nil {
		return nil, err
	}
	userinfo := new(model.UserRedis)
	err = json.Unmarshal(result, userinfo)
	if err != nil {
		return nil, err
	}
	return userinfo, nil
}

func IsFollowKeyExist(userId int64) (bool, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()
	key := "follow:" + strconv.FormatInt(userId, 10)
	result, err := redis.Strings(redisConn.Do("keys", key))
	if err != nil {
		return false, err
	}
	if len(result) == 0 {
		return false, nil
	}
	return true, nil
}

func IsFanKeyExist(userId int64) (bool, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()
	key := "fan:" + strconv.FormatInt(userId, 10)
	result, err := redis.Strings(redisConn.Do("keys", key))
	if err != nil {
		return false, err
	}
	if len(result) == 0 {
		return false, nil
	}
	return true, nil
}
