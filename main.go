package main

import (
	"github.com/gin-gonic/gin"
)

func Home(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"Message": "pong",
	})
}

func Html(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("views/login.html")
	r.GET("/", Home)
	r.GET("/login", Html)
	r.Run("127.0.0.1:5678")
}
