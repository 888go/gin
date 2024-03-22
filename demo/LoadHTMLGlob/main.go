package main

import (
	"github.com/888go/gin"
	"net/http"
)

// Gin框架官方文档示例：HTML渲染（中文版）https://gin-gonic.com/zh-cn/docs/examples/html-rendering/
func main() {
	router := gin类.X创建默认对象()
	router.X加载HTML模板目录("./demo/LoadHTMLGlob/*")
	// router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
// 加载并缓存多个HTML模板文件，这些文件位于"templates"目录下，分别是"template1.html"和"template2.html"。
	router.X绑定GET("/index", func(c *gin类.Context) {
		c.X输出html模板(http.StatusOK, "index.tmpl", gin类.H{
			"title": "Main website",
		})
	})
	router.X监听(":8080")
}
