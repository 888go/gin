package main

import (
	"log"
	"net/http"
	"time"
	
	"github.com/888go/gin/gin-contrib/timeout"
	"github.com/888go/gin"
)

func testResponse(c *gin类.Context) {
	c.X输出文本(http.StatusRequestTimeout, "timeout")
}

func timeoutMiddleware() gin类.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(500*time.Millisecond),
		timeout.WithHandler(func(c *gin类.Context) {
			c.X中间件继续()
		}),
		timeout.WithResponse(testResponse),
	)
}

func main() {
	r := gin类.X创建()
	r.X中间件(timeoutMiddleware())
	r.X绑定GET("/slow", func(c *gin类.Context) {
		time.Sleep(800 * time.Millisecond)
		c.X设置状态码(http.StatusOK)
	})
	if err := r.X监听(":8080"); err != nil {
		log.Fatal(err)
	}
}
