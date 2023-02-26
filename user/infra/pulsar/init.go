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
	ctx = context.Background()
	//follow_user producer
	p_follow_user pulsar.Producer

	FollowUserSchemaDef = "{\"type\":\"record\",\"name\":\"FollowUser\",\"namespace\":\"douyin_prod\"," +
		"\"fields\":[{\"name\":\"UserID\",\"type\":\"int\"},{\"name\":\"FollowID\",\"type\":\"int\"},{\"name\":\"ActionType\",\"type\":\"int\"}]}"
)

type FollowUserJSON struct {
	UserID     int64 `json:"userId"`
	FollowID   int64 `json:"followId"`
	ActionType int   `json:"actionType"`
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

	FollowUserJS := pulsar.NewJSONSchema(FollowUserSchemaDef, nil)

	p_follow_user, err = client.CreateProducer(pulsar.ProducerOptions{
		Topic:  constant.FollowUserTopic,
		Schema: FollowUserJS,
	})

	if err != nil {
		klog.Error(err)
	}
	//consume follow user message, and write to mysql
	go func() {
		err := FollowUserConsume(ctx, client)
		if err != nil {
			klog.Error(err)
			return
		}
	}()
}
