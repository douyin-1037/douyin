package bizdto

type Message struct {
	ID         int64  `json:"id"`          // 消息id
	Content    string `json:"content"`     // 消息内容
	CreateTime string `json:"create_time"` // 消息发送时间，格式 yyyy-MM-dd HH:MM:ss
}
