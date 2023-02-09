package service

import (
	"context"
	"douyin/code_gen/kitex_gen/videoproto"
	"douyin/video/infra/dal"
)

type CreateVideoService struct {
	ctx context.Context
}

func NewCreateVideoService(ctx context.Context) *CreateVideoService {
	return &CreateVideoService{ctx: ctx}
}

func (s *CreateVideoService) CreateVideo(req *videoproto.CreateVideoReq) error {
	// 如果添加失败，返回error
	return dal.CreateVideo(s.ctx, req.VideoBaseInfo.UserId, req.VideoBaseInfo.Title, req.VideoBaseInfo.PlayUrl, req.VideoBaseInfo.CoverUrl)
}
