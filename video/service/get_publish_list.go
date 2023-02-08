package service

import (
	"context"
	"douyin/code_gen/kitex_gen/videoproto"
	"douyin/video/infra/dal"
	"douyin/video/pack"
)

type MGetVideoByUserIdService struct {
	ctx context.Context
}

func NewMGetVideoByUserIdService(ctx context.Context) *MGetVideoByUserIdService {
	return &MGetVideoByUserIdService{ctx: ctx}
}

// MGetVideo 通过UserID从DAO层获取视频基本信息，并查出当前用户是否点赞，组装后返回
func (s *MGetVideoByUserIdService) MGetVideo(req *videoproto.GetVideoListByUserIdReq) ([]*videoproto.VideoInfo, error) {
	videoModels, err := dal.MGetVideoByUserID(s.ctx, req.UserId)
	// 只能得到视频id,uid,title，play_url,cover_url,created_time
	if err != nil {
		return nil, err
	}
	videos := pack.Videos(videoModels) // 做类型转换：视频id、base_info、点赞数、评论数已经得到，还需要判断是否点赞
	// 把视频的其他信息进行绑定
	appUserID := req.AppUserId
	for i := 0; i < len(videos); i++ {
		vid := videos[i].VideoId
		isFavorite, err := dal.IsFavorite(s.ctx, vid, appUserID)
		if err != nil {
			return nil, err
		}
		videos[i].IsFavorite = isFavorite
	}
	return videos, nil
}
