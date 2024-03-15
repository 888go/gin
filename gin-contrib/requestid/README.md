# # 请求ID

[![运行测试](https://github.com/gin-contrib/requestid/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/gin-contrib/requestid/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/requestid/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/requestid)
[![Go 代码质量报告](https://goreportcard.com/badge/github.com/gin-contrib/requestid)](https://goreportcard.com/report/github.com/gin-contrib/requestid)
[![GoDoc](https://godoc.org/github.com/gin-contrib/requestid?status.svg)](https://godoc.org/github.com/gin-contrib/requestid)
[![加入Gitter聊天](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin)

这是 Gin 框架的请求ID中间件。通过 `X-Request-ID` 头部为响应添加一个标识符。如果请求头部中包含 `X-Request-ID` 值，则将其传递回调用者。
## Usage

### # 开始使用

下载并安装：

```sh
go get github.com/gin-contrib/requestid
```

在你的代码中导入并使用它：

```go
import "github.com/gin-contrib/requestid"
```
## # 配置

定义你的自定义生成器函数：

```go
func main() {

  // 创建一个新的Gin实例
  r := gin.New()

  // 使用requestid中间件，传入自定义生成器和自定义Header键
  r.Use(
    requestid.New(
      requestid.WithGenerator(func() string {
        return "test"
      }),
      requestid.WithCustomHeaderStrKey("your-customer-key"),
    ),
  )

  // 示例ping请求处理
  r.GET("/ping", func(c *gin.Context) {
    c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
  })

  // 在0.0.0.0:8080监听并启动服务
  r.Run(":8080")
}
```

这段代码是用Go语言编写的，主要内容是创建一个基于Gin框架的HTTP服务器，并使用名为`requestid`的中间件。其中，通过`WithGenerator`参数设置了一个自定义的请求ID生成器函数，该函数始终返回字符串"test"；通过`WithCustomHeaderStrKey`参数设置了将请求ID放置到HTTP响应头中的自定义键名。此外，还定义了一个简单的GET请求处理器，用于处理"/ping"路径的请求，返回包含当前时间戳的"pong"响应。最后，服务器在0.0.0.0:8080端口启动并开始监听请求。
## # 示例

```go
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func main() {

	// 创建一个新的gin实例
	r := gin.New()

	// 使用requestid中间件
	r.Use(requestid.New())

	// 示例ping请求
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// 在0.0.0.0:8080监听并启动服务
	r.Run(":8080")
}

// 如何获取请求标识符：

// 示例根路径请求
r.GET("/", func(c *gin.Context) {
	// 从上下文中获取请求ID
	c.String(http.StatusOK, "id:"+requestid.Get(c))
})
```

这段代码是用Go语言编写的，它创建了一个基于Gin框架的HTTP服务器。在主函数中，首先初始化了一个新的Gin引擎，并添加了`requestid`中间件以自动为每个请求生成一个唯一标识符。

然后定义了一个处理"/ping" GET请求的路由，该路由会返回当前时间戳和字符串"pong"。

最后，服务器在0.0.0.0:8080端口上运行。

另外还展示了如何在处理"/" GET请求时，从请求上下文中获取并返回由`requestid`中间件生成的请求标识符。
