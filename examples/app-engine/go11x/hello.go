package main
import (
	"log"
	"net/http"
	"os"
	
	"e.coding.net/gogit/go/gin"
	)
func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

// 启动一个没有中间件的新Gin实例
	r := gin.New()

// 定义处理程序
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

// 在定义的端口上侦听和服务
	log.Printf("Listening on port %s", port)
	r.Run(":" + port)
}
