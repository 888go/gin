package routes

import (
	"github.com/888go/gin"
)

var router = gin类.X创建默认对象()

// Run will start the server
func X监听() {
	getRoutes()
	router.X监听(":5000")
}

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
func getRoutes() {
	v1 := router.X创建分组路由("/v1")
	addUserRoutes(v1)
	addPingRoutes(v1)

	v2 := router.X创建分组路由("/v2")
	addPingRoutes(v2)
}
