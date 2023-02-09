package service

import (
	"context"
	"douyin/code_gen/kitex_gen/userproto"
	"douyin/user/infra/dal"
)

type GetFriendListService struct {
	ctx context.Context
}

func NewGetFriendListService(ctx context.Context) *GetFriendListService {
	return &GetFriendListService{ctx: ctx}
}

func (s *GetFriendListService) GetFriendList(req *userproto.GetFriendListReq) ([]*userproto.UserInfo, error) {
	appUserId := req.AppUserId
	userId := req.UserId

	//查看当前用户的好友列表
	uids, err := dal.GetFriendList(s.ctx, userId)
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
