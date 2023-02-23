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
	//follow:userId
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
		constant.FanCountRedisPrefix, userCntInfo.FanCnt,
		constant.WorkCountRedisPrefix, userCntInfo.WorkCnt,
		constant.FavoriteCountRedisPrefix, userCntInfo.FavoriteCnt)
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

func DeleteMessageLatestTime(userId int64) error {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := constant.MessageLatestTimeRedisPrefix + strconv.FormatInt(userId, 10)
	_, err := redisConn.Do("del", key)
	return err
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

// SetCheckFailCounter 登录失败计数器，缓存key user_login_fail_counter:username:yyyyMMddHHmm
func SetCheckFailCounter(username string) error {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	now := time.Now()
	key := constant.LoginFailCounterRedisPrefix + username + ":" + now.Format("200601021504")
	//设置过期时间为10分钟
	expireTime := 600
	_, getErr := redis.Int64(redisConn.Do("get", key))

	//incr 命令将 key 中储存的数字值增一。如果 key 不存在，那么 key 的值会先被初始化为 0 ，然后再执行 INCR 操作。
	_, err := redis.Int64(redisConn.Do("incr", key))

	if err != nil {
		return err
	}

	if getErr != nil {
		//getErr为redis.ErrNil的情况，即key之前不存在，然后进行了incr操作，然后进行设置过期时间
		if errors.Is(getErr, redis.ErrNil) {
			_, err := redisConn.Do("expire", key, expireTime)
			if err != nil {
				return err
			}
		} else {
			return getErr
		}
	}

	//往前检查10分钟，统计失败次数
	var tenMinuteKeys [10]string
	oneMin, _ := time.ParseDuration("-1m")
	for i := 0; i < 10; i++ {
		tenMinuteKeys[i] = constant.LoginFailCounterRedisPrefix + username + ":" + now.Format("200601021504")
		now = now.Add(oneMin)
	}

	var counts [10]int
	for i := 0; i < 10; i++ {
		counts[i], getErr = redis.Int(redisConn.Do("get", tenMinuteKeys[i]))
	}

	total := 0
	for _, value := range counts {
		total += value
	}
	if total >= 5 {
		Lock(username)
		//因为锁定半小时，所以十分钟内的计数器都可以主动删除了
		redisConn.Send("MULTI")
		for i, _ := range tenMinuteKeys {
			redisConn.Send("DEL", tenMinuteKeys[i])
		}
		redisConn.Do("EXEC")
	}
	return nil
}

// DeleteLoginFailCounter 移除最近十分钟计数器
func DeleteLoginFailCounter(username string) error {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	now := time.Now()
	//往前检查10分钟
	var tenMinuteKeys [10]string
	oneMin, _ := time.ParseDuration("-1m")
	for i := 0; i < 10; i++ {
		tenMinuteKeys[i] = constant.LoginFailCounterRedisPrefix + username + ":" + now.Format("200601021504")
		now = now.Add(oneMin)
	}
	redisConn.Send("MULTI")
	for i, _ := range tenMinuteKeys {
		redisConn.Send("DEL", tenMinuteKeys[i])
	}
	_, err := redisConn.Do("EXEC")
	return err
}

// Lock 失败达到一定一定次数 锁定30分钟
func Lock(username string) error {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := constant.LoginFailLockRedisPrefix + username
	//设置过期时间为30分钟
	expirTime := 1800
	_, err := redisConn.Do("set", key, 1, "ex", expirTime)
	if err != nil {
		return err
	}
	return nil
}
