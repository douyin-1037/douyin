package coredto

import (
	errno "douyin/common/code"
	"douyin/types"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Success = types.BaseResp{
	Code: 0,
	Msg:  "success",
}

func Error(c *gin.Context, rawErr error) {
	err := errno.ConvertErr(rawErr)
	s, _ := json.Marshal(err)
	fmt.Println(string(s))
	c.JSON(err.StatusCode(), err)
}

func Send(c *gin.Context, resp interface{}) {
	s, _ := json.Marshal(resp)
	fmt.Println(string(s))
	c.JSON(http.StatusOK, resp)
}

func OK(c *gin.Context) {
	fmt.Println("resp success")
	c.JSON(http.StatusOK, Success)
}
