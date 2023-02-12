package redis

import (
	"douyin/common/constant"
	"douyin/user/infra/redis/model"
	"encoding/json"
	"strconv"
	"time"
)

func AddRelation(userId int64, toUserId int64) error {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := constant.FollowRedisPrefix + strconv.FormatInt(userId, 10)

	now := time.Now()
	time := now.Unix()

	_, err := redisConn.Do("zadd", key, time, toUserId)
	if err != nil {
		redisConn.Do("del", key)
		return err
	}
	_, err = redisConn.Do("expire", key, expireTimeUtil.GetRandTime())
	if err != nil {
		return err
	}

	key = constant.FanRedisPrefix + strconv.FormatInt(toUserId, 10)
	_, err = redisConn.Do("zadd", key, time, userId)
	if err != nil {
		redisConn.Do("del", key)
		return err
	}
	_, err = redisConn.Do("expire", key, expireTimeUtil.GetRandTime())
	if err != nil {
		return err
	}
	userInfoKey := constant.UserInfoRedisPrefix + strconv.FormatInt(userId, 10)
	redisConn.Do("del", userInfoKey)
	userInfoKey = constant.UserInfoRedisPrefix + strconv.FormatInt(toUserId, 10)
	redisConn.Do("del", userInfoKey)

	return nil
}

func DeleteRelation(userId int64, toUserId int64) error {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := constant.FollowRedisPrefix + strconv.FormatInt(userId, 10)

	_, err := redisConn.Do("zrem", key, toUserId)
	if err != nil {
		redisConn.Do("del", key)
		return err
	}
	_, err = redisConn.Do("expire", key, expireTimeUtil.GetRandTime())
	if err != nil {
		return err
	}

	key = constant.FanRedisPrefix + strconv.FormatInt(toUserId, 10)
	_, err = redisConn.Do("zrem", key, userId)
	if err != nil {
		redisConn.Do("del", key)
		return err
	}
	_, err = redisConn.Do("expire", key, expireTimeUtil.GetRandTime())
	if err != nil {
		return err
	}
	userInfoKey := constant.UserInfoRedisPrefix + strconv.FormatInt(userId, 10)
	redisConn.Do("del", userInfoKey)
	userInfoKey = constant.UserInfoRedisPrefix + strconv.FormatInt(toUserId, 10)
	redisConn.Do("del", userInfoKey)

	return nil
}

func AddUserInfo(userinfo model.UserRedis) error {

	ub, err := json.Marshal(userinfo)
	if err != nil {
		return err
	}

	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := constant.UserInfoRedisPrefix + strconv.FormatInt(userinfo.UserId, 10)
	_, err = redisConn.Do("set", key, ub, "ex", expireTimeUtil.GetRandTime())
	if err != nil {
		return err
	}

	return nil
}

func AddFollowList(userId int64, FollowIdList []int64) error {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := constant.FollowRedisPrefix + strconv.FormatInt(userId, 10)

	l := len(FollowIdList)

	for i, followId := range FollowIdList {
		_, err := redisConn.Do("zadd", key, l-i, followId)
		if err != nil {
			redisConn.Do("del", key)
			return err
		}
	}

	_, err := redisConn.Do("expire", key, expireTimeUtil.GetRandTime())
	if err != nil {
		return err
	}

	return nil
}

func AddFanList(userId int64, FanIdList []int64) error {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := constant.FanRedisPrefix + strconv.FormatInt(userId, 10)

	l := len(FanIdList)

	for i, fanId := range FanIdList {
		_, err := redisConn.Do("zadd", key, l-i, fanId)
		if err != nil {
			redisConn.Do("del", key)
			return err
		}
	}

	_, err := redisConn.Do("expire", key, expireTimeUtil.GetRandTime())
	if err != nil {
		return err
	}

	return nil
}
