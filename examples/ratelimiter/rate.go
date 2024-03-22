package main

import (
	"flag"
	"log"
	"time"
	
	"github.com/fatih/color"
	"github.com/888go/gin"
	"go.uber.org/ratelimit"
)

var (
	limit ratelimit.Limiter
	rps   = flag.Int("rps", 100, "request per second")
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("[GIN] ")
	log.SetOutput(gin类.DefaultWriter)
}

func leakBucket() gin类.HandlerFunc {
	prev := time.Now()
	return func(ctx *gin类.Context) {
		now := limit.Take()
		log.Print(color.CyanString("%v", now.Sub(prev)))
		prev = now
	}
}

func ginRun(rps int) {
	limit = ratelimit.New(rps)

	app := gin类.X创建默认对象()
	app.X中间件(leakBucket())

	app.X绑定GET("/rate", func(ctx *gin类.Context) {
		ctx.X输出JSON(200, "rate limiting test")
	})

	log.Printf(color.CyanString("Current Rate Limit: %v requests/s", rps))
	app.X监听(":8080")
}

func main() {
	flag.Parse()
	ginRun(*rps)
}
