package model

type VideoRedis struct {
	VideoId  uint
	UserId   int64
	Title    string
	PlayUrl  string
	CoverUrl string
}
