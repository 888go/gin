package main

import (
	"log"
	"net/http"
	
	"github.com/888go/gin/gin-contrib/location"
	"github.com/888go/gin"
)

func main() {
	router := gin类.X创建默认对象()

// 配置为自动检测 scheme 和 host，同时提供回退方案：https://foo.com/base
// - 当无法确定默认 scheme 时使用 https
// - 当无法确定默认 host 时使用 foo.com
// - 将 /base 包含在路径中
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
