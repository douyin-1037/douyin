package main

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"

	"douyin/gateway/application"

	"go.uber.org/fx"
)

var fxApp *fx.App

var Module = fx.Options(
	fx.Provide(application.NewCommentAppService),
	fx.Provide(application.NewMessageAppService),
	fx.Provide(application.NewUserAppService),
	fx.Provide(application.NewVideoAppService),

	fx.Populate(&application.CommentAppIns),
	fx.Populate(&application.MessageAppIns),
	fx.Populate(&application.UserAppIns),
	fx.Populate(&application.VideoAppIns),
)

func InitInjectModule() {
	fxApp = fx.New(
		Module,
		fx.StartTimeout(time.Second*30),
		fx.StopTimeout(time.Second*30),
	)
	err := fxApp.Err()
	if err != nil {
		klog.Fatal(err)
	}

	startCtx, cancel := context.WithTimeout(context.Background(), fxApp.StartTimeout())
	defer cancel()
	if err = fxApp.Start(startCtx); err != nil {
		klog.Fatal(err)
	}
	return
}
