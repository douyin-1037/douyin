package service

import (
	"context"
	"time"

	"douyin/code_gen/kitex_gen/messageproto"
	"douyin/common/util"
	"douyin/message/infra/dal"
	"douyin/message/infra/redis"
	"douyin/message/infra/redis/model"

	"github.com/cloudwego/kitex/pkg/klog"
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
	exists, err := redis.IsMessageKeyExist(req.UserId, req.ToUserId)
	if err != nil {
		return err
	}

	if !exists {
		// fetch messagelist into cache if not exists
		messages, err := dal.GetMessageList(s.ctx, req.UserId, req.ToUserId)
		if err != nil {
			return err
		}
		err = redis.AddMessageList(req.UserId, req.ToUserId, messages)
		if err != nil {
			return err
		}
	}

	uuid, err := util.GenSnowFlake(0)
	if err != nil {
		klog.Error("Failed to generate UUID" + err.Error())
		return err
	}

	message := model.MessageRedis{
		FromUserId: req.UserId,
		ToUserId:   req.ToUserId,
		Content:    req.Content,
		MessageId:  int64(uuid),
		CreateTime: time.Now().Unix(),
	}
	err = redis.AddMessage(req.UserId, req.ToUserId, message)
	go dal.CreateMessage(s.ctx, req.UserId, req.ToUserId, req.Content)
	return err
}
