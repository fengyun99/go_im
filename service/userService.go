package service

import (
	"dialog/models"
	"dialog/utils"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// TODO:添加状态码,struct

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
	// 使用当前时间作为UUID的一部分
	currentTime := time.Now()

	// 将名字和当前时间拼接起来作为UUID的输入数据
	data := user.Name + currentTime.String()
	user.Uid = uuid.NewSHA1(uuid.NameSpaceURL, []byte(data)).String()

	// 简单加密
	salt := fmt.Sprintf("%06d", rand.Int31())
	user.Salt = salt
	if password == repassword {
		user.PassWord = utils.Encrypt(password, salt)
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

// UpdateUserByName
// @Summary 通过名字修改用户
// @Tags 用户模块
// @Param name query string true "用户名"
// @Param password query string true "密码"
// @Param repassword query string true "确认密码"
// @Accept json
// @Produce json
// @Success 200 {string} json{"code","message"}
// @Router /user/updateUserByName [get]
func UpdateUserByName(c *gin.Context) {
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
			models.UpdateUserByName(user)
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
// @Router /user/deleteUserByName [get]
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

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @Param id query string false "id"
// @Accept json
// @Produce json
// @Success 200 {string} json{"code","message"}
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(200, gin.H{
		"message": "success",
	})
}

// UpdateUser
// @Summary 修改用户信息
// @Tags 用户模块
// @Param id formData string false "id"
// @Param name formData string false "name"
// @Param password formData string false "password"
// @Param phone formData string false "phone"
// @Param email formData string false "email"
// @Success 200 {string} json{"code","message"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")
	user.Phone = c.PostForm("phone") // TODO:发送验证码校验
	user.Email = c.PostForm("email") // TODO:发送邮件校验
	// 校验
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err) //输出日志会报错为什么?
		c.JSON(200, gin.H{
			"message": "修改参数有问题",
		})
	} else {
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

	}

}

// FindUserByNameAndPwd
// @Summary 登录
// @Tags 用户模块
// @Param name query string false "name"
// @Param password query string false "password"
// @Accept json
// @Produce json
// @Success 200 {string} json{"code","message"}
// @Router /user/findUserByNameAndPwd [post]
func FindUserByNameAndPwd(c *gin.Context) {
	username := c.Query("name")
	plainPassword := c.Query("password")
	// 加密密码
	user := models.FindUserByName(username)
	// 判断用户是否存在---是否需要?安全角度将直接提示登录失败就行
	if user.Name == "" {
		c.JSON(-1, gin.H{
			"message": "用户不存在",
		})
		return
	}
	salt := user.Salt
	password := utils.Encrypt(plainPassword, salt)
	// TODO:校验登录次数和间隔时间
	if models.FindUserByNameAndPwd(username, password) {
		// 判断token是否存在---TODO:redis添加过期时间
		if user.Identity == "" {
			// token加密
			str := fmt.Sprintf("%d", time.Now().Unix())
			temp := utils.Md5Encode(str)
			models.GenToken(user, temp)
		}
		c.JSON(200, gin.H{
			"message": "登录成功",
		})
	} else {
		c.JSON(-1, gin.H{
			"message": "登录失败",
		})
	}
}

// 防止跨域站点伪造请求
var upGrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// SendMsg 为什么没打印日志
func SendMsg(c *gin.Context) {
	ws, err := upGrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ws)

	MsgHandler(ws, c)
}

func MsgHandler(ws *websocket.Conn, c *gin.Context) {
	msg, err := utils.Subscribe(c, utils.PublishKey)
	if err != nil {
		fmt.Println(err)
	}
	NowTime := time.Now().Format("2006-01-02 15:04:52")
	m := fmt.Sprintf("[ws][%s]:%s", NowTime, msg)
	err = ws.WriteMessage(1, []byte(m))
	if err != nil {
		fmt.Println(err)
	}
}
