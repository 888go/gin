package main

import (
	"fmt"
	"github.com/888go/gin"
	"net/http"
)

// gin的helloWorld

func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin类.X创建默认对象()
	// 限制表单上传大小 8MB，默认为32MB
	r.X最大Multipart内存 = 8 << 20
	r.X绑定POST("/upload", func(c *gin类.Context) {
		form, err := c.X取表单multipart对象()
		if err != nil {
			c.X输出文本(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
		}
		// 获取所有图片
		files := form.File["files"]
		// 遍历所有图片
		for _, file := range files {
			// 逐个存
			if err := c.X保存上传文件(file, ".\\demo\\MultipartForm\\"+file.Filename); err != nil {
				c.X输出文本(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
				return
			}
		}
		c.X输出文本(200, fmt.Sprintf("upload ok %d files", len(files)))
	})
	//默认端口号是8080
	r.X监听(":8000")
}
