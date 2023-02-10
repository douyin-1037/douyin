package dal

// @path: comment/infra/dal/orm.go
// @description: DAL of comment
// @auth: wan-nan <wan_nan@foxmail.com>
import (
	"context"
	"douyin/code_gen/kitex_gen/commentproto"
	"douyin/comment/infra/dal/model"
	"douyin/comment/pack"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

// CreateComment
// create a comment by (userID, videoId, content)
// userID is the ID of the *author* of this comment
func CreateComment(ctx context.Context, userID int64, videoId int64, content string) (*commentproto.CommentInfo, error) {
	comment := model.Comment{
		UserId:   userID,
		VideoId:  videoId,
		Contents: content,
	}

	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Table("comment").Create(&comment).Error // 通过数据的指针来创建，所以要用&comment
		if err != nil {
			klog.Error("create comment fail " + err.Error())
			return err
		}
		err = tx.Table("video").Where("id = ?", comment.VideoId).Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error
		if err != nil {
			klog.Error("AddCommentCount error " + err.Error())
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return pack.Comment(&comment), nil
}

// DeleteComment
// delete a comment by commentID
func DeleteComment(ctx context.Context, commentID int64) error {
	//err := DB.WithContext(ctx).Delete(&model.Comment{}, commentID).Error
	//if err != nil {
	//	klog.Error("delete comment fail: " + err.Error())
	//	return err
	//}
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Delete(&model.Comment{}, commentID).Error
		if err != nil {
			klog.Error("delete comment fail: " + err.Error())
			return err
		}
		err = tx.Table("video").Where("id = ?", 0).Update("comment_count", gorm.Expr("comment_count - ?", 1)).Error
		if err != nil {
			klog.Error("SubCommentCount error " + err.Error())
			return err
		}
		return nil
	})
	return err
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
