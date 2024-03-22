package routes

import (
	"github.com/888go/gin"
)

var router = gin类.X创建默认对象()

// Run 将启动服务器
func X监听() {
	getRoutes()
	router.X监听(":5000")
}

// getRoutes 将创建我们整个应用程序的路由
// 这样，每个路由组都可以在自己的文件中定义
// 从而避免此文件过于杂乱
func getRoutes() {
	v1 := router.X创建分组路由("/v1")
	addUserRoutes(v1)
	addPingRoutes(v1)

	v2 := router.X创建分组路由("/v2")
	addPingRoutes(v2)
}
