package router

import (
	"douyin/api/auth"
	"douyin/api/handlers"
	"douyin/common/conf"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	if conf.Server.RunMode == "debug" {
		r.Use(gin.Logger(), gin.Recovery())
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	/*
		unauthGroup := r.Group("/douyin")
		{
		}
	*/
	authGroup := r.Group("/douyin")
	authGroup.Use(auth.JWT())
	{
		authGroup.POST("/favorite/action/", handlers.LikeAction)
		authGroup.GET("/favorite/list/", handlers.LikeList)
		authGroup.POST("/comment/action/", handlers.CommentAction)
		authGroup.GET("/comment/list/", handlers.CommentList)
		authGroup.GET("/user/", handlers.GetUserInfo)
	}
	return r
}
