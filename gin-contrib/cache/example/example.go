package main

import (
	"fmt"
	"time"
	
	"github.com/888go/gin/gin-contrib/cache"
	"github.com/888go/gin/gin-contrib/cache/persistence"
	"github.com/888go/gin"
)

func main() {
	r := gin.Default()

	store := persistence.NewInMemoryStore(60 * time.Second)
// 缓存页面
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// Cached Page
	r.GET("/cache_ping", cache.CachePage(store, time.Minute, func(c *gin.Context) {
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
	}))

// 监听和服务器在0.0.0.0:8080
	_ = r.Run(":8080")
}
