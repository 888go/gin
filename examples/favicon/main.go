package main
import (
	"net/http"
	
	"e.coding.net/gogit/go/gin"
	)
func main() {
	app := gin.Default()

// 从相对于主的位置提供静态图标文件
// go directory app.StaticFile("/favicon.ico"， "./.assets/favicon.ico")
	app.StaticFile("/favicon.ico", "./favicon.ico")

	app.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello favicon.")
	})

	app.Run(":8080")
}
