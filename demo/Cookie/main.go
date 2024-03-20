package main

import (
	"fmt"
	"github.com/888go/gin"
)

// 参考 Gin 框架官方文档（中文版）：https://gin-gonic.com/zh-cn/docs/examples/cookie/
func main() {

	router := gin.Default()

	router.GET("/cookie", func(c *gin.Context) {

		cookie, err := c.Cookie("gin_cookie")

		if err != nil {
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}

		fmt.Printf("Cookie value: %s \n", cookie)
	})

	router.Run()
}
