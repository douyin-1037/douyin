package dal

import (
	"context"
	"douyin/user/infra/dal/model"
	"github.com/cloudwego/kitex/pkg/klog"
)

// QueryUserByIds Get user information based on user id
func QueryUserByIds(ctx context.Context, userIds []int64) ([]*model.User, error) {
	var users []*model.User
	err := DB.WithContext(ctx).Where("id in (?)", userIds).Find(&users).Error
	if err != nil {
		klog.Error("query user by ids fail " + err.Error())
		return nil, err
	}
	return users, nil
}

// GetUserByName needs to query user information by name
func GetUserByName(ctx context.Context, userName string) (*model.User, error) {
	var user model.User
	if err := DB.WithContext(ctx).Where("name = ?", userName).First(&user).Error; err != nil {
		//klog.Error("get user by name fail " + err.Error())
		return nil, err
	}
	return &user, nil
}

// CreateUser Create a user based on the given user information, and return the user ID
func CreateUser(ctx context.Context, username string, password string) (int64, error) {
	user := &model.User{
		Name:          username,
		Password:      password,
		FollowCount:   0,
		FollowerCount: 0,
	}
	err := DB.WithContext(ctx).Create(&user).Error
	if err != nil {
		klog.Error("create user data fail " + err.Error())
		return 0, err
	}
	return int64(user.ID), nil
}
