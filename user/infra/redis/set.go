package redis

import (
	"douyin/user/infra/redis/model"
	"encoding/json"
	"strconv"
	"time"
)

func AddRelation(userId int64, toUserId int64) error {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := "follow:" + strconv.FormatInt(userId, 10)

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

	key = "fan:" + strconv.FormatInt(toUserId, 10)
	_, err = redisConn.Do("zadd", key, time, userId)
	if err != nil {
		redisConn.Do("del", key)
		return err
	}
	_, err = redisConn.Do("expire", key, expireTimeUtil.GetRandTime())
	if err != nil {
		return err
	}

	return nil
}

func DeleteRelation(userId int64, toUserId int64) error {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := "follow:" + strconv.FormatInt(userId, 10)

	_, err := redisConn.Do("zrem", key, toUserId)
	if err != nil {
		redisConn.Do("del", key)
		return err
	}
	_, err = redisConn.Do("expire", key, expireTimeUtil.GetRandTime())
	if err != nil {
		return err
	}

	key = "fan:" + strconv.FormatInt(toUserId, 10)
	_, err = redisConn.Do("zrem", key, userId)
	if err != nil {
		redisConn.Do("del", key)
		return err
	}
	_, err = redisConn.Do("expire", key, expireTimeUtil.GetRandTime())
	if err != nil {
		return err
	}

	return nil
}

func AddUserInfo(userinfo model.UserRedis) error {

	ub, err := json.Marshal(userinfo)
	if err != nil {
		return err
	}

	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := "userinfo:" + strconv.FormatInt(userinfo.UserId, 10)
	_, err = redisConn.Do("set", key, ub, "ex", expireTimeUtil.GetRandTime())
	if err != nil {
		return err
	}

	return nil
}
