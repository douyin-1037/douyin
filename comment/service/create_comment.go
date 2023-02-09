package service

// @path: comment/service/create_comment.go
// @description: CreateComment service of comment
// @auth: wan-nan <wan_nan@foxmail.com>
import (
	"context"
	"douyin/code_gen/kitex_gen/commentproto"
	"douyin/comment/infra/dal"
)

type CreateCommentService struct {
	ctx context.Context
}

func NewCreateCommentService(ctx context.Context) *CreateCommentService {
	return &CreateCommentService{
		ctx: ctx,
	}
}

func (s *CreateCommentService) CreateComment(req *commentproto.CreateCommentReq) error {
	return dal.CreateComment(s.ctx, req.UserId, req.VideoId, req.Content)
}
