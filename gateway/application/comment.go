package application

import (
	"context"

	"douyin/types/bizdto"
)

var CommentAppIns *CommentAppService

type CommentAppService struct {
}

func NewCommentAppService() *CommentAppService {
	return &CommentAppService{}
}

func (c CommentAppService) CreateComment(ctx context.Context, appUserID int64, videoID int64, content string) (comment *bizdto.Comment, err error) {
	//TODO implement me
	panic("implement me")
}

func (c CommentAppService) DeleteComment(ctx context.Context, commentID int64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (c CommentAppService) GetCommentList(ctx context.Context, appUserID int64, videoID int64) (commentList []*bizdto.Comment, err error) {
	//TODO implement me
	panic("implement me")
}
