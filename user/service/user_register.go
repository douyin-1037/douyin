package service

import (
	"context"
	"douyin/code_gen/kitex_gen/userproto"
	"douyin/common/constant"
	"douyin/pkg/code"
	"douyin/user/infra/dal"
	"douyin/user/infra/redis"
	"errors"

	"github.com/dlclark/regexp2"
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
	id, err := dal.CreateUser(s.ctx, req.UserAccount.Username, string(encryptedPassword))
	redis.AddBloomKey(constant.UserInfoRedisPrefix, id)
	return id, err
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
