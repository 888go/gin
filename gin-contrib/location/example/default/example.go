package main

import (
	"log"
	"net/http"
	
	"github.com/888go/gin/gin-contrib/location"
	"github.com/888go/gin"
)

func main() {
	router := gin.Default()

	// configure to automatically detect scheme and host
	// - use http when default scheme cannot be determined
	// - use localhost:8080 when default host cannot be determined
	router.Use(location.Default())

	router.GET("/", func(c *gin.Context) {
		url := location.Get(c)
		c.String(http.StatusOK, url.String())
	})

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}
