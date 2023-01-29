package main

import (
	"context"
	videoproto "douyin/code_gen/kitex_gen/videoproto"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// CreateVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CreateVideo(ctx context.Context, req *videoproto.CreateVideoReq) (resp *videoproto.CreateVideoResp, err error) {
	// TODO: Your code here...
	return
}

// GetVideoListByUserId implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideoListByUserId(ctx context.Context, req *videoproto.GetVideoListByUserIdReq) (resp *videoproto.GetVideoListByUserIdResp, err error) {
	// TODO: Your code here...
	return
}

// GetVideoListByTime implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideoListByTime(ctx context.Context, req *videoproto.GetVideoListByTimeReq) (resp *videoproto.GetVideoListByTimeResp, err error) {
	// TODO: Your code here...
	return
}

// LikeVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) LikeVideo(ctx context.Context, req *videoproto.LikeVideoReq) (resp *videoproto.LikeVideoResp, err error) {
	// TODO: Your code here...
	return
}

// UnLikeVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) UnLikeVideo(ctx context.Context, req *videoproto.UnLikeVideoReq) (resp *videoproto.UnLikeVideoResp, err error) {
	// TODO: Your code here...
	return
}

// GetLikeVideoList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetLikeVideoList(ctx context.Context, req *videoproto.GetLikeVideoListReq) (resp *videoproto.GetLikeVideoListResp, err error) {
	// TODO: Your code here...
	return
}
