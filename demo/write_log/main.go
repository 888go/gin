package main

import (
	"github.com/888go/gin"
	"io"
	"os"
)

// 参考 Gin 框架官方文档（中文版）：
// https://gin-gonic.com/zh-cn/docs/examples/write-log/
// 由于您给出的代码片段仅为一个链接引用，并无实际的 Go 语言代码，故以上为对此链接内容的中文描述。如果需要对具体代码进行翻译，请提供相应的 Go 语言代码段。
func main() {
	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	gin.DisableConsoleColor()

	// 记录到文件。
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":8080")
}
