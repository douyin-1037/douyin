package main

// @path: video/main.go
// @description: set config and Run() server of video service
// @author: Chongzhi <dczdcz2001@aliyun.com>
import (
	videoproto "douyin/code_gen/kitex_gen/videoproto/videoservice"
	"douyin/common/conf"
	"douyin/common/constant"
	"douyin/pkg/middleware"
	"douyin/pkg/tracer"
	"douyin/video/infra/dal"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"net"
)

func Init() {
	conf.InitConfig()
	dal.Init()
	tracer.InitJaeger(constant.VideoDomainServiceName)
}

func main() {
	Init()
	r, err := etcd.NewEtcdRegistry([]string{conf.Server.EtcdAddress}) // r should not be reused.
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", conf.Server.VideoServiceAddr)
	if err != nil {
		panic(err)
	}
	svr := videoproto.NewServer(new(VideoServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constant.VideoDomainServiceName}), // server name
		server.WithMiddleware(middleware.CommonMiddleware),                                                   // middleWare
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		server.WithSuite(trace.NewDefaultServerSuite()),                    // tracer
		server.WithRegistry(r),                                             // registry
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
