package handler

import (
	"blog-server/database"
	"blog-server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginResponse struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Uid   int    `json:"uid"`
	Token string `json:"token"`
}

func Login(c *gin.Context) {
	name := c.PostForm("user")
	pass := c.PostForm("pass")

	if len(name) == 0 {
		c.JSON(200, gin.H{
			"message": "用户名不能为空",
		})
	}
	if len(pass) == 0 {
		c.JSON(200, gin.H{
			"message": "密码不能为空",
		})
	}
	user := database.GetUserByName(name)
	if user == nil {
		c.JSON(1001, gin.H{
			"message": "用户名不存在",
		})
	}
	if user.PassWd != pass {
		c.JSON(1002, gin.H{
			"message": "用户名或密码不正确",
		})
	}
	utils.LogRus.Infof("user %s(%d) login", name, user.Id)
	c.JSON(http.StatusOK, LoginResponse{
		Code:  http.StatusOK,
		Msg:   "ok",
		Uid:   user.Id,
		Token: "",
	})
}
