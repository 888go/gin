package main

import (
	"github.com/888go/gin"
	"net/http"
)

// 访问 gin-gonic.com/zh-cn/docs/examples/serving-static-files/
// 以获取 Gin 框架关于服务静态文件的中文文档示例

func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	// 监听并在 0.0.0.0:8080 上启动服务
	router.Run(":8080")
}
