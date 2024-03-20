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
	r := gin.Default()
	r.Use(sentry.Recovery(raven.DefaultClient, false))
// 仅发送崩溃报告
// r.Use(sentry.Recovery(raven.DefaultClient, true)) 
// 
// （翻译为：）
// 
// 只启用崩溃报告发送功能
// r.Use(sentry提供的恢复中间件，使用raven的默认客户端，并开启true参数以捕获所有panic信息）
	r.Run(":8080")
}
