package service

import (
	"context"
	"douyin/code_gen/kitex_gen/userproto"
	"douyin/pkg/code"
	"douyin/user/infra/dal"
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

func (s *CheckUserService) CheckUser(req *userproto.CheckUserReq) (int64, error) {
	user, err := dal.GetUserByName(s.ctx, req.UserAccount.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // 如果没找到
			return 0, code.LoginErr
		}
		return 0, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.UserAccount.Password))
	if err != nil {
		return 0, code.LoginErr
	}
	return int64(user.ID), nil
}
