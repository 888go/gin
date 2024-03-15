package main

import (
	"log"
	
	"github.com/888go/gin/gin-contrib/secure"
	"github.com/888go/gin"
)

func main() {
	router := gin.Default()

	securityConfig := secure.DefaultConfig()
	securityConfig.AllowedHosts = []string{"example.com", "ssl.example.com"}
	securityConfig.SSLHost = "ssl.example.com"
	router.Use(secure.New(securityConfig))

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}
