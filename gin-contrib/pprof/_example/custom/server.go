package main
import (
	"net/http"
	
	"e.coding.net/gogit/go/gin/gin-contrib/pprof"
	"e.coding.net/gogit/go/gin"
	)
func main() {
	router := gin.Default()
	adminGroup := router.Group("/admin", func(c *gin.Context) {
		if c.Request.Header.Get("Authorization") != "foobar" {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	})
	pprof.RouteRegister(adminGroup, "pprof")
	router.Run(":8080")
}
