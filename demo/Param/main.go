package main

import (
	"github.com/888go/gin"
	"net/http"
	"strings"
)

// 参考链接：https://topgoer.com/gin框架/gin路由/api参数.html
func main() {
	r := gin类.X创建默认对象()
	r.X绑定GET("/user/:name/*action", func(c *gin类.Context) {
		name := c.X取API参数值("name")
		action := c.X取API参数值("action")
		//截取/
		action = strings.Trim(action, "/")
		c.X输出文本(http.StatusOK, name+" is "+action)
	})
	//默认为监听8080端口
	r.X监听(":8000")
}
