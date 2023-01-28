package api

import (
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
