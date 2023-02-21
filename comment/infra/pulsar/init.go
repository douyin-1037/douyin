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
	p_comment pulsar.Producer

	CommentSchemaDef = "{\"type\":\"record\",\"name\":\"Comment\",\"namespace\":\"douyin_prod\"," +
		"\"fields\":[{\"name\":\"UserId\",\"type\":\"int\"},{\"name\":\"VideoId\",\"type\":\"int\"},{\"name\":\"Content\",\"type\":\"string\"},{\"name\":\"CommentUUID\",\"type\":\"int\"},{\"name\":\"CreateTime\",\"type\":\"int\"},{\"name\":\"ActionType\",\"type\":\"int\"}]}"
)

type CommentJSON struct {
	UserId      int64  `json:"userId"`
	VideoId     int64  `json:"videoId"`
	Content     string `json:"content"`
	CommentUUID int64  `json:"commentUUID"`
	CreateTime  int64  `json:"createTime"`
	ActionType  int    `json:"actionType"`
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

	commentJS := pulsar.NewJSONSchema(CommentSchemaDef, nil)

	p_comment, err = client.CreateProducer(pulsar.ProducerOptions{
		Topic:  constant.CommentTopic,
		Schema: commentJS,
	})

	if err != nil {
		klog.Error(err)
	}
	//consume like video message, and write to mysql
	go func() {
		err := CommentConsume(ctx, client)
		if err != nil {
			klog.Error(err)
			return
		}
	}()
}
