package service

import (
	"context"
	"douyin/code_gen/kitex_gen/userproto"
	"douyin/user/infra/dal"
	"douyin/user/infra/redis"
	"github.com/pkg/errors"
	"gorm.io/gorm"
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

	err := redis.DeleteRelation(userId, followId)
	if err != nil {
		return err
	}

	go func() {
		dal.UnFollowUser(s.ctx, userId, followId)
	}()
	return nil
}
