package pulsar

import (
	"context"
	"douyin/common/constant"
	"douyin/user/infra/dal"
	"douyin/user/infra/redis"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/cloudwego/kitex/pkg/klog"
)

func FollowUserConsume(ctx context.Context, client pulsar.Client) error {
	//listen the channel
	channel := make(chan pulsar.ConsumerMessage, 100)
	var followUserJS FollowUserJSON
	consumerJS := pulsar.NewJSONSchema(FollowUserSchemaDef, nil)
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            constant.FollowUserTopic,
		SubscriptionName: "sub-1",
		Schema:           consumerJS,
		Type:             pulsar.Shared,
		MessageChannel:   channel,
	})
	if err != nil {
		return err
	}
	defer consumer.Close()

	for cm := range channel {
		consumer := cm.Consumer
		msg := cm.Message
		err = msg.GetSchemaValue(&followUserJS)
		if err != nil {
			klog.Error(err)
		}
		err = consumer.Ack(msg)
		if err != nil {
			klog.Error(err)
		}
		switch followUserJS.ActionType {
		case constant.FollowUser:
			if err := dal.FollowUser(ctx, followUserJS.UserID, followUserJS.FollowID); err != nil {
				klog.Error("mysql error:", err)
				err = redis.DeleteRelationKey(followUserJS.UserID, followUserJS.FollowID)
				if err != nil {
					klog.Error("del redis key err", err)
				}
			}
			break
		case constant.UnFollowUser:
			if err := dal.UnFollowUser(ctx, followUserJS.UserID, followUserJS.FollowID); err != nil {
				klog.Error("mysql error:", err)
				err = redis.DeleteRelationKey(followUserJS.UserID, followUserJS.FollowID)
				if err != nil {
					klog.Error("del redis key err", err)
				}
			}
			break
		}

	}

	return nil
}
