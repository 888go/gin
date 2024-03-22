package main

import (
	"net/http"
	
	"github.com/888go/gin"
)

var db = make(map[string]string)

func setupRouter() *gin类.Engine {
// 禁用控制台颜色
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
	// 在0.0.0.0:8080监听并服务
	r.X监听(":8080")
}
