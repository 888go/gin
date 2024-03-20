package main

import (
	"log"
	"net/http"
	"os"
	
	"github.com/888go/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	// 初始化一个新的Gin实例，不包含中间件
	r := gin.New()

	// Define handlers
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// 在定义的端口上监听并提供服务
	log.Printf("Listening on port %s", port)
	r.Run(":" + port)
}
