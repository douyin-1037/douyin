package model

type MessageRedis struct {
	MessageId  int64
	FromUserId int64
	ToUserId   int64
	Content    string
	CreateTime int64
}
