package service

import (
	"context"
	"douyin/code_gen/kitex_gen/userproto"
	"douyin/user/infra/dal"
	"douyin/user/infra/redis"
	"github.com/cloudwego/kitex/pkg/klog"
)

type GetFollowListService struct {
	ctx context.Context
}

func NewGetFollowListService(ctx context.Context) *GetFollowListService {
	return &GetFollowListService{ctx: ctx}
}

func (s *GetFollowListService) GetFollowList(appUserId int64, userId int64) ([]*userproto.UserInfo, error) {

	followIdList, redisErr := redis.GetFollowList(userId)

	if redisErr != nil || followIdList == nil || len(followIdList) <= 0 {
		klog.Error("get follow list Redis missed ", redisErr)
	} else {
		return GetFollowListMakeList(s, appUserId, followIdList)
	}

	followIdDalList, err := dal.GetFollowList(s.ctx, userId)
	if err != nil {
		return nil, err
	}

	go func() {
		redis.AddFollowList(userId, followIdDalList)
	}()

	return GetFollowListMakeList(s, appUserId, followIdDalList)
}

func GetFollowListMakeList(s *GetFollowListService, appUserId int64, usersId []int64) ([]*userproto.UserInfo, error) {
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
