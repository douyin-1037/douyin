package application

import (
	"douyin/types/bizdto"
	"github.com/gin-gonic/gin"
)

type UserAppService interface {
	GetUser(c *gin.Context, appUserID int64, userID int64) (user *bizdto.User, err error)
	//CreateUser(c *gin.Context, username string, password string) (userID int64, err error)
	//CheckUser(c *gin.Context, username string, password string) (userID int64, err error)
	//FollowUser(c *gin.Context, fanID int64, toUserID int64) (err error)
	//UnFollowUser(c *gin.Context, fanID int64, toUserID int64) (err error)
	//GetFollowList(c *gin.Context, appUserID int64, userID int64) (userList []*bizdto.User, err error)
	//GetFanList(c *gin.Context, appUserID int64, userID int64) (userList []*bizdto.User, err error)
	//GetFriendList(c *gin.Context, appUserID int64, userID int64) (userList []*bizdto.User, err error)
}
