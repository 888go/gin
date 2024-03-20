package main

import (
	"log"
	
	"github.com/888go/gin/gin-contrib/static"
	"github.com/888go/gin"
)

func main() {
	r := gin.Default()

// 如果允许目录索引
// r.Use(static.Serve("/", static.LocalFile("/tmp", true)))
// 设置前缀
// r.Use(static.Serve("/static", static.LocalFile("/tmp", true)))
// 
// 翻译成中文：
// 
// 如果允许目录索引功能
// 使用 r.Use 方法，将本地目录 "/tmp" 的内容通过根路径 "/" 提供静态服务，并启用目录索引
// 设置访问前缀为 "/static"
// 使用 r.Use 方法，将本地目录 "/tmp" 的内容通过 "/static" 路径提供静态服务，并启用目录索引

	r.Use(static.Serve("/", static.LocalFile("/tmp", false)))
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "test")
	})
	// 在0.0.0.0:8080监听并服务
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
