package utils

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	DB    *gorm.DB
	REDIS *redis.Client
)

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
	// 添加日志,打印SQL语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢sql的阈值
			LogLevel:      logger.Info, // 日志级别
			Colorful:      true,        // 彩色
		},
	)

	dbUsername := viper.GetString("mysql.username")
	dbPassword := viper.GetString("mysql.password")
	dbHost := viper.GetString("mysql.host")
	dbPort := viper.GetInt("mysql.port")
	dbName := viper.GetString("mysql.database")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUsername, dbPassword, dbHost, dbPort, dbName)
	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database")
	}

	DB = _db
}

func InitRedis() {
	host := viper.GetString("redis.host")
	port := viper.GetInt("redis.port")
	password := viper.GetString("redis.password")
	db := viper.GetInt("redis.db")
	poolSize := viper.GetInt("redis.poolSize")
	minIdleConn := viper.GetInt("redis.minIdleConn")

	rdb := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprint(host) + ":" + fmt.Sprint(port),
		Password:     password,
		DB:           db,
		MinIdleConns: minIdleConn,
		PoolSize:     poolSize,
	})

	// 创建一个上下文对象
	ctx := context.Background()

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Redis init success...", pong)
		REDIS = rdb
	}
}

const PublishKey = "websocket"

// Publish 发布消息到Redis
func Publish(ctx context.Context, channel string, msg string) error {
	var err error
	fmt.Println("Publish...", msg)
	err = REDIS.Publish(ctx, channel, msg).Err()
	return err
}

// Subscribe 订阅Redis消息
func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := REDIS.Subscribe(ctx, channel)
	msg, err := sub.ReceiveMessage(ctx)
	fmt.Println("Subscribe...", msg.Payload)
	return msg.Payload, err
}
