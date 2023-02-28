package redis

import (
	"douyin/common/constant"
	"douyin/pkg/mapreduce"
	"douyin/user/infra/redis/model"
	"encoding/json"
	"errors"
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

// GetUserInfo Use mapreduce to get different fields of user information in parallel
func GetUserInfo(userId int64) (*model.UserRedis, error) {

	userInfo := new(model.UserInfoRedis)
	var followCount, fanCount, workCount, favoriteCount int64
	cntKey := constant.UserInfoCntRedisPrefix + strconv.FormatInt(userId, 10)
	var err error
	expireTime := expireTimeUtil.GetRandTime()
	err = mapreduce.Finish(func() error {
		redisConn := redisPool.Get()
		defer redisConn.Close()

		key := constant.UserInfoRedisPrefix + strconv.FormatInt(userId, 10)
		result, err := redis.Bytes(redisConn.Do("get", key))
		if err != nil {
			if errors.Is(err, redis.ErrNil) {
				return nil
			}
			klog.Error("get user info err:", err)
			return err
		}
		err = json.Unmarshal(result, userInfo)
		if err != nil {
			return err
		}
		_, err = redisConn.Do("expire", key, expireTime)
		if err != nil {
			klog.Error(err)
		}
		return nil
	}, func() error {
		redisConn := redisPool.Get()
		defer redisConn.Close()

		var cntErr error
		followCount, cntErr = redis.Int64(redisConn.Do("hget",
			cntKey, constant.FollowCountRedisPrefix))
		if cntErr != nil {
			if errors.Is(err, redis.ErrNil) {
				return nil
			}
			return cntErr
		}
		return nil
	}, func() error {
		redisConn := redisPool.Get()
		defer redisConn.Close()

		var cntErr error
		fanCount, cntErr = redis.Int64(redisConn.Do("hget",
			cntKey, constant.FanCountRedisPrefix))
		if cntErr != nil {
			if errors.Is(err, redis.ErrNil) {
				return nil
			}
			return cntErr
		}
		return nil
	}, func() error {
		redisConn := redisPool.Get()
		defer redisConn.Close()

		var cntErr error
		workCount, cntErr = redis.Int64(redisConn.Do("hget",
			cntKey, constant.WorkCountRedisPrefix))
		if cntErr != nil {
			if errors.Is(err, redis.ErrNil) {
				return nil
			}
			return cntErr
		}
		return nil
	}, func() error {
		redisConn := redisPool.Get()
		defer redisConn.Close()

		var cntErr error
		favoriteCount, cntErr = redis.Int64(redisConn.Do("hget",
			cntKey, constant.FavoriteCountRedisPrefix))
		if cntErr != nil {
			if errors.Is(err, redis.ErrNil) {
				return nil
			}
			return cntErr
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	user := &model.UserRedis{
		UserId:      userId,
		UserName:    userInfo.UserName,
		FollowCnt:   followCount,
		FanCnt:      fanCount,
		WorkCnt:     workCount,
		FavoriteCnt: favoriteCount,
	}
	redisConn := redisPool.Get()
	defer redisConn.Close()

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

	if bloomOpen == false {
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

// IsLock 判断是否被登录锁定
func IsLock(username string) (bool, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := constant.LoginFailLockRedisPrefix + username
	result, err := redis.Strings(redisConn.Do("keys", key))
	if err != nil {
		return false, err
	}
	//结果不存在，代表没有lock，则返回false
	if len(result) == 0 {
		return false, nil
	}
	return true, nil
}

// GetUnlockTime 获取解锁的时间
func GetUnlockTime(username string) (int, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()
	key := constant.LoginFailLockRedisPrefix + username
	//获取key的过期时间 单位默认是秒
	ttl, err := redis.Int(redisConn.Do("TTL", key))
	if err != nil {
		return 0, err
	}
	//转化为分钟
	return ttl / 60, nil
}
