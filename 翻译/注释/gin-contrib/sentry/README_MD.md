
<原文开始>
sentry

[![Build Status](https://travis-ci.org/gin-contrib/sentry.svg?branch=master)](https://travis-ci.org/gin-contrib/sentry)
[![Go Report Card](https://goreportcard.com/badge/github.com/gin-contrib/sentry)](https://goreportcard.com/report/github.com/gin-contrib/sentry)
[![GoDoc](https://godoc.org/github.com/gin-contrib/sentry?status.svg)](https://godoc.org/github.com/gin-contrib/sentry)
[![Join the chat at https://gitter.im/gin-gonic/gin](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin)

---

> The `sentry` middleware is no longer maintained and was superseded by the `sentry-go` SDK.
> Learn more about the project on [GitHub](https://github.com/getsentry/sentry-go) and check out the new [gin middleware](https://github.com/getsentry/sentry-go/tree/master/gin).

---


<原文结束>

# <翻译开始>
# sentry

[![构建状态](https://travis-ci.org/gin-contrib/sentry.svg?branch=master)](https://travis-ci.org/gin-contrib/sentry)
[![Go 代码质量报告](https://goreportcard.com/badge/github.com/gin-contrib/sentry)](https://goreportcard.com/report/github.com/gin-contrib/sentry)
[![GoDoc 文档](https://godoc.org/github.com/gin-contrib/sentry?status.svg)](https://godoc.org/github.com/gin-contrib/sentry)
[![加入聊天室](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin)

---

> 中间件`sentry`已停止维护，并被`sentry-go` SDK取代。
> 您可以在[GitHub](https://github.com/getsentry/sentry-go)上了解更多关于该项目的信息，并查看新的[gin中间件](https://github.com/getsentry/sentry-go/tree/master/gin)。

# <翻译结束>


<原文开始>
Example

See the [example](example/main.go)

[embedmd]:
<原文结束>

# <翻译开始>
# 示例

请参阅 [example](example/main.go)

[embedmd]:

# <翻译结束>


<原文开始>
(example/main.go go)
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
	r.Use(sentry.Recovery(raven.DefaultClient, false))
	// only send crash reporting
	// r.Use(sentry.Recovery(raven.DefaultClient, true))
	r.Run(":8080")
}
```

<原文结束>

# <翻译开始>
# ```go
// (example/main.go Go 代码文件)

package main

import (
    "github.com/getsentry/raven-go" // 引入raven错误报告库
    "github.com/gin-contrib/sentry"  // 引入gin框架的sentry中间件
    "github.com/gin-gonic/gin"       // 引入gin框架
)

func init() {
    raven.SetDSN("https://<key>:<secret>@app.getsentry.com/<project>") // 设置Sentry DSN（数据源名称）
}

func main() {
    r := gin.Default() // 初始化gin引擎

// 使用sentry中间件进行恢复处理，不发送非崩溃报告
    r.Use(sentry.Recovery(raven.DefaultClient, false))

// 若要同时发送崩溃报告和非崩溃报告，可以使用如下代码：
// r.Use(sentry.Recovery(raven.DefaultClient, true))

    r.Run(":8080") // 启动gin服务在8080端口
}
```

这段Go代码示例展示了如何在gin框架中集成sentry错误报告服务。首先通过`raven-go`库设置Sentry的数据源名称（DSN），然后在gin引擎中使用`sentry`中间件进行异常恢复处理，并配置是否发送非崩溃报告。最后启动服务器在8080端口监听请求。其中`<key>:<secret>`和`<project>`需要替换为实际的Sentry密钥、密钥秘密和项目ID。

# <翻译结束>

