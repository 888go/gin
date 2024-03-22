package main

import (
	"fmt"
	"github.com/888go/gin"
)

// 官方示例：将map作为查询字符串或POST表单参数
// 参考网址：https://gin-gonic.com/zh-cn/docs/examples/map-as-querystring-or-postform/

// POST 请求，URL 为 /post，并在查询参数中携带键值对 ids[a]=1234 和 ids[b]=hello，使用 HTTP/1.1 协议版本
// 设置请求头 Content-Type 为 application/x-www-form-urlencoded，表示请求体内容采用 URL 编码格式
//
// 请求体数据为 names[first]=thinkerou&names[second]=tianou，其中包含两个键值对：names[first] 和 names[second]

// 返回 ids: map[b:hello a:1234], names: map[second:tianou first:thinkerou]
func main() {
	router := gin类.X创建默认对象()

	router.X绑定POST("/post", func(c *gin类.Context) {

		ids := c.X取URL参数Map值("ids")
		names := c.X取表单参数Map值("names")

		fmt.Printf("ids: %v; names: %v", ids, names)
	})
	router.X监听(":8080")
}
