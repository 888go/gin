// 构建go1.16

package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
	
	"github.com/888go/gin"
)

func main() {
// 创建上下文，监听来自操作系统的中断信号
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(10 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

// 在运行例程中初始化服务器，使其不会阻塞下面的正常关机处理
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

// 监听中断信号
	<-ctx.Done()

// 恢复中断信号的默认行为，并通知用户关机
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

// 上下文用于通知服务器，它有5秒的时间来完成当前正在处理的请求
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
