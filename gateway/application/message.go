package application

import (
	"context"
	"douyin/code_gen/kitex_gen/messageproto"
	"douyin/gateway/rpc"
	"github.com/pkg/errors"
	"time"

	"douyin/types/bizdto"
)

var MessageAppIns *MessageAppService

type MessageAppService struct {
}

func NewMessageAppService() *MessageAppService {
	return &MessageAppService{}
}

func (m MessageAppService) CreateMessage(ctx context.Context, appUserID int64, toUserID int64, content string) (err error) {
	req := &messageproto.CreateMessageReq{
		UserId:   appUserID,
		ToUserId: toUserID,
		Content:  content,
	}
	err = rpc.CreateMessage(ctx, req)
	if err != nil {
		return errors.Wrapf(err, "CreateMessage rpc failed, appUserID: %v, toUserID: %v, content: %v", appUserID, toUserID, content)
	}
	return nil
}

func (m MessageAppService) GetMessageList(ctx context.Context, appUserID int64, toUserID int64) (messageList []*bizdto.Message, err error) {
	us, err := rpc.GetMessageList(ctx, &messageproto.GetMessageListReq{
		UserId:   appUserID,
		ToUserId: toUserID,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "GetMessageList rpc failed, appUserID: %v, toUserID: %v", appUserID, toUserID)
	}
	return toMessageDTOs(us), nil
}

func toMessageDTO(message *messageproto.MessageInfo) *bizdto.Message {
	if message == nil {
		return nil
	}
	return &bizdto.Message{
		ID:      message.MessageId,
		Content: message.Content,
		//CreateTime: message.CreateTime,
		CreateTime: time.Unix(message.CreateTime, 0).Format("2006-01-02 15:04:05"),
	}
}

func toMessageDTOs(messages []*messageproto.MessageInfo) []*bizdto.Message {
	us := make([]*bizdto.Message, 0, len(messages))
	for _, user := range messages {
		us = append(us, toMessageDTO(user))
	}
	return us
}
