package handlers

import (
	app "douyin/application"
	"douyin/application/impl"
)

var UserService app.UserAppService
var videoService app.VideoAppService
var commentService app.CommentAppService
var messageService app.MessageAppService

func InjectAppService(userServiceImpl *impl.UserAppServiceImpl) {
	UserService = userServiceImpl
}
