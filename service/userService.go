package service

import (
	"dialog/models"
	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {
	//userData := make([]*models.UserBasic, 10)
	userData := models.GetUserList()

	c.JSON(200, gin.H{
		"message": userData,
	})
}
