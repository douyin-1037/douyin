package service

import (
	"context"
	"douyin/code_gen/kitex_gen/userproto"
	"douyin/user/infra/dal"
	"douyin/user/infra/redis"
	"errors"
)

type UnFollowUserService struct {
	ctx context.Context
}

// NewUnFollowUserService new UnFollowUserService
func NewUnFollowUserService(ctx context.Context) *UnFollowUserService {
	return &UnFollowUserService{
		ctx: ctx,
	}
}

// UnFollowUser unFollow user by id
func (s *UnFollowUserService) UnFollowUser(req *userproto.UnFollowUserReq) error {
	if req.FanUserId == req.FollowedUserId {
		return errors.New("can't unfollow yourself")
	}
	err := UnFollowUserRedisCheck(int64(req.FanUserId), int64(req.FollowedUserId))
	if err != nil {
		return err
	}
	return dal.UnFollowUser(s.ctx, int64(req.FanUserId), int64(req.FollowedUserId))
}

func UnFollowUserRedisCheck(userId int64, toUserId int64) error {
	followed, err := redis.GetIsFollowById(int64(userId), int64(toUserId))
	if err != nil {
		return err
	}
	if !followed {
		return errors.New("this User does not follow toUser")
	}
	return redis.DeleteRelation(int64(userId), int64(toUserId))
}
