package api

import (
	"douyin/common/code"
	"douyin/common/constant"
	"douyin/types/bizdto"
	respond "douyin/types/coredto"

	"github.com/gin-gonic/gin"
)

func Follow(c *gin.Context) {
	appUserID := c.GetInt64(constant.IdentityKey)
	param := new(bizdto.FollowOperationReq)
	if err := c.ShouldBind(param); err != nil {
		respond.Error(c, err)
		return
	}
	switch param.ActionType {
	case 1: // follow
		err := userService.FollowUser(c, appUserID, param.ToUserId)
		if err != nil {
			respond.Error(c, err)
			return
		}
		respond.OK(c)
	case 2: // unfollow
		err := userService.UnFollowUser(c, appUserID, param.ToUserId)
		if err != nil {
			respond.Error(c, err)
			return
		}
		respond.OK(c)
	default:
		respond.Error(c, code.ParamErr)
	}
}

func FollowList(c *gin.Context) {
	appUserID := c.GetInt64(constant.IdentityKey)
	param := new(bizdto.FollowListReq)
	if err := c.ShouldBind(param); err != nil {
		respond.Error(c, err)
		return
	}
	users, err := userService.GetFollowList(c, appUserID, param.UserId)
	if err != nil {
		respond.Error(c, err)
		return
	}
	resp := &bizdto.FollowListResp{
		BaseResp: respond.Success,
		UserList: users,
	}
	respond.Send(c, resp)
}

func FanList(c *gin.Context) {
	appUserID := c.GetInt64(constant.IdentityKey)
	param := new(bizdto.FanListReq)
	if err := c.ShouldBind(param); err != nil {
		respond.Error(c, err)
		return
	}
	users, err := userService.GetFanList(c, appUserID, param.UserId)
	if err != nil {
		respond.Error(c, err)
		return
	}
	resp := &bizdto.FanListResp{
		BaseResp: respond.Success,
		UserList: users,
	}
	respond.Send(c, resp)
}

func FriendList(c *gin.Context) {
	appUserID := c.GetInt64(constant.IdentityKey)
	param := new(bizdto.FriendListReq)
	if err := c.ShouldBind(param); err != nil {
		respond.Error(c, err)
		return
	}
	users, err := userService.GetFriendList(c, appUserID, param.UserId)
	if err != nil {
		respond.Error(c, err)
		return
	}
	resp := &bizdto.FriendListResp{
		BaseResp: respond.Success,
		UserList: users,
	}
	respond.Send(c, resp)
}
