package main

import (
	"fmt"
	"github.com/888go/gin"
)

//https://gin-gonic.com/zh-cn/docs/examples/map-as-querystring-or-postform/

// POST /post?ids[a]=1234&ids[b]=hello HTTP/1.1
// Content-Type: application/x-www-form-urlencoded
//
// names[first]=thinkerou&names[second]=tianou

// 返回 ids: map[b:hello a:1234], names: map[second:tianou first:thinkerou]
func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {

		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		fmt.Printf("ids: %v; names: %v", ids, names)
	})
	router.Run(":8080")
}
