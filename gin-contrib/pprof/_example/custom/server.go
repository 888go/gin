package main

import (
	"net/http"
	
	"github.com/888go/gin/gin-contrib/pprof"
	"github.com/888go/gin"
)

func main() {
	router := gin类.X创建默认对象()
	adminGroup := router.X创建分组路由("/admin", func(c *gin类.Context) {
		if c.X请求.Header.Get("Authorization") != "foobar" {
			c.X停止并带状态码(http.StatusForbidden)
			return
		}
		c.X中间件继续()
	})
	pprof.RouteRegister(adminGroup, "pprof")
	router.X监听(":8080")
}
