package application

import (
	"context"

	"douyin/code_gen/kitex_gen/userproto"
	"douyin/code_gen/kitex_gen/videoproto"
	"douyin/common/conf"
	"douyin/gateway/rpc"
	"douyin/types/bizdto"
)

var VideoAppIns *VideoAppService

type VideoAppService struct {
}

func NewVideoAppService() *VideoAppService {
	return &VideoAppService{}
}

func (v VideoAppService) PublishVideo(ctx context.Context, appUserID int64, title string) (err error) {
	// need oss implementation
	/*
		fileHeader, err := ctx.FormFile("data")
		if err != nil {
			return err
		}
		file, err := fileHeader.Open()
		if err != nil {
			return err
		}
		defer file.Close()
		ossUploadReq := &oss.Video{
			Title:    title,
			Filename: "/simple-douyin/" + fileHeader.Filename,
			File:     file,
		}
		ossVideoID, err := oss.Upload(ossUploadReq)
		if err != nil {
			return err
		}
	*/
	req := &videoproto.CreateVideoReq{
		VideoBaseInfo: &videoproto.VideoBaseInfo{
			UserId: appUserID,
			//OssVideoId: ossVideoID,
			OssVideoId: "",
			Title:      title,
		},
	}
	if err := rpc.CreateVideo(ctx, req); err != nil {
		return err
	}
	return nil
}

func (v VideoAppService) LikeVideo(ctx context.Context, appUserID int64, videoID int64) (err error) {
	req := &videoproto.LikeVideoReq{
		UserId:  appUserID,
		VideoId: videoID,
	}
	if err := rpc.LikeVideo(ctx, req); err != nil {
		return err
	}
	return nil
}

func (v VideoAppService) UnLikeVideo(ctx context.Context, appUserID int64, videoID int64) (err error) {
	req := &videoproto.UnLikeVideoReq{
		UserId:  appUserID,
		VideoId: videoID,
	}
	if err := rpc.UnLikeVideo(ctx, req); err != nil {
		return err
	}
	return nil
}

func (v VideoAppService) GetVideoList(ctx context.Context, appUserID int64, userID int64) (videoList []*bizdto.Video, err error) {
	req := &videoproto.GetVideoListByUserIdReq{
		AppUserId: appUserID,
		UserId:    userID,
	}
	videos, err := rpc.GetVideoListByUserId(ctx, req)
	if err != nil {
		return nil, err
	}
	// get authors
	n := len(videos)
	authors := make([]*userproto.UserInfo, n)
	for i := 0; i < n; i++ {
		subReq := &userproto.GetUserReq{
			AppUserId: appUserID,
			UserId:    videos[i].VideoBaseInfo.UserId,
		}
		authors[i], err = rpc.GetUser(ctx, subReq)
		if err != nil {
			return nil, err
		}
	}
	// pack videos and authors
	packedVideos, err := toVideoDTOs(videos, authors)
	if err != nil {
		return nil, err
	}
	return packedVideos, nil
}

func (v VideoAppService) GetLikeVideoList(ctx context.Context, appUserID int64, userID int64) (userList []*bizdto.Video, err error) {
	req := &videoproto.GetLikeVideoListReq{
		AppUserId: appUserID,
		UserId:    userID,
	}
	videos, err := rpc.GetLikeVideoList(ctx, req)
	if err != nil {
		return nil, err
	}
	n := len(videos)
	authors := make([]*userproto.UserInfo, n)
	for i := 0; i < n; i++ {
		subReq := &userproto.GetUserReq{
			AppUserId: appUserID,
			UserId:    videos[i].VideoBaseInfo.UserId,
		}
		authors[i], err = rpc.GetUser(ctx, subReq)
		if err != nil {
			return nil, err
		}
	}
	packedVideos, err := toVideoDTOs(videos, authors)
	if err != nil {
		return nil, err
	}
	return packedVideos, nil
}

func (v VideoAppService) Feed(ctx context.Context, appUserID int64, latestTime int64) (videoList []*bizdto.Video, nextTime int64, err error) {
	req := &videoproto.GetVideoListByTimeReq{
		AppUserId:  appUserID,
		LatestTime: latestTime,
		Count:      conf.Server.FeedCount,
	}
	videos, nextTime, err := rpc.GetVideoListByTime(ctx, req)
	if err != nil {
		return nil, 0, err
	}
	n := len(videos)
	authors := make([]*userproto.UserInfo, n)
	for i := 0; i < n; i++ {
		subReq := &userproto.GetUserReq{
			AppUserId: appUserID,
			UserId:    videos[i].VideoBaseInfo.UserId,
		}
		authors[i], err = rpc.GetUser(ctx, subReq)
		if err != nil {
			return nil, 0, err
		}
	}
	packedVideos, err := toVideoDTOs(videos, authors)
	if err != nil {
		return nil, 0, err
	}
	return packedVideos, nextTime, nil
}

// toVideoDTO
// transform one videoproto.VideoInfo into one bizdto.Video with author information
func toVideoDTO(v *videoproto.VideoInfo, author *userproto.UserInfo) (*bizdto.Video, error) {
	// need redis implementation
	/*
			playURL, err := cache.GetPlayURL(v.VideoBaseInfo.OssVideoId)
			if err != nil {
				return nil, err
			}
			coverURL, err := cache.GetCoverURL(v.VideoBaseInfo.OssVideoId)
			if err != nil {
				return nil, err
			}
			// coverURL := "https://tva1.sinaimg.cn/large/e6c9d24ely1h2wrrikp8uj20tc1io422.jpg"
		}
	*/
	return &bizdto.Video{
		ID:     v.VideoId,
		Author: toUserDTO(author),
		/*
			PlayAddr:     playURL,
			CoverAddr:    coverURL,
		*/
		PlayAddr:     "",
		CoverAddr:    "",
		LikeCount:    v.LikeCount,
		CommentCount: v.CommentCount,
		IsFavorite:   v.IsFavorite,
		Title:        v.VideoBaseInfo.Title,
	}, nil
}

// toVideoDTOs
// apply toVideoDTO to an array of videoproto.VideoInfo
func toVideoDTOs(vs []*videoproto.VideoInfo, authors []*userproto.UserInfo) ([]*bizdto.Video, error) {
	videos := make([]*bizdto.Video, len(vs))
	var err error
	for i, v := range vs {
		videos[i], err = toVideoDTO(v, authors[i])
		if err != nil {
			return nil, err
		}
	}
	return videos, nil
}
