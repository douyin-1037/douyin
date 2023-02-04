package pack

import (
	"douyin/code_gen/kitex_gen/userproto"
	"douyin/pkg/code"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *userproto.BaseResp {
	return baseResp(code.ConvertErr(err))
}

func baseResp(err code.ErrNo) *userproto.BaseResp {
	return &userproto.BaseResp{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
