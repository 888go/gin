package main

import (
	"log"
	"net/http"
	
	limits "github.com/888go/gin/gin-contrib/size"
	"github.com/888go/gin"
)

func handler(ctx *gin类.Context) {
	val := ctx.X取表单参数值("b")
	if len(ctx.X错误s) > 0 {
		return
	}
	ctx.X输出文本(http.StatusOK, "got %s\n", val)
}

func main() {
	r := gin类.X创建默认对象()
	r.X中间件(limits.RequestSizeLimiter(10))
	r.X绑定POST("/", handler)
	if err := r.X监听(":8080"); err != nil {
		log.Fatal(err)
	}
}
