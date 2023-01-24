package handlers

import (
	"douyin/cmd/inject"
)

var userService = inject.UserAppService
var videoService = inject.VideoAppService
var commentService = inject.CommentAppService
var messageService = inject.MessageAppService
