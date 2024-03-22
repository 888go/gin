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
	gin类.X关闭控制台颜色()

	// 记录到文件。
	f, _ := os.Create("gin.log")
	gin类.DefaultWriter = io.MultiWriter(f)

	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin类.X创建默认对象()
	router.X绑定GET("/ping", func(c *gin类.Context) {
		c.X输出文本(200, "pong")
	})

	router.X监听(":8080")
}
