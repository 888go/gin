package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	
	"github.com/888go/gin/gin-contrib/requestid"
	"github.com/888go/gin"
)

func main() {
	r := gin类.X创建()

	r.X中间件(requestid.New())

	// Example ping request.
	r.X绑定GET("/ping", func(c *gin类.Context) {
		c.X输出文本(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// Example / request.
	r.X绑定GET("/", func(c *gin类.Context) {
		c.X输出文本(http.StatusOK, "id:"+requestid.Get(c))
	})

	// 在0.0.0.0:8080监听并服务
	if err := r.X监听(":8080"); err != nil {
		log.Fatal(err)
	}
}
