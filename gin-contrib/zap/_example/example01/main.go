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

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	r.X中间件(ginzap.Ginzap(logger, time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	r.X中间件(ginzap.RecoveryWithZap(logger, true))

	// Example ping request.
	r.X绑定GET("/ping", func(c *gin类.Context) {
		c.X输出文本(200, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// Example when panic happen.
	r.X绑定GET("/panic", func(c *gin类.Context) {
		panic("An unexpected error happen!")
	})

	// Listen and Server in 0.0.0.0:8080
	r.X监听(":8080")
}
