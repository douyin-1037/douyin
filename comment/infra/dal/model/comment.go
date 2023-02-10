package model

import "gorm.io/gorm"

// Comment Gorm Data Structures
type Comment struct {
	gorm.Model
	UserId   int64  `gorm:"column:user_id;not null;index:fk_user_comment"`   //comment 作者的id
	VideoId  int64  `gorm:"column:video_id;not null;index:fk_video_comment"` //comment 所在视频的id
	Contents string `grom:"column:contents;type:varchar(255);not null"`      //comment 的内容
}

func (Comment) TableName() string {
	return "comment"
}
