package service

// @path: comment/service/get_comment_list.go
// @description: GetCommentList service of comment
// @auth: wan-nan <wan_nan@foxmail.com>
import (
	"context"
	"douyin/code_gen/kitex_gen/commentproto"
	"douyin/comment/infra/dal"
	"douyin/comment/pack"
)

type GetCommentListService struct {
	ctx context.Context
}

func NewGetCommentListService(ctx context.Context) *GetCommentListService {
	return &GetCommentListService{
		ctx: ctx,
	}
}

func (s *GetCommentListService) GetCommentList(req *commentproto.GetCommentListReq) ([]*commentproto.CommentInfo, error) {
	comments, err := dal.GetCommentList(s.ctx, req.VideoId)
	if err != nil {
		return nil, err
	}
	return pack.Comments(comments), nil
}
