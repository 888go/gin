package main

import (
	"net/http"
	
	"github.com/888go/gin"
)

func main() {
	router := gin.Default()

	// version 1
	apiV1 := router.Group("/v1")

	apiV1.GET("users", func(c *gin.Context) {
		c.JSON(http.StatusOK, "List Of V1 Users")
	})

	// 只有经过授权的人员才能添加用户
	authV1 := apiV1.Group("/", AuthMiddleWare())

	authV1.POST("users/add", AddV1User)

	// version 2
	apiV2 := router.Group("/v2")

	apiV2.GET("users", func(c *gin.Context) {
		c.JSON(http.StatusOK, "List Of V2 Users")
	})

	// 只有经过授权的人员才能添加用户
	authV2 := apiV2.Group("/", AuthMiddleWare())

	authV2.POST("users/add", AddV2User)

	_ = router.Run(":8081")
}


// ff:
// c:
func AddV1User(c *gin.Context) {
	// AddUser

	c.JSON(http.StatusOK, "V1 User added")
}


// ff:
// c:
func AddV2User(c *gin.Context) {
	// AddUser

	c.JSON(http.StatusOK, "V2 User added")
}


// ff:
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 在这里，您可以添加您的认证方法以授权用户。
		username := c.PostForm("user")
		password := c.PostForm("password")

		if username == "foo" && password == "bar" {
			return
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
