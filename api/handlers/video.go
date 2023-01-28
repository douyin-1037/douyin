package handlers

import (
	"douyin/api/auth"
	"douyin/common/constant"
	"douyin/types/bizdto"
	respond "douyin/types/coredto"
	"github.com/gin-gonic/gin"
	"time"
)

// Feed @title    Video Streaming (GET)
// @description   unlimited login status, return the list of videos in reverse chronological order of submission, the number of videos is controlled by the server, up to 30 in a single time
func Feed(c *gin.Context) {
	appUserID, err := auth.GetUserID(c)
	if err != nil { // Cases in which the user is not logged in
		appUserID = -1
	}

	param := new(bizdto.VideoFeedReq)
	if err := c.ShouldBind(param); err != nil {
		respond.Error(c, err)
		return
	}
	if param.LatestTime <= 0 {
		param.LatestTime = time.Now().Unix()
	}

	videoList, nextTime, err := videoService.Feed(c, appUserID, param.LatestTime)
	if err != nil {
		respond.Error(c, err)
		return
	}

	resp := &bizdto.VideoFeedResp{
		BaseResp:  respond.Success,
		NextTime:  nextTime,
		VideoList: videoList,
	}
	respond.Send(c, resp)
}

// Upload @title upload video (POST)
// @description Login user to select video upload
func Upload(c *gin.Context) {
	appUserID := c.GetInt64(constant.IdentityKey)
	param := new(bizdto.VideoUploadReq)
	if err := c.ShouldBind(param); err != nil {
		respond.Error(c, err)
		return
	}

	if err := videoService.PublishVideo(c, appUserID, param.Title); err != nil {
		respond.Error(c, err)
		return
	}
	respond.OK(c)
}

// List @title upload list (GET)
// @description Log in to the user's video posting list and directly list all the videos that the user has contributed to
func List(c *gin.Context) {
	appUserID := c.GetInt64(constant.IdentityKey)
	param := new(bizdto.VideoListReq)
	if err := c.ShouldBind(param); err != nil {
		respond.Error(c, err)
		return
	}

	videoList, err := videoService.GetVideoList(c, appUserID, param.UserId)
	if err != nil {
		respond.Error(c, err)
		return
	}

	resp := &bizdto.VideoListResp{
		BaseResp:  respond.Success,
		VideoList: videoList,
	}
	respond.Send(c, resp)
}
