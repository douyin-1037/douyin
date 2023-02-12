package dal

// @path: video/infra/dal/orm.go
// @description: DAL of video
// @author: Chongzhi <dczdcz2001@aliyun.com>
import (
	"context"
	"douyin/video/infra/dal/model"
	"math"
	"time"
)

// CreateVideo 创建视频
func CreateVideo(ctx context.Context, userID int64, title string, playUrl string, coverUrl string) error {
	video := &model.Video{
		UserId:   userID,
		Title:    title,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
	}
	return DB.WithContext(ctx).Create(&video).Error
}

// MGetVideoByUserID 根据用户id查视频
func MGetVideoByUserID(ctx context.Context, userID int64) ([]*model.Video, error) {
	var videos []*model.Video
	if err := DB.WithContext(ctx).Where("user_id = ?", userID).Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

// GetLikeCount 返回视频点赞数
func GetLikeCount(ctx context.Context, videoID int64) (int64, error) {
	var video model.Video
	if err := DB.WithContext(ctx).Where("ID = ?", videoID).First(&video).Error; err != nil {
		return 0, err
	}
	return video.FavoriteCount, nil
}

// GetCommentCount 返回视频评论数
func GetCommentCount(ctx context.Context, videoID int64) (int64, error) {
	var video model.Video
	if err := DB.WithContext(ctx).Where("ID = ?", videoID).First(&video).Error; err != nil {
		return 0, err
	}
	return video.CommentCount, nil
}

// IsFavorite 返回是否点赞
func IsFavorite(ctx context.Context, videoID int64, userID int64) (bool, error) {
	var favorites []*model.Favorite
	result := DB.WithContext(ctx).Where("user_id = ? AND video_id = ?", userID, videoID).Find(&favorites)
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}

// MGetVideoByTime 根据时间戳返回最近count个视频,还需要返回next time
func MGetVideoByTime(ctx context.Context, latestTime time.Time, count int64) ([]*model.Video, int64, error) {
	var videos []*model.Video
	if err := DB.WithContext(ctx).Where("created_at < ?", latestTime).Limit(int(count)).Order("created_at DESC").Find(&videos).Error; err != nil {
		return nil, 0, err
	}
	var nextTime int64 = math.MaxInt32
	if len(videos) != 0 { // 查到了新视频
		nextTime = videos[0].CreatedAt.Unix()
	}
	return videos, nextTime, nil
}

// LikeVideo 点赞视频
func LikeVideo(ctx context.Context, userID int64, videoID int64) error {
	isFavorite, err := IsFavorite(ctx, videoID, userID)
	if err != nil {
		return err
	}
	if isFavorite == true {
		return nil
	}
	favorite := &model.Favorite{
		UserId:  userID,
		VideoId: videoID,
	}
	if err := DB.WithContext(ctx).Create(&favorite).Error; err != nil {
		return err
	}
	var video model.Video
	if err := DB.WithContext(ctx).Where("ID = ?", videoID).First(&video).Error; err != nil {
		return err
	}
	video.FavoriteCount++
	DB.WithContext(ctx).Save(&video)
	return nil
}

// UnLikeVideo 取消点赞视频
func UnLikeVideo(ctx context.Context, userID int64, videoID int64) error {
	isFavorite, err := IsFavorite(ctx, videoID, userID)
	if err != nil {
		return err
	}
	if isFavorite == false {
		return nil
	}
	err = DB.WithContext(ctx).Where("user_id = ? AND video_id = ?", userID, videoID).Delete(&model.Favorite{}).Error
	if err != nil {
		return err
	}
	var video model.Video
	if err := DB.WithContext(ctx).Where("ID = ?", videoID).First(&video).Error; err != nil {
		return err
	}
	video.FavoriteCount--
	DB.WithContext(ctx).Save(&video)
	return nil
}

// MGetLikeList 通过用户ID获取用户点赞的视频ID数组
func MGetLikeList(ctx context.Context, userID int64) ([]int64, error) {
	var favorites []*model.Favorite
	if err := DB.WithContext(ctx).Where("user_id = ?", userID).Find(&favorites).Error; err != nil {
		return nil, err
	}
	var likeList []int64
	for _, favorite := range favorites {
		likeList = append(likeList, favorite.VideoId)
	}
	return likeList, nil
}

// MGetVideoInfo 通过视频ID查询得到model.Video信息
func MGetVideoInfo(ctx context.Context, videoID int64) (*model.Video, error) {
	var videoInfo *model.Video
	DB.WithContext(ctx).Where("ID = ?", videoID).First(&videoInfo)
	return videoInfo, nil
}
