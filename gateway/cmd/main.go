package main

import (
	"github.com/gin-gonic/gin"

	"douyin/common/conf"
	"douyin/gateway/api/auth"
	"douyin/gateway/rpc"

	"github.com/cloudwego/kitex/pkg/klog"
)

func Init() {
	conf.InitConfig()
	auth.Init()
	InitInjectModule()
	rpc.Init()
}

func main() {
	Init()
	r := gin.New()
	if conf.Server.RunMode == "debug" {
		r.Use(gin.Logger(), gin.Recovery())
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	register(r)
	if err := r.Run(conf.Server.HttpPort); err != nil {
		klog.Fatal(err)
	}
}
