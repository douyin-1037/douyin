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
	p_like_video pulsar.Producer
	//consumer
	//c pulsar.Consumer

	client pulsar.Client

	LikeVideoSchemaDef = "{\"type\":\"record\",\"name\":\"LikeVideo\",\"namespace\":\"douyin_test\"," +
		"\"fields\":[{\"name\":\"UserID\",\"type\":\"int\"},{\"name\":\"VideoID\",\"type\":\"int\"},{\"name\":\"ActionType\",\"type\":\"int\"}]}"
)

type LikeVideoJSON struct {
	UserID     int64 `json:"userId"`
	VideoID    int64 `json:"videoID"`
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

	likeVideoJS := pulsar.NewJSONSchema(LikeVideoSchemaDef, nil)

	p_like_video, err = client.CreateProducer(pulsar.ProducerOptions{
		Topic:  constant.LikeVideoTopic,
		Schema: likeVideoJS,
	})

	if err != nil {
		klog.Error(err)
	}
	//consume like video message, and write to mysql
	go func() {
		err := LikeVideoConsume(ctx, client)
		if err != nil {
			klog.Error(err)
			return
		}
	}()

}
