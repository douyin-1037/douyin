package redis

import (
	"context"
	"douyin/common/util"
	"douyin/pkg/code"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"os"
	"strconv"
	"time"
)

type DistributedLock struct {
	TTL             int
	RandomValue     uint64
	Key             string
	TryLockInterval time.Duration
	watchDog        chan bool
}

func (l *DistributedLock) TryLock(redisConn redis.Conn) error {

	_, err := redis.String(redisConn.Do("set", l.Key, l.RandomValue, "ex", l.TTL, "nx"))

	// 加锁失败
	if err == redis.ErrNil {
		return code.DistributedLockErr
	}
	if err != nil {
		return err
	}
	// 加锁成功
	go l.startWatchDog()
	return nil
}

func (l *DistributedLock) Unlock() error {
	redisConn := redisPool.Get()
	defer redisConn.Close()
	script, err := os.ReadFile("unlock.lua")
	if err != nil {
		return err
	}
	lua := redis.NewScript(1, string(script))
	_, err = lua.Do(redisConn, l.Key, l.RandomValue)
	close(l.watchDog)
	return err
}

func (l *DistributedLock) Lock(ctx context.Context) error {
	// 尝试加锁
	redisConn := redisPool.Get()
	defer redisConn.Close()
	err := l.TryLock(redisConn)
	if err == nil {
		return nil
	}
	if !errors.Is(err, code.DistributedLockErr) {
		return err
	}
	// 加锁失败，不断尝试
	ticker := time.NewTicker(l.TryLockInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			// 超时
			return code.LockTimeOutErr
		case <-ticker.C:
			// 重新尝试加锁
			err := l.TryLock(redisConn)
			if err == nil {
				return nil
			}
			if !errors.Is(err, code.DistributedLockErr) {
				return err
			}
		}
	}
}

func (l *DistributedLock) startWatchDog() {
	delteTime := time.Duration(l.TTL / 3)
	ticker := time.NewTicker(delteTime)
	defer ticker.Stop()
	redisConn := redisPool.Get()
	defer redisConn.Close()
	for {
		select {
		case <-ticker.C:
			// 延长锁的过期时间
			ok, err := redis.Int(redisConn.Do("expire", l.Key, l.TTL))
			// 异常或锁已经不存在则不再续期
			if (err != nil) || ok == 0 {
				return
			}
		case <-l.watchDog:
			// 已经解锁
			return
		}
	}
}

func NewUserKeyLock(userId int64, prefix string) DistributedLock {
	value, _ := util.GenSnowFlake(0)
	deltTime := time.Duration(100)
	watchDog := make(chan bool)
	return DistributedLock{
		Key:             prefix + strconv.FormatInt(userId, 10),
		RandomValue:     value,
		TTL:             1,
		TryLockInterval: deltTime,
		watchDog:        watchDog,
	}
}
