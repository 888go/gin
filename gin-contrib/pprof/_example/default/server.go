package main
import (
	"e.coding.net/gogit/go/gin/gin-contrib/pprof"
	"e.coding.net/gogit/go/gin"
	)
func main() {
	router := gin.Default()
	pprof.Register(router)
	router.Run(":8080")
}
