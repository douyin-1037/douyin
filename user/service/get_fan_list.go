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

func (s *GetFanListService) GetFanList(appUserId int64, userId int64) ([]*userproto.UserInfo, error) {

	fanIdList, redisErr := redis.GetFollowList(userId)
	if redisErr != nil || fanIdList == nil || len(fanIdList) <= 0 {
		klog.Error("get fan list Redis missed " + redisErr.Error())
	} else {
		return GetFanListMakeList(s, appUserId, fanIdList)
	}

	fanIdDalList, err := dal.GetFanList(s.ctx, userId)
	if err != nil {
		return nil, err
	}

	go func() {
		redis.AddFanList(userId, fanIdDalList)
	}()

	return GetFanListMakeList(s, appUserId, fanIdDalList)
}

func GetFanListMakeList(s *GetFanListService, appUserId int64, usersId []int64) ([]*userproto.UserInfo, error) {
	if len(usersId) == 0 {
		return make([]*userproto.UserInfo, 0), nil
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
