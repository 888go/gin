package hello

import (
	"net/http"
	
	"github.com/888go/gin"
)

// 这个函数的名称是必须的。App Engine 使用它来正确地驱动请求。
func init() {
	// 初始化一个新的Gin实例，不包含中间件
	r := gin.New()

	// Define your handlers
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// 使用net/http处理所有请求
	http.Handle("/", r)
}
