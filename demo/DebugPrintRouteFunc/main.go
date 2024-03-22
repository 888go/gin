package main

import (
	"github.com/888go/gin"
	"log"
	"net/http"
)

// 参考 Gin 框架官方文档（中文版）：
// https://gin-gonic.com/zh-cn/docs/examples/define-format-for-the-log-of-routes/
// 此代码行是引用Gin框架的官方文档链接，用于定义路由日志格式的示例。由于该行本身已经是注释，无需额外翻译或解释，直接提供对应的中文链接即可。
func main() {
	r := gin类.X创建默认对象()
	gin类.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	r.X绑定POST("/foo", func(c *gin类.Context) {
		c.X输出JSON(http.StatusOK, "foo")
	})

	r.X绑定GET("/bar", func(c *gin类.Context) {
		c.X输出JSON(http.StatusOK, "bar")
	})

	r.X绑定GET("/status", func(c *gin类.Context) {
		c.X输出JSON(http.StatusOK, "ok")
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.X监听()
}
