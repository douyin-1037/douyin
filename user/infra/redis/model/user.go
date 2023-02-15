package model

type UserInfoRedis struct {
	UserId   int64
	UserName string
}
type UserCntRedis struct {
	FollowCnt int64
	FanCnt    int64
}

type UserRedis struct {
	UserId    int64
	UserName  string
	FollowCnt int64
	FanCnt    int64
}
