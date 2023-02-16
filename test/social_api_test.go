package test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestRelation(t *testing.T) {
	e := newExpect(t)

	//用户A、用户B登录
	userIdA, tokenA := getTestUserToken(testUserA, e)
	userIdB, tokenB := getTestUserToken(testUserB, e)

	//用户A关注用户B
	relationResp := e.POST("/douyin/relation/action/").
		WithQuery("token", tokenA).WithQuery("to_user_id", userIdB).WithQuery("action_type", 1).
		WithFormField("token", tokenA).WithFormField("to_user_id", userIdB).WithFormField("action_type", 1).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	relationResp.Value("status_code").Number().IsEqual(0)

	//获取用户A关注列表
	followListResp := e.GET("/douyin/relation/follow/list/").
		WithQuery("token", tokenA).WithQuery("user_id", userIdA).
		WithFormField("token", tokenA).WithFormField("user_id", userIdA).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	followListResp.Value("status_code").Number().IsEqual(0)

	//用户A关注列表应当包含用户B
	containTestUserB := false
	for _, element := range followListResp.Value("user_list").Array().Iter() {
		user := element.Object()
		user.ContainsKey("id")
		if int(user.Value("id").Number().Raw()) == userIdB {
			containTestUserB = true
		}
	}
	assert.True(t, containTestUserB, "Follow test user failed")

	//获取用户B粉丝列表
	followerListResp := e.GET("/douyin/relation/follower/list/").
		WithQuery("token", tokenB).WithQuery("user_id", userIdB).
		WithFormField("token", tokenB).WithFormField("user_id", userIdB).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	followerListResp.Value("status_code").Number().IsEqual(0)

	//用户B粉丝列表应当包含用户A
	containTestUserA := false
	for _, element := range followerListResp.Value("user_list").Array().Iter() {
		user := element.Object()
		user.ContainsKey("id")
		if int(user.Value("id").Number().Raw()) == userIdA {
			containTestUserA = true
		}
	}
	assert.True(t, containTestUserA, "Follower test user failed")

	//用户A取消关注用户B
	relationRespAunfollowB := e.POST("/douyin/relation/action/").
		WithQuery("token", tokenA).WithQuery("to_user_id", userIdB).WithQuery("action_type", 2).
		WithFormField("token", tokenA).WithFormField("to_user_id", userIdB).WithFormField("action_type", 2).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	relationRespAunfollowB.Value("status_code").Number().IsEqual(0)

	//获取用户A关注列表
	followListResp = e.GET("/douyin/relation/follow/list/").
		WithQuery("token", tokenA).WithQuery("user_id", userIdA).
		WithFormField("token", tokenA).WithFormField("user_id", userIdA).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	followListResp.Value("status_code").Number().IsEqual(0)

	//用户A关注列表应当不包含用户B
	containTestUserB = false
	for _, element := range followListResp.Value("user_list").Array().Iter() {
		user := element.Object()
		user.ContainsKey("id")
		if int(user.Value("id").Number().Raw()) == userIdB {
			containTestUserB = true
		}
	}
	assert.False(t, containTestUserB, "Follow test user failed")

	//获取用户B粉丝列表
	followerListResp = e.GET("/douyin/relation/follower/list/").
		WithQuery("token", tokenB).WithQuery("user_id", userIdB).
		WithFormField("token", tokenB).WithFormField("user_id", userIdB).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	followerListResp.Value("status_code").Number().IsEqual(0)

	//用户B粉丝列表应当不包含用户A
	containTestUserA = false
	for _, element := range followerListResp.Value("user_list").Array().Iter() {
		user := element.Object()
		user.ContainsKey("id")
		if int(user.Value("id").Number().Raw()) == userIdA {
			containTestUserA = true
		}
	}
	assert.False(t, containTestUserA, "Follower test user failed")
}

func TestFriend(t *testing.T) {
	e := newExpect(t)

	//用户A、用户B登录
	userIdA, tokenA := getTestUserToken(testUserA, e)
	userIdB, tokenB := getTestUserToken(testUserB, e)

	//用户A取消关注用户B
	relationRespAunfollowB := e.POST("/douyin/relation/action/").
		WithQuery("token", tokenA).WithQuery("to_user_id", userIdB).WithQuery("action_type", 2).
		WithFormField("token", tokenA).WithFormField("to_user_id", userIdB).WithFormField("action_type", 2).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	relationRespAunfollowB.Value("status_code").Number().IsEqual(0)

	//用户B取消关注用户A
	relationRespBunfollowA := e.POST("/douyin/relation/action/").
		WithQuery("token", tokenB).WithQuery("to_user_id", userIdA).WithQuery("action_type", 2).
		WithFormField("token", tokenB).WithFormField("to_user_id", userIdA).WithFormField("action_type", 2).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	relationRespBunfollowA.Value("status_code").Number().IsEqual(0)

	//用户A关注用户B
	relationRespAB := e.POST("/douyin/relation/action/").
		WithQuery("token", tokenA).WithQuery("to_user_id", userIdB).WithQuery("action_type", 1).
		WithFormField("token", tokenA).WithFormField("to_user_id", userIdB).WithFormField("action_type", 1).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	relationRespAB.Value("status_code").Number().IsEqual(0)

	//获取用户A好友列表
	friendListRespA1 := e.GET("/douyin/relation/friend/list/").
		WithQuery("token", tokenA).WithQuery("user_id", userIdA).
		WithFormField("token", tokenA).WithFormField("user_id", userIdA).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	friendListRespA1.Value("status_code").Number().IsEqual(0)

	//用户A好友列表应当不包含用户B
	containTestUserB1 := false
	for _, element := range friendListRespA1.Value("user_list").Array().Iter() {
		user := element.Object()
		user.ContainsKey("id")
		if int(user.Value("id").Number().Raw()) == userIdB {
			containTestUserB1 = true
		}
	}
	assert.False(t, containTestUserB1, "Follow test user failed")

	//获取用户B好友列表
	friendListRespB1 := e.GET("/douyin/relation/friend/list/").
		WithQuery("token", tokenB).WithQuery("user_id", userIdB).
		WithFormField("token", tokenB).WithFormField("user_id", userIdB).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	friendListRespB1.Value("status_code").Number().IsEqual(0)

	//用户B好友列表应当不包含用户A
	containTestUserA1 := false
	for _, element := range friendListRespB1.Value("user_list").Array().Iter() {
		user := element.Object()
		user.ContainsKey("id")
		if int(user.Value("id").Number().Raw()) == userIdA {
			containTestUserA1 = true
		}
	}
	assert.False(t, containTestUserA1, "Follow test user failed")

	//用户B关注用户A
	relationRespBA := e.POST("/douyin/relation/action/").
		WithQuery("token", tokenB).WithQuery("to_user_id", userIdA).WithQuery("action_type", 1).
		WithFormField("token", tokenB).WithFormField("to_user_id", userIdA).WithFormField("action_type", 1).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	relationRespBA.Value("status_code").Number().IsEqual(0)

	//获取用户A好友列表
	friendListRespA2 := e.GET("/douyin/relation/friend/list/").
		WithQuery("token", tokenA).WithQuery("user_id", userIdA).
		WithFormField("token", tokenA).WithFormField("user_id", userIdA).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	friendListRespA2.Value("status_code").Number().IsEqual(0)

	//用户A好友列表应当包含用户B
	containTestUserB2 := false
	for _, element := range friendListRespA2.Value("user_list").Array().Iter() {
		user := element.Object()
		user.ContainsKey("id")
		if int(user.Value("id").Number().Raw()) == userIdB {
			containTestUserB2 = true
		}
	}
	assert.True(t, containTestUserB2, "Follow test user failed")

	//获取用户B好友列表
	friendListRespB2 := e.GET("/douyin/relation/friend/list/").
		WithQuery("token", tokenB).WithQuery("user_id", userIdB).
		WithFormField("token", tokenB).WithFormField("user_id", userIdB).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	friendListRespB2.Value("status_code").Number().IsEqual(0)

	//用户B好友列表应当包含用户A
	containTestUserA2 := false
	for _, element := range friendListRespB2.Value("user_list").Array().Iter() {
		user := element.Object()
		user.ContainsKey("id")
		if int(user.Value("id").Number().Raw()) == userIdA {
			containTestUserA2 = true
		}
	}
	assert.True(t, containTestUserA2, "Follow test user failed")
}

func TestChat(t *testing.T) {
	e := newExpect(t)

	userIdA, tokenA := getTestUserToken(testUserA, e)
	userIdB, tokenB := getTestUserToken(testUserB, e)

	messageResp := e.POST("/douyin/message/action/").
		WithQuery("token", tokenA).WithQuery("to_user_id", userIdB).WithQuery("action_type", 1).WithQuery("content", "Send to UserB").
		WithFormField("token", tokenA).WithFormField("to_user_id", userIdB).WithFormField("action_type", 1).WithQuery("content", "Send to UserB").
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	messageResp.Value("status_code").Number().IsEqual(0)

	chatResp := e.GET("/douyin/message/chat/").
		WithQuery("token", tokenA).WithQuery("to_user_id", userIdB).
		WithFormField("token", tokenA).WithFormField("to_user_id", userIdB).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	chatResp.Value("status_code").Number().IsEqual(0)
	chatResp.Value("message_list").Array().Length().Gt(0)

	chatResp = e.GET("/douyin/message/chat/").
		WithQuery("token", tokenB).WithQuery("to_user_id", userIdA).
		WithFormField("token", tokenB).WithFormField("to_user_id", userIdA).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	chatResp.Value("status_code").Number().IsEqual(0)
	chatResp.Value("message_list").Array().Length().Gt(0)
}
