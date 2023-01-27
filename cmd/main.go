package main

import (
	"douyin/api/auth"
	"douyin/cmd/router"
	"douyin/common/conf"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Init() {
	conf.InitConfig()
	auth.Init()
}

func main() {
	Init()
	r := router.NewRouter()
	if err := r.Run(conf.Server.HttpPort); err != nil {
		klog.Fatal(err)
	}
}
