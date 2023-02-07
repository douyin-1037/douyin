package service

import (
	"context"
	"douyin/code_gen/kitex_gen/videoproto"
	"douyin/video/infra/dal"
)

type LikeVideoService struct {
	ctx context.Context
}

func NewLikeVideoService(ctx context.Context) *LikeVideoService {
	return &LikeVideoService{ctx: ctx}
}

func (s *LikeVideoService) LikeVideo(req *videoproto.LikeVideoReq) error {
	// 如果插入错误，返回error
	return dal.LikeVideo(s.ctx, req.UserId, req.VideoId)
}
