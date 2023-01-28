package rpc

import (
	"context"
	"douyin/code_gen/kitex_gen/messageproto"
	"douyin/code_gen/kitex_gen/messageproto/messageservice"
	errno "douyin/common/code"
	config "douyin/common/conf"
	"douyin/common/constant"
	"douyin/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"time"
)

var messageClient messageservice.Client

func initMessageRPC() {
	r, err := etcd.NewEtcdResolver([]string{config.Server.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := messageservice.NewClient(
		constant.MessageDomainServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(time.Minute),                // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	messageClient = c
}

func CreateMessage(ctx context.Context, req *messageproto.CreateMessageReq) error {
	resp, err := messageClient.CreateMessage(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return nil
}

func GetMessageList(ctx context.Context, req *messageproto.GetMessageListReq) ([]*messageproto.MessageInfo, error) {
	resp, err := messageClient.GetMessageList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.MessageInfos, nil
}
