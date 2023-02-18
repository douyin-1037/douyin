package pulsar

import (
	"context"
	"douyin/common/constant"
	"github.com/apache/pulsar-client-go/pulsar"
)

func FollowUserProduce(ctx context.Context, userId int64, followId int64) error {
	_, err := p_follow_user.Send(ctx, &pulsar.ProducerMessage{
		Value: &FollowUserJSON{
			UserID:     userId,
			FollowID:   followId,
			ActionType: constant.FollowUser,
		},
	})
	return err
}

func UnFollowUserProduce(ctx context.Context, userId int64, followId int64) error {
	_, err := p_follow_user.Send(ctx, &pulsar.ProducerMessage{
		Value: &FollowUserJSON{
			UserID:     userId,
			FollowID:   followId,
			ActionType: constant.UnFollowUser,
		},
	})
	return err
}
