package main
import (
	"fmt"
	"log"
	"net/http"
	"time"
	
	"e.coding.net/gogit/go/gin/gin-contrib/requestid"
	"e.coding.net/gogit/go/gin"
	)
func main() {
	r := gin.New()

	r.Use(requestid.New())

	// Example ping request.
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// Example / request.
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "id:"+requestid.Get(c))
	})

	// Listen and Server in 0.0.0.0:8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
