package main
import (
	"log"
	
	"e.coding.net/gogit/go/gin/gin-contrib/rollbar"
	"e.coding.net/gogit/go/gin"
	
	roll "github.com/rollbar/rollbar-go"
	)
func main() {
	roll.SetToken("MY_TOKEN")
	// roll.SetEnvironment("production") // defaults to "development"

	r := gin.Default()
	r.Use(rollbar.Recovery(true))

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
