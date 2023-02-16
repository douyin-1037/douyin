package redis

import (
	"douyin/common/constant"
	"douyin/video/infra/dal/model"
	redismodel "douyin/video/infra/redis/model"
	"encoding/json"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"time"
)

func AddPublishList(videoListp []*model.Video, userId int64) error {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	if videoListp == nil || len(videoListp) == 0 {
		return nil
	}

	var videoList []redismodel.VideoRedis

	for i := range videoListp {
		videoList = append(videoList, redismodel.VideoRedis{
			VideoId:  videoListp[i].ID,
			UserId:   videoListp[i].UserId,
			Title:    videoListp[i].Title,
			PlayUrl:  videoListp[i].PlayUrl,
			CoverUrl: videoListp[i].CoverUrl,
		})
	}
	ub, err := json.Marshal(videoList)
	if err != nil {
		return err
	}

	key := constant.PublishRedisPrefix + strconv.FormatInt(userId, 10)
	expireTime := expireTimeUtil.GetRandTime()
	_, err = redisConn.Do("set", key, ub, "ex", expireTime)
	if err != nil {
		redisConn.Do("del", key)
		return err
	}

	redisConn.Send("multi")
	for i := range videoListp {
		countKey := constant.VideoInfoCntRedisPrefix + strconv.FormatInt(int64(videoListp[i].ID), 10)
		redisConn.Send("hset", countKey,
			constant.LikeCountRedisPrefix, videoListp[i].FavoriteCount,
			constant.CommentCountRedisPrefix, videoListp[i].CommentCount)
		redisConn.Send("expire", countKey, expireTime)
	}
	_, err = redisConn.Do("exec")
	if err != nil {
		return err
	}
	return nil
}

func DelPublishList(userId int64) error {

	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := constant.PublishRedisPrefix + strconv.FormatInt(userId, 10)
	_, err := redisConn.Do("del", key)
	if err != nil {
		return err
	}

	return nil
}

func AddLikeList(userId int64, likeList []int64) error {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	if likeList == nil || len(likeList) == 0 {
		return nil
	}

	key := constant.LikeRedisPrefix + strconv.FormatInt(userId, 10)

	l := len(likeList)

	redisConn.Send("multi")
	for i, likeId := range likeList {
		redisConn.Send("zadd", key, l-i, likeId)
	}
	_, err := redisConn.Do("exec")
	if err != nil {
		redisConn.Do("del", key)
		return err
	}

	_, err = redisConn.Do("expire", key, expireTimeUtil.GetRandTime())
	if err != nil {
		klog.Error(err)
	}

	return nil
}

func AddLike(userId int64, videoId int64) error {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := constant.LikeRedisPrefix + strconv.FormatInt(userId, 10)

	now := time.Now()
	time := now.Unix()

	_, err := redisConn.Do("zadd", key, time, videoId)
	if err != nil {
		redisConn.Do("del", key)
		return err
	}

	expireTime := expireTimeUtil.GetRandTime()
	_, err = redisConn.Do("expire", key, expireTime)
	if err != nil {
		klog.Error(err)
	}

	cntKey := constant.VideoInfoCntRedisPrefix + strconv.FormatInt(videoId, 10)
	err = incrCount(redisConn, cntKey, constant.LikeCountRedisPrefix, 1, expireTime)
	if err != nil {
		klog.Error(err)
	}

	cntKey = constant.UserInfoCntRedisPrefix + strconv.FormatInt(userId, 10)
	err = incrCount(redisConn, cntKey, constant.FavoriteCountRedisPrefix, 1, expireTime)
	if err != nil {
		klog.Error(err)
	}

	return nil
}

func DeleteLike(userId int64, videoId int64) error {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := constant.LikeRedisPrefix + strconv.FormatInt(userId, 10)

	_, err := redisConn.Do("zrem", key, videoId)
	if err != nil {
		redisConn.Do("del", key)
		return err
	}

	videoInfoKey := constant.VideoInfoRedisPrefix + strconv.FormatInt(int64(videoId), 10)
	redisConn.Do("del", videoInfoKey)

	expireTime := expireTimeUtil.GetRandTime()
	_, err = redisConn.Do("expire", key, expireTime)
	if err != nil {
		klog.Error(err)
	}

	cntKey := constant.VideoInfoCntRedisPrefix + strconv.FormatInt(int64(videoId), 10)
	err = incrCount(redisConn, cntKey, constant.LikeCountRedisPrefix, -1, expireTime)
	if err != nil {
		klog.Error(err)
	}

	cntKey = constant.UserInfoCntRedisPrefix + strconv.FormatInt(userId, 10)
	err = incrCount(redisConn, cntKey, constant.FavoriteCountRedisPrefix, -1, expireTime)
	if err != nil {
		klog.Error(err)
	}

	return nil
}

func AddVideoInfo(video model.Video) error {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	videoRedisInfo := redismodel.VideoRedis{
		VideoId:  video.ID,
		UserId:   video.UserId,
		Title:    video.Title,
		PlayUrl:  video.PlayUrl,
		CoverUrl: video.CoverUrl,
	}
	ub, err := json.Marshal(videoRedisInfo)
	if err != nil {
		return err
	}

	key := constant.VideoInfoRedisPrefix + strconv.FormatInt(int64(video.ID), 10)
	expireTime := expireTimeUtil.GetRandTime()
	_, err = redisConn.Do("set", key, ub, "ex", expireTime)
	if err != nil {
		return err
	}

	countKey := constant.VideoInfoCntRedisPrefix + strconv.FormatInt(int64(video.ID), 10)
	_, err = redisConn.Do("hset", countKey,
		constant.LikeCountRedisPrefix, video.FavoriteCount,
		constant.CommentCountRedisPrefix, video.CommentCount)
	if err != nil {
		return err
	}
	_, err = redisConn.Do("expire", countKey, expireTime)
	if err != nil {
		klog.Error(err)
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
