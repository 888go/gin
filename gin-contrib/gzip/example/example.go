package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	
	"github.com/888go/gin/gin-contrib/gzip"
	"github.com/888go/gin"
)

func main() {
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// 在0.0.0.0:8080监听并服务
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
