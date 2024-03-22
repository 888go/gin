package main

import (
	"log"
	"path/filepath"
	
	"github.com/888go/gin/gin-contrib/multitemplate"
	"github.com/888go/gin"
)

func main() {
	router := gin类.X创建默认对象()
	router.HTMLRender = loadTemplates("./templates")
	router.X绑定GET("/", func(c *gin类.Context) {
		c.X输出html模板(200, "index.html", gin类.H{
			"title": "Welcome!",
		})
	})
	router.X绑定GET("/article", func(c *gin类.Context) {
		c.X输出html模板(200, "article.html", gin类.H{
			"title": "Html5 Article Engine",
		})
	})

	if err := router.X监听(":8080"); err != nil {
		log.Fatal(err)
	}
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/includes/*.html")
	if err != nil {
		panic(err.Error())
	}

	// 从layouts/和includes/目录生成我们的templates映射
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}
