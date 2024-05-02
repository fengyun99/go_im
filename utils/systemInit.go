package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	//fmt.Println("config app", viper.Get("app"))
	//fmt.Println("config mysql", viper.Get("mysql"))
}

func InitDB() {
	dbUsername := viper.GetString("mysql.username")
	dbPassword := viper.GetString("mysql.password")
	dbHost := viper.GetString("mysql.host")
	dbPort := viper.GetInt("mysql.port")
	dbName := viper.GetString("mysql.database")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUsername, dbPassword, dbHost, dbPort, dbName)
	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = _db
}
