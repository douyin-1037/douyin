package pack

import (
	"douyin/code_gen/kitex_gen/messageproto"
	"douyin/message/infra/dal/model"
	"time"
)

func Message(message *model.Message) *messageproto.MessageInfo {
	return &messageproto.MessageInfo{
		MessageId:  int64(message.ID),
		Content:    message.Contents,
		CreateTime: time.Unix(message.CreatedAt.Unix(), 0).Format("2006-01-02 15:04:05"),
	}
}

func Messages(messages []*model.Message) []*messageproto.MessageInfo {
	messageInfos := make([]*messageproto.MessageInfo, len(messages))
	for i, message := range messages {
		messageInfos[i] = Message(message)
	}
	return messageInfos
}
