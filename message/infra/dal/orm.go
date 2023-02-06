package dal

import (
	"context"
	"douyin/message/infra/dal/model"
	"github.com/cloudwego/kitex/pkg/klog"
)

func CreateMessage(ctx context.Context, userID int64, toUserID int64, content string) error {
	message := model.Message{
		FromUserId: userID,
		ToUserId:   toUserID,
		Contents:   content,
	}
	err := DB.WithContext(ctx).Create(&message).Error
	if err != nil {
		klog.Error("create message fail " + err.Error())
		return err
	}
	return nil
}

func GetMessageList(ctx context.Context, userID int64, toUserID int64) ([]*model.Message, error) {
	var messages []*model.Message
	err := DB.WithContext(ctx).Where("from_user_id = ? AND to_user_id = ?", userID, toUserID).Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}
