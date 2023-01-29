package main

import (
	"context"
	commentproto "douyin/code_gen/kitex_gen/commentproto"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CreateComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CreateComment(ctx context.Context, req *commentproto.CreateCommentReq) (resp *commentproto.CreateCommentResp, err error) {
	// TODO: Your code here...
	return
}

// DeleteComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) DeleteComment(ctx context.Context, req *commentproto.DeleteCommentReq) (resp *commentproto.DeleteCommentResp, err error) {
	// TODO: Your code here...
	return
}

// GetCommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) GetCommentList(ctx context.Context, req *commentproto.GetCommentListReq) (resp *commentproto.GetCommentListResp, err error) {
	// TODO: Your code here...
	return
}
