package main

import (
	"fmt"
	"github.com/888go/gin"
	"net/http"
)

// https://topgoer.com/gin%E6%A1%86%E6%9E%B6/gin%E8%B7%AF%E7%94%B1/url%E5%8F%82%E6%95%B0.html
func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		//指定默认值
		//http://localhost:8080/user 才会打印出来默认的值
		name := c.DefaultQuery("name", "枯藤")
		c.String(http.StatusOK, fmt.Sprintf("hello %s", name))

	})
	r.Run()
}
