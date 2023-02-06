package dal

// @path: comment/infra/dal/orm.go
// @description: DAL of comment
// @auth: wan-nan <wan_nan@foxmail.com>
import (
	"context"
	"douyin/comment/infra/dal/model"
	"github.com/cloudwego/kitex/pkg/klog"
)

// CreateComment
// create a comment by (userID, videoId, content)
func CreateComment(ctx context.Context, userID int64, videoId int64, content string) error {
	comment := model.Comment{
		UserId:   userID,
		VideoId:  videoId,
		Contents: content,
	}
	err := DB.WithContext(ctx).Create(&comment).Error
	if err != nil {
		klog.Error("create comment fail: " + err.Error())
		return err
	}
	return nil
}

// DeleteComment
// delete a comment by commentID
func DeleteComment(ctx context.Context, commentID int64) error {
	err := DB.WithContext(ctx).Delete(&model.Comment{}, commentID).Error
	if err != nil {
		klog.Error("delete comment fail: " + err.Error())
		return err
	}
	return nil
}

// GetCommentList
// get comment list by videoID
func GetCommentList(ctx context.Context, videoID int64) ([]*model.Comment, error) {
	var comments []*model.Comment
	err := DB.WithContext(ctx).Where("video_id = ?", videoID).Find(&comments).Error
	if err != nil {
		klog.Error("get comment list fail: " + err.Error())
		return nil, err
	}
	return comments, nil
}
