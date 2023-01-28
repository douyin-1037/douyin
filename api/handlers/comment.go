package handlers

import (
	"douyin/common/code"
	"douyin/common/constant"
	"douyin/types/bizdto"
	"douyin/types/coredto"
	"github.com/gin-gonic/gin"
)

// CommentAction 评论操作(POST)：登录用户对视频的评论操作和对特定评论的删除操作
func CommentAction(c *gin.Context) {
	appUserID := c.GetInt64(constant.IdentityKey)
	param := new(bizdto.CommentOperationReq)
	if err := c.ShouldBind(param); err != nil {
		coredto.Error(c, err)
		return
	}
	switch param.ActionType {
	case 1: // 评论
		comment, err := commentService.CreateComment(c, appUserID, param.VideoId, param.CommentText)
		if err != nil {
			coredto.Error(c, err)
			return
		}
		author, err := UserService.GetUser(c, appUserID, comment.User.ID)
		if err != nil {
			coredto.Error(c, err)
			return
		}
		resp := &bizdto.CreateCommentResp{
			BaseResp: coredto.Success,
			Comment: &bizdto.Comment{
				ID:         comment.ID,
				User:       author,
				Content:    comment.Content,
				CreateDate: comment.CreateDate,
			},
		}
		coredto.Send(c, resp)
	case 2: // 删除评论
		if err := commentService.DeleteComment(c, param.CommentId); err != nil {
			coredto.Error(c, err)
			return
		}
		coredto.OK(c)
	default:
		coredto.Error(c, code.ParamErr)
	}
}

// CommentList 评论列表(GET)：获取登录用户的所有评论
func CommentList(c *gin.Context) {
	appUserID := c.GetInt64(constant.IdentityKey)
	param := new(bizdto.CommentListReq)
	if err := c.ShouldBind(param); err != nil {
		coredto.Error(c, err)
		return
	}
	comments, err := commentService.GetCommentList(c, appUserID, param.VideoId)
	if err != nil {
		coredto.Error(c, err)
		return
	}
	n := len(comments)
	authors := make([]*bizdto.User, n)
	for i := 0; i < n; i++ {
		authors[i], err = UserService.GetUser(c, appUserID, comments[i].User.ID)
		if err != nil {
			coredto.Error(c, err)
			return
		}
	}
	resp := &bizdto.CommentListResp{
		BaseResp:    coredto.Success,
		CommentList: comments,
	}
	coredto.Send(c, resp)
}
