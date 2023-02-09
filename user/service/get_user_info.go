package service

import (
	"context"
	"douyin/code_gen/kitex_gen/userproto"
	"douyin/pkg/code"
	"douyin/user/infra/dal"
	"douyin/user/infra/redis"
	"errors"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type GetUserService struct {
	ctx context.Context
}

func NewGetUserService(ctx context.Context) *GetUserService {
	return &GetUserService{ctx: ctx}
}

func (s *GetUserService) GetUser(req *userproto.GetUserReq) (*userproto.UserInfo, error) {
	return s.GetUserInfoByID(req.AppUserId, req.UserId)
}

// GetUserInfoByID  查询userId的信息 封装为UserInfo返回，appUserId主要用于判断当前用户是否关注了userId用户
func (s *GetUserService) GetUserInfoByID(appUserId, userId int64) (*userproto.UserInfo, error) {
	ruser, rerr := redis.GetUserInfo(userId)
	if rerr != nil {
		klog.Error("get user info by id Redis missed " + rerr.Error())
	}
	if ruser != nil {
		isfollow, ierr := dal.IsFollowByID(s.ctx, appUserId, userId)
		if ierr != nil {
			return nil, ierr
		}
		userInfo := &userproto.UserInfo{
			UserId:        int64(ruser.UserId),
			Username:      ruser.UserName,
			FollowCount:   ruser.FollowCount,
			FollowerCount: ruser.FollowerCount,
			IsFollow:      isfollow,
		}
		return userInfo, nil
	}

	user, err := dal.GetUserByID(s.ctx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // 如果没找到
			return nil, code.UserNotExistErr
		}
		return nil, err
	}
	isfollow, ierr := dal.IsFollowByID(s.ctx, appUserId, userId)
	if ierr != nil {
		return nil, ierr
	}
	userInfo := &userproto.UserInfo{
		UserId:        int64(user.ID),
		Username:      user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      isfollow,
	}
	return userInfo, nil
}
