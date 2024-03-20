package main

import (
	"fmt"
	"github.com/888go/gin"
	"net/http"
)

// 参考链接：https://topgoer.com/gin框架/简介.html
// 此段代码为Go语言中的单行注释，其内容是对一个URL的描述，翻译成中文即为：
// ```go
// 参考链接：https://topgoer.com/gin框架简介.html
func main() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		fmt.Println(c.FullPath())
		c.String(http.StatusOK, "hello World!")
	})
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")
}
