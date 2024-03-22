package routes

import (
	"net/http"
	
	"github.com/888go/gin"
)

func addUserRoutes(rg *gin类.RouterGroup) {
	users := rg.X创建分组路由("/users")

	users.X绑定GET("/", func(c *gin类.Context) {
		c.X输出JSON(http.StatusOK, "users")
	})
	users.X绑定GET("/comments", func(c *gin类.Context) {
		c.X输出JSON(http.StatusOK, "users comments")
	})
	users.X绑定GET("/pictures", func(c *gin类.Context) {
		c.X输出JSON(http.StatusOK, "users pictures")
	})
}
