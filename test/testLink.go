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

	user := models.UserBasic{}

	//db.AutoMigrate(&user)

	//user.Name = "默默"
	//
	//db.Create(&user)
	//
	//db.First(&user, 1)
	//
	//db.Model(&user).Update("PassWord", 1234)
	user.Name = "fengyun"
	user.PassWord = "1"
	db.Model(&user).Where("name = ?", "默默").Updates(map[string]interface{}{"name": "默默", "pass_word": "1"})
	//print(db.Row())
}
