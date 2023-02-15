package service

import (
	"context"
	"time"

	"douyin/code_gen/kitex_gen/messageproto"
	"douyin/message/infra/dal"
	"douyin/message/infra/dal/model"
	"douyin/message/infra/redis"
	"douyin/message/pack"

	"github.com/cloudwego/kitex/pkg/klog"
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
	exists, err := redis.IsMessageKeyExist(req.UserId, req.ToUserId)
	if err != nil {
		klog.Error(err)
	}
	userId := req.UserId
	toUserId := req.ToUserId
	latestTime, err := redis.GetMessageLatestTime(userId, toUserId)
	if err != nil {
		latestTime = 0
	}

	nowTime := time.Now().Unix()
	defer redis.AddMessageLatestTime(userId, toUserId, nowTime)

	if exists {
		messagesInRedis, err := redis.GetMessageList(req.UserId, req.ToUserId, latestTime, nowTime)
		if err == nil {
			messages := make([]*model.Message, len(messagesInRedis))
			for i, msg := range messagesInRedis {
				messages[i] = pack.MessageFromRedisModel(&msg)
			}
			return pack.Messages(messages), nil
		} else {
			klog.Warn("Redis GetMessageList error: " + err.Error())
			return nil, err
		}
	}

	messages, err := dal.GetMessageList(s.ctx, req.UserId, req.ToUserId, latestTime)
	if err != nil {
		return nil, err
	}
	// cache messagelist in redis
	err = redis.AddMessageList(req.UserId, req.ToUserId, messages)
	if err != nil {
		klog.Error(err)
	}
	return pack.Messages(messages), nil
}
