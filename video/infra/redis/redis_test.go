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
func TestDeletePublishList(t *testing.T) {
	testInit()
	err := DelPublishList(1234)
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
	videoList := result
	for _, v := range videoList {
		fmt.Println(v)
	}
}

func TestDelPublishList(t *testing.T) {
	testInit()
	err := DelPublishList(5)
	if err != nil {
		fmt.Println(err)
	}
}

func TestAddLikeList(t *testing.T) {
	testInit()
	var likeList = []int64{11, 20}
	err := AddLikeList(1, likeList)
	if err != nil {
		fmt.Println(err)
	}
}

func TestGetLikeList(t *testing.T) {
	testInit()
	likeList, err := GetLikeList(3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(likeList)
}

func TestGetIsLikeById(t *testing.T) {
	testInit()
	isLikeById, err := GetIsLikeById(2, 13)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(isLikeById)
}

func TestAddVideoInfo(t *testing.T) {
	testInit()
	v1 := model.Video{
		UserId:        5,
		Title:         "tst",
		PlayUrl:       "123",
		CoverUrl:      "123",
		FavoriteCount: 3,
		CommentCount:  2,
	}
	v1.ID = 9
	err := AddVideoInfo(v1)
	if err != nil {
		fmt.Println(err)
	}
}

func TestGetVideoInfo(t *testing.T) {
	testInit()
	videoInfo, err := GetVideoInfo(9)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(videoInfo)
}
