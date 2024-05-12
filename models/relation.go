package models

import (
	"gorm.io/gorm"
)

// Relation 好友关系
type Relation struct {
	gorm.Model
	OwnerId  int64 // 谁的关系信息
	TargetId int64 // 对应的谁
	Type     int   //关系类型 0 1 2
	Desc     string
}

func (table *Relation) TableName() string {
	return "relation"
}
