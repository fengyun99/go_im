package main

import (
	"dialog/router"
	"dialog/utils"
)

func main() {
	utils.InitConfig()
	utils.InitDB()

	r := router.Router()
	r.Run(":8081")
}
