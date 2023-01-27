package application

import (
	"douyin/application/impl"
	"go.uber.org/fx"
)

var Module = fx.Module("app",
	fx.Provide(impl.NewUserAppService),
)
