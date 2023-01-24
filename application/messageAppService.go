package application

import (
	"douyin/types/bizdto"
	"github.com/gin-gonic/gin"
)

type MessageAppService interface {
	CreateMessage(c *gin.Context, appUserID int64, toUserID int64, content string) (err error)
	GetMessageList(c *gin.Context, appUserID int64, toUserID int64) (messageList []*bizdto.Message, err error)
}
