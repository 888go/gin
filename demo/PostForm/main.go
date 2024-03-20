package main

import (
	"fmt"
	"github.com/888go/gin"
	"net/http"
)

// 参考链接：https://topgoer.com/gin框架/gin路由/表单参数.html
func main() {
	r := gin.Default()
	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		username := c.PostForm("username")
		password := c.PostForm("userpassword")
		// c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
// 使用c.String方法向客户端返回HTTP状态码为200的响应，并在响应体中输出格式化的字符串，内容为：username（用户名称）、password（密码）和type（类型），其中username、password、types分别替换为对应的变量值。
		c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
	})
	r.Run()
}
