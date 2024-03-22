package main

import (
	"log"
	
	"github.com/888go/gin/gin-contrib/secure"
	"github.com/888go/gin"
)

func main() {
	router := gin类.X创建默认对象()

	router.X中间件(secure.New(secure.Config{
		AllowedHosts:          []string{"example.com", "ssl.example.com"},
		SSLRedirect:           true,
		SSLHost:               "ssl.example.com",
		STSSeconds:            315360000,
		STSIncludeSubdomains:  true,
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "default-src 'self'",
		IENoOpen:              true,
		ReferrerPolicy:        "strict-origin-when-cross-origin",
		SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
	}))

	router.X绑定GET("/ping", func(c *gin类.Context) {
		c.X输出文本(200, "pong")
	})

	// Listen and Server in 0.0.0.0:8080
	if err := router.X监听(); err != nil {
		log.Fatal(err)
	}
}
