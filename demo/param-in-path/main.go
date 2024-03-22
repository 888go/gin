package main

import (
	"github.com/888go/gin"
	"net/http"
)

// 参考 Gin 框架官方文档（中文版）：
// https://gin-gonic.com/zh-cn/docs/examples/param-in-path/
func main() {
	router := gin类.X创建默认对象()

	// 此 handler 将匹配 /user/john 但不会匹配 /user/ 或者 /user
	router.X绑定GET("/user/:name", func(c *gin类.Context) {
		name := c.X取API参数值("name")
		c.X输出文本(http.StatusOK, "Hello %s", name)
	})

	// 此 handler 将匹配 /user/john/ 和 /user/john/send
	// 如果没有其他路由匹配 /user/john，它将重定向到 /user/john/
	router.X绑定GET("/用户名/:名称/*动作", func(c *gin类.Context) {
		name := c.X取API参数值("名称")
		action := c.X取API参数值("动作")
		message := name + " is " + action
		c.X输出文本(http.StatusOK, message)
	})

	router.X监听(":8080")
}
