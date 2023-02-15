package service

import (
	"context"
	"douyin/code_gen/kitex_gen/userproto"
	"douyin/common/constant"
	"douyin/user/infra/dal"
	"douyin/user/infra/redis"
	"gorm.io/gorm"

	"github.com/cloudwego/kitex/pkg/klog"
)

type GetFriendListService struct {
	ctx context.Context
}

func NewGetFriendListService(ctx context.Context) *GetFriendListService {
	return &GetFriendListService{ctx: ctx}
}

func (s *GetFriendListService) GetFriendList(appUserId int64, userId int64) ([]*userproto.UserInfo, error) {

	friendIdList, redisErr := redis.GetFollowList(userId)
	if redisErr != nil || friendIdList == nil || len(friendIdList) <= 0 {
		klog.Error("get fan list Redis missed " + redisErr.Error())
	} else {
		return GetFriendListMakeList(s, appUserId, friendIdList)
	}

	isExist, _ := redis.IsKeyExistByBloom(constant.FollowRedisPrefix, userId)
	if isExist == false {
		return nil, gorm.ErrRecordNotFound
	}
	uids, err := dal.GetFriendList(s.ctx, userId)
	if err != nil {
		return nil, err
	}
	go func() {
		followIds, _ := dal.GetFollowList(s.ctx, userId)
		redis.AddFollowList(userId, followIds)
		fanIds, _ := dal.GetFanList(s.ctx, userId)
		redis.AddFanList(userId, fanIds)
	}()
	return GetFriendListMakeList(s, appUserId, uids)
}

func GetFriendListMakeList(s *GetFriendListService, appUserId int64, usersId []int64) ([]*userproto.UserInfo, error) {
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
