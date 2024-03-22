package main

import (
	"github.com/888go/gin"
	"net/http"
)

// 该链接为Gin框架教程中关于单个文件上传的部分：
// https://topgoer.com/gin框架/gin路由/上传单个文件.html
// 此段Go代码中的注释仅包含一个URL链接，链接指向的是一个关于Gin框架的在线教程页面，具体讲解了如何在Gin框架中处理单个文件的上传功能。
func main() {
	r := gin类.X创建默认对象()
	//限制上传最大尺寸
	r.X最大Multipart内存 = 8 << 20
	r.X绑定POST("/upload", func(c *gin类.Context) {
		file, err := c.X取表单上传文件("file")
		if err != nil {
			c.X输出文本(500, "上传图片出错")
		}
		// c.JSON(200, gin.H{"message": file.Header.Context})
// 使用c.JSON发送HTTP响应，状态码为200，并返回一个JSON对象。其中JSON对象包含一个键值对："message"对应file.Header.Context的值。
		c.X保存上传文件(file, ".\\demo\\FormFile\\"+file.Filename)
		c.X输出文本(http.StatusOK, ".\\demo\\FormFile\\"+file.Filename)
	})
	r.X监听()
}
