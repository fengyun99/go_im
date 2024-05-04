package service

import (
	"dialog/models"
	"github.com/gin-gonic/gin"
)

// GetUserList
// @Summary 用户列表
// @Tags 用户模块
// @Accept json
// @Produce json
// @Success 200 {string} json{"code","message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	//userData := make([]*models.UserBasic, 10)
	userData := models.GetUserList()

	c.JSON(200, gin.H{
		"message": userData,
	})
}

// CreateUser
// @Summary 新增用户
// @Tags 用户模块
// @Param name query string true "用户名"
// @Param password query string true "密码"
// @Param repassword query string true "确认密码"
// @Accept json
// @Produce json
// @Success 200 {string} json{"code","message"}
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")
	if password == repassword {
		user.PassWord = password
		if !models.CheckUserExist(user) {
			models.CreateUser(user)
			c.JSON(200, gin.H{
				"message": "success",
			})

		} else {
			c.JSON(-1, gin.H{
				"message": "用户已存在，更换名字",
			})
		}
	} else {
		c.JSON(-1, gin.H{
			"message": "两次密码不一致",
		})
	}
}

// UpdateUser
// @Summary 更新用户
// @Tags 用户模块
// @Param name query string true "用户名"
// @Param password query string true "密码"
// @Param repassword query string true "确认密码"
// @Accept json
// @Produce json
// @Success 200 {string} json{"code","message"}
// @Router /user/updateUser [get]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")
	if password == repassword {
		user.PassWord = password
		// 先查询数据库中是否存在当前名字
		if !models.CheckUserExist(user) {
			c.JSON(-1, gin.H{
				"message": "用户不存在",
			})
		} else {
			models.UpdateUser(user)
			c.JSON(200, gin.H{
				"message": "success",
			})

		}

	} else {
		c.JSON(-1, gin.H{
			"message": "两次密码不一致",
		})
	}
}

// DeleteUserByName
// @Summary 通过名字删除用户
// @Tags 用户模块
// @Param name query string true "用户名"
// @Accept json
// @Produce json
// @Success 200 {string} json{"code","message"}
// @Router /user/deleteUser [get]
func DeleteUserByName(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	if !models.CheckUserExist(user) {
		c.JSON(-1, gin.H{
			"message": "用户不存在",
		})
	} else {
		models.DeleteUserByName(user)
		c.JSON(200, gin.H{
			"message": "success",
		})
	}

}
