package redis

import (
	"douyin/common/conf"
	"douyin/common/util"
	"douyin/message/infra/dal/model"
	redisModel "douyin/message/infra/redis/model"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"testing"
	"time"
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

func TestAddMessage(t *testing.T) {
	testInit()
	m3 := redisModel.MessageRedis{
		MessageId:  3,
		FromUserId: 4,
		ToUserId:   3,
		Content:    "m3:4 to 3",
		CreateTime: time.Now().Unix() + 10,
	}

	err := AddMessage(4, 3, m3)
	if err != nil {
		fmt.Println(err)
	}
}

func TestAddMessageList(t *testing.T) {
	testInit()
	var messageListp []*model.Message
	m1 := model.Message{
		MessageUUId: 1,
		FromUserId:  3,
		ToUserId:    4,
		Contents:    "m1:3 to 4",
		CreateTime:  time.Now().Unix(),
	}

	m2 := model.Message{
		MessageUUId: 2,
		FromUserId:  4,
		ToUserId:    3,
		Contents:    "m2:4 to 3",
		CreateTime:  time.Now().Unix() + 10,
	}
	messageListp = append(messageListp, &m1)
	messageListp = append(messageListp, &m2)
	err := AddMessageList(3, 4, messageListp)
	if err != nil {
		fmt.Println(err)
	}
}

func TestGetMessageList(t *testing.T) {
	testInit()
	nowTime := time.Now().Unix()
	result, err := GetMessageList(38, 40, 0, nowTime)
	if err != nil {
		fmt.Println(err)
	}
	messageList := result
	fmt.Printf("%v\n", messageList)
}

func TestAddMessageLatestTime(t *testing.T) {
	testInit()
	err := AddMessageLatestTime(1, 2, 10)
	fmt.Println(err)
}

func TestGetMessageLatestTime(t *testing.T) {
	testInit()
	result, err := GetMessageLatestTime(1, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
