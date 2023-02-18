package pulsar

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
)

func CreateMessageProduce(ctx context.Context, userId int64, toUserId int64, content string, createTime int64) error {
	_, err := p_create_message.Send(ctx, &pulsar.ProducerMessage{
		Value: &CreateMessageJSON{
			UserId:     userId,
			ToUserId:   toUserId,
			Content:    content,
			CreateTime: createTime,
		},
	})
	return err
}
