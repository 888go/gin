package main

import (
	"log"
	
	"github.com/888go/gin/gin-contrib/multitemplate"
	"github.com/888go/gin"
)

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("index", "templates/base.html", "templates/index.html")
	r.AddFromFiles("article", "templates/base.html", "templates/index.html", "templates/article.html")
	return r
}

func main() {
	router := gin类.X创建默认对象()
	router.HTMLRender = createMyRender()
	router.X绑定GET("/", func(c *gin类.Context) {
		c.X输出html模板(200, "index", gin类.H{
			"title": "Html5 Template Engine",
		})
	})
	router.X绑定GET("/article", func(c *gin类.Context) {
		c.X输出html模板(200, "article", gin类.H{
			"title": "Html5 Article Engine",
		})
	})

	if err := router.X监听(":8080"); err != nil {
		log.Fatal(err)
	}
}
