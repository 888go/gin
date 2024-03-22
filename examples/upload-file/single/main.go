package main

import (
	"net/http"
	"path/filepath"
	
	"github.com/888go/gin"
)

func main() {
	router := gin类.X创建默认对象()
	// 设置multipart表单的较低内存限制（默认为32 MiB）
	router.X最大Multipart内存 = 8 << 20 // 8 MiB
	router.X绑定静态文件目录("/", "./public")
	router.X绑定POST("/upload", func(c *gin类.Context) {
		name := c.X取表单参数值("name")
		email := c.X取表单参数值("email")

		// Source
		file, err := c.X取表单上传文件("file")
		if err != nil {
			c.X输出文本(http.StatusBadRequest, "get form err: %s", err.Error())
			return
		}

		filename := filepath.Base(file.Filename)
		if err := c.X保存上传文件(file, filename); err != nil {
			c.X输出文本(http.StatusBadRequest, "upload file err: %s", err.Error())
			return
		}

		c.X输出文本(http.StatusOK, "File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, name, email)
	})
	router.X监听(":8080")
}
