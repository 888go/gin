package main

import (
	"github.com/888go/gin"
	"net/http"
)

// https://gin-gonic.com/zh-cn/docs/examples/html-rendering/
func main() {
	router := gin.Default()
	router.LoadHTMLGlob("./demo/LoadHTMLGlob/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	router.Run(":8080")
}
