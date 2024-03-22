package main

import (
	"net/http"
	
	"github.com/888go/gin"
)

func main() {
	router := gin类.X创建默认对象()

	// version 1
	apiV1 := router.X创建分组路由("/v1")

	apiV1.X绑定GET("users", func(c *gin类.Context) {
		c.X输出JSON(http.StatusOK, "List Of V1 Users")
	})

	// 只有经过授权的人员才能添加用户
	authV1 := apiV1.X创建分组路由("/", AuthMiddleWare())

	authV1.X绑定POST("users/add", AddV1User)

	// version 2
	apiV2 := router.X创建分组路由("/v2")

	apiV2.X绑定GET("users", func(c *gin类.Context) {
		c.X输出JSON(http.StatusOK, "List Of V2 Users")
	})

	// 只有经过授权的人员才能添加用户
	authV2 := apiV2.X创建分组路由("/", AuthMiddleWare())

	authV2.X绑定POST("users/add", AddV2User)

	_ = router.X监听(":8081")
}

func AddV1User(c *gin类.Context) {
	// AddUser

	c.X输出JSON(http.StatusOK, "V1 User added")
}

func AddV2User(c *gin类.Context) {
	// AddUser

	c.X输出JSON(http.StatusOK, "V2 User added")
}

func AuthMiddleWare() gin类.HandlerFunc {
	return func(c *gin类.Context) {
		// 在这里，您可以添加您的认证方法以授权用户。
		username := c.X取表单参数值("user")
		password := c.X取表单参数值("password")

		if username == "foo" && password == "bar" {
			return
		} else {
			c.X停止并带状态码(http.StatusUnauthorized)
		}
	}
}
