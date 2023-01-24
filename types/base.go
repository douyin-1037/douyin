package types

type BaseResp struct {
	Code int64  `json:"status_code"`
	Msg  string `json:"status_msg"`
}
