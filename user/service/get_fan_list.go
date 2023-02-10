package service

import (
	"context"
	"douyin/code_gen/kitex_gen/userproto"
	"douyin/user/infra/dal"
	"douyin/user/infra/redis"

	"github.com/cloudwego/kitex/pkg/klog"
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

	users, rerr := redis.GetFollowList(userId)
	if rerr != nil || users == nil {
		klog.Error("get follow list Redis missed " + rerr.Error())
	}
	if users != nil {
		return GetFanListMakeList(s, appUserId, users)
	}

	//查看当前用户的粉丝列表uids
	uids, err := dal.GetFanList(s.ctx, userId)
	if err != nil {
		return nil, err
	}
	return GetFanListMakeList(s, appUserId, uids)
	/*
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
	*/
}

func GetFanListMakeList(s *GetFanListService, appUserId int64, usersId []int64) ([]*userproto.UserInfo, error) {
	if len(usersId) == 0 {
		return nil, nil
	}
	userInfos := make([]*userproto.UserInfo, len(usersId))

	for i, uid := range usersId {
		userInfo, err := NewGetUserService(s.ctx).GetUserInfoByID(appUserId, uid)
		if err != nil {
			return nil, err
		}
		userInfos[i] = userInfo
	}

	return userInfos, nil
}
