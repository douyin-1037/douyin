package api

import (
	"douyin/common/code"
	"douyin/gateway/api/auth"
	"douyin/gateway/application"
	"douyin/types/bizdto"
	"douyin/types/coredto"

	"github.com/gin-gonic/gin"
)

// Handle the GET request of /message/chat/, return the message list between appUser and toUser
func GetMessageList(c *gin.Context) {
	appUserID, err := auth.GetUserID(c)
	if err != nil {
		coredto.Error(c, err)
		return
	}

	param := new(bizdto.MessageListReq)
	if err := c.ShouldBind(param); err != nil {
		coredto.Error(c, err)
		return
	}
	messageList, err := application.MessageAppIns.GetMessageList(c, appUserID, param.ToUserId)
	if err != nil {
		coredto.Error(c, err)
	}

	response := bizdto.MessageListResp{
		BaseResp:    coredto.Success,
		MessageList: messageList,
	}
	coredto.Send(c, &response)
}

// Handle the POST request of /message/action/, currently only support message sending
func HandleMessageAction(c *gin.Context) {
	appUserID, err := auth.GetUserID(c)
	if err != nil {
		coredto.Error(c, err)
		return
	}

	param := new(bizdto.MessageOperationReq)
	if err := c.ShouldBind(param); err != nil {
		coredto.Error(c, err)
		return
	}

	switch param.ActionType {
	case 1:
		err := application.MessageAppIns.CreateMessage(c, appUserID, param.ToUserId, param.Content)
		if err != nil {
			coredto.Error(c, err)
		} else {
			coredto.OK(c)
		}
	default:
		coredto.Error(c, code.ParamErr)
	}
}
