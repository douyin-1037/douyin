package api

// @path: gateway/api/comment.go
// @description: api layer of comment
// @auth: wan-nan <wan_nan@foxmail.com>
import (
	"github.com/gin-gonic/gin"

	"douyin/common/code"
	"douyin/common/constant"
	"douyin/gateway/application"
	"douyin/types/bizdto"
	"douyin/types/coredto"
)

// CommentAction (POST): create comment on video or delete one comment
func CommentAction(c *gin.Context) {
	appUserID := c.GetInt64(constant.IdentityKey)
	param := new(bizdto.CommentOperationReq)
	if err := c.ShouldBind(param); err != nil {
		coredto.Error(c, err)
		return
	}
	switch param.ActionType {
	case 1: // create comment on a video
		comment, err := application.CommentAppIns.CreateComment(c, appUserID, param.VideoId, param.CommentText)
		if err != nil {
			coredto.Error(c, err)
			return
		}
		//author, err := application.UserAppIns.GetUser(c, appUserID, comment.User.ID)
		//if err != nil {
		//	coredto.Error(c, err)
		//	return
		//}
		resp := &bizdto.CreateCommentResp{
			BaseResp: coredto.Success,
			Comment:  comment,
		}
		coredto.Send(c, resp)
	case 2: // delete one comment
		if err := application.CommentAppIns.DeleteComment(c, param.CommentId, param.VideoId); err != nil {
			coredto.Error(c, err)
			return
		}
		coredto.OK(c)
	default:
		coredto.Error(c, code.ParamErr)
	}
}

// CommentList (GET): get comment list of one video
func CommentList(c *gin.Context) {
	appUserID := c.GetInt64(constant.IdentityKey)
	param := new(bizdto.CommentListReq)
	if err := c.ShouldBind(param); err != nil {
		coredto.Error(c, err)
		return
	}
	comments, err := application.CommentAppIns.GetCommentList(c, appUserID, param.VideoId)
	if err != nil {
		coredto.Error(c, err)
		return
	}
	resp := &bizdto.CommentListResp{
		BaseResp:    coredto.Success,
		CommentList: comments,
	}
	coredto.Send(c, resp)
}
