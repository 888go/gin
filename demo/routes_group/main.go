package main

import (
	"fmt"
	"github.com/888go/gin"
)

// 参考链接：https://topgoer.com/gin框架/gin路由/routesgroup.html

func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin类.X创建默认对象()
	// 路由组1 ，处理GET请求
	v1 := r.X创建分组路由("/v1")
	{ // {} 是书写规范
		v1.X绑定GET("/login", login)
		v1.X绑定GET("submit", submit)
	}

	v2 := r.X创建分组路由("/v2")
	{
		v2.X绑定POST("/login", login)
		v2.X绑定POST("/submit", submit)
	}
	r.X监听(":8000")
}

func login(c *gin类.Context) {
	name := c.X取URL参数值并带默认("name", "jack")
	c.X输出文本(200, fmt.Sprintf("hello %s\n", name))
}

func submit(c *gin类.Context) {
	name := c.X取URL参数值并带默认("name", "lily")
	c.X输出文本(200, fmt.Sprintf("hello %s\n", name))
}
