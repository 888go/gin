package routes
import (
	"net/http"
	
	"e.coding.net/gogit/go/gin"
	)
func addUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users")
	})
	users.GET("/comments", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users comments")
	})
	users.GET("/pictures", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users pictures")
	})
}
