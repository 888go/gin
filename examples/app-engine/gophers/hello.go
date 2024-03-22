package hello

import (
	"net/http"
	
	"github.com/888go/gin"
)

// 这个函数的名称是必须的。App Engine 使用它来正确地驱动请求。
func init() {
	// 初始化一个新的Gin实例，不包含中间件
	r := gin类.X创建()

	// Define your handlers
	r.X绑定GET("/", func(c *gin类.Context) {
		c.X输出文本(http.StatusOK, "Hello World!")
	})
	r.X绑定GET("/ping", func(c *gin类.Context) {
		c.X输出文本(http.StatusOK, "pong")
	})

	// 使用net/http处理所有请求
	http.Handle("/", r)
}
