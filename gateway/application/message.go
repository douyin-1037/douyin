package application

import (
	"context"

	"douyin/types/bizdto"
)

var MessageAppIns *MessageAppService

type MessageAppService struct {
}

func NewMessageAppService() *MessageAppService {
	return &MessageAppService{}
}

func (m MessageAppService) CreateMessage(ctx context.Context, appUserID int64, toUserID int64, content string) (err error) {
	//TODO implement me
	panic("implement me")
}

func (m MessageAppService) GetMessageList(ctx context.Context, appUserID int64, toUserID int64) (messageList []*bizdto.Message, err error) {
	//TODO implement me
	panic("implement me")
}
