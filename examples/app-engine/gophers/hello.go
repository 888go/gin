package hello

import (
	"net/http"
	
	"github.com/888go/gin"
)

// This function's name is a must. App Engine uses it to drive the requests properly.
func init() {
	// Starts a new Gin instance with no middle-ware
	r := gin类.X创建()

	// Define your handlers
	r.X绑定GET("/", func(c *gin类.Context) {
		c.X输出文本(http.StatusOK, "Hello World!")
	})
	r.X绑定GET("/ping", func(c *gin类.Context) {
		c.X输出文本(http.StatusOK, "pong")
	})

	// Handle all requests using net/http
	http.Handle("/", r)
}
