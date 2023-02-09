package service

import (
	"context"
	"douyin/code_gen/kitex_gen/userproto"
	"douyin/user/infra/dal"
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
	return dal.FollowUser(s.ctx, int64(req.FanUserId), int64(req.FollowedUserId))
}
