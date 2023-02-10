package service

// @path: comment/service/delete_comment.go
// @description: DeleteComment service of comment
// @auth: wan-nan <wan_nan@foxmail.com>
import (
	"context"
	"douyin/code_gen/kitex_gen/commentproto"
	"douyin/comment/infra/dal"
)

type DeleteCommentService struct {
	ctx context.Context
}

func NewDeleteCommentService(ctx context.Context) *DeleteCommentService {
	return &DeleteCommentService{
		ctx: ctx,
	}
}

func (s *DeleteCommentService) DeleteComment(req *commentproto.DeleteCommentReq) error {
	return dal.DeleteComment(s.ctx, req.CommentId, req.VideoId)
}
