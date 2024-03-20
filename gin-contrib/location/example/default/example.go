package main

import (
	"log"
	"net/http"
	
	"github.com/888go/gin/gin-contrib/location"
	"github.com/888go/gin"
)

func main() {
	router := gin.Default()

// 配置自动检测 scheme（协议）和 host（主机）
// - 当无法确定默认 scheme 时，使用 http
// - 当无法确定默认 host 时，使用 localhost:8080
	router.Use(location.Default())

	router.GET("/", func(c *gin.Context) {
		url := location.Get(c)
		c.String(http.StatusOK, url.String())
	})

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}
