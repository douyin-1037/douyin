package pack

// @path: video/pack/video.go
// @description: pack []*model.Video into []*videoproto.VideoInfo
// @author: Chongzhi <dczdcz2001@aliyun.com>
import (
	"douyin/code_gen/kitex_gen/videoproto"
	"douyin/video/infra/dal/model"
)

// Video pack video info : video to videoproto.VideoInfo
func Video(m *model.Video) *videoproto.VideoInfo {
	if m == nil {
		return nil
	}
	return &videoproto.VideoInfo{
		VideoBaseInfo: &videoproto.VideoBaseInfo{
			UserId:   int64(m.UserId),
			PlayUrl:  m.PlayUrl,
			CoverUrl: m.CoverUrl,
			Title:    m.Title,
		},
		VideoId:      int64(m.ID),
		LikeCount:    m.FavoriteCount,
		CommentCount: m.CommentCount,
	}
}

func Videos(ms []*model.Video) []*videoproto.VideoInfo {
	videos := make([]*videoproto.VideoInfo, len(ms))
	for i, m := range ms {
		if n := Video(m); n != nil {
			videos[i] = n
		}
	}
	return videos
}
