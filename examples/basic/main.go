package main
import (
	"net/http"
	
	"e.coding.net/gogit/go/gin"
	)
var db = make(map[string]string)

func setupRouter() *gin.Engine {
// 禁用控制台颜色
// disableconsolecolor ()
	r := gin.Default()

// Ping测试
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

// 获取用户价值
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

// 授权组(使用gin. basicauth()中间件)与:Authorized:= r.Group("/") Authorized . use (gin. basicauth())相同
// 凭证{"foo";bar";manu"; "123"}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // 用户:foo密码:酒吧
		"manu": "123", // 用户:马努密码:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http:// localhost: 8080 / admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

// 解析JSON
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
// 监听和服务器在0.0.0.0:8080
	r.Run(":8080")
}
