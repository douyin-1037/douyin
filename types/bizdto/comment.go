package bizdto

import "douyin/types"

type Comment struct {
	ID         int64  `json:"id"`          // 评论id
	User       *User  `json:"user"`        // 评论用户信息
	Content    string `json:"content"`     // 评论内容
	CreateDate string `json:"create_date"` // 评论发布日期，格式 mm-dd
}

// 评论操作
type CommentOperationReq struct {
	VideoId     int64  `form:"video_id" json:"video_id" binding:"required"`
	ActionType  int    `form:"action_type" json:"action_type" binding:"required" msg:"1-发布评论，2-删除评论"`
	CommentText string `form:"comment_text" json:"comment_text" msg:"action_type==1时使用"`
	CommentId   int64  `form:"comment_id" json:"comment_id" msg:"要删除的评论id，action_type==2时使用"`
}

type CreateCommentResp struct {
	types.BaseResp
	Comment *Comment `json:"comment,omitempty"`
}

// 评论列表
type CommentListReq struct {
	VideoId int64 `form:"video_id" json:"video_id" binding:"required"`
}

type CommentListResp struct {
	types.BaseResp
	CommentList []*Comment `json:"comment_list,omitempty"`
}
