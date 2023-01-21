package bizdto

import "douyin/types"

type Video struct {
	ID           int64  `json:"id"`
	Author       *User  `json:"author"`
	PlayAddr     string `json:"play_url"`
	CoverAddr    string `json:"cover_url"`
	LikeCount    int64  `json:"favorite_count"`
	CommentCount int64  `json:"comment_count"`
	IsFavorite   bool   `json:"is_favorite"`
	Title        string `json:"title"`
}

// 视频流
type VideoFeedReq struct {
	LatestTime int64 `form:"latest_time" json:"latest_time"`
}

type VideoFeedResp struct {
	types.BaseResp
	NextTime  int64    `json:"next_time"`
	VideoList []*Video `json:"video_list,omitempty"`
}

// 点赞操作
type LikeOperationReq struct {
	VideoId    int64 `form:"video_id" json:"video_id" binding:"required"`
	ActionType int   `form:"action_type" json:"action_type" binding:"required" msg:"1-点赞，2-取消点赞"`
}

// 点赞列表
type LikeListReq struct {
	UserId int64 `form:"user_id" json:"user_id" binding:"required,max=32"`
}

type LikeListResp struct {
	types.BaseResp
	VideoList []*Video `json:"video_list,omitempty"`
}

// 投稿接口
type VideoUploadReq struct {
	Title string `form:"title" json:"title" binding:"required"`
}

// 发布列表
type VideoListReq struct {
	UserId int64 `form:"user_id" json:"user_id" binding:"required,max=32"`
}

type VideoListResp struct {
	types.BaseResp
	VideoList []*Video `json:"video_list,omitempty"`
}
