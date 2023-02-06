package model

import "gorm.io/gorm"

// Message Gorm Data Structures
type Message struct {
	gorm.Model
	FromUserId int64  `gorm:"column:from_user_id;not null;index:fk_user_message_from"`
	ToUserId   int64  `gorm:"column:to_user_id;not null;index:fk_user_message_to"`
	Contents   string `grom:"column:contents;type:varchar(255);not null"`
}

func (Message) TableName() string {
	return "message"
}
