package service

import (
	"context"
	"douyin/code_gen/kitex_gen/userproto"
	"douyin/user/infra/dal"
)

type GetFanListService struct {
	ctx context.Context
}

func NewGetFanListService(ctx context.Context) *GetFanListService {
	return &GetFanListService{ctx: ctx}
}

func (s *GetFanListService) GetFanList(req *userproto.GetFanListReq) ([]*userproto.UserInfo, error) {
	appUserId := req.AppUserId
	userId := req.UserId

	//查看当前用户的粉丝列表uids
	uids, err := dal.GetFanList(s.ctx, userId)
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
