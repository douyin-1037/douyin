package model

type UserRedis struct {
	UserId        int64
	UserName      string
	FollowCount   int64
	FollowerCount int64
}
