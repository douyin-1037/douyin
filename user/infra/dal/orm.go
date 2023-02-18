package dal

import (
	"context"
	"douyin/user/infra/dal/model"
	"errors"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

// GetUserByName needs to query user information by name
func GetUserByName(ctx context.Context, userName string) (*model.User, error) {
	var user model.User
	err := DB.WithContext(ctx).Where("name = ?", userName).First(&user).Error
	if err != nil {
		klog.Error("get user by name fail " + err.Error())
		return nil, err
	}
	return &user, nil
}

// GetUserByID needs to query user information by name
func GetUserByID(ctx context.Context, userId int64) (*model.User, error) {
	var user model.User
	err := DB.WithContext(ctx).Where("id = ?", userId).First(&user).Error
	if err != nil {
		klog.Error("get user by id fail " + err.Error())
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

// IsFollowByID get appuser whether follow user or not
func IsFollowByID(ctx context.Context, appUserID, userId int64) (bool, error) {
	//var user model.User
	//err := DB.WithContext(ctx).Table("relation").Where("to_user_id = ?", userId).Find(&followers).Error
	var rel model.Relation
	err := DB.WithContext(ctx).Table("relation").Where("user_id = ? AND to_user_id = ?", appUserID, userId).First(&rel).Error
	/* if err != nil {
		klog.Error("get user by id fail " + err.Error())
		return false, err
	}
	if rel.UserId == appUserID && rel.ToUserId == userId {
		return true, nil
	}
	return false, nil */
	if err == nil {
		return true, nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return false, err
}

// FollowUser perform <A Follow B> operation, based on the given user id
func FollowUser(ctx context.Context, fanID, userId int64) error {
	if fanID == userId {
		return errors.New("you can't follow yourself")
	}
	follow := model.Relation{
		UserId:   fanID,
		ToUserId: userId,
	}
	var temp model.Relation
	terr := DB.Table("relation").Where("user_id = ? AND to_user_id = ?", fanID, userId).First(&temp).Error
	if temp.UserId == fanID && temp.ToUserId == userId {
		return errors.New("you have followed the user")
	}
	if !errors.Is(terr, gorm.ErrRecordNotFound) {
		return terr
	}
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var err error
		err = tx.Table("relation").Create(&follow).Error
		if err != nil {
			klog.Error("create relation record fail " + err.Error())
			return err
		}
		err = tx.Table("user").Where("id = ?", fanID).Update("follow_count", gorm.Expr("follow_count + ?", 1)).Error
		if err != nil {
			klog.Error("update user record follow_count fail " + err.Error())
			return err
		}
		err = tx.Table("user").Where("id = ?", userId).Update("follower_count", gorm.Expr("follower_count + ?", 1)).Error
		if err != nil {
			klog.Error("update user record follower_count fail " + err.Error())
			return err
		}
		return nil
	})
	if err != nil {
		klog.Error("follow user fail " + err.Error())
		return err
	}
	return nil
}

// UnFollowUser perform <A UnFollow B> operation, based on the given user id
func UnFollowUser(ctx context.Context, fanID, userId int64) error {
	if fanID == userId {
		return errors.New("you can't unfollow yourself")
	}
	follow := model.Relation{
		UserId:   fanID,
		ToUserId: userId,
	}
	var temp model.Relation
	terr := DB.Table("relation").Where("user_id = ? AND to_user_id = ?", fanID, userId).First(&temp).Error

	if errors.Is(terr, gorm.ErrRecordNotFound) {
		return terr
	}
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var err error

		err = tx.Table("user").Where("id = ?", fanID).Update("follow_count", gorm.Expr("follow_count - ?", 1)).Error
		if err != nil {
			klog.Error("sub user record follow_count fail " + err.Error())
			return err
		}
		err = tx.Table("user").Where("id = ?", userId).Update("follower_count", gorm.Expr("follower_count - ?", 1)).Error
		if err != nil {
			klog.Error("sub user record follower_count fail " + err.Error())
			return err
		}
		err = tx.Table("relation").Where("user_id = ? AND to_user_id = ?", fanID, userId).Delete(&follow).Error
		if err != nil {
			klog.Error("delete relation record fail " + err.Error())
			return err
		}
		return nil
	})
	if err != nil {
		klog.Error("unfollow user fail " + err.Error())
		return err
	}
	return nil
}

// GetFanList get the follower id list of the user based on user id
func GetFanList(ctx context.Context, userId int64) ([]int64, error) {
	var followers []*model.Relation
	err := DB.WithContext(ctx).Table("relation").Where("to_user_id = ?", userId).Find(&followers).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		klog.Error(err)
		return make([]int64, 0), err
	}
	if err != nil {
		klog.Error("find fan list fail " + err.Error())
		return nil, err
	}
	userIds := make([]int64, len(followers))
	for i, fan := range followers {
		userIds[i] = int64(fan.UserId)
	}
	return userIds, nil
}

// GetFollowList get the follow id list of the user based on user id
func GetFollowList(ctx context.Context, userId int64) ([]int64, error) {
	var follows []*model.Relation
	err := DB.WithContext(ctx).Table("relation").Where("user_id = ?", userId).Find(&follows).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		klog.Error(err)
		return make([]int64, 0), err
	}
	if err != nil {
		klog.Error("find follow list fail " + err.Error())
		return nil, err
	}
	userIds := make([]int64, len(follows))
	for i, following := range follows {
		userIds[i] = int64(following.ToUserId)
	}
	return userIds, nil
}

// GetFriendList get the friend id list of the user based on user id
func GetFriendList(ctx context.Context, userId int64) ([]int64, error) {
	var follows []*model.Relation
	var friend []*model.Relation
	var friends []*model.Relation
	//err := DB.WithContext(ctx).Table("relation").Where("user_id = ?", userId).Find(&friends).Error
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var err error
		err = tx.Table("relation").Where("user_id = ?", userId).Find(&follows).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				klog.Error(err)
				return err
			}
			klog.Error("find follow list in GetFriendList() fail " + err.Error())
			return err
		}
		for _, follow := range follows {
			//err = tx.Table("relation").Where(&model.Relation{UserId: follow.ToUserId, ToUserId: userId}).First(&friend).Error
			err = tx.Table("relation").Where("user_id = ? AND to_user_id = ?", follow.ToUserId, follow.UserId).First(&friend).Error
			if err != nil {
				klog.Error("find friend list in GetFriendList() fail " + err.Error())
				break
			}
			friends = append(friends, friend[0])
		}
		return nil
	})
	if errors.Is(err, gorm.ErrRecordNotFound) {
		klog.Error(err)
		return make([]int64, 0), err
	}
	if err != nil {
		klog.Error("find friend list fail " + err.Error())
		return nil, err
	}
	userIds := make([]int64, len(friends))
	for i, friend := range friends {
		userIds[i] = int64(friend.UserId)
	}
	return userIds, nil
}
