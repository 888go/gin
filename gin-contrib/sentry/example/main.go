package main

import (
	"github.com/getsentry/raven-go"
	"github.com/888go/gin/gin-contrib/sentry"
	"github.com/888go/gin"
)

func init() {
	raven.SetDSN("https://<key>:<secret>@app.getsentry.com/<project>")
}

func main() {
	r := gin类.X创建默认对象()
	r.X中间件(sentry.Recovery(raven.DefaultClient, false))
	// only send crash reporting
	// r.Use(sentry.Recovery(raven.DefaultClient, true))
	r.X监听(":8080")
}
