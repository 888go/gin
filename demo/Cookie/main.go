package main

import (
	"fmt"
	"github.com/888go/gin"
)

// 参考 Gin 框架官方文档（中文版）：https://gin-gonic.com/zh-cn/docs/examples/cookie/
func main() {

	router := gin类.X创建默认对象()

	router.X绑定GET("/cookie", func(c *gin类.Context) {

		cookie, err := c.X取cookie值("gin_cookie")

		if err != nil {
			cookie = "NotSet"
			c.X设置cookie值("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}

		fmt.Printf("Cookie value: %s \n", cookie)
	})

	router.X监听()
}
