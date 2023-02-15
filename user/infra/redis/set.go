package redis

import (
	"douyin/common/constant"
	redisModel "douyin/user/infra/redis/model"
	"encoding/json"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
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

	expireTime := expireTimeUtil.GetRandTime()
	_, err = redisConn.Do("expire", key, expireTime)
	if err != nil {
		return err
	}

	key = constant.FanRedisPrefix + strconv.FormatInt(toUserId, 10)
	_, err = redisConn.Do("zadd", key, time, userId)
	if err != nil {
		redisConn.Do("del", key)
		return err
	}
	_, err = redisConn.Do("expire", key, expireTime)
	if err != nil {
		return err
	}

	followCntKey := constant.UserInfoCntRedisPrefix + strconv.FormatInt(userId, 10)
	err = incrCount(redisConn, followCntKey, constant.FollowCountRedisPrefix, 1, expireTime)
	if err != nil {
		klog.Error(err)
	}
	fanCntKey := constant.UserInfoCntRedisPrefix + strconv.FormatInt(toUserId, 10)
	err = incrCount(redisConn, fanCntKey, constant.FanCountRedisPrefix, 1, expireTime)
	if err != nil {
		klog.Error(err)
	}
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

	expireTime := expireTimeUtil.GetRandTime()
	_, err = redisConn.Do("expire", key, expireTime)
	if err != nil {
		return err
	}

	key = constant.FanRedisPrefix + strconv.FormatInt(toUserId, 10)
	_, err = redisConn.Do("zrem", key, userId)
	if err != nil {
		redisConn.Do("del", key)
		return err
	}
	_, err = redisConn.Do("expire", key, expireTime)
	if err != nil {
		return err
	}

	followCntKey := constant.UserInfoCntRedisPrefix + strconv.FormatInt(userId, 10)
	err = incrCount(redisConn, followCntKey, constant.FollowCountRedisPrefix, -1, expireTime)
	if err != nil {
		klog.Error(err)
	}
	fanCntKey := constant.UserInfoCntRedisPrefix + strconv.FormatInt(toUserId, 10)
	err = incrCount(redisConn, fanCntKey, constant.FanCountRedisPrefix, -1, expireTime)
	if err != nil {
		klog.Error(err)
	}

	return nil
}

func AddUserInfo(userInfo redisModel.UserInfoRedis, userCntInfo redisModel.UserCntRedis) error {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	ub, err := json.Marshal(userInfo)
	if err != nil {
		return err
	}

	key := constant.UserInfoRedisPrefix + strconv.FormatInt(userInfo.UserId, 10)
	expireTime := expireTimeUtil.GetRandTime()
	_, err = redisConn.Do("set", key, ub, "ex", expireTime)
	if err != nil {
		return err
	}

	countKey := constant.UserInfoCntRedisPrefix + strconv.FormatInt(userInfo.UserId, 10)
	_, err = redisConn.Do("hset", countKey,
		constant.FollowCountRedisPrefix, userCntInfo.FollowCnt,
		constant.FanCountRedisPrefix, userCntInfo.FanCnt)
	if err != nil {
		return err
	}
	_, err = redisConn.Do("expire", countKey, expireTime)
	if err != nil {
		klog.Error(err)
	}
	return nil
}

func AddFollowList(userId int64, FollowIdList []int64) error {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	if FollowIdList == nil || len(FollowIdList) == 0 {
		return nil
	}

	key := constant.FollowRedisPrefix + strconv.FormatInt(userId, 10)

	l := len(FollowIdList) //用于计分倒序
	redisConn.Send("multi")
	for i, followId := range FollowIdList {
		redisConn.Send("zadd", key, l-i, followId)
	}

	_, err := redisConn.Do("exec")
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

func AddFanList(userId int64, FanIdList []int64) error {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	if FanIdList == nil || len(FanIdList) == 0 {
		return nil
	}

	key := constant.FanRedisPrefix + strconv.FormatInt(userId, 10)

	l := len(FanIdList) //用于计分倒序
	redisConn.Send("multi")
	for i, fanId := range FanIdList {
		redisConn.Send("zadd", key, l-i, fanId)
	}

	_, err := redisConn.Do("exec")
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

func AddBloomKey(prefix string, keyId int64) error {
	redisConn := redisPool.Get()
	defer redisConn.Close()
	_, err := redisConn.Do("bf.add", prefix, keyId)
	if err != nil {
		klog.Error(err)
		return err
	}
	return nil
}

func incrCount(redisConn redis.Conn, cntKey string, filed string, v int, expireTime int) error {
	result, err := redis.Strings(redisConn.Do("keys", cntKey))
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return nil
	}
	if v < 0 {
		cnt, cntErr := redis.Int64(redisConn.Do("hget", cntKey, filed))
		if cntErr != nil {
			redisConn.Do("del", cntKey)
			return cntErr
		}
		if cnt <= 0 {
			redisConn.Do("del", cntKey)
			errMsg := "del CountKey: " + cntKey + "error: like count can not lower than 0"
			return errors.New(errMsg)
		}
	}
	_, err = redisConn.Do("hincrby", cntKey, filed, v)
	if err != nil {
		redisConn.Do("del", cntKey)
		return err
	}
	_, err = redisConn.Do("expire", cntKey, expireTime)
	if err != nil {
		klog.Error(err)
	}
	return nil
}
