package main
import (
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	
	"e.coding.net/gogit/go/gin"
	)
type BindFile struct {
	Name  string                `form:"name" binding:"required"`
	Email string                `form:"email" binding:"required"`
	File  *multipart.FileHeader `form:"file" binding:"required"`
}

func main() {
	router := gin.Default()
// 为多部分表单设置较低的内存限制(默认为32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Static("/", "./public")
	router.POST("/upload", func(c *gin.Context) {
		var bindFile BindFile

// 绑定文件
		if err := c.ShouldBind(&bindFile); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
			return
		}

// 保存上传的文件
		file := bindFile.File
		dst := filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, bindFile.Name, bindFile.Email))
	})
	router.Run(":8080")
}
