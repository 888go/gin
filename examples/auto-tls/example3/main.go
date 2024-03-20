package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	
	"github.com/gin-gonic/autotls"
	"github.com/888go/gin"
)

func main() {
	// 创建一个上下文，用于监听来自操作系统的中断信号。
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	r := gin.Default()

	// Ping handler
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	log.Fatal(autotls.RunWithContext(ctx, r, "example1.com", "example2.com"))
}
