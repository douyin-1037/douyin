package service

import (
	"context"
	"douyin/code_gen/kitex_gen/videoproto"
	"douyin/video/infra/dal"
	"douyin/video/pack"
)

type MGetLikeVideoService struct {
	ctx context.Context
}

func NewMGetLikeVideoService(ctx context.Context) *MGetLikeVideoService {
	return &MGetLikeVideoService{ctx: ctx}
}

// MGetLikeVideo 通过用户ID从DAO层获取喜欢视频的基本信息，并查出当前用户是否点赞，组装后返回
func (s *MGetLikeVideoService) MGetLikeVideo(req *videoproto.GetLikeVideoListReq) ([]*videoproto.VideoInfo, error) {
	favorites, err := dal.MGetLikeList(s.ctx, req.AppUserId)
	if err != nil {
		return nil, err
	}
	// 只能得到视频id，uid,title,play_url,cover_url,created_time
	videoModels, err := dal.MGetLikeVideo(s.ctx, favorites)
	if err != nil {
		return nil, err
	}
	videos := pack.Videos(videoModels) // 做类型转换：视频id、base_info、点赞数、评论数已经得到，还需要判断是否点赞
	// 把视频的其他信息进行绑定
	for i := 0; i < len(videos); i++ {
		videos[i].IsFavorite = true // 返回的视频都是已经点赞的
	}
	return videos, nil
}
