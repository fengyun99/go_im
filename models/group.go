package models

import (
	"gorm.io/gorm"
)

// Group 群信息
type Group struct {
	gorm.Model
	Name    string
	OwnerId int64
	Icon    string
	Type    string
	Desc    string
}

func (table *Group) TableName() string {
	return "group"
}
