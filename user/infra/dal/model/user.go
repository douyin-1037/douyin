package model

import (
	"gorm.io/gorm"
)

// User Gorm Data structures
type User struct {
	gorm.Model
	Name          string `gorm:"column:name;index:uni_name,unique;type:varchar(32);not null"`
	Password      string `gorm:"column:password;type:varchar(255);not null"`
	FollowCount   int64  `gorm:"column:follow_count;default:0"`
	FollowerCount int64  `gorm:"column:follower_count;default:0"`
	WorkCount     int64  `gorm:"column:work_count;default:0"`
	FavoriteCount int64  `gorm:"column:favorite_count;default:0"`
}

func (User) TableName() string {
	return "user"
}
