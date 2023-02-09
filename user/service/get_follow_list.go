package service

import (
	"context"
	"douyin/code_gen/kitex_gen/userproto"
	"douyin/user/infra/dal"
)

type GetFollowListService struct {
	ctx context.Context
}

func NewGetFollowListService(ctx context.Context) *GetFollowListService {
	return &GetFollowListService{ctx: ctx}
}

func (s *GetFollowListService) GetFollowList(req *userproto.GetFollowListReq) ([]*userproto.UserInfo, error) {
	appUserId := req.AppUserId
	userId := req.UserId

	//查看当前用户的关注列表
	uids, err := dal.GetFollowList(s.ctx, userId)
	if err != nil {
		return nil, err
	}

	if len(uids) == 0 {
		return nil, nil
	}
	userInfos := make([]*userproto.UserInfo, len(uids))

	for i, uid := range uids {
		userInfo, err := NewGetUserService(s.ctx).GetUserInfoByID(appUserId, uid)
		if err != nil {
			return nil, err
		}
		userInfos[i] = userInfo
	}

	return userInfos, nil
}
