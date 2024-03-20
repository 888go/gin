package main

import (
	"log"
	"net/http"
	"time"
	
	"github.com/888go/gin/gin-contrib/timeout"
	"github.com/888go/gin"
)

func emptySuccessResponse(c *gin.Context) {
	time.Sleep(200 * time.Microsecond)
	c.String(http.StatusOK, "")
}

func main() {
	r := gin.New()

	r.GET("/", timeout.New(
		timeout.WithTimeout(100*time.Microsecond),
		timeout.WithHandler(emptySuccessResponse),
	))

	// 在0.0.0.0:8080监听并服务
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
