package statuserr

import (
	"fmt"

	"github.com/pkg/errors"
)

const defaultCode = 50000

type StatusErr struct {
	Code    int64  `json:"status_code"`
	Message string `json:"status_msg"`
}

func (e StatusErr) Error() string {
	return fmt.Sprintf("err_code=%v, err_msg=%v", e.Code, e.Message)
}

func New(code int64, msg string) error {
	return StatusErr{code, msg}
}

func Newf(code int64, format string, args ...interface{}) error {
	return StatusErr{code, fmt.Sprintf(format, args...)}
}

func Code(err error) int64 {
	if err == nil {
		return 0
	}
	e := new(StatusErr)
	if errors.As(err, e) {
		return e.Code
	}
	return defaultCode
}
