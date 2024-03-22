package main

import (
	"log"
	"net/http"
	
	"github.com/888go/gin/gin-contrib/location"
	"github.com/888go/gin"
)

func main() {
	router := gin类.X创建默认对象()

	// configure to automatically detect scheme and host with
	// fallback to https://foo.com/base
	// - use https when default scheme cannot be determined
	// - use foo.com when default host cannot be determined
	// - include /base as the path
	router.X中间件(location.New(location.Config{
		Scheme:  "https",
		Host:    "foo.com",
		Base:    "/base",
		Headers: location.Headers{Scheme: "X-Forwarded-Proto", Host: "X-Forwarded-For"},
	}))

	router.X绑定GET("/", func(c *gin类.Context) {
		url := location.Get(c)
		c.X输出文本(http.StatusOK, url.String())
	})

	if err := router.X监听(); err != nil {
		log.Fatal(err)
	}
}
