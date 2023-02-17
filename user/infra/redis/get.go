package redis

import (
	"douyin/common/constant"
	"douyin/user/infra/redis/model"
	"encoding/json"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gomodule/redigo/redis"
	"strconv"
)

func GetFollowList(userId int64) ([]int64, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := constant.FollowRedisPrefix + strconv.FormatInt(userId, 10)
	result, err := redis.Int64s(redisConn.Do("zrevrange", key, 0, -1))
	if err != nil {
		return nil, err
	}

	expireTime := expireTimeUtil.GetRandTime()
	_, err = redisConn.Do("expire", key, expireTime)
	if err != nil {
		return result, err
	}
	return result, nil

}

func GetFanList(userId int64) ([]int64, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := constant.FanRedisPrefix + strconv.FormatInt(userId, 10)
	result, err := redis.Int64s(redisConn.Do("zrevrange", key, 0, -1))
	if err != nil {
		return nil, err
	}

	expireTime := expireTimeUtil.GetRandTime()
	_, err = redisConn.Do("expire", key, expireTime)
	if err != nil {
		return result, err
	}

	return result, nil
}

func GetFriendList(userId int64) ([]int64, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()
	followkey := constant.FollowRedisPrefix + strconv.FormatInt(userId, 10)
	fankey := constant.FanRedisPrefix + strconv.FormatInt(userId, 10)
	result, err := redis.Int64s(redisConn.Do("zinter", 2, followkey, fankey))
	if err != nil {
		return nil, err
	}

	expireTime := expireTimeUtil.GetRandTime()
	_, err = redisConn.Do("expire", followkey, expireTime)
	if err != nil {
		return result, err
	}
	_, err = redisConn.Do("expire", fankey, expireTime)
	if err != nil {
		return result, err
	}

	return result, nil
}

func GetIsFollowById(userId int64, followId int64) (bool, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := constant.FollowRedisPrefix + strconv.FormatInt(userId, 10)
	followkey := strconv.FormatInt(followId, 10)
	result, err := redisConn.Do("zscore", key, followkey)
	if err != nil {
		return false, err
	}
	expireTime := expireTimeUtil.GetRandTime()
	if result == nil {
		_, err = redisConn.Do("expire", key, expireTime)
		if err != nil {
			return false, err
		}
		return false, nil
	}
	_, err = redisConn.Do("expire", key, expireTime)
	if err != nil {
		return true, err
	}
	return true, nil
}

func GetUserInfo(userId int64) (*model.UserRedis, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := constant.UserInfoRedisPrefix + strconv.FormatInt(userId, 10)
	result, err := redis.Bytes(redisConn.Do("get", key))
	if err != nil {
		return nil, err
	}
	userInfo := new(model.UserInfoRedis)
	err = json.Unmarshal(result, userInfo)
	if err != nil {
		return nil, err
	}

	expireTime := expireTimeUtil.GetRandTime()
	_, err = redisConn.Do("expire", key, expireTime)
	if err != nil {
		klog.Error(err)
	}

	cntKey := constant.UserInfoCntRedisPrefix + strconv.FormatInt(userId, 10)
	cnt, cntErr := redis.Int64s(redisConn.Do("hmget",
		cntKey, constant.FollowCountRedisPrefix, constant.FanCountRedisPrefix,
		constant.WorkCountRedisPrefix, constant.FavoriteCountRedisPrefix))
	if cntErr != nil {
		return nil, cntErr
	}
	if cnt == nil || len(cnt) == 0 {
		return nil, redis.ErrNil
	}
	user := &model.UserRedis{
		UserId:      userId,
		UserName:    userInfo.UserName,
		FollowCnt:   cnt[0],
		FanCnt:      cnt[1],
		WorkCnt:     cnt[2],
		FavoriteCnt: cnt[3],
	}
	redisConn.Do("expire", cntKey, expireTime)

	return user, nil
}

func IsFollowKeyExist(userId int64) (bool, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()
	key := constant.FollowRedisPrefix + strconv.FormatInt(userId, 10)
	result, err := redis.Strings(redisConn.Do("keys", key))
	if err != nil {
		return false, err
	}

	expireTime := expireTimeUtil.GetRandTime()
	if len(result) == 0 {
		_, err = redisConn.Do("expire", key, expireTime)
		if err != nil {
			return false, err
		}
		return false, nil
	}
	_, err = redisConn.Do("expire", key, expireTime)
	if err != nil {
		return true, err
	}
	return true, nil
}

func IsFanKeyExist(userId int64) (bool, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()
	key := constant.FanRedisPrefix + strconv.FormatInt(userId, 10)
	result, err := redis.Strings(redisConn.Do("keys", key))

	if err != nil {
		return false, err
	}
	expireTime := expireTimeUtil.GetRandTime()
	if len(result) == 0 {
		_, err = redisConn.Do("expire", key, expireTime)
		if err != nil {
			return false, err
		}
		return false, nil
	}

	_, err = redisConn.Do("expire", key, expireTime)
	if err != nil {
		return true, err
	}
	return true, nil
}

func IsKeyExistByBloom(prefix string, keyId int64) (bool, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	if bloomKeyOpen == false {
		return true, nil
	}

	result, err := redis.Int(redisConn.Do("bf.exists", prefix, keyId))
	if err != nil {
		return true, err
	}
	if result == 0 {
		return false, nil
	}
	return true, nil
}
