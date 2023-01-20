package cmd

import (
	app "douyin/application"
	appImpl "douyin/application/impl"
	"go.uber.org/fx"
)

var UserAppService app.UserAppService

func inject() {
	fx.New(
		fx.Provide(appImpl.NewUserAppService()),
		fx.Populate(&UserAppService))

}
