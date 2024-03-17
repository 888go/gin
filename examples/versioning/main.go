package main

import (
	"net/http"
	
	"github.com/888go/gin"
)

func main() {
	router := gin.Default()

// 版本1
	apiV1 := router.Group("/v1")

	apiV1.GET("users", func(c *gin.Context) {
		c.JSON(http.StatusOK, "List Of V1 Users")
	})

// 用户只能由授权人员添加
	authV1 := apiV1.Group("/", AuthMiddleWare())

	authV1.POST("users/add", AddV1User)

// 版本2
	apiV2 := router.Group("/v2")

	apiV2.GET("users", func(c *gin.Context) {
		c.JSON(http.StatusOK, "List Of V2 Users")
	})

// 用户只能由授权人员添加
	authV2 := apiV2.Group("/", AuthMiddleWare())

	authV2.POST("users/add", AddV2User)

	_ = router.Run(":8081")
}


// ff:
// c:

// ff:
// c:
func AddV1User(c *gin.Context) {
// 添加用户

	c.JSON(http.StatusOK, "V1 User added")
}


// ff:
// c:

// ff:
// c:
func AddV2User(c *gin.Context) {
// 添加用户

	c.JSON(http.StatusOK, "V2 User added")
}


// ff:

// ff:
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
// 在这里，您可以添加身份验证方法来授权用户
		username := c.PostForm("user")
		password := c.PostForm("password")

		if username == "foo" && password == "bar" {
			return
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
