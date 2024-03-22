package main

import (
	"log"
	
	"github.com/888go/gin/gin-contrib/static"
	"github.com/888go/gin"
)

func main() {
	r := gin类.X创建默认对象()

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

	r.X中间件(static.Serve("/", static.LocalFile("/tmp", false)))
	r.X绑定GET("/ping", func(c *gin类.Context) {
		c.X输出文本(200, "test")
	})
	// 在0.0.0.0:8080监听并服务
	if err := r.X监听(":8080"); err != nil {
		log.Fatal(err)
	}
}
