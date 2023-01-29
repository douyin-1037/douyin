package application

import (
	"context"

	"douyin/types/bizdto"
)

var VideoAppIns *VideoAppService

type VideoAppService struct {
}

func NewVideoAppService() *VideoAppService {
	return &VideoAppService{}
}

func (v VideoAppService) PublishVideo(ctx context.Context, appUserID int64, title string) (err error) {
	//TODO implement me
	panic("implement me")
}

func (v VideoAppService) LikeVideo(ctx context.Context, appUserID int64, videoID int64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (v VideoAppService) UnLikeVideo(ctx context.Context, appUserID int64, videoID int64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (v VideoAppService) GetVideoList(ctx context.Context, appUserID int64, userID int64) (videoList []*bizdto.Video, err error) {
	//TODO implement me
	panic("implement me")
}

func (v VideoAppService) GetLikeVideoList(ctx context.Context, appUserID int64, userID int64) (userList []*bizdto.Video, err error) {
	//TODO implement me
	panic("implement me")
}

func (v VideoAppService) Feed(ctx context.Context, appUserID int64, latestTime int64) (videoList []*bizdto.Video, nextTime int64, err error) {
	//TODO implement me
	panic("implement me")
}
