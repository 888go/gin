# # sentry

[![构建状态](https://travis-ci.org/gin-contrib/sentry.svg?branch=master)](https://travis-ci.org/gin-contrib/sentry)
[![Go 代码质量报告](https://goreportcard.com/badge/github.com/gin-contrib/sentry)](https://goreportcard.com/report/github.com/gin-contrib/sentry)
[![GoDoc](https://godoc.org/github.com/gin-contrib/sentry?status.svg)](https://godoc.org/github.com/gin-contrib/sentry)
[![加入聊天室](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin)

---

> `sentry` 中间件已停止维护，并已被 `sentry-go` SDK 替代。
> 在 [GitHub](https://github.com/getsentry/sentry-go) 上了解更多该项目信息，并查看新的 [gin 中间件](https://github.com/getsentry/sentry-go/tree/master/gin)。
## # 示例

参见 [example](example/main.go)

[embedmd]:
# # 以下是翻译后的内容：

```go
package main

import (
	"github.com/getsentry/raven-go"
	"github.com/gin-contrib/sentry"
	"github.com/gin-gonic/gin"
)

func init() {
	raven.SetDSN("https://<key>:<secret>@app.getsentry.com/<project>")
}

func main() {
	r := gin.Default()
	r.Use(sentry.Recovery(raven.DefaultClient, false)) // 仅发送崩溃报告
	// r.Use(sentry.Recovery(raven.DefaultClient, true)) // 发送崩溃报告并启用panic捕获
	r.Run(":8080")
}
```

注意：在实际使用时，请将`<key>:<secret>`和`<project>`替换为您的Sentry项目对应的密钥、密钥 secret 和项目ID。
