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
	c.JSON(http.StatusOK, BaseResp{
		Code: int64(code.HTTPCoder(statuserr.Code(err))),
		Msg:  err.Error(),
	})
}

func ErrorBaseResp(err error) BaseResp {
	return BaseResp{
		Code: int64(code.HTTPCoder(statuserr.Code(err))),
		Msg:  err.Error(),
	}
}

func Send(c *gin.Context, resp interface{}) {
	c.JSON(http.StatusOK, resp)
}

func OK(c *gin.Context) {
	c.JSON(http.StatusOK, Success)
}
