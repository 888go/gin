package main

import (
	"github.com/888go/gin"
	"net/http"
)

// Gin框架官方文档示例：ASCII JSON（中文版）

func main() {
	r := gin类.X创建默认对象()

	r.X绑定GET("/someJSON", func(c *gin类.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.X输出JSON并按ASCII(http.StatusOK, data)
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.X监听(":8080")
}
