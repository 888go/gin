package main

import (
	"net/http"
	
	"github.com/888go/gin"
)

func main() {
	app := gin类.X创建默认对象()

// 从main.go目录的相对位置提供静态favicon文件
// app.StaticFile("/favicon.ico", "./.assets/favicon.ico")
	app.X绑定静态单文件("/favicon.ico", "./favicon.ico")

	app.X绑定GET("/ping", func(c *gin类.Context) {
		c.X输出文本(http.StatusOK, "Hello favicon.")
	})

	app.X监听(":8080")
}
