package service

import (
	"context"
	config "douyin/common/conf"
	"douyin/user/infra/dal"
	"douyin/user/infra/redis"
	"fmt"
	"testing"
)

func Init() {
	config.InitConfig()
	dal.Init()
	redis.Init()
}
func TestGetFollowList(t *testing.T) {
	Init()
	ctx := context.Background()

	follows, err := NewGetFollowListService(ctx).GetFollowList(40, 39)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(follows)
}
