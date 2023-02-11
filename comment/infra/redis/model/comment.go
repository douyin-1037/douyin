package model

type CommentRedis struct {
	CommentId  int64
	VideoId    int64
	UserId     int64
	Content    string
	CreateTime int64
}
