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
)

type CheckUserService struct {
	ctx context.Context
}

func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{ctx: ctx}
}

func (s *CheckUserService) CheckUser(req *userproto.CheckUserReq) (userId int64, err error) {
	username := req.UserAccount.Username
	//Verify that the user is locked out
	lock, _ := redis.IsLock(username)

	if lock {
		//Get Expiration Time
		min, _ := redis.GetUnlockTime(username)
		return 0, code.NewLoginFailedTooManyErr(min)
	}

	//Check if the user entered the password correctly
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
		err = redis.SetCheckFailCounter(username)
		if err != nil {
			return 0, code.ServiceErr
		}
		return 0, code.LoginErr
	}
	//Login successful Remove failure counter
	err = redis.DeleteLoginFailCounter(username)
	if err != nil {
		return 0, code.ServiceErr
	}

	err = redis.DeleteMessageLatestTime(int64(user.ID))
	if err != nil {
		return 0, code.ServiceErr
	}
	return int64(user.ID), nil
}
