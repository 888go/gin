package main

import (
	"log"
	
	"github.com/888go/gin/gin-contrib/expvar"
	"github.com/888go/gin"
)

func main() {
	r := gin类.X创建默认对象()

	r.X绑定GET("/debug/vars", expvar.Handler())

	if err := r.X监听(":8080"); err != nil {
		log.Fatal(err)
	}
}
