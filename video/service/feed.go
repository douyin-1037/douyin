package service

import (
	"context"
	"douyin/code_gen/kitex_gen/videoproto"
	"douyin/video/infra/dal"
	"douyin/video/infra/redis"
	"douyin/video/pack"
	"time"
)

type MGetVideoByTimeService struct {
	ctx context.Context
}

func NewMGetVideoByTimeService(ctx context.Context) *MGetVideoByTimeService {
	return &MGetVideoByTimeService{ctx: ctx}
}

// MGetVideoByTime 通过指定latestTime和count，从DAO层获取视频基本信息，并查出当前用户是否点赞，组装后返回
func (s *MGetVideoByTimeService) MGetVideoByTime(req *videoproto.GetVideoListByTimeReq) ([]*videoproto.VideoInfo, int64, error) {
	videoModels, nextTime, err := dal.MGetVideoByTime(s.ctx, time.Unix(req.LatestTime, 0), req.Count)
	// 只能得到视频id，uid，title，play_url,cover_url,created_time
	if err != nil {
		return nil, 0, err
	}
	videos := pack.Videos(videoModels) // 类型转换：视频id、base_info、点赞数、评论数已经得到，还需要判断是否点赞
	// 把视频的其他信息进行绑定
	appUserID := req.AppUserId
	isLikeKeyExist, err := redis.IsLikeKeyExist(appUserID)
	if err != nil {
		return nil, 0, err
	}
	if isLikeKeyExist == false {
		// 如果redis没有appUserID的记录，则去mysql查询一次点赞列表进行缓存
		likeList, err := dal.MGetLikeList(s.ctx, appUserID)
		if err != nil {
			return nil, 0, err
		}
		if err := redis.AddLikeList(appUserID, likeList); err != nil {
			return nil, 0, err
		}
	}
	for i := 0; i < len(videos); i++ {
		if appUserID > 0 { // 判断是否进行了登陆
			isFavorite, err := redis.GetIsLikeById(appUserID, videos[i].VideoId)
			if err != nil {
				return nil, 0, err
			}
			videos[i].IsFavorite = isFavorite
		} else { // 如果没有登陆，则点赞直接返回false
			videos[i].IsFavorite = false
		}
	}
	return videos, nextTime, nil
}
