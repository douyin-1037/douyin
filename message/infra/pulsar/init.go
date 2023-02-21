package pulsar

import (
	"context"
	"douyin/common/conf"
	"douyin/common/constant"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)

var (
	ctx context.Context
	//video_like producer
	p_create_message pulsar.Producer

	CreateMessageSchemaDef = "{\"type\":\"record\",\"name\":\"CreateMessage\",\"namespace\":\"douyin_prod\"," +
		"\"fields\":[{\"name\":\"UserId\",\"type\":\"int\"},{\"name\":\"ToUserId\",\"type\":\"int\"},{\"name\":\"Content\",\"type\":\"string\"}, {\"name\":\"CreateTime\",\"type\":\"int\"}]}"
)

type CreateMessageJSON struct {
	UserId     int64  `json:"userId"`
	ToUserId   int64  `json:"toUserId"`
	Content    string `json:"content"`
	CreateTime int64  `json:"createTime"`
}

func Init() {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		// pulsar://localhost:6650
		URL: conf.Pulsar.URL,
		//Producer-create, subscribe and unsubscribe operations will be retried until this interval,
		//after which the operation will be marked as failed
		OperationTimeout: time.Duration(conf.Pulsar.OperationTimeout) * time.Second,
		//Timeout for the establishment of a TCP connection
		ConnectionTimeout: time.Duration(conf.Pulsar.ConnectionTimeout) * time.Second,
	})
	if err != nil {
		klog.Errorf("Could not instantiate Pulsar client: %v", err)
	}

	CreateMessageJS := pulsar.NewJSONSchema(CreateMessageSchemaDef, nil)

	p_create_message, err = client.CreateProducer(pulsar.ProducerOptions{
		Topic:  constant.CreateMessageTopic,
		Schema: CreateMessageJS,
	})

	if err != nil {
		klog.Error(err)
	}
	//consume create message message, and write to mysql
	go func() {
		err := CreateMessageConsume(ctx, client)
		if err != nil {
			klog.Error(err)
			return
		}
	}()
}
