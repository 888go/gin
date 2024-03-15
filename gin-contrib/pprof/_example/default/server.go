package main

import (
	"github.com/888go/gin/gin-contrib/pprof"
	"github.com/888go/gin"
)

func main() {
	router := gin.Default()
	pprof.Register(router)
	router.Run(":8080")
}
