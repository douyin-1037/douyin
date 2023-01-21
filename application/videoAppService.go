package application

import (
	"douyin/types/bizdto"
	"github.com/gin-gonic/gin"
)

type VideoAppService interface {
	PublishVideo(c *gin.Context, appUserID int64, title string) (err error)
	LikeVideo(c *gin.Context, appUserID int64, videoID int64) (err error)
	UnLikeVideo(c *gin.Context, appUserID int64, videoID int64) (err error)
	GetVideoList(c *gin.Context, appUserID int64, userID int64) (videoList []*bizdto.Video, err error)
	GetLikeVideoList(c *gin.Context, appUserID int64, userID int64) (userList []*bizdto.Video, err error)
	Feed(c *gin.Context, appUserID int64, latestTime int64) (videoList []*bizdto.Video, err error)
}
