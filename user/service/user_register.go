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

type UserRegisterService struct {
	ctx context.Context
}

func NewUserRegisterService(ctx context.Context) *UserRegisterService {
	return &UserRegisterService{
		ctx: ctx,
	}
}

// CreateUser create user info.
func (s *UserRegisterService) CreateUser(req *userproto.CreateUserReq) (int64, error) {
	_, err := dal.GetUserByName(s.ctx, req.UserAccount.Username)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// name exists
		if err == nil {
			return 0, code.UserAlreadyExistErr
		}
		//other error type
		return 0, err
	}
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.UserAccount.Password), bcrypt.DefaultCost)
	return dal.CreateUser(s.ctx, req.UserAccount.Username, string(encryptedPassword))
}
