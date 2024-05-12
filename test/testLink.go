package main

import (
	"dialog/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:Awsl711!@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//user := models.UserBasic{}
	//
	//// 初始化Table
	//db.AutoMigrate(&user)
	db.AutoMigrate(&models.Message{})
	db.AutoMigrate(&models.Relation{})
	db.AutoMigrate(&models.Group{})
}
