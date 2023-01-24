package main

import (
	"douyin/cmd/router"
	"douyin/common/conf"
	"github.com/cloudwego/kitex/pkg/klog"
	"net/http"
)

func main() {
	conf.InitConfig()
	r := router.NewRouter()
	if err := http.ListenAndServe(conf.Server.HttpPort, r); err != nil {
		klog.Fatal(err)
	}
}
