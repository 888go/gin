package main

import (
	"log"
	"net/http"
	
	"github.com/888go/gin/gin-contrib/location"
	"github.com/888go/gin"
)

func main() {
	router := gin类.X创建默认对象()

	// configure to automatically detect scheme and host
	// - use http when default scheme cannot be determined
	// - use localhost:8080 when default host cannot be determined
	router.X中间件(location.Default())

	router.X绑定GET("/", func(c *gin类.Context) {
		url := location.Get(c)
		c.X输出文本(http.StatusOK, url.String())
	})

	if err := router.X监听(); err != nil {
		log.Fatal(err)
	}
}
