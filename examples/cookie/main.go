package main

import (
	"net/http"
	
	"github.com/888go/gin"
)

func CookieTool() gin类.HandlerFunc {
	return func(c *gin类.Context) {
		// Get cookie
		if cookie, err := c.X取cookie值("label"); err == nil {
			if cookie == "ok" {
				c.X中间件继续()
				return
			}
		}

		// Cookie verification failed
		c.X输出JSON(http.StatusForbidden, gin类.H{"error": "Forbidden with no cookie"})
		c.X停止()
	}
}

func main() {
	route := gin类.X创建默认对象()

	route.X绑定GET("/login", func(c *gin类.Context) {
		// Set cookie {"label": "ok" }, maxAge 30 seconds.
		c.X设置cookie值("label", "ok", 30, "/", "localhost", false, true)
		c.X输出文本(200, "Login success!")
	})

	route.X绑定GET("/home", CookieTool(), func(c *gin类.Context) {
		c.X输出JSON(200, gin类.H{"data": "Your home page"})
	})

	route.X监听(":8080")
}
