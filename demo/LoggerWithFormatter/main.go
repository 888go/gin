package main

import (
	"fmt"
	"github.com/888go/gin"
	"time"
)

// Gin框架官方文档示例：自定义中间件（中文版）

func main() {
	router := gin类.X创建()
	// LoggerWithFormatter 中间件会写入日志到 gin.DefaultWriter
	// 默认 gin.DefaultWriter = os.Stdout
	router.X中间件(gin类.X中间件函数_自定义日志格式(func(param gin类.LogFormatterParams) string {
		// 你的自定义格式
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.X客户端IP,
			param.X响应时间.Format(time.RFC1123),
			param.X方法,
			param.X路径,
			param.X请求.Proto,
			param.X状态码,
			param.X时长,
			param.X请求.UserAgent(),
			param.X错误信息,
		)
	}))
	router.X中间件(gin类.Recovery())
	router.X绑定GET("/ping", func(c *gin类.Context) {
		c.X输出文本(200, "pong")
	})
	router.X监听(":8080")
}
