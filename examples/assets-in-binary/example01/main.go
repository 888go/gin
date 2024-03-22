package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
	
	"github.com/888go/gin"
)

func main() {
	r := gin类.X创建()
	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	r.X设置Template模板(t)
	r.X绑定GET("/", func(c *gin类.Context) {
		c.X输出html模板(http.StatusOK, "/html/index.tmpl", gin类.H{
			"Foo": "World",
		})
	})
	r.X绑定GET("/bar", func(c *gin类.Context) {
		c.X输出html模板(http.StatusOK, "/html/bar.tmpl", gin类.H{
			"Bar": "World",
		})
	})
	r.X监听(":8080")
}

func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".tmpl") {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
