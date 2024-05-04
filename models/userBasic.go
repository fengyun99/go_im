package models

import (
	"dialog/utils"
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string
	Email         string
	Identity      string
	ClientIP      string
	ClientPort    string
	LoginTime     time.Time `gorm:"default:null;"`
	HeartbeatTime time.Time `gorm:"default:null;"`
	LogoutTime    time.Time `gorm:"default:null;"`
	IsLogin       bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	return data
}

func CreateUser(user UserBasic) *gorm.DB {
	return utils.DB.Create(&user)
}

func UpdateUser(user UserBasic) *gorm.DB {
	return utils.DB.Model(&user).Where("name = ?", user.Name).Updates(&user)
}

func DeleteUserByName(user UserBasic) *gorm.DB {
	return utils.DB.Where("name = ?", user.Name).Delete(&UserBasic{})
}

func CheckUserExist(user UserBasic) bool {
	result := utils.DB.Where("name = ?", user.Name).First(&user)
	if result.Error != nil {
		// 用户不存在
		return false
	} else {
		// 用户存在
		return true
	}
}
