package main

import (
	"blog-server/handler"
	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func BlogList(c *gin.Context) {
	c.HTML(200, "blog.html", nil)
}
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./views/login.html", "./views/blog.html")

	r.GET("/login", login)
	r.GET("/blog/list", BlogList)

	// 提交接口submit
	r.POST("/login/submit", handler.Login)
	r.Run("127.0.0.1:5678")
}
