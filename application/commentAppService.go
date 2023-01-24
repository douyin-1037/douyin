package application

import (
	"douyin/types/bizdto"
	"github.com/gin-gonic/gin"
)

type CommentAppService interface {
	CreateComment(c *gin.Context, appUserID int64, videoID int64, content string) (comment *bizdto.Comment, err error)
	DeleteComment(c *gin.Context, commentID int64) (err error)
	GetCommentList(c *gin.Context, appUserID int64, videoID int64) (commentList []*bizdto.Comment, err error)
}
