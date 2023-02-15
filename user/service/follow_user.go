package service

import (
	"context"
	"douyin/code_gen/kitex_gen/userproto"
	"douyin/common/constant"
	"douyin/user/infra/dal"
	"douyin/user/infra/redis"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type FollowUserService struct {
	ctx context.Context
}

// NewFollowUserService new FollowUserService
func NewFollowUserService(ctx context.Context) *FollowUserService {
	return &FollowUserService{
		ctx: ctx,
	}
}

// FollowUser Follow user by id
func (s *FollowUserService) FollowUser(req *userproto.FollowUserReq) error {
	if req.FanUserId == req.FollowedUserId {
		return errors.New("can't follow yourself")
	}
	userId := req.FanUserId
	followId := req.FollowedUserId

	if exist, _ := redis.IsFollowKeyExist(userId); exist == false {
		followIdDalList, err := dal.GetFollowList(s.ctx, userId)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		redis.AddFollowList(userId, followIdDalList)
	}
	if exist, _ := redis.IsFanKeyExist(followId); exist == false {
		fanIdDalList, err := dal.GetFollowList(s.ctx, followId)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		redis.AddFollowList(followId, fanIdDalList)
	}
	err := redis.AddRelation(userId, followId)
	if err != nil {
		return err
	}

	go func() {
		dal.FollowUser(s.ctx, userId, followId)
	}()
	redis.AddBloomKey(constant.FollowRedisPrefix, userId)
	redis.AddBloomKey(constant.FanRedisPrefix, followId)
	return nil
}
