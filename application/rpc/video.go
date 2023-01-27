package rpc

import (
	"context"
	"douyin/code_gen/kitex_gen/videoproto"
	"douyin/code_gen/kitex_gen/videoproto/videoservice"
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

var videoClient videoservice.Client

func initVideoRPC() {
	r, err := etcd.NewEtcdResolver([]string{config.Server.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := videoservice.NewClient(
		constant.VideoDomainServiceName,
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
	videoClient = c
}

func CreateVideo(ctx context.Context, req *videoproto.CreateVideoReq) error {
	resp, err := videoClient.CreateVideo(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return nil
}

func GetVideoListByUserId(ctx context.Context, req *videoproto.GetVideoListByUserIdReq) ([]*videoproto.VideoInfo, error) {
	resp, err := videoClient.GetVideoListByUserId(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.VideoInfos, nil
}

func GetVideoListByTime(ctx context.Context, req *videoproto.GetVideoListByTimeReq) ([]*videoproto.VideoInfo, int64, error) {
	resp, err := videoClient.GetVideoListByTime(ctx, req)
	if err != nil {
		return nil, 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.VideoInfos, resp.NextTime, nil
}

func LikeVideo(ctx context.Context, req *videoproto.LikeVideoReq) error {
	resp, err := videoClient.LikeVideo(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return nil
}

func UnLikeVideo(ctx context.Context, req *videoproto.UnLikeVideoReq) error {
	resp, err := videoClient.UnLikeVideo(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return nil
}

func GetLikeVideoList(ctx context.Context, req *videoproto.GetLikeVideoListReq) ([]*videoproto.VideoInfo, error) {
	resp, err := videoClient.GetLikeVideoList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.VideoInfos, nil
}
