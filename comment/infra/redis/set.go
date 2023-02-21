package redis

import (
	"douyin/comment/infra/dal/model"
	redisModel "douyin/comment/infra/redis/model"
	"douyin/common/constant"
	"encoding/json"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"strconv"
)

func AddComment(commentRedis redisModel.CommentRedis) error {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := constant.CommentRedisPrefix + strconv.FormatInt(commentRedis.VideoId, 10)

	_, err := redisConn.Do("zadd", key, commentRedis.CreateTime, commentRedis.CommentId)
	if err != nil {
		redisConn.Do("del", key)
		return err
	}

	expireTime := expireTimeUtil.GetRandTime()
	_, err = redisConn.Do("expire", key, expireTime)
	if err != nil {
		return err
	}

	commentInfoKey := constant.CommentInfoRedisPrefix + strconv.FormatInt(int64(commentRedis.CommentId), 10)
	ub, err := json.Marshal(commentRedis)
	if err != nil {
		return err
	}
	_, err = redisConn.Do("set", commentInfoKey, ub, "ex", expireTime)
	if err != nil {
		return err
	}

	cntKey := constant.VideoInfoCntRedisPrefix + strconv.FormatInt(commentRedis.VideoId, 10)
	err = incrCount(redisConn, cntKey, constant.CommentCountRedisPrefix, 1, expireTime)
	if err != nil {
		klog.Error(err)
	}

	return nil
}

func DeleteComment(commentId int64, videoId int64) error {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := constant.CommentRedisPrefix + strconv.FormatInt(videoId, 10)

	_, err := redisConn.Do("zrem", key, commentId)
	if err != nil {
		redisConn.Do("del", key)
		return err
	}

	commentInfoKey := constant.CommentInfoRedisPrefix + strconv.FormatInt(int64(commentId), 10)
	redisConn.Do("del", commentInfoKey)

	expireTime := expireTimeUtil.GetRandTime()
	_, err = redisConn.Do("expire", key, expireTime)
	if err != nil {
		return err
	}

	cntKey := constant.VideoInfoCntRedisPrefix + strconv.FormatInt(videoId, 10)
	err = incrCount(redisConn, cntKey, constant.CommentCountRedisPrefix, 1, expireTime)
	if err != nil {
		klog.Error(err)
	}
	return nil
}

func AddCommentList(commentListp []*model.Comment) error {
	redisConn := redisPool.Get()
	defer redisConn.Close()
	if commentListp == nil || len(commentListp) == 0 {
		return nil
	}
	var key string
	expireTime := expireTimeUtil.GetRandTime()

	redisConn.Send("multi")
	for i := range commentListp {
		commentRedis := redisModel.CommentRedis{
			CommentId:  commentListp[i].CommentUUID,
			VideoId:    commentListp[i].VideoId,
			UserId:     commentListp[i].UserId,
			Content:    commentListp[i].Contents,
			CreateTime: commentListp[i].CreateTime,
		}
		key = constant.CommentRedisPrefix + strconv.FormatInt(commentRedis.VideoId, 10)
		redisConn.Send("zadd", key, commentRedis.CreateTime, commentRedis.CommentId)

		commentInfoKey := constant.CommentInfoRedisPrefix + strconv.FormatInt(int64(commentRedis.CommentId), 10)
		ub, _ := json.Marshal(commentRedis)
		redisConn.Send("set", commentInfoKey, ub, "ex", expireTime)
	}
	_, err := redisConn.Do("exec")
	if err != nil {
		redisConn.Do("del", key)
		return err
	}
	_, err = redisConn.Do("expire", key, expireTime)
	if err != nil {
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
