package routes

import (
	"github.com/888go/gin"
)

var router = gin.Default()

// Run 将启动服务器
func Run() {
	getRoutes()
	router.Run(":5000")
}

// getRoutes 将创建我们整个应用程序的路由
// 这样，每个路由组都可以在自己的文件中定义
// 从而避免此文件过于杂乱
func getRoutes() {
	v1 := router.Group("/v1")
	addUserRoutes(v1)
	addPingRoutes(v1)

	v2 := router.Group("/v2")
	addPingRoutes(v2)
}
