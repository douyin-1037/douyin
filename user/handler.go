package main

import (
	"context"
	userproto "douyin/code_gen/kitex_gen/userproto"
	"douyin/pkg/code"
	"douyin/user/pack"
	"douyin/user/service"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *userproto.CreateUserReq) (resp *userproto.CreateUserResp, err error) {
	// TODO: Your code here... done
	resp = new(userproto.CreateUserResp)

	if len(req.UserAccount.Username) == 0 || len(req.UserAccount.Password) == 0 || len(req.UserAccount.Username) > 32 || len(req.UserAccount.Password) > 32 {
		resp.BaseResp = pack.BuildBaseResp(code.ParamErr)
		return resp, nil
	}

	userID, err := service.NewUserRegisterService(ctx).CreateUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(code.Success)
	resp.UserId = userID
	return resp, nil
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *userproto.GetUserReq) (resp *userproto.GetUserResp, err error) {
	// TODO: Your code here...
	return
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *userproto.CheckUserReq) (resp *userproto.CheckUserResp, err error) {
	// TODO: Your code here...
	resp = new(userproto.CheckUserResp)

	if len(req.UserAccount.Username) == 0 || len(req.UserAccount.Password) == 0 || len(req.UserAccount.Username) > 32 || len(req.UserAccount.Password) > 32 {
		resp.BaseResp = pack.BuildBaseResp(code.ParamErr)
		return resp, nil
	}

	uid, err := service.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.UserId = uid
	resp.BaseResp = pack.BuildBaseResp(code.Success)
	return resp, nil
}

// FollowUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) FollowUser(ctx context.Context, req *userproto.FollowUserReq) (resp *userproto.FollowUserResp, err error) {
	// TODO: Your code here...
	resp = new(userproto.FollowUserResp)

	if req.FollowedUserId < 0 || req.FanUserId < 0 {
		resp.BaseResp = pack.BuildBaseResp(code.ParamErr)
		return resp, nil
	}

	err = service.NewFollowUserService(ctx).FollowUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(code.Success)
	return resp, nil
}

// UnFollowUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) UnFollowUser(ctx context.Context, req *userproto.UnFollowUserReq) (resp *userproto.UnFollowUserResp, err error) {
	// TODO: Your code here...
	return
}

// GetFollowList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetFollowList(ctx context.Context, req *userproto.GetFollowListReq) (resp *userproto.GetFollowListResp, err error) {
	// TODO: Your code here...
	return
}

// GetFanList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetFanList(ctx context.Context, req *userproto.GetFanListReq) (resp *userproto.GetFanListResp, err error) {
	// TODO: Your code here...
	return
}

// GetFriendList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetFriendList(ctx context.Context, req *userproto.GetFriendListReq) (resp *userproto.GetFriendListResp, err error) {
	// TODO: Your code here...
	return
}
