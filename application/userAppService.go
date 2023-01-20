package application

import (
	"douyin/types/bizdto"
	"github.com/gin-gonic/gin"
)

type UserAppService interface {
	GetUser(c *gin.Context, appUserID int64, userID int64) (user *bizdto.User, err error)
	CreateUser(c *gin.Context, appUserID int64, userID int64) (user *bizdto.User, err error)
}
