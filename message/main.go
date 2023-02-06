package main

import (
	messageproto "douyin/code_gen/kitex_gen/messageproto/messageservice"
	config "douyin/common/conf"
	"douyin/common/constant"
	"douyin/message/infra/dal"
	"douyin/pkg/middleware"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"net"
)

func Init() {
	config.InitConfig()
	dal.Init()
}

func main() {
	Init()
	r, err := etcd.NewEtcdRegistry([]string{config.Server.EtcdAddress}) // r should not be reused.
	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", config.Server.MessageServiceAddr)
	if err != nil {
		panic(err)
	}

	svr := messageproto.NewServer(new(MessageServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constant.MessageDomainServiceName}), // server name
		server.WithMiddleware(middleware.CommonMiddleware),                                                     // middleWare
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
