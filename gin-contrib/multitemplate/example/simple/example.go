package main
import (
	"log"
	
	"e.coding.net/gogit/go/gin/gin-contrib/multitemplate"
	"e.coding.net/gogit/go/gin"
	)
func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("index", "templates/base.html", "templates/index.html")
	r.AddFromFiles("article", "templates/base.html", "templates/index.html", "templates/article.html")
	return r
}

func main() {
	router := gin.Default()
	router.HTMLRender = createMyRender()
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index", gin.H{
			"title": "Html5 Template Engine",
		})
	})
	router.GET("/article", func(c *gin.Context) {
		c.HTML(200, "article", gin.H{
			"title": "Html5 Article Engine",
		})
	})

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
