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

// ConfigRuntime 设置操作系统的线程数量。

// ff:
func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

// StartWorkers 通过 goroutine 启动 starsWorker。

// ff:
func StartWorkers() {
	go statsWorker()
}

// StartGin 通过设置路由启动 Gin Web 服务器。

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
