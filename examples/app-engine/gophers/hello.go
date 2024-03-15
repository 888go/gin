package hello

import (
	"net/http"
	
	"github.com/888go/gin"
)

// 这个函数的名字是必须的
// App Engine使用它来正确地驱动请求
func init() {
// 启动一个没有中间件的新Gin实例
	r := gin.New()

// 定义处理程序
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

// 使用net/http处理所有请求
	http.Handle("/", r)
}
