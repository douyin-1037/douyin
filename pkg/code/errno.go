package code

import (
	"errors"
	"fmt"
	"net/http"
)

const (
	SuccessCode               = 0
	ServiceErrCode            = 10001
	ParamErrCode              = 10002
	LoginErrCode              = 10003
	UserNotExistErrCode       = 10004
	UserAlreadyExistErrCode   = 10005
	UnauthorizedErrCode       = 10006
	LoginFailedTooManyErrCode = 10007
	UsernameCheckErrCode      = 10008
	PasswordCheckErrCode      = 10009
)

type ErrNo struct {
	ErrCode int64  `json:"status_code"`
	ErrMsg  string `json:"status_msg"`
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{code, msg}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success             = NewErrNo(SuccessCode, "Success")
	ServiceErr          = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	ParamErr            = NewErrNo(ParamErrCode, "Wrong Parameter has been given")
	LoginErr            = NewErrNo(LoginErrCode, "Wrong username or password")
	UserNotExistErr     = NewErrNo(UserNotExistErrCode, "User does not exists")
	UserAlreadyExistErr = NewErrNo(UserAlreadyExistErrCode, "User already exists")
	UnauthorizedErr     = NewErrNo(UnauthorizedErrCode, "JWT Token Unauthorized")
	UsernameCheckErr    = NewErrNo(UsernameCheckErrCode, "用户名应当使用字母,数字,下划线,减号,且长度4-32位")
	PasswordCheckErr    = NewErrNo(PasswordCheckErrCode, "密码应当包含大写字母,小写字母,数字,且长度5-32位")
)

// ConvertErr convert error to ErrNo
func ConvertErr(err error) ErrNo {
	if err == nil {
		return Success
	}
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}

func (e ErrNo) StatusCode() int {
	return http.StatusOK
}
