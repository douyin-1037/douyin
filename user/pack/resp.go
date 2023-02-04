package pack

import (
	"douyin/code_gen/kitex_gen/userproto"
	"douyin/pkg/errno"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *userproto.BaseResp {
	return baseResp(errno.ConvertErr(err))
}

func baseResp(err errno.ErrNo) *userproto.BaseResp {
	return &userproto.BaseResp{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
