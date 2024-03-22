package main

import (
	"log"
	
	"github.com/888go/gin/gin-contrib/secure"
	"github.com/888go/gin"
)

func main() {
	router := gin类.X创建默认对象()

	securityConfig := secure.DefaultConfig()
	securityConfig.AllowedHosts = []string{"example.com", "ssl.example.com"}
	securityConfig.SSLHost = "ssl.example.com"
	router.X中间件(secure.New(securityConfig))

	router.X绑定GET("/ping", func(c *gin类.Context) {
		c.X输出文本(200, "pong")
	})

	if err := router.X监听(); err != nil {
		log.Fatal(err)
	}
}
