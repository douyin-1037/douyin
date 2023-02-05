package model

import (
	"gorm.io/gorm"
)

// Video Gorm Data Structures
type Video struct {
	gorm.Model
	UserId        int64  `gorm:"column:user_id;not null;index:fk_user_video"`
	Title         string `gorm:"column:title;type:varchar(128);not null"`
	PlayUrl       string `gorm:"column:play_url;varchar(128);not null"`
	CoverUrl      string `gorm:"column:cover_url;varchar(128);not null"`
	FavoriteCount int64  `gorm:"column:favorite_count;default:0"`
	CommentCount  int64  `gorm:"column:comment_count;default:0"`
}

func (v *Video) TableName() string {
	return "video"
}
