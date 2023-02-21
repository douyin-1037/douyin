package application

import (
	"context"

	"github.com/pkg/errors"

	"douyin/code_gen/kitex_gen/userproto"
	"douyin/gateway/rpc"
	"douyin/types/bizdto"
)

var UserAppIns *UserAppService

type UserAppService struct {
}

func NewUserAppService() *UserAppService {
	return &UserAppService{}
}

func (i *UserAppService) GetUser(ctx context.Context, appUserID int64, userId int64) (user *bizdto.User, err error) {
	u, err := rpc.GetUser(ctx, &userproto.GetUserReq{
		AppUserId: appUserID,
		UserId:    userId,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "GetUser rpc failed, appUserID: %v, userId: %v", appUserID, userId)
	}
	return toUserDTO(u), nil
}

func (i *UserAppService) CreateUser(ctx context.Context, username string, password string) (userId int64, err error) {
	//TODO implement me done
	req := &userproto.CreateUserReq{
		UserAccount: &userproto.UserAccount{
			Username: username,
			Password: password,
		},
	}
	userId, err = rpc.CreateUser(ctx, req)
	if err != nil {
		return 0, errors.Wrapf(err, "CreateUser rpc failed, username: %v", username)
	}
	return userId, nil
}

func (i *UserAppService) CheckUser(ctx context.Context, username string, password string) (userId int64, err error) {
	//TODO implement me
	req := &userproto.CheckUserReq{
		UserAccount: &userproto.UserAccount{
			Username: username,
			Password: password,
		},
	}
	userId, err = rpc.CheckUser(ctx, req)
	if err != nil {
		return 0, errors.Wrapf(err, "CheckUser rpc failed, username: %v", username)
	}
	return userId, nil
}

func (i *UserAppService) FollowUser(ctx context.Context, fanID int64, toUserID int64) (err error) {
	err = rpc.FollowUser(ctx, &userproto.FollowUserReq{
		FanUserId:      fanID,
		FollowedUserId: toUserID,
	})
	if err != nil {
		return errors.Wrapf(err, "FollowUser rpc failed, fanID: %v, toUserID: %v", fanID, toUserID)
	}
	return nil
}

func (i *UserAppService) UnFollowUser(ctx context.Context, fanID int64, toUserID int64) (err error) {
	err = rpc.UnFollowUser(ctx, &userproto.UnFollowUserReq{
		FanUserId:      fanID,
		FollowedUserId: toUserID,
	})
	if err != nil {
		return errors.Wrapf(err, "UnFollowUser rpc failed, fanID: %v, toUserID: %v", fanID, toUserID)
	}
	return nil
}

func (i *UserAppService) GetFollowList(ctx context.Context, appUserID int64, userId int64) (userList []*bizdto.User, err error) {
	us, err := rpc.GetFollowList(ctx, &userproto.GetFollowListReq{
		AppUserId: appUserID,
		UserId:    userId,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "GetFollowList rpc failed, appUserID: %v, UserID: %v", appUserID, userId)
	}
	return toUserDTOs(us), nil
}

func (i *UserAppService) GetFanList(ctx context.Context, appUserID int64, userId int64) (userList []*bizdto.User, err error) {
	us, err := rpc.GetFanList(ctx, &userproto.GetFanListReq{
		AppUserId: appUserID,
		UserId:    userId,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "GetFanList rpc failed, appUserID: %v, UserID: %v", appUserID, userId)
	}
	return toUserDTOs(us), nil
}

func (i *UserAppService) GetFriendList(ctx context.Context, appUserID int64, userId int64) (userList []*bizdto.User, err error) {
	us, err := rpc.GetFriendList(ctx, &userproto.GetFriendListReq{
		AppUserId: appUserID,
		UserId:    userId,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "GetFriendList rpc failed, appUserID: %v, UserID: %v", appUserID, userId)
	}
	return toUserDTOs(us), nil
}

func toUserDTO(user *userproto.UserInfo) *bizdto.User {
	if user == nil {
		return nil
	}
	return &bizdto.User{
		ID:            user.UserId,
		Name:          user.Username,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      user.IsFollow,
		WorkCount:     user.WorkCount,
		FavoriteCount: user.FavoriteCount,
	}
}

func toUserDTOs(users []*userproto.UserInfo) []*bizdto.User {
	us := make([]*bizdto.User, 0, len(users))
	for _, user := range users {
		us = append(us, toUserDTO(user))
	}
	return us
}
