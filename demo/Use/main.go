package main

import (
	"github.com/888go/gin"
	"log"
	"time"
)

// Gin框架官方文档示例：自定义中间件（中文版）

// 定义Logger中间件

// ff:
func Logger() gin类.HandlerFunc {
	return func(c *gin类.Context) {
		t := time.Now()

		// 设置 example 变量
		c.X设置值("example", "12345")

		// 请求前

		c.X中间件继续()

		// 请求后
		latency := time.Since(t)
		log.Print(latency)

		// 获取发送的 status
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	r := gin类.X创建()
	r.X中间件(Logger()) //调用中间件

	r.X绑定GET("/test", func(c *gin类.Context) {
		example := c.X取值PANI("example").(string)

		// 打印："12345"
		log.Println(example)
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.X监听(":8080")
}
