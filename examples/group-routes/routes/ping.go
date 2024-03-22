package routes

import (
	"net/http"
	
	"github.com/888go/gin"
)

func addPingRoutes(rg *gin类.RouterGroup) {
	ping := rg.X创建分组路由("/ping")

	ping.X绑定GET("/", func(c *gin类.Context) {
		c.X输出JSON(http.StatusOK, "pong")
	})
}
