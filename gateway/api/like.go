package api

// @path: gateway/api/like.go
// @description: api layer of like
// @author: Chongzhi <dczdcz2001@aliyun.com>
import (
	"github.com/gin-gonic/gin"

	"douyin/common/code"
	"douyin/common/constant"
	"douyin/gateway/application"
	"douyin/types/bizdto"
	"douyin/types/coredto"
)

// LikeAction (POST)
// Like and unlike operations for videos by logged-in users
func LikeAction(c *gin.Context) {
	appUserID := c.GetInt64(constant.IdentityKey)
	param := new(bizdto.LikeOperationReq)
	if err := c.ShouldBind(param); err != nil {
		coredto.Error(c, err)
		return
	}
	switch param.ActionType {
	case 1: // 点赞
		if err := application.VideoAppIns.LikeVideo(c, appUserID, param.VideoId); err != nil {
			coredto.Error(c, err)
			return
		}
		coredto.OK(c)
	case 2: // 取消点赞
		if err := application.VideoAppIns.UnLikeVideo(c, appUserID, param.VideoId); err != nil {
			coredto.Error(c, err)
			return
		}
		coredto.OK(c)
	default:
		coredto.Error(c, code.ParamErr)
	}
}

// LikeList (GET)
// All liked videos by logged-in users
func LikeList(c *gin.Context) {
	appUserID := c.GetInt64(constant.IdentityKey)
	param := new(bizdto.LikeListReq)
	if err := c.ShouldBind(param); err != nil {
		coredto.Error(c, err)
		return
	}
	videos, err := application.VideoAppIns.GetLikeVideoList(c, appUserID, param.UserId)
	if err != nil {
		coredto.Error(c, err)
		return
	}
	resp := &bizdto.LikeListResp{
		BaseResp:  coredto.Success,
		VideoList: videos,
	}
	coredto.Send(c, resp)
}
