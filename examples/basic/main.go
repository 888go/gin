package main

import (
	"net/http"
	
	"github.com/888go/gin"
)

var db = make(map[string]string)

func setupRouter() *gin类.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin类.X创建默认对象()

	// Ping test
	r.X绑定GET("/ping", func(c *gin类.Context) {
		c.X输出文本(http.StatusOK, "pong")
	})

	// Get user value
	r.X绑定GET("/user/:name", func(c *gin类.Context) {
		user := c.X参数.ByName("name")
		value, ok := db[user]
		if ok {
			c.X输出JSON(http.StatusOK, gin类.H{"user": user, "value": value})
		} else {
			c.X输出JSON(http.StatusOK, gin类.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.X创建分组路由("/", gin类.X中间件函数_简单认证(gin类.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.X绑定POST("admin", func(c *gin类.Context) {
		user := c.X取值PANI(gin类.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.X取参数到指针PANI(&json) == nil {
			db[user] = json.Value
			c.X输出JSON(http.StatusOK, gin类.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.X监听(":8080")
}
