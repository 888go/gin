package main
import (
	"log"
	"net/http"
	
	"e.coding.net/gogit/go/gin/gin-contrib/location"
	"e.coding.net/gogit/go/gin"
	)
func main() {
	router := gin.Default()

	// configure to automatically detect scheme and host with
	// fallback to https://foo.com/base
	// - use https when default scheme cannot be determined
	// - use foo.com when default host cannot be determined
	// - include /base as the path
	router.Use(location.New(location.Config{
		Scheme:  "https",
		Host:    "foo.com",
		Base:    "/base",
		Headers: location.Headers{Scheme: "X-Forwarded-Proto", Host: "X-Forwarded-For"},
	}))

	router.GET("/", func(c *gin.Context) {
		url := location.Get(c)
		c.String(http.StatusOK, url.String())
	})

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}
