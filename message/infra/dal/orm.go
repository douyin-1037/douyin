package dal

import (
	"context"
	"douyin/common/util"
	"douyin/message/infra/dal/model"

	"github.com/cloudwego/kitex/pkg/klog"
)

func CreateMessage(ctx context.Context, userID int64, toUserID int64, content string, createTime int64) error {
	uuid, err := util.GenSnowFlake(0)
	if err != nil {
		klog.Error("Failed to generate UUID" + err.Error())
		return err
	}

	message := model.Message{
		FromUserId:  userID,
		ToUserId:    toUserID,
		Contents:    content,
		MessageUUId: int64(uuid),
		CreateTime:  createTime,
	}
	err = DB.WithContext(ctx).Create(&message).Error
	if err != nil {
		klog.Error("create message fail " + err.Error())
		return err
	}
	return nil
}

func GetMessageList(ctx context.Context, userID int64, toUserID int64) ([]*model.Message, error) {
	var messages []*model.Message
	err := DB.WithContext(ctx).Where("from_user_id = ? AND to_user_id = ?", userID, toUserID).Or(
		"from_user_id = ? AND to_user_id = ?", toUserID, userID,
	).Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}
