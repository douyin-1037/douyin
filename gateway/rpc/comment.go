package rpc

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"

	"douyin/code_gen/kitex_gen/commentproto"
	"douyin/code_gen/kitex_gen/commentproto/commentservice"
	config "douyin/common/conf"
	"douyin/common/constant"
	"douyin/pkg/middleware"
	errno "douyin/pkg/statuserr"
)

var commentClient commentservice.Client

func initCommentRPC() {
	r, err := etcd.NewEtcdResolver([]string{config.Server.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := commentservice.NewClient(
		constant.CommentDomainServiceName,
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
	commentClient = c
}

func CreateComment(ctx context.Context, req *commentproto.CreateCommentReq) (*commentproto.CommentInfo, error) {
	resp, err := commentClient.CreateComment(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.New(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.CommentInfo, nil
}

func DeleteComment(ctx context.Context, req *commentproto.DeleteCommentReq) error {
	resp, err := commentClient.DeleteComment(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.New(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return nil
}

func GetCommentList(ctx context.Context, req *commentproto.GetCommentListReq) ([]*commentproto.CommentInfo, error) {
	resp, err := commentClient.GetCommentList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.New(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.CommentInfos, nil
}
