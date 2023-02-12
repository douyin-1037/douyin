package pack

// @path: video/pack/resp.go
// @description: build base response from error
// @author: Chongzhi <dczdcz2001@aliyun.com>
import (
	"douyin/code_gen/kitex_gen/videoproto"
	"douyin/pkg/code"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *videoproto.BaseResp {
	return baseResp(code.ConvertErr(err))
}

func baseResp(err code.ErrNo) *videoproto.BaseResp {
	return &videoproto.BaseResp{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
