package main

import (
	"github.com/888go/gin"
	"net/http"
	"strings"
)

// 参考链接：https://topgoer.com/gin框架/gin路由/api参数.html
func main() {
	r := gin.Default()
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		//截取/
		action = strings.Trim(action, "/")
		c.String(http.StatusOK, name+" is "+action)
	})
	//默认为监听8080端口
	r.Run(":8000")
}
