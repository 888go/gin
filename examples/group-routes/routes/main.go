package routes

import (
	"github.com/888go/gin"
)

var router = gin.Default()

// Run将启动服务器

// ff:

// ff:

// ff:

// ff:

// ff:
func Run() {
	getRoutes()
	router.Run(":5000")
}

// getRoutes将以这种方式创建我们整个应用程序的路由，每一组路由都可以在它们自己的文件中定义，所以这个不会那么混乱
func getRoutes() {
	v1 := router.Group("/v1")
	addUserRoutes(v1)
	addPingRoutes(v1)

	v2 := router.Group("/v2")
	addPingRoutes(v2)
}
