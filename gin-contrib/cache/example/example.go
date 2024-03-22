package main

import (
	"fmt"
	"time"
	
	"github.com/888go/gin"
	"github.com/888go/gin/gin-contrib/cache"
	"github.com/888go/gin/gin-contrib/cache/persistence"
)

func main() {
	r := gin类.X创建默认对象()

	store := persistence.NewInMemoryStore(60 * time.Second)
	// Non-cached Page
	r.X绑定GET("/ping", func(c *gin类.Context) {
		c.X输出文本(200, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// Cached Page
	r.X绑定GET("/cache_ping", cache.CachePage(store, time.Minute, func(c *gin类.Context) {
		c.X输出文本(200, "pong "+fmt.Sprint(time.Now().Unix()))
	}))

	// 在0.0.0.0:8080监听并服务
	_ = r.X监听(":8080")
}
