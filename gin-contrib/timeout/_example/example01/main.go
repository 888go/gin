package main

import (
	"log"
	"net/http"
	"time"
	
	"github.com/888go/gin/gin-contrib/timeout"
	"github.com/888go/gin"
)

func emptySuccessResponse(c *gin类.Context) {
	time.Sleep(200 * time.Microsecond)
	c.X输出文本(http.StatusOK, "")
}

func main() {
	r := gin类.X创建()

	r.X绑定GET("/", timeout.New(
		timeout.WithTimeout(100*time.Microsecond),
		timeout.WithHandler(emptySuccessResponse),
	))

	// 在0.0.0.0:8080监听并服务
	if err := r.X监听(":8080"); err != nil {
		log.Fatal(err)
	}
}
