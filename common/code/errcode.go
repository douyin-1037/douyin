package code

import (
	"net/http"

	"douyin/pkg/statuserr"
)

const (
	SuccessCode             = 0
	ServiceErrCode          = 10001
	ParamErrCode            = 10002
	LoginErrCode            = 10003
	UserNotExistErrCode     = 10004
	UserAlreadyExistErrCode = 10005
	UnauthorizedErrCode     = 10006
)

var (
	Success             = statuserr.New(SuccessCode, "Success")
	ServiceErr          = statuserr.New(ServiceErrCode, "Service is unable to start successfully")
	ParamErr            = statuserr.New(ParamErrCode, "Wrong Parameter has been given")
	LoginErr            = statuserr.New(LoginErrCode, "Wrong username or password")
	UserNotExistErr     = statuserr.New(UserNotExistErrCode, "User does not exists")
	UserAlreadyExistErr = statuserr.New(UserAlreadyExistErrCode, "User already exists")
	UnauthorizedErr     = statuserr.New(UnauthorizedErrCode, "JWT Token Unauthorized")
)

var mapper = map[int64]int{
	SuccessCode:             http.StatusOK,
	ServiceErrCode:          http.StatusInternalServerError,
	ParamErrCode:            http.StatusBadRequest,
	LoginErrCode:            http.StatusBadRequest,
	UserNotExistErrCode:     http.StatusBadRequest,
	UserAlreadyExistErrCode: http.StatusBadRequest,
	UnauthorizedErrCode:     http.StatusUnauthorized,
}

func HTTPCoder(code int64) int {
	if http, ok := mapper[code]; ok {
		return http
	}
	return http.StatusInternalServerError
}
