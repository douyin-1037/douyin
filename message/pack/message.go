package pack

import (
	"douyin/code_gen/kitex_gen/messageproto"
	"douyin/message/infra/dal/model"
)

func Message(message *model.Message) *messageproto.MessageInfo {
	return &messageproto.MessageInfo{
		MessageId:  int64(message.ID),
		FromUserId: int64(message.FromUserId),
		ToUserId:   int64(message.ToUserId),
		Content:    message.Contents,
		CreateTime: message.CreateTime,
	}
}

func Messages(messages []*model.Message) []*messageproto.MessageInfo {
	messageInfos := make([]*messageproto.MessageInfo, len(messages))
	for i, message := range messages {
		messageInfos[i] = Message(message)
	}
	return messageInfos
}
