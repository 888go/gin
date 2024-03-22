package main

import (
	"log"
	"net/http"
	
	"github.com/888go/gin/gin-contrib/location"
	"github.com/888go/gin"
)

func main() {
	router := gin类.X创建默认对象()

// 配置自动检测 scheme（协议）和 host（主机）
// - 当无法确定默认 scheme 时，使用 http
// - 当无法确定默认 host 时，使用 localhost:8080
	router.X中间件(location.Default())

	router.X绑定GET("/", func(c *gin类.Context) {
		url := location.Get(c)
		c.X输出文本(http.StatusOK, url.String())
	})

	if err := router.X监听(); err != nil {
		log.Fatal(err)
	}
}
