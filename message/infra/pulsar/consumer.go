package pulsar

import (
	"context"
	"douyin/common/constant"
	"douyin/message/infra/dal"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/cloudwego/kitex/pkg/klog"
)

func CreateMessageConsume(ctx context.Context, client pulsar.Client) error {
	//listen the channel
	channel := make(chan pulsar.ConsumerMessage, 100)
	var createMessageJS CreateMessageJSON
	consumerJS := pulsar.NewJSONSchema(CreateMessageSchemaDef, nil)
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            constant.CreateMessageTopic,
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
		err = msg.GetSchemaValue(&createMessageJS)
		if err != nil {
			klog.Error(err)
			return err
		}
		err = consumer.Ack(msg)
		if err != nil {
			klog.Error(err)
			return err
		}

		if err := dal.CreateMessage(ctx, createMessageJS.UserId, createMessageJS.ToUserId, createMessageJS.Content, createMessageJS.CreateTime); err != nil {
			klog.Error("mysql error:", err)
			return err
		}
	}

	return nil
}
