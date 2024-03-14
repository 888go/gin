package main
import (
	"log"
	
	"e.coding.net/gogit/go/gin/gin-contrib/expvar"
	"e.coding.net/gogit/go/gin"
	)
func main() {
	r := gin.Default()

	r.GET("/debug/vars", expvar.Handler())

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
