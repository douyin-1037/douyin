package redis

import (
	"douyin/common/constant"
	"douyin/video/infra/dal/model"
	redismodel "douyin/video/infra/redis/model"
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"strconv"
)

func GetPublishList(userId int64) ([]*model.Video, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := constant.PublishRedisPrefix + strconv.FormatInt(userId, 10)
	result, err := redis.Bytes(redisConn.Do("get", key))
	if err != nil {
		return nil, err
	}
	videoRedisListp := new([]redismodel.VideoRedis)
	err = json.Unmarshal(result, videoRedisListp)
	if err != nil {
		return nil, err
	}
	expireTime := expireTimeUtil.GetRandTime()
	_, err = redisConn.Do("expire", key, expireTime)
	if err != nil {
		return nil, err
	}
	videoRedisList := *videoRedisListp
	videoList := make([]*model.Video, len(videoRedisList))

	for i := range videoRedisList {
		videoId := videoRedisList[i].VideoId
		likeCntKey := constant.LikeCountRedisPrefix + strconv.FormatInt(int64(videoId), 10)
		likeCnt, likeCntErr := redis.Int64(redisConn.Do("get", likeCntKey))
		if likeCntErr != nil {
			return nil, likeCntErr
		}
		_, err = redisConn.Do("expire", likeCntKey, expireTime)
		if err != nil {
			return nil, err
		}

		commentCntKey := constant.CommentCountRedisPrefix + strconv.FormatInt(int64(videoId), 10)
		commentCnt, commentCntErr := redis.Int64(redisConn.Do("get", commentCntKey))
		if commentCntErr != nil {
			return nil, commentCntErr
		}
		_, err = redisConn.Do("expire", commentCntKey, expireTime)
		if err != nil {
			return nil, err
		}

		videoList[i] = &model.Video{
			UserId:        videoRedisList[i].UserId,
			Title:         videoRedisList[i].Title,
			PlayUrl:       videoRedisList[i].PlayUrl,
			CoverUrl:      videoRedisList[i].CoverUrl,
			FavoriteCount: likeCnt,
			CommentCount:  commentCnt,
		}
		videoList[i].ID = videoId
	}
	return videoList, nil
}

func GetVideoInfo(videoId int64) (*model.Video, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := constant.VideoInfoRedisPrefix + strconv.FormatInt(videoId, 10)
	result, err := redis.Bytes(redisConn.Do("get", key))
	if err != nil {
		return nil, err
	}
	videoInfo := new(model.Video)
	err = json.Unmarshal(result, videoInfo)
	if err != nil {
		return nil, err
	}

	expireTime := expireTimeUtil.GetRandTime()
	_, err = redisConn.Do("expire", key, expireTime)
	if err != nil {
		return nil, err
	}
	return videoInfo, nil
}

func GetLikeList(userId int64) ([]int64, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := constant.LikeRedisPrefix + strconv.FormatInt(userId, 10)
	result, err := redis.Int64s(redisConn.Do("zrevrange", key, 0, -1))
	if err != nil {
		return nil, err
	}

	expireTime := expireTimeUtil.GetRandTime()
	_, err = redisConn.Do("expire", key, expireTime)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetIsLikeById(userId int64, videoId int64) (bool, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := constant.LikeRedisPrefix + strconv.FormatInt(userId, 10)
	likeKey := strconv.FormatInt(videoId, 10)
	result, err := redisConn.Do("zscore", key, likeKey)
	if err != nil {
		return false, err
	}
	expireTime := expireTimeUtil.GetRandTime()
	_, err = redisConn.Do("expire", key, expireTime)
	if err != nil {
		return false, err
	}
	if result == nil {
		return false, nil
	}
	return true, nil
}

func IsLikeKeyExist(userId int64) (bool, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()
	key := constant.LikeRedisPrefix + strconv.FormatInt(userId, 10)
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

func GetLikeCountById(videoId int64) (int64, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := constant.LikeCountRedisPrefix + strconv.FormatInt(int64(videoId), 10)
	result, err := redis.Int64(redisConn.Do("get", key))
	if err != nil {
		if errors.Is(err, redis.ErrNil) {
			return -1, err
		}
		return 0, err
	}
	expireTime := expireTimeUtil.GetRandTime()
	_, err = redisConn.Do("expire", key, expireTime)
	if err != nil {
		return result, err
	}
	return result, err
}

func GetCommentCountById(videoId int64) (int64, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := constant.CommentCountRedisPrefix + strconv.FormatInt(int64(videoId), 10)
	result, err := redis.Int64(redisConn.Do("get", key))
	if err != nil {
		if errors.Is(err, redis.ErrNil) {
			return -1, err
		}
		return 0, err
	}
	expireTime := expireTimeUtil.GetRandTime()
	_, err = redisConn.Do("expire", key, expireTime)
	if err != nil {
		return result, err
	}
	return result, err
}
