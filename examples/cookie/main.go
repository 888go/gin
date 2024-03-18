package main

import (
	"net/http"
	
	"github.com/888go/gin"
)


// ff:

// ff:

// ff:

// ff:

// ff:

// ff:
func CookieTool() gin.HandlerFunc {
	return func(c *gin.Context) {
// 把饼干
		if cookie, err := c.Cookie("label"); err == nil {
			if cookie == "ok" {
				c.Next()
				return
			}
		}

// Cookie验证失败
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden with no cookie"})
		c.Abort()
	}
}

func main() {
	route := gin.Default()

	route.GET("/login", func(c *gin.Context) {
// 设置cookie {"label";}， maxAge 30秒
		c.SetCookie("label", "ok", 30, "/", "localhost", false, true)
		c.String(200, "Login success!")
	})

	route.GET("/home", CookieTool(), func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "Your home page"})
	})

	route.Run(":8080")
}
