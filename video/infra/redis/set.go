package redis

import (
	"douyin/common/constant"
	"douyin/video/infra/dal/model"
	redismodel "douyin/video/infra/redis/model"
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"strconv"
	"time"
)

func AddPublishList(videoListp []*model.Video, userId int64) error {
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

	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := constant.PublishRedisPrefix + strconv.FormatInt(userId, 10)
	expireTime := expireTimeUtil.GetRandTime()
	_, err = redisConn.Do("set", key, ub, "ex", expireTime)
	if err != nil {
		return err
	}
	for i := range videoListp {
		likeCountKey := constant.LikeCountRedisPrefix + strconv.FormatInt(int64(videoListp[i].ID), 10)
		_, err = redisConn.Do("set", likeCountKey, videoListp[i].FavoriteCount, "ex", expireTime)
		if err != nil {
			return err
		}
		CommentCountKey := constant.CommentCountRedisPrefix + strconv.FormatInt(int64(videoListp[i].ID), 10)
		_, err = redisConn.Do("set", CommentCountKey, videoListp[i].CommentCount, "ex", expireTime)
		if err != nil {
			return err
		}
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

	key := constant.LikeRedisPrefix + strconv.FormatInt(userId, 10)

	l := len(likeList)

	for i, likeId := range likeList {
		_, err := redisConn.Do("zadd", key, l-i, likeId)
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
		return err
	}

	LikeCountKey := constant.LikeCountRedisPrefix + strconv.FormatInt(int64(videoId), 10)
	count, cnterr := GetLikeCountById(videoId)
	if cnterr != nil {
		if errors.Is(cnterr, redis.ErrNil) {
			return nil
		}
		return cnterr
	}
	count++
	_, err = redisConn.Do("set", LikeCountKey, count, "ex", expireTime)
	if err != nil {
		return err
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
		return err
	}

	LikeCountKey := constant.LikeCountRedisPrefix + strconv.FormatInt(int64(videoId), 10)
	count, cnterr := GetLikeCountById(videoId)
	if cnterr != nil {
		if errors.Is(err, redis.ErrNil) {
			return nil
		}
		return cnterr
	}
	count--
	if count < 0 {
		return nil
	}
	_, err = redisConn.Do("set", LikeCountKey, count, "ex", expireTime)
	if err != nil {
		return err
	}
	return nil
}

func AddVideoInfo(video model.Video) error {
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

	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := constant.VideoInfoRedisPrefix + strconv.FormatInt(int64(video.ID), 10)
	expireTime := expireTimeUtil.GetRandTime()
	_, err = redisConn.Do("set", key, ub, "ex", expireTime)
	if err != nil {
		return err
	}

	likeCountKey := constant.LikeCountRedisPrefix + strconv.FormatInt(int64(video.ID), 10)
	_, err = redisConn.Do("set", likeCountKey, video.FavoriteCount, "ex", expireTime)
	if err != nil {
		return err
	}
	CommentCountKey := constant.CommentCountRedisPrefix + strconv.FormatInt(int64(video.ID), 10)
	_, err = redisConn.Do("set", CommentCountKey, video.CommentCount, "ex", expireTime)
	if err != nil {
		return err
	}

	return nil
}
