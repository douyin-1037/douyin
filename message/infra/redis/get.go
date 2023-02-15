package redis

import (
	"douyin/common/constant"
	redisModel "douyin/message/infra/redis/model"
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"strconv"
)

func GetMessageList(userId int64, toUserId int64, latestTime int64, nowTime int64) ([]redisModel.MessageRedis, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	var key string
	if userId < toUserId {
		key = constant.MessageRedisPrefix + strconv.FormatInt(userId, 10) + ":" + strconv.FormatInt(toUserId, 10)
	} else {
		key = constant.MessageRedisPrefix + strconv.FormatInt(toUserId, 10) + ":" + strconv.FormatInt(userId, 10)
	}

	result, err := redis.Strings(redisConn.Do("zrangebyscore", key, latestTime, nowTime))
	if err != nil {
		if errors.Is(err, redis.ErrNil) {
			return make([]redisModel.MessageRedis, 0), nil
		}
		return nil, err
	}

	messageList := make([]redisModel.MessageRedis, len(result))
	for i := range result {
		messageInfo := new(redisModel.MessageRedis)
		data := []byte(result[i])
		err = json.Unmarshal(data, messageInfo)
		if err != nil {
			return nil, err
		}
		messageList[i] = *messageInfo
	}

	expireTime := expireTimeUtil.GetRandTime()
	_, err = redisConn.Do("expire", key, expireTime)
	if err != nil {
		return nil, err
	}
	return messageList, nil
}

func IsMessageKeyExist(userId int64, toUserId int64) (bool, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	var key string
	if userId < toUserId {
		key = constant.MessageRedisPrefix + strconv.FormatInt(userId, 10) + ":" + strconv.FormatInt(toUserId, 10)
	} else {
		key = constant.MessageRedisPrefix + strconv.FormatInt(toUserId, 10) + ":" + strconv.FormatInt(userId, 10)
	}
	result, err := redis.Strings(redisConn.Do("keys", key))
	if err != nil {
		return false, err
	}
	if len(result) == 0 {
		return false, nil
	}
	expireTime := expireTimeUtil.GetRandTime()
	_, err = redisConn.Do("expire", key, expireTime)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetMessageLatestTime(userId int64, toUserId int64) (int64, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := constant.MessageLatestTimeRedisPrefix + strconv.FormatInt(userId, 10)
	result, err := redis.Int64(redisConn.Do("hget", key, toUserId))
	if err != nil {
		if errors.Is(err, redis.ErrNil) {
			return 0, nil
		}
		return 0, err
	}
	return result, nil
}
