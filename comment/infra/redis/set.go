package redis

import (
	"douyin/comment/infra/dal/model"
	redisModel "douyin/comment/infra/redis/model"
	"douyin/common/constant"
	"encoding/json"
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

	commentCountKey := constant.CommentCountRedisPrefix + strconv.FormatInt(int64(commentRedis.VideoId), 10)
	count, cnterr := getCommentCountById(commentRedis.VideoId)
	if cnterr != nil {
		if errors.Is(err, redis.ErrNil) {
			return nil
		}
		return cnterr
	}
	count++
	_, err = redisConn.Do("set", commentCountKey, count, "ex", expireTime)
	if err != nil {
		return err
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

	commentCountKey := constant.CommentCountRedisPrefix + strconv.FormatInt(int64(videoId), 10)
	count, cnterr := getCommentCountById(videoId)
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
	_, err = redisConn.Do("set", commentCountKey, count, "ex", expireTime)
	if err != nil {
		return err
	}

	return nil
}

func AddCommentList(commentListp []*model.Comment) error {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	var key string
	expireTime := expireTimeUtil.GetRandTime()
	for i := range commentListp {
		commentRedis := redisModel.CommentRedis{
			CommentId:  commentListp[i].CommentUUId,
			VideoId:    commentListp[i].VideoId,
			UserId:     commentListp[i].UserId,
			Content:    commentListp[i].Contents,
			CreateTime: commentListp[i].CreateTime,
		}
		key = constant.CommentRedisPrefix + strconv.FormatInt(commentRedis.VideoId, 10)

		_, err := redisConn.Do("zadd", key, commentRedis.CreateTime, commentRedis.CommentId)
		if err != nil {
			redisConn.Do("del", key)
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
	}

	_, err := redisConn.Do("expire", key, expireTime)
	if err != nil {
		return err
	}
	return nil
}
