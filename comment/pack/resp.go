package pack

// @path: comment/pack/resp.go
// @description: build base response from error
// @auth: wan-nan <wan_nan@foxmail.com>
import (
	"douyin/code_gen/kitex_gen/commentproto"
	"douyin/pkg/code"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *commentproto.BaseResp {
	return baseResp(code.ConvertErr(err))
}

func baseResp(err code.ErrNo) *commentproto.BaseResp {
	return &commentproto.BaseResp{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
