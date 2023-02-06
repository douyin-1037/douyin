package service

import (
	"context"
	"douyin/code_gen/kitex_gen/messageproto"
	"douyin/message/infra/dal"
)

type CreateMessageService struct {
	ctx context.Context
}

func NewCreateMessageService(ctx context.Context) *CreateMessageService {
	return &CreateMessageService{
		ctx: ctx,
	}
}

func (s *CreateMessageService) CreateMessage(req *messageproto.CreateMessageReq) error {
	return dal.CreateMessage(s.ctx, req.UserId, req.ToUserId, req.Content)
}
