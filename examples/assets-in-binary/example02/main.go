package main

import (
	"embed"
	"html/template"
	"net/http"
	
	"github.com/888go/gin"
)

//go:embed assets/* templates/*
var f embed.FS

func main() {
	router := gin类.X创建默认对象()
	templ := template.Must(template.New("").ParseFS(f, "templates/*.tmpl", "templates/foo/*.tmpl"))
	router.X设置Template模板(templ)

	// example: /public/assets/images/example.png
	router.X绑定静态文件目录FS("/public", http.FS(f))

	router.X绑定GET("/", func(c *gin类.Context) {
		c.X输出html模板(http.StatusOK, "index.tmpl", gin类.H{
			"title": "Main website",
		})
	})

	router.X绑定GET("/foo", func(c *gin类.Context) {
		c.X输出html模板(http.StatusOK, "bar.tmpl", gin类.H{
			"title": "Foo website",
		})
	})

	router.X绑定GET("favicon.ico", func(c *gin类.Context) {
		file, _ := f.ReadFile("assets/favicon.ico")
		c.X输出字节集(
			http.StatusOK,
			"image/x-icon",
			file,
		)
	})

	router.X监听(":8080")
}
