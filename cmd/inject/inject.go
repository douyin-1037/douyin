package inject

import (
	"douyin/api/handlers"
	appImpl "douyin/application/impl"
	"go.uber.org/fx"
)

var UserAppService appImpl.UserAppServiceImpl

//var VideoAppService app.VideoAppService
//var CommentAppService app.CommentAppService
//var MessageAppService app.MessageAppService

func Inject() {
	fx.New(appImpl.Module,
		fx.Invoke(handlers.InjectAppService),
	)
}
