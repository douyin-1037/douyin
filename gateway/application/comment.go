package application

// @path: gateway/application/comment.go
// @description: application layer of comment
// @auth: wan-nan <wan_nan@foxmail.com>
import (
	"context"
	"douyin/code_gen/kitex_gen/commentproto"
	"douyin/code_gen/kitex_gen/userproto"
	"douyin/gateway/rpc"
	"github.com/pkg/errors"

	"douyin/types/bizdto"
)

var CommentAppIns *CommentAppService

type CommentAppService struct {
}

func NewCommentAppService() *CommentAppService {
	return &CommentAppService{}
}

// CreateComment
// create a comment
func (c CommentAppService) CreateComment(ctx context.Context, appUserID int64, videoID int64, content string) (comment *bizdto.Comment, err error) {
	//panic("implement me")
	req := &commentproto.CreateCommentReq{
		UserId:  appUserID,
		VideoId: videoID,
		Content: content,
	}
	commentInfo, err := rpc.CreateComment(ctx, req)
	if err != nil {
		return nil, errors.Wrapf(err, "CreateComment rpc failed, appUserID: %v, videoID: %v, content: %s", appUserID, videoID, content)
	}
	author, err := rpc.GetUser(ctx, &userproto.GetUserReq{
		AppUserId: appUserID,
		UserId:    commentInfo.UserId,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "GetUser rpc failed, appUserID: %v, userId: %v", appUserID, commentInfo.UserId)
	}
	return toCommentDTO(commentInfo, toUserDTO(author)), nil
}

// DeleteComment
// delete a comment
func (c CommentAppService) DeleteComment(ctx context.Context, commentID int64, videoID int64) (err error) {
	//panic("implement me")
	err = rpc.DeleteComment(ctx, &commentproto.DeleteCommentReq{CommentId: commentID, VideoId: videoID})
	if err != nil {
		return errors.Wrapf(err, "DeleteComment rpc failed, commentID: %v", commentID)
	}
	return nil
}

// GetCommentList
// get comment list by video's id
func (c CommentAppService) GetCommentList(ctx context.Context, appUserID int64, videoID int64) (commentList []*bizdto.Comment, err error) {
	//panic("implement me")
	commentInfos, err := rpc.GetCommentList(ctx, &commentproto.GetCommentListReq{VideoId: videoID})
	if err != nil {
		return nil, errors.Wrapf(err, "GetCommentList rpc failed, appUserID: %v, videoID: %v", appUserID, videoID)
	}

	n := len(commentInfos)
	authors := make([]*bizdto.User, n)
	for i := 0; i < n; i++ {
		authorInfo, err := rpc.GetUser(ctx, &userproto.GetUserReq{
			AppUserId: appUserID,
			UserId:    commentInfos[i].UserId, //获取评论的作者id
		})
		if err != nil {
			return nil, errors.Wrapf(err, "GetUser rpc failed, appUserID: %v, userId: %v", appUserID, commentInfos[i].UserId)
		}
		authors[i] = toUserDTO(authorInfo)
	}
	return toCommentDTOs(commentInfos, authors), nil
}

func toCommentDTO(commentInfo *commentproto.CommentInfo, user *bizdto.User) *bizdto.Comment {
	if commentInfo == nil {
		return nil
	}
	return &bizdto.Comment{
		ID:         commentInfo.CommentId,
		User:       user,
		Content:    commentInfo.Content,
		CreateDate: commentInfo.CreateDate,
	}
}

func toCommentDTOs(commentInfos []*commentproto.CommentInfo, authors []*bizdto.User) []*bizdto.Comment {
	n := len(commentInfos)
	comments := make([]*bizdto.Comment, n)
	for i := 0; i < n; i++ {
		comments[i] = &bizdto.Comment{
			ID:         commentInfos[i].CommentId,
			User:       authors[i],
			Content:    commentInfos[i].Content,
			CreateDate: commentInfos[i].CreateDate,
		}
	}
	return comments
}
