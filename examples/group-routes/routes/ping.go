package routes
import (
	"net/http"
	
	"e.coding.net/gogit/go/gin"
	)
func addPingRoutes(rg *gin.RouterGroup) {
	ping := rg.Group("/ping")

	ping.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
}
