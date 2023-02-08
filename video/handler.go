package main

import (
	"context"
	"douyin/code_gen/kitex_gen/videoproto"
	"douyin/pkg/code"
	"douyin/video/pack"
	"douyin/video/service"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// CreateVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CreateVideo(ctx context.Context, req *videoproto.CreateVideoReq) (resp *videoproto.CreateVideoResp, err error) {
	resp = new(videoproto.CreateVideoResp)

	if req.VideoBaseInfo.UserId <= 0 || len(req.VideoBaseInfo.CoverUrl) == 0 || len(req.VideoBaseInfo.PlayUrl) == 0 || len(req.VideoBaseInfo.Title) == 0 {
		resp.BaseResp = pack.BuildBaseResp(code.ParamErr)
		return resp, nil
	}
	err = service.NewCreateVideoService(ctx).CreateVideo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(code.Success)
	return resp, nil
}

// GetVideoListByUserId implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideoListByUserId(ctx context.Context, req *videoproto.GetVideoListByUserIdReq) (resp *videoproto.GetVideoListByUserIdResp, err error) {
	resp = new(videoproto.GetVideoListByUserIdResp)

	if req.UserId < 0 || req.AppUserId < 0 {
		resp.BaseResp = pack.BuildBaseResp(code.ParamErr)
		return resp, nil
	}

	videos, err := service.NewMGetVideoByUserIdService(ctx).MGetVideo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(code.Success)
	resp.VideoInfos = videos
	return resp, nil
}

// GetVideoListByTime implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideoListByTime(ctx context.Context, req *videoproto.GetVideoListByTimeReq) (resp *videoproto.GetVideoListByTimeResp, err error) {
	resp = new(videoproto.GetVideoListByTimeResp)
	if (req.AppUserId < 0 && req.AppUserId != -1) || req.Count > 1000 || req.LatestTime < 0 { // count限制小于1000，避免查询过多数据
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	videos, nextTime, err := service.NewMGetVideoByTimeService(ctx).MGetVideoByTime(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(code.Success)
	resp.VideoInfos = videos
	resp.NextTime = nextTime
	return resp, nil
}

// LikeVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) LikeVideo(ctx context.Context, req *videoproto.LikeVideoReq) (resp *videoproto.LikeVideoResp, err error) {
	resp = new(videoproto.LikeVideoResp)

	if req.UserId < 0 || req.VideoId < 0 {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	err = service.NewLikeVideoService(ctx).LikeVideo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(code.Success)
	return resp, nil
}

// UnLikeVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) UnLikeVideo(ctx context.Context, req *videoproto.UnLikeVideoReq) (resp *videoproto.UnLikeVideoResp, err error) {
	resp = new(videoproto.UnLikeVideoResp)

	if req.UserId < 0 || req.VideoId < 0 {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	err = service.NewUnLikeVideoService(ctx).UnLikeVideo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(code.Success)
	return resp, nil
}

// GetLikeVideoList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetLikeVideoList(ctx context.Context, req *videoproto.GetLikeVideoListReq) (resp *videoproto.GetLikeVideoListResp, err error) {
	resp = new(videoproto.GetLikeVideoListResp)
	if req.AppUserId < 0 || req.AppUserId != req.UserId {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	videos, err := service.NewMGetLikeVideoService(ctx).MGetLikeVideo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(code.Success)
	resp.VideoInfos = videos
	return resp, nil
}
