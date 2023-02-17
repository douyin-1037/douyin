package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

// Relation Gorm Data Structures
type Relation struct {
	UserId    int64 `gorm:"column:user_id;not null;index:fk_user_relation"`
	ToUserId  int64 `gorm:"column:to_user_id;not null;index:fk_user_relation_to"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt soft_delete.DeletedAt
}

func (Relation) TableName() string {
	return "relation"
}
