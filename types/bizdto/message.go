package bizdto

import (
	"douyin/types/coredto"
)

type Message struct {
	ID         int64  `json:"id"`           // 消息id
	UserID     int64  `json:"from_user_id"` // 消息发送者id
	ToUserId   int64  `json:"to_user_id"`   // 消息接收者id
	Content    string `json:"content"`      // 消息内容
	CreateTime int64  `json:"create_time"`  // 消息发送时间，格式为UNIX时间戳
}

// 消息操作
type MessageOperationReq struct {
	ToUserId   int64  `form:"to_user_id" json:"to_user_id" binding:"required,max=32"`
	ActionType int    `form:"action_type" json:"action_type" binding:"required" msg:"1-发送消息"`
	Content    string `form:"content" json:"content" msg:"action_type==1时使用"`
}

// 评论列表
type MessageListReq struct {
	ToUserId int64 `form:"to_user_id" json:"to_user_id" binding:"required,max=32"`
}

type MessageListResp struct {
	coredto.BaseResp
	MessageList []*Message `json:"message_list,omitempty"`
}
