package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	
	"github.com/888go/gin/gin-contrib/requestid"
	"github.com/888go/gin"
)

func main() {
	r := gin.New()

	r.Use(
		requestid.New(
			requestid.WithGenerator(func() string {
				return "test"
			}),
			requestid.WithCustomHeaderStrKey("your-customer-key"),
			requestid.WithHandler(func(c *gin.Context, requestID string) {
				log.Printf("RequestID: %s", requestID)
			}),
		),
	)

	// Example ping request.
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// Example / request.
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "id:"+requestid.Get(c))
	})

	// 在0.0.0.0:8080监听并服务
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
