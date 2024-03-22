package main

import (
	"fmt"
	"time"
	
	ginzap "github.com/888go/gin/gin-contrib/zap"
	"github.com/888go/gin"
	"go.uber.org/zap"
)

func main() {
	r := gin类.X创建()

	logger, _ := zap.NewProduction()

// 添加一个ginzap中间件，其功能包括：
//   - 记录所有请求，类似于综合访问日志和错误日志。
//   - 将日志输出到标准输出（stdout）。
//   - 使用UTC时间格式并遵循RFC3339规范。
	r.X中间件(ginzap.Ginzap(logger, time.RFC3339, true))

// 将所有 panic 记录到错误日志中
//   - stack 表示是否输出堆栈信息。
	r.X中间件(ginzap.RecoveryWithZap(logger, true))

	// Example ping request.
	r.X绑定GET("/ping", func(c *gin类.Context) {
		c.X输出文本(200, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// 示例：发生panic的情况。
	r.X绑定GET("/panic", func(c *gin类.Context) {
		panic("An unexpected error happen!")
	})

	// 在0.0.0.0:8080监听并服务
	r.X监听(":8080")
}
