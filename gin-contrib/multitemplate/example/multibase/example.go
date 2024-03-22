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
	router.X绑定GET("/admin", func(c *gin类.Context) {
		c.X输出html模板(200, "admin.html", gin类.H{
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

	articleLayouts, err := filepath.Glob(templatesDir + "/layouts/article-base.html")
	if err != nil {
		panic(err.Error())
	}

	articles, err := filepath.Glob(templatesDir + "/articles/*.html")
	if err != nil {
		panic(err.Error())
	}

	// 从 articleLayouts/ 和 articles/ 目录生成我们的 templates 映射
	for _, article := range articles {
		layoutCopy := make([]string, len(articleLayouts))
		copy(layoutCopy, articleLayouts)
		files := append(layoutCopy, article)
		r.AddFromFiles(filepath.Base(article), files...)
	}

	adminLayouts, err := filepath.Glob(templatesDir + "/layouts/admin-base.html")
	if err != nil {
		panic(err.Error())
	}

	admins, err := filepath.Glob(templatesDir + "/admins/*.html")
	if err != nil {
		panic(err.Error())
	}

	// 从adminLayouts/和admins/目录生成我们的templates映射
	for _, admin := range admins {
		layoutCopy := make([]string, len(adminLayouts))
		copy(layoutCopy, adminLayouts)
		files := append(layoutCopy, admin)
		r.AddFromFiles(filepath.Base(admin), files...)
	}
	return r
}
