package main

import (
	"log"
	
	"github.com/888go/gin/gin-contrib/rollbar"
	"github.com/888go/gin"
	
	roll "github.com/rollbar/rollbar-go"
)

func main() {
	roll.SetToken("MY_TOKEN")
	// roll.SetEnvironment("production") // defaults to "development"

	r := gin类.X创建默认对象()
	r.X中间件(rollbar.Recovery(true))

	if err := r.X监听(":8080"); err != nil {
		log.Fatal(err)
	}
}
