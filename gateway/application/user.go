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

func (i *UserAppService) GetUser(ctx context.Context, appUserID int64, userID int64) (user *bizdto.User, err error) {
	u, err := rpc.GetUser(ctx, &userproto.GetUserReq{
		AppUserId: appUserID,
		UserId:    userID,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "GetUser rpc failed, appUserID: %v, userID: %v", appUserID, userID)
	}
	return toDTO(u), nil
}

func (i *UserAppService) CreateUser(ctx context.Context, username string, password string) (userID int64, err error) {
	//TODO implement me done
	req := &userproto.CreateUserReq{
		UserAccount: &userproto.UserAccount{
			Username: username,
			Password: password,
		},
	}
	userID, err = rpc.CreateUser(ctx, req)
	if err != nil {
		return 0, errors.Wrapf(err, "CreateUser rpc failed, username: %v", username)
	}
	return userID, nil
}

func (i *UserAppService) CheckUser(ctx context.Context, username string, password string) (userID int64, err error) {
	//TODO implement me
	req := &userproto.CheckUserReq{
		UserAccount: &userproto.UserAccount{
			Username: username,
			Password: password,
		},
	}
	userID, err = rpc.CheckUser(ctx, req)
	if err != nil {
		return 0, errors.Wrapf(err, "CheckUser rpc failed, username: %v", username)
	}
	return userID, nil
}

func (i *UserAppService) FollowUser(ctx context.Context, fanID int64, toUserID int64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (i *UserAppService) UnFollowUser(ctx context.Context, fanID int64, toUserID int64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (i *UserAppService) GetFollowList(ctx context.Context, appUserID int64, userID int64) (userList []*bizdto.User, err error) {
	//TODO implement me
	panic("implement me")
}

func (i *UserAppService) GetFanList(ctx context.Context, appUserID int64, userID int64) (userList []*bizdto.User, err error) {
	//TODO implement me
	panic("implement me")
}

func (i *UserAppService) GetFriendList(ctx context.Context, appUserID int64, userID int64) (userList []*bizdto.User, err error) {
	//TODO implement me
	panic("implement me")
}

func toDTO(user *userproto.UserInfo) *bizdto.User {
	if user == nil {
		return nil
	}
	return &bizdto.User{
		ID:            user.UserId,
		Name:          user.Username,
		FollowCount:   user.FollowerCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      user.IsFollow,
	}
}

func toDTOs(users []*userproto.UserInfo) []*bizdto.User {
	us := make([]*bizdto.User, 0, len(users))
	for _, user := range users {
		us = append(us, toDTO(user))
	}
	return us
}
