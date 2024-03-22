package main

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	
	"github.com/888go/gin"
)

type BindFile struct {
	Name  string                `form:"name" binding:"required"`
	Email string                `form:"email" binding:"required"`
	File  *multipart.FileHeader `form:"file" binding:"required"`
}

func main() {
	router := gin类.X创建默认对象()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.X最大Multipart内存 = 8 << 20 // 8 MiB
	router.X绑定静态文件目录("/", "./public")
	router.X绑定POST("/upload", func(c *gin类.Context) {
		var bindFile BindFile

		// Bind file
		if err := c.X取参数到指针(&bindFile); err != nil {
			c.X输出文本(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
			return
		}

		// Save uploaded file
		file := bindFile.File
		dst := filepath.Base(file.Filename)
		if err := c.X保存上传文件(file, dst); err != nil {
			c.X输出文本(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		c.X输出文本(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, bindFile.Name, bindFile.Email))
	})
	router.X监听(":8080")
}
