package main

import (
	"github.com/888go/gin"
	"net/http"
)

// Gin框架官方文档示例：HTML渲染（中文版）https://gin-gonic.com/zh-cn/docs/examples/html-rendering/
func main() {
	router := gin.Default()
	router.LoadHTMLGlob("./demo/LoadHTMLGlob/*")
	// router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
// 加载并缓存多个HTML模板文件，这些文件位于"templates"目录下，分别是"template1.html"和"template2.html"。
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	router.Run(":8080")
}
