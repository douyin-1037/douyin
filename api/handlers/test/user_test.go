package test

import (
	"douyin/api/auth"
	appImpl "douyin/application/impl"
	"douyin/cmd/inject"
	"douyin/cmd/router"
	"douyin/common/conf"
	"douyin/types/bizdto"
	"douyin/types/coredto"
	"encoding/json"
	"fmt"
	. "github.com/bytedance/mockey"
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
	inject.Inject()
	r := router.NewRouter()
	//向注册的路有发起请求
	req, _ := http.NewRequest("GET", "/douyin/user/", nil)
	params := make(url.Values)
	params.Add("user_id", "1")
	params.Add("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiaWF0IjoxNjc0NzQ4NzYxLCJleHAiOjE2NzQ4MzUxOTksImF1ZCI6IiIsImlzcyI6IiIsInN1YiI6IiJ9.f9Kd1EFGCyk05Y4lzpmWIJtOgbDYLFVm1x0kiAvEu2k")
	req.URL.RawQuery = params.Encode()
	w := httptest.NewRecorder()
	user_test := &bizdto.User{
		ID:            1,
		Name:          "yui",
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	}
	PatchConvey("TestUser", t, func() {
		fmt.Println("is nil?")
		Mock(appImpl.UserAppServiceImpl.GetUser).Return(user_test, nil).Build() // mock方法
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
