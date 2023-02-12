package service

import (
	"context"
	"douyin/code_gen/kitex_gen/userproto"
	"douyin/pkg/code"
	"douyin/user/infra/dal"
	"douyin/user/infra/redis"
	redisModel "douyin/user/infra/redis/model"
	"douyin/user/pack"
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
	var userInfo *userproto.UserInfo
	var isFollow bool
	userInfoRedis, redisErr := redis.GetUserInfo(userId)
	if redisErr != nil || userInfoRedis == nil {
		klog.Error("get user info by id Redis missed " + redisErr.Error())
		userInfoDal, err := dal.GetUserByID(s.ctx, userId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) { // 如果没找到
				return nil, code.UserNotExistErr
			}
			return nil, err
		}
		go func() {
			redis.AddUserInfo(redisModel.UserRedis{
				UserId:        int64(userInfoDal.ID),
				UserName:      userInfoDal.Name,
				FollowCount:   userInfoDal.FollowCount,
				FollowerCount: userInfoDal.FollowerCount,
			})
		}()
		userInfo = pack.PackUserDal(userInfoDal)
	} else {
		userInfo = pack.PackUserRedis(userInfoRedis)
	}

	if exist, _ := redis.IsFollowKeyExist(appUserId); exist == false {
		followList, err := dal.GetFollowList(s.ctx, appUserId)
		if err != nil {
			klog.Error("dal get relation err: ", err)
			return nil, err
		}
		redis.AddFollowList(appUserId, followList)
	}
	isFollow, redisErr = redis.GetIsFollowById(appUserId, userId)
	if redisErr != nil {
		klog.Error("get isFollowById Redis missed " + redisErr.Error())
		var dalErr error
		userInfo.IsFollow, dalErr = dal.IsFollowByID(s.ctx, appUserId, userId)
		if dalErr != nil {
			klog.Error("dal get isFollowByID err: ", dalErr)
			return nil, dalErr
		}
	} else {
		userInfo.IsFollow = isFollow
	}
	return userInfo, nil
}
