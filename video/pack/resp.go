package pack

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
