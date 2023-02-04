package model

import "gorm.io/gorm"

// Relation Gorm Data Structures
type Relation struct {
	gorm.Model
	UserId   int64 `gorm:"column:user_id;not null;index:fk_user_relation"`
	ToUserId int64 `gorm:"column:to_user_id;not null;index:fk_user_relation_to"`
}

func (Relation) TableName() string {
	return "relation"
}
