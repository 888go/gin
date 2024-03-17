package main

import (
	"github.com/888go/gin"
	"net/http"
	"strings"
)

// https://topgoer.com/gin%E6%A1%86%E6%9E%B6/gin%E8%B7%AF%E7%94%B1/api%E5%8F%82%E6%95%B0.html
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
