package main

import (
	"github.com/888go/gin"
	"log"
	"net/http"
	"time"
)

// https://gin-gonic.com/zh-cn/docs/examples/custom-middleware/

// 定义Logger中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 设置 example 变量
		c.Set("example", "12345")

		// 请求前

		c.Next()

		// 请求后
		latency := time.Since(t)
		log.Print(latency)

		// 获取发送的 status
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	r := gin.New()
	r.Use(Logger()) //调用中间件

	r.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}
