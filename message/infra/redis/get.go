package redis

import (
	"douyin/common/constant"
	redisModel "douyin/message/infra/redis/model"
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"strconv"
)

func GetMessageList(userId int64, toUserId int64) ([]redisModel.MessageRedis, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	var key string
	if userId < toUserId {
		key = constant.MessageRedisPrefix + strconv.FormatInt(userId, 10) + ":" + strconv.FormatInt(toUserId, 10)
	} else {
		key = constant.MessageRedisPrefix + strconv.FormatInt(toUserId, 10) + ":" + strconv.FormatInt(userId, 10)
	}

	result, err := redis.Strings(redisConn.Do("zrevrange", key, 0, -1))
	if err != nil {
		return nil, err
	}
	if len(result) <= 0 {
		return nil, redis.ErrNil
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
