package model

import (
	"gorm.io/gorm"
)

// Favorite Gorm Data Structures
type Favorite struct {
	gorm.Model
	UserId  int64 `gorm:"column:user_id;not null;index:fk_user_favorite"`
	VideoId int64 `gorm:"column:video_id;not null;index:fk_video_favorite"`
}

func (f *Favorite) TableName() string {
	return "favorite"
}
