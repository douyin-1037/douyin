package service

import (
	"context"
	"douyin/code_gen/kitex_gen/messageproto"
	"douyin/message/infra/dal"
	"douyin/message/pack"
)

type GetMessageListService struct {
	ctx context.Context
}

func NewGetMessageListService(ctx context.Context) *GetMessageListService {
	return &GetMessageListService{
		ctx: ctx,
	}
}

func (s *GetMessageListService) GetMessageList(req *messageproto.GetMessageListReq) ([]*messageproto.MessageInfo, error) {
	messages, err := dal.GetMessageList(s.ctx, req.UserId, req.ToUserId)
	if err != nil {
		return nil, err
	}
	return pack.Messages(messages), nil
}
