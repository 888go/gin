package main

import (
	"github.com/888go/gin"
	"net/http"
)

// 访问 gin-gonic.com 的中文文档中“使用基本认证中间件”示例部分
// （由于您提供的代码片段仅包含一个链接，故此翻译为对链接内容的描述性注释）

// 模拟一些私人数据
var secrets = gin类.H{
	"foo":    gin类.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin类.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin类.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main() {
	r := gin类.X创建默认对象()

	// 路由组使用 gin.BasicAuth() 中间件
	// gin.Accounts 是 map[string]string 的一种快捷方式
	authorized := r.X创建分组路由("/admin", gin类.X中间件函数_简单认证(gin类.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	// /admin/secrets 端点
	// 触发 "localhost:8080/admin/secrets
	authorized.X绑定GET("/secrets", func(c *gin类.Context) {
		// 获取用户，它是由 BasicAuth 中间件设置的
		user := c.X取值PANI(gin类.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.X输出JSON(http.StatusOK, gin类.H{"user": user, "secret": secret})
		} else {
			c.X输出JSON(http.StatusOK, gin类.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.X监听(":8080")
}
