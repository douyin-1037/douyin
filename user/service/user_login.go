package service

import (
	"context"
	"douyin/code_gen/kitex_gen/userproto"
	"douyin/pkg/code"
	"douyin/user/infra/dal"
	"douyin/user/infra/redis"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strconv"
)

type CheckUserService struct {
	ctx context.Context
}

func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{ctx: ctx}
}

func (s *CheckUserService) CheckUser(req *userproto.CheckUserReq) (int64, error) {
	username := req.UserAccount.Username
	//1，验证用户是否被登录锁定
	lock, _ := redis.IsLock(username)

	if lock {
		//获取过期时间
		t, _ := redis.GetUnlockTime(username)
		return 0, code.NewErrNo(code.LoginFailedTooManyErrCode, "Login verification failed too many times, please try again after "+strconv.Itoa(t)+" minutes!")
	}

	//检查用户输入密码是否正确
	user, err := dal.GetUserByName(s.ctx, req.UserAccount.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // 如果没找到
			return 0, code.LoginErr
		}
		return 0, err
	}
	//param1 hashedPassword(stored) param2 password req param
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.UserAccount.Password))
	if err != nil {
		redis.SetCheckFailCounter(username)
		return 0, code.LoginErr
	}
	//登录成功 移除失败计数器
	redis.DeleteLoginFailCounter(username)

	redis.DeleteMessageLatestTime(int64(user.ID))
	return int64(user.ID), nil
}
