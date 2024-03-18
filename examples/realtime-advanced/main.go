package main

import (
	"fmt"
	"runtime"
	
	"github.com/888go/gin"
)

func main() {
	ConfigRuntime()
	StartWorkers()
	StartGin()
}

// ConfigRuntime设置操作系统线程数

// ff:

// ff:

// ff:

// ff:

// ff:

// ff:

// ff:
func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

// StartWorkers通过例程启动starsWorker

// ff:

// ff:

// ff:

// ff:

// ff:

// ff:

// ff:
func StartWorkers() {
	go statsWorker()
}

// 通过设置路由器启动web服务器

// ff:

// ff:

// ff:

// ff:

// ff:

// ff:

// ff:
func StartGin() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(rateLimit, gin.Recovery())
	router.LoadHTMLGlob("resources/*.templ.html")
	router.Static("/static", "resources/static")
	router.GET("/", index)
	router.GET("/room/:roomid", roomGET)
	router.POST("/room-post/:roomid", roomPOST)
	router.GET("/stream/:roomid", streamRoom)

	router.Run(":80")
}
