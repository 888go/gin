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
	// Non-cached Page
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// Cached Page
	r.GET("/cache_ping", cache.CachePage(store, time.Minute, func(c *gin.Context) {
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
	}))

	// Listen and Server in 0.0.0.0:8080
	_ = r.Run(":8080")
}
