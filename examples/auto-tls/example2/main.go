package main

import (
	"log"
	
	"github.com/gin-gonic/autotls"
	"github.com/888go/gin"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	r := gin类.X创建默认对象()

	// Ping handler
	r.X绑定GET("/ping", func(c *gin类.Context) {
		c.X输出文本(200, "pong")
	})

	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("example1.com", "example2.com"),
		Cache:      autocert.DirCache("/var/www/.cache"),
	}

	log.Fatal(autotls.RunWithManager(r, &m))
}
