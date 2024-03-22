package main

import (
	"fmt"
	"github.com/888go/gin"
	"net/http"
)

// 参考链接：https://topgoer.com/gin框架/gin路由/url参数.html
func main() {
	r := gin类.X创建默认对象()
	r.X绑定GET("/user", func(c *gin类.Context) {
		//指定默认值
		//http://localhost:8080/user 才会打印出来默认的值
		name := c.X取URL参数值并带默认("name", "枯藤")
		c.X输出文本(http.StatusOK, fmt.Sprintf("hello %s", name))
	})
	r.X监听()
}
