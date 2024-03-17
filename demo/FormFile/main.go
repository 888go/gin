package main

import (
	"github.com/888go/gin"
	"net/http"
)

// https://topgoer.com/gin%E6%A1%86%E6%9E%B6/gin%E8%B7%AF%E7%94%B1/%E4%B8%8A%E4%BC%A0%E5%8D%95%E4%B8%AA%E6%96%87%E4%BB%B6.html
func main() {
	r := gin.Default()
	//限制上传最大尺寸
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(500, "上传图片出错")
		}
		// c.JSON(200, gin.H{"message": file.Header.Context})
		c.SaveUploadedFile(file, ".\\demo\\FormFile\\"+file.Filename)
		c.String(http.StatusOK, ".\\demo\\FormFile\\"+file.Filename)
	})
	r.Run()
}
