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
	r := gin类.X创建默认对象()
	r.X中间件(gzip.Gzip(gzip.DefaultCompression))
	r.X绑定GET("/ping", func(c *gin类.Context) {
		c.X输出文本(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// Listen and Server in 0.0.0.0:8080
	if err := r.X监听(":8080"); err != nil {
		log.Fatal(err)
	}
}
