//go:build go1.8
// +build go1.8

package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	
	"github.com/888go/gin"
)

func main() {
	router := gin类.X创建默认对象()
	router.X绑定GET("/", func(c *gin类.Context) {
		time.Sleep(5 * time.Second)
		c.X输出文本(http.StatusOK, "Welcome Gin Server")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		log.Println("receive interrupt signal")
		if err := server.Close(); err != nil {
			log.Fatal("Server Close:", err)
		}
	}()

	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("Server closed under request")
		} else {
			log.Fatal("Server closed unexpect")
		}
	}

	log.Println("Server exiting")
}
