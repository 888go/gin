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
func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

// StartWorkers 通过 goroutine 启动 starsWorker。
func StartWorkers() {
	go statsWorker()
}

// StartGin 通过设置路由启动 Gin Web 服务器。
func StartGin() {
	gin类.X设置运行模式(gin类.X常量_运行模式_发布)

	router := gin类.X创建()
	router.X中间件(rateLimit, gin类.Recovery())
	router.X加载HTML模板目录("resources/*.templ.html")
	router.X绑定静态文件目录("/static", "resources/static")
	router.X绑定GET("/", index)
	router.X绑定GET("/room/:roomid", roomGET)
	router.X绑定POST("/room-post/:roomid", roomPOST)
	router.X绑定GET("/stream/:roomid", streamRoom)

	router.X监听(":80")
}
