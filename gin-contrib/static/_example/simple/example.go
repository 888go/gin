package main

import (
	"log"
	
	"github.com/888go/gin/gin-contrib/static"
	"github.com/888go/gin"
)

func main() {
	r := gin类.X创建默认对象()

	// if Allow DirectoryIndex
	// r.Use(static.Serve("/", static.LocalFile("/tmp", true)))
	// set prefix
	// r.Use(static.Serve("/static", static.LocalFile("/tmp", true)))

	r.X中间件(static.Serve("/", static.LocalFile("/tmp", false)))
	r.X绑定GET("/ping", func(c *gin类.Context) {
		c.X输出文本(200, "test")
	})
	// Listen and Server in 0.0.0.0:8080
	if err := r.X监听(":8080"); err != nil {
		log.Fatal(err)
	}
}
