package main

import (
	"context"
	userproto "douyin/code_gen/kitex_gen/userproto"
	"douyin/pkg/code"
	"douyin/user/pack"
	"douyin/user/service"

	"github.com/dlclark/regexp2"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *userproto.CreateUserReq) (resp *userproto.CreateUserResp, err error) {
	resp = new(userproto.CreateUserResp)

	if len(req.UserAccount.Username) == 0 || len(req.UserAccount.Password) == 0 || len(req.UserAccount.Username) > 32 || len(req.UserAccount.Password) > 32 {
		resp.BaseResp = pack.BuildBaseResp(code.ParamErr)
		return resp, nil
	}

	userId, err := service.NewUserRegisterService(ctx).CreateUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(code.Success)
	resp.UserId = userId
	return resp, nil
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *userproto.GetUserReq) (resp *userproto.GetUserResp, err error) {
	resp = new(userproto.GetUserResp)

	if req.UserId < 0 {
		resp.BaseResp = pack.BuildBaseResp(code.ParamErr)
		return resp, nil
	}

	userInfo, err := service.NewGetUserService(ctx).GetUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(code.Success)
	resp.UserInfo = userInfo
	return resp, nil
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *userproto.CheckUserReq) (resp *userproto.CheckUserResp, err error) {
	resp = new(userproto.CheckUserResp)
	/*
		if len(req.UserAccount.Username) == 0 || len(req.UserAccount.Password) == 0 || len(req.UserAccount.Username) > 32 || len(req.UserAccount.Password) > 32 {
			resp.BaseResp = pack.BuildBaseResp(code.ParamErr)
			return resp, nil
		}
	*/
	name_check := CheckUsername(req.UserAccount.Username)
	password_check := CheckPassword(req.UserAccount.Password)
	// "用户名应当使用字母,数字,下划线,减号,且长度4-32位"
	// "密码应当包含大写字母,小写字母,数字,且长度5-32位"
	if name_check == "" {
		resp.BaseResp = pack.BuildBaseResp(code.UsernameCheckErr)
		return resp, nil
	}
	if password_check == "" {
		resp.BaseResp = pack.BuildBaseResp(code.PasswordCheckErr)
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
	resp = new(userproto.UnFollowUserResp)

	if req.FollowedUserId < 0 || req.FanUserId < 0 {
		resp.BaseResp = pack.BuildBaseResp(code.ParamErr)
		return resp, nil
	}

	err = service.NewUnFollowUserService(ctx).UnFollowUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(code.Success)
	return resp, nil
}

// GetFollowList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetFollowList(ctx context.Context, req *userproto.GetFollowListReq) (resp *userproto.GetFollowListResp, err error) {
	resp = new(userproto.GetFollowListResp)

	if req.UserId < 0 || req.AppUserId < 0 {
		resp.BaseResp = pack.BuildBaseResp(code.ParamErr)
		return resp, nil
	}

	follows, err := service.NewGetFollowListService(ctx).GetFollowList(req.AppUserId, req.UserId)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.UserInfos = follows
	resp.BaseResp = pack.BuildBaseResp(code.Success)
	return resp, nil
}

// GetFanList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetFanList(ctx context.Context, req *userproto.GetFanListReq) (resp *userproto.GetFanListResp, err error) {
	resp = new(userproto.GetFanListResp)

	if req.UserId < 0 || req.AppUserId < 0 {
		resp.BaseResp = pack.BuildBaseResp(code.ParamErr)
		return resp, nil
	}

	fans, err := service.NewGetFanListService(ctx).GetFanList(req.AppUserId, req.UserId)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.UserInfos = fans
	resp.BaseResp = pack.BuildBaseResp(code.Success)
	return resp, nil
}

// GetFriendList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetFriendList(ctx context.Context, req *userproto.GetFriendListReq) (resp *userproto.GetFriendListResp, err error) {
	resp = new(userproto.GetFriendListResp)

	if req.UserId < 0 || req.AppUserId < 0 {
		resp.BaseResp = pack.BuildBaseResp(code.ParamErr)
		return resp, nil
	}

	fans, err := service.NewGetFriendListService(ctx).GetFriendList(req.AppUserId, req.UserId)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.UserInfos = fans
	resp.BaseResp = pack.BuildBaseResp(code.Success)
	return resp, nil
}

func CheckUsername(str string) string {
	//expr := `^(?![0-9a-zA-Z]+$)(?![a-zA-Z!@#$%^&*]+$)(?![0-9!@#$%^&*]+$)[0-9A-Za-z!@#$%^&*]{8,16}$`
	expr := `^[a-zA-Z0-9_-]{4,32}$`
	reg, _ := regexp2.Compile(expr, 0)
	m, _ := reg.FindStringMatch(str)
	if m != nil {
		res := m.String()
		return res
	}
	return ""
}

func CheckPassword(str string) string {
	//expr := `^(?![0-9a-zA-Z]+$)(?![a-zA-Z!@#$%^&*]+$)(?![0-9!@#$%^&*]+$)[0-9A-Za-z!@#$%^&*]{8,16}$`
	expr := `^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])[a-zA-Z0-9]{5,32}$`
	reg, _ := regexp2.Compile(expr, 0)
	m, _ := reg.FindStringMatch(str)
	if m != nil {
		res := m.String()
		return res
	}
	return ""
}
