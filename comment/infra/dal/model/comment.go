package model

import "gorm.io/gorm"

// Comment Gorm Data Structures
type Comment struct {
	gorm.Model
	UserId   int64  `gorm:"column:user_id;not null;index:fk_user_comment"`
	VideoId  int64  `gorm:"column:video_id;not null;index:fk_video_comment"`
	Contents string `grom:"column:contents;type:varchar(255);not null"`
}

func (Comment) TableName() string {
	return "comment"
}
