package handlers

import (
	"douyin/api/auth"
	. "douyin/application/mock"
	"douyin/common/conf"
	"douyin/types/bizdto"
	"douyin/types/coredto"
	"encoding/json"
	"fmt"
	. "github.com/bytedance/mockey"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestUser(t *testing.T) {
	conf.InitConfig()
	auth.Init()
	r := gin.New()
	authGroup := r.Group("/douyin")
	authGroup.Use(auth.JWT())
	{
		authGroup.GET("/user/", GetUserInfo)
	}
	ctrl := gomock.NewController(t)
	m := NewMockUserAppService(ctrl)
	//向注册的路有发起请求
	req, _ := http.NewRequest("GET", "/douyin/user/", nil)
	params := make(url.Values)
	params.Add("user_id", "1")
	params.Add("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiaWF0IjoxNjc0OTA1MTQ4LCJleHAiOjE2Nzc1MTM1OTksImF1ZCI6IiIsImlzcyI6IiIsInN1YiI6IiJ9.mRcRqdwU62uEacqanqTZNl5pZ4B0ebqoknpz7mfJ7eI")
	req.URL.RawQuery = params.Encode()
	w := httptest.NewRecorder()
	user_test := &bizdto.User{
		ID:            1,
		Name:          "yui",
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	}

	m.EXPECT().GetUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(user_test, nil)
	UserService = m

	PatchConvey("TestUser", t, func() {

		r.ServeHTTP(w, req)
		result := w.Result()
		defer result.Body.Close()

		// 读取响应body
		body, _ := io.ReadAll(result.Body)
		fmt.Printf("response:%v\n", string(body))
		// 解析响应，判断响应是否与预期一致
		response := &bizdto.UserQueryResp{}
		if err := json.Unmarshal(body, response); err != nil {
			t.Errorf("解析响应出错，err:%v\n", err)
		}

		So(response.User, ShouldResemble, user_test)
		So(response.BaseResp, ShouldResemble, coredto.Success)

	})
	//模拟http服务处理请求

}
