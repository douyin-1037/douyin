package pulsar

import (
	"context"
	"douyin/common/constant"
	"github.com/apache/pulsar-client-go/pulsar"
)

func CreateCommentProduce(ctx context.Context, userId int64, videoId int64, content string, commentUUID int64, createTime int64) error {
	_, err := p_comment.Send(ctx, &pulsar.ProducerMessage{
		Value: &CommentJSON{
			UserId:      userId,
			VideoId:     videoId,
			Content:     content,
			CommentUUID: commentUUID,
			CreateTime:  createTime,
			ActionType:  constant.CreateComment,
		},
	})
	return err
}

func DeleteCommentProduce(ctx context.Context, commentUUID int64, videoId int64) error {
	_, err := p_comment.Send(ctx, &pulsar.ProducerMessage{
		Value: &CommentJSON{
			VideoId:     videoId,
			CommentUUID: commentUUID,
			ActionType:  constant.DeleteComment,
		},
	})
	return err
}
