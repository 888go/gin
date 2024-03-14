package main
import (
	"log"
	
	"github.com/gin-gonic/autotls"
	"e.coding.net/gogit/go/gin"
	"golang.org/x/crypto/acme/autocert"
	)
func main() {
	r := gin.Default()

// 平处理程序
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("example1.com", "example2.com"),
		Cache:      autocert.DirCache("/var/www/.cache"),
	}

	log.Fatal(autotls.RunWithManager(r, &m))
}
