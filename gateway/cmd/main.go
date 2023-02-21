package main

import (
	"douyin/pkg/cos"
	"github.com/gin-gonic/gin"
	"io"
	"os"

	"douyin/common/conf"
	"douyin/gateway/api/auth"
	"douyin/gateway/rpc"

	"github.com/cloudwego/kitex/pkg/klog"
)

func Init() {
	conf.InitConfig()
	auth.Init()
	cos.Init()
	InitInjectModule()
	rpc.Init()
}

func main() {
	Init()
	r := gin.New()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
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
