package main
import (
	"github.com/getsentry/raven-go"
	"e.coding.net/gogit/go/gin/gin-contrib/sentry"
	"e.coding.net/gogit/go/gin"
	)
func init() {
	raven.SetDSN("https://<key>:<secret>@app.getsentry.com/<project>")
}

func main() {
	r := gin.Default()
	r.Use(sentry.Recovery(raven.DefaultClient, false))
	// only send crash reporting
	// r.Use(sentry.Recovery(raven.DefaultClient, true))
	r.Run(":8080")
}
