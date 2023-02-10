package redis

import (
	"douyin/common/conf"
	"douyin/common/util"
	"douyin/video/infra/dal/model"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"testing"
)

func testInit() {
	conf.InitConfig()
	expireTimeUtil = util.ExpireTimeUtil{
		ExpireTime:     conf.Redis.ExpireTime,
		MaxRandAddTime: conf.Redis.MaxRandAddTime,
	}
	redisPool = &redis.Pool{
		MaxIdle:   conf.Redis.MaxIdle,
		MaxActive: conf.Redis.MaxActive,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", conf.Redis.Address)
		},
	}
}

func TestGetLikeCountById(t *testing.T) {
	testInit()
	result, err := GetLikeCountById(2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func TestAddPublishList(t *testing.T) {
	testInit()
	var videoListp []*model.Video
	v1 := model.Video{
		UserId:        5,
		Title:         "tst",
		PlayUrl:       "123",
		CoverUrl:      "123",
		FavoriteCount: 3,
		CommentCount:  2,
	}
	v1.ID = 9

	v2 := model.Video{
		UserId:        5,
		Title:         "tst",
		PlayUrl:       "123",
		CoverUrl:      "123",
		FavoriteCount: 6,
		CommentCount:  7,
	}
	v2.ID = 10
	videoListp = append(videoListp, &v1)
	videoListp = append(videoListp, &v2)
	err := AddPublishList(videoListp, 5)
	if err != nil {
		fmt.Println(err)
	}
}

func TestGetPublishList(t *testing.T) {
	testInit()
	result, err := GetPublishList(5)
	if err != nil {
		fmt.Println(err)
	}
	videoList := *result
	fmt.Println(videoList)
}
