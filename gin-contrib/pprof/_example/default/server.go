package main

import (
	"github.com/888go/gin/gin-contrib/pprof"
	"github.com/888go/gin"
)

func main() {
	router := gin类.X创建默认对象()
	pprof.Register(router)
	router.X监听(":8080")
}
