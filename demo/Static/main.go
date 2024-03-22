package main

import (
	"github.com/888go/gin"
	"net/http"
)

// 访问 gin-gonic.com/zh-cn/docs/examples/serving-static-files/
// 以获取 Gin 框架关于服务静态文件的中文文档示例

func main() {
	router := gin类.X创建默认对象()
	router.X绑定静态文件目录("/assets", "./assets")
	router.X绑定静态文件目录FS("/more_static", http.Dir("my_file_system"))
	router.X绑定静态单文件("/favicon.ico", "./resources/favicon.ico")

	// 监听并在 0.0.0.0:8080 上启动服务
	router.X监听(":8080")
}
