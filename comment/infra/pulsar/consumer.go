package pulsar

import (
	"context"
	"douyin/comment/infra/dal"
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
		Type:             pulsar.Exclusive,
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
			return err
		}
		err = consumer.Ack(msg)
		if err != nil {
			return err
		}
		fmt.Println(commentJS)
		switch commentJS.ActionType {
		case constant.CreateComment:
			if _, err := dal.CreateComment(ctx, commentJS.UserId, commentJS.VideoId, commentJS.Content, commentJS.CommentUUID, commentJS.CreateTime); err != nil {
				klog.Error("mysql error:", err)
				return err
			}
			break
		case constant.DeleteComment:
			if err := dal.DeleteComment(ctx, commentJS.CommentUUID, commentJS.VideoId); err != nil {
				klog.Error("mysql error:", err)
				return err
			}
			break
		}
	}

	return nil
}
