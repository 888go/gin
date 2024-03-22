package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
	
	"github.com/888go/gin"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d%02d/%02d", year, month, day)
}

func main() {
	router := gin类.X创建默认对象()
	router.X设置模板分隔符("{[{", "}]}")
	router.X设置Template模板函数(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	router.X加载HTML模板文件("./testdata/raw.tmpl")

	router.X绑定GET("/raw", func(c *gin类.Context) {
		c.X输出html模板(http.StatusOK, "raw.tmpl", gin类.H{
			"now": time.Date(2017, 0o7, 0o1, 0, 0, 0, 0, time.UTC),
		})
	})

	router.X监听(":8080")
}
