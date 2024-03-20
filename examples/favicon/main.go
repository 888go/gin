package main

import (
	"net/http"
	
	"github.com/888go/gin"
)

func main() {
	app := gin.Default()

// 从main.go目录的相对位置提供静态favicon文件
// app.StaticFile("/favicon.ico", "./.assets/favicon.ico")
	app.StaticFile("/favicon.ico", "./favicon.ico")

	app.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello favicon.")
	})

	app.Run(":8080")
}
