package main

import (
	"douyin/gateway/api"
	"douyin/gateway/api/auth"

	"github.com/gin-gonic/gin"
)

func register(r *gin.Engine) {
	unAuthGroup := r.Group("/douyin")
	{
		unAuthGroup.GET("/feed", api.Feed)
	}

	authGroup := r.Group("/douyin")
	authGroup.Use(auth.JWT())
	{
		authGroup.POST("/publish/action/", api.Upload)
		authGroup.GET("/publish/list/", api.List)
		authGroup.POST("/favorite/action/", api.LikeAction)
		authGroup.GET("/favorite/list/", api.LikeList)
		authGroup.POST("/comment/action/", api.CommentAction)
		authGroup.GET("/comment/list/", api.CommentList)
		authGroup.GET("/user/", api.GetUserInfo)
	}
}
