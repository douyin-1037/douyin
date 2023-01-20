package bizdto

import "douyin/types"

type User struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
}

// 用户信息
type UserQueryReq struct {
	UserId int64 `form:"user_id" json:"user_id" binding:"required,max=32"`
}

type UserQueryResp struct {
	types.BaseResp
	User *User `json:"user,omitempty"`
}
