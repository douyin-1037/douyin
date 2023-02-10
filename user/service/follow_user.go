package service

import (
	"context"
	"douyin/code_gen/kitex_gen/userproto"
	"douyin/user/infra/dal"
	"douyin/user/infra/redis"
	"errors"
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
	err := FollowUserRedisCheck(int64(req.FanUserId), int64(req.FollowedUserId))
	if err != nil {
		return err
	}
	return dal.FollowUser(s.ctx, int64(req.FanUserId), int64(req.FollowedUserId))
}

func FollowUserRedisCheck(userId int64, toUserId int64) error {
	followed, err := redis.GetIsFollowById(int64(userId), int64(toUserId))
	if err != nil {
		return err
	}
	if followed {
		return errors.New("this User has followed toUser")
	}
	return redis.AddRelation(int64(userId), int64(toUserId))
}
