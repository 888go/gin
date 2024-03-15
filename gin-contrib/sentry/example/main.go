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
// （翻译后）
// 只启用崩溃报告发送功能
// r.Use(sentry.Recovery(raven.DefaultClient, true))
// 此处代码的含义是：在处理HTTP请求时，使用sentry库提供的Recovery中间件，并通过raven.DefaultClient（默认的Raven客户端）来捕获和发送错误信息。参数true表示在发生panic时，不仅记录并恢复panic，还会向Sentry服务端发送崩溃报告。
	r.Run(":8080")
}
