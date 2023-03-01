package pulsar

import (
	"context"
	"douyin/comment/infra/dal"
	"douyin/comment/infra/redis"
	"douyin/common/constant"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/cloudwego/kitex/pkg/klog"
)

func CommentConsume(ctx context.Context, client pulsar.Client) error {
	//listen the channel
	channel := make(chan pulsar.ConsumerMessage, 100)
	var commentJS CommentJSON
	consumerJS := pulsar.NewJSONSchema(CommentSchemaDef, nil)
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            constant.CommentTopic,
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
		err = msg.GetSchemaValue(&commentJS)
		if err != nil {
			klog.Error(err)
		}
		err = consumer.Ack(msg)
		if err != nil {
			klog.Error(err)
		}
		fmt.Println("CommentConsume", commentJS)
		switch commentJS.ActionType {
		case constant.CreateComment:
			if _, err := dal.CreateComment(ctx, commentJS.UserId, commentJS.VideoId, commentJS.Content, commentJS.CommentUUID, commentJS.CreateTime); err != nil {
				klog.Error("mysql error:", err)
				err = redis.DeleteCommentKey(commentJS.VideoId, commentJS.CommentUUID)
				if err != nil {
					klog.Error("redis del key err", err)
				}
			}
			break
		case constant.DeleteComment:
			if err := dal.DeleteComment(ctx, commentJS.CommentUUID, commentJS.VideoId); err != nil {
				klog.Error("mysql error:", err)
				err = redis.DeleteCommentKey(commentJS.VideoId, commentJS.CommentUUID)
				if err != nil {
					klog.Error("redis del key err", err)
				}
			}
			break
		}
	}

	return nil
}
