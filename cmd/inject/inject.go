package inject

import (
	"douyin/api/handlers"
	"douyin/application"
	appImpl "douyin/application/impl"
	"go.uber.org/fx"
)

var UserAppService appImpl.UserAppServiceImpl

//var VideoAppService app.VideoAppService
//var CommentAppService app.CommentAppService
//var MessageAppService app.MessageAppService

func Inject() {
	fx.New(application.Module,
		fx.Invoke(handlers.InjectAppService),
	)
}
