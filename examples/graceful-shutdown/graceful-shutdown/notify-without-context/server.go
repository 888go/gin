//go:build go1.8
// +build go1.8

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	
	"github.com/888go/gin"
)

func main() {
	router := gin类.X创建默认对象()
	router.X绑定GET("/", func(c *gin类.Context) {
		time.Sleep(5 * time.Second)
		c.X输出文本(http.StatusOK, "Welcome Gin Server")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

// 在一个goroutine中初始化服务器，以便于
// 不会阻塞下面的优雅关闭处理流程
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

// 等待中断信号以优雅地关闭服务器，超时时间为5秒。
	quit := make(chan os.Signal, 1)
// (无参数) kill 默认发送 syscall.SIGTERM
// kill -2 等同于 syscall.SIGINT
// kill -9 等同于 syscall.SIGKILL，但无法被捕获，因此不需要添加它
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

// 上下文用于通知服务器，它有5秒钟的时间来完成当前正在处理的请求
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
