package main

import (
	"net/http"
	
	"github.com/888go/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
// 禁用控制台颜色
// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

// 授权分组（使用gin.BasicAuth()中间件）
// 等同于：
// authorized := r.Group("/")
// authorized.Use(gin.BasicAuth(gin.Credentials{
//	  "foo":  "bar", // 用户名：密码
//	  "manu": "123",
// }))
// 
// 这段Go注释翻译成中文后的大致意思是：
// 
// 此处定义一个授权访问的分组，该分组将采用gin.BasicAuth()中间件进行身份验证。
// 这与以下代码功能相同：
// 首先创建一个名为authorized的路由分组，并将其根路径设置为"/"。
// 然后在该分组中使用gin.BasicAuth()中间件进行基本认证，其中包含如下用户名和密码凭据：
// 用户名 "foo" 对应的密码是 "bar"
// 用户名 "manu" 对应的密码是 "123"
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
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
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	r := setupRouter()
	// 在0.0.0.0:8080监听并服务
	r.Run(":8080")
}
