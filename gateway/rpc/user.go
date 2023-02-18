package rpc

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"

	"douyin/code_gen/kitex_gen/userproto"
	"douyin/code_gen/kitex_gen/userproto/userservice"
	config "douyin/common/conf"
	"douyin/common/constant"
	"douyin/pkg/middleware"
	errno "douyin/pkg/statuserr"
)

var userClient userservice.Client

func initUserRPC() {
	r, err := etcd.NewEtcdResolver([]string{config.Server.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		constant.UserDomainServiceName,
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
	userClient = c
}

func CreateUser(ctx context.Context, req *userproto.CreateUserReq) (int64, error) {
	resp, err := userClient.CreateUser(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.New(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.UserId, nil
}

func CheckUser(ctx context.Context, req *userproto.CheckUserReq) (int64, error) {
	resp, err := userClient.CheckUser(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.New(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.UserId, nil
}

func GetUser(ctx context.Context, req *userproto.GetUserReq) (*userproto.UserInfo, error) {
	resp, err := userClient.GetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.New(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.UserInfo, nil
}

func FollowUser(ctx context.Context, req *userproto.FollowUserReq) error {
	resp, err := userClient.FollowUser(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.New(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return nil
}

func UnFollowUser(ctx context.Context, req *userproto.UnFollowUserReq) error {
	resp, err := userClient.UnFollowUser(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.New(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return nil
}

func GetFollowList(ctx context.Context, req *userproto.GetFollowListReq) ([]*userproto.UserInfo, error) {
	resp, err := userClient.GetFollowList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.New(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.UserInfos, nil
}

func GetFanList(ctx context.Context, req *userproto.GetFanListReq) ([]*userproto.UserInfo, error) {
	resp, err := userClient.GetFanList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.New(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.UserInfos, nil
}

func GetFriendList(ctx context.Context, req *userproto.GetFriendListReq) ([]*userproto.UserInfo, error) {
	resp, err := userClient.GetFriendList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.New(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.UserInfos, nil
}
