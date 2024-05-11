package main

import (
	"dialog/router"
	"dialog/utils"
)

func main() {
	utils.InitConfig()
	utils.InitDB()
	utils.InitRedis()

	r := router.Router()
	r.Run(":8081")
}
