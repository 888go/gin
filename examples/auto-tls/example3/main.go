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
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	r := gin类.X创建默认对象()

	// Ping handler
	r.X绑定GET("/ping", func(c *gin类.Context) {
		c.X输出文本(http.StatusOK, "pong")
	})

	log.Fatal(autotls.RunWithContext(ctx, r, "example1.com", "example2.com"))
}
