package main

import (
	"log"
	
	"github.com/gin-gonic/autotls"
	"github.com/888go/gin"
)

func main() {
	r := gin.Default()

// 平处理程序
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	log.Fatal(autotls.Run(r, "example1.com", "example2.com"))
}
