package coredto

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"douyin/common/code"
	"douyin/pkg/statuserr"
)

type BaseResp struct {
	Code int64  `json:"status_code"`
	Msg  string `json:"status_msg"`
}

var Success = BaseResp{
	Code: 0,
	Msg:  "success",
}

func Error(c *gin.Context, err error) {
	c.AbortWithStatusJSON(code.HTTPCoder(statuserr.Code(err)), err)
}

func Send(c *gin.Context, resp interface{}) {
	c.JSON(http.StatusOK, resp)
}

func OK(c *gin.Context) {
	c.JSON(http.StatusOK, Success)
}
