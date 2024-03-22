package main

import (
	"fmt"
	"log"
	
	"github.com/888go/gin"
)

const (
	Addr = "127.0.0.1:2003"
)

func main() {
	r := gin类.X创建默认对象()
	r.X绑定GET("/:path", func(c *gin类.Context) {
		// 在这个处理程序中，我们只是简单地向代理响应发送一些基本信息。
		req := c.X请求
		urlPath := fmt.Sprintf("http://%s%s", Addr, req.URL.Path)
		realIP := fmt.Sprintf("RemoteAddr=%s,X-Forwarded-For=%v,X-Real-Ip=%v", req.RemoteAddr, req.Header.Get("X-Forwarded-For"), req.Header.Get("X-Real-Ip"))
		c.X输出JSON(200, gin类.H{
			"path": urlPath,
			"ip":   realIP,
		})
	})

	if err := r.X监听(Addr); err != nil {
		log.Printf("Error: %v", err)
	}
}
