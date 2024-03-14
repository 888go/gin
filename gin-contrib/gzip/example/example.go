package main
import (
	"fmt"
	"log"
	"net/http"
	"time"
	
	"e.coding.net/gogit/go/gin/gin-contrib/gzip"
	"e.coding.net/gogit/go/gin"
	)
func main() {
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// Listen and Server in 0.0.0.0:8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
