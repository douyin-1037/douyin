package service

import (
	"context"
	"douyin/code_gen/kitex_gen/userproto"
	"douyin/user/infra/dal"
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
	return dal.UnFollowUser(s.ctx, int64(req.FanUserId), int64(req.FollowedUserId))
}
