package inject

import (
	app "douyin/application"
	appImpl "douyin/application/impl"
	"go.uber.org/fx"
)

var UserAppService app.UserAppService
var VideoAppService app.VideoAppService
var CommentAppService app.CommentAppService
var MessageAppService app.MessageAppService

func inject() {
	fx.New(
		fx.Provide(appImpl.NewUserAppService()),
		fx.Populate(&UserAppService))
	fx.New(
		fx.Provide(appImpl.NewVideoAppService()),
		fx.Populate(&VideoAppService))
	fx.New(
		fx.Provide(appImpl.NewCommentAppService()),
		fx.Populate(&CommentAppService))
	fx.New(
		fx.Provide(appImpl.NewMessageAppService()),
		fx.Populate(&MessageAppService))
}
