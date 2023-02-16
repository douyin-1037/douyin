package api

import (
	"douyin/gateway/api/auth"
	"github.com/gin-gonic/gin"

	"douyin/common/constant"
	"douyin/gateway/application"
	"douyin/types/bizdto"
	respond "douyin/types/coredto"
)

func GetUserInfo(c *gin.Context) {
	appUserID := c.GetInt64(constant.IdentityKey)
	param := new(bizdto.UserQueryReq)
	if err := c.ShouldBind(param); err != nil {
		respond.Error(c, err)
		return
	}
	//req := &userproto.GetUserReq{
	//	AppUserId: appUserID,
	//	UserId:    param.UserId,
	//}
	//user, err := rpc.GetUser(c, req)

	//调用app层接口
	user, err := application.UserAppIns.GetUser(c, appUserID, param.UserId)
	if err != nil {
		respond.Error(c, err)
		return
	}
	resp := &bizdto.UserQueryResp{
		BaseResp: respond.Success,
		User:     user,
	}
	respond.Send(c, resp)
}

// Create User registration interface
func Create(c *gin.Context) {
	param := new(bizdto.UserRegisterReq)
	if err := c.ShouldBind(param); err != nil {
		respond.Send(c, &bizdto.UserRegisterResp{BaseResp: respond.ErrorBaseResp(err)})
		return
	}

	userID, err := application.UserAppIns.CreateUser(c, param.Username, param.Password)
	if err != nil {
		respond.Send(c, &bizdto.UserRegisterResp{BaseResp: respond.ErrorBaseResp(err)})
		return
	}

	token, err := auth.GenerateToken(userID)
	if err != nil {
		respond.Send(c, &bizdto.UserRegisterResp{BaseResp: respond.ErrorBaseResp(err)})
		return
	}
	resp := &bizdto.UserRegisterResp{
		BaseResp: respond.Success,
		UserID:   userID,
		Token:    token,
	}
	respond.Send(c, resp)
}

// Check User login interface
func Check(c *gin.Context) {
	param := new(bizdto.UserLoginReq)
	if err := c.ShouldBind(param); err != nil {
		respond.Error(c, err)
		return
	}

	userID, err := application.UserAppIns.CheckUser(c, param.Username, param.Password)
	if err != nil {
		respond.Error(c, err)
		return
	}
	token, err := auth.GenerateToken(userID)
	if err != nil {
		respond.Error(c, err)
		return
	}
	resp := &bizdto.UserLoginResp{
		BaseResp: respond.Success,
		UserID:   userID,
		Token:    token,
	}
	respond.Send(c, resp)
}
