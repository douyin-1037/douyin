package pack

import (
	"douyin/code_gen/kitex_gen/messageproto"
	"douyin/pkg/code"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *messageproto.BaseResp {
	return baseResp(code.ConvertErr(err))
}

func baseResp(err code.ErrNo) *messageproto.BaseResp {
	return &messageproto.BaseResp{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
