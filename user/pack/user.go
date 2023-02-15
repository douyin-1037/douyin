package pack

import (
	"douyin/code_gen/kitex_gen/userproto"
	"douyin/user/infra/dal/model"
	redisModel "douyin/user/infra/redis/model"
)

func PackUserRedis(userRedis *redisModel.UserRedis) *userproto.UserInfo {
	if userRedis == nil {
		return nil
	}
	return &userproto.UserInfo{
		UserId:        userRedis.UserId,
		Username:      userRedis.UserName,
		FollowCount:   userRedis.FollowCnt,
		FollowerCount: userRedis.FanCnt,
	}
}

func PackUserDal(user *model.User) *userproto.UserInfo {
	if user == nil {
		return nil
	}
	return &userproto.UserInfo{
		UserId:        int64(user.ID),
		Username:      user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
	}
}
