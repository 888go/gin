
<原文开始>
RequestID

[![Run Tests](https://github.com/gin-contrib/requestid/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/gin-contrib/requestid/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/requestid/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/requestid)
[![Go Report Card](https://goreportcard.com/badge/github.com/gin-contrib/requestid)](https://goreportcard.com/report/github.com/gin-contrib/requestid)
[![GoDoc](https://godoc.org/github.com/gin-contrib/requestid?status.svg)](https://godoc.org/github.com/gin-contrib/requestid)
[![Join the chat at https://gitter.im/gin-gonic/gin](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin)

Request ID middleware for Gin Framework. Adds an indentifier to the response using the `X-Request-ID` header. Passes the `X-Request-ID` value back to the caller if it's sent in the request headers.


<原文结束>

# <翻译开始>
# 请求ID

[![运行测试](https://github.com/gin-contrib/requestid/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/gin-contrib/requestid/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/requestid/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/requestid)
[![Go 代码质量报告](https://goreportcard.com/badge/github.com/gin-contrib/requestid)](https://goreportcard.com/report/github.com/gin-contrib/requestid)
[![GoDoc](https://godoc.org/github.com/gin-contrib/requestid?status.svg)](https://godoc.org/github.com/gin-contrib/requestid)
[![加入聊天 https://gitter.im/gin-gonic/gin](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin)

这是一个为 Gin 框架设计的请求 ID 中间件。它通过 `X-Request-ID` 头部添加一个标识符到响应中。如果请求头部包含了 `X-Request-ID` 值，该中间件会将其传递回调用者。

# <翻译结束>


<原文开始>
Start using it

Download and install it.

```sh
go get github.com/gin-contrib/requestid
```

Import it in your code, then use it:

```go
import "github.com/gin-contrib/requestid"
```


<原文结束>

# <翻译开始>
# 开始使用

下载并安装：

```sh
go get github.com/gin-contrib/requestid
```

在代码中导入并使用：

```go
import "github.com/gin-contrib/requestid"
```

# <翻译结束>


<原文开始>
Config

define your custom generator function:

```go
func main() {

  r := gin.New()

  r.Use(
    requestid.New(
      requestid.WithGenerator(func() string {
        return "test"
      }),
      requestid.WithCustomHeaderStrKey("your-customer-key"),
    ),
  )

  // Example ping request.
  r.GET("/ping", func(c *gin.Context) {
    c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
  })

  // Listen and Server in 0.0.0.0:8080
  r.Run(":8080")
}
```


<原文结束>

# <翻译开始>
# 配置

定义你的自定义生成器函数：

```go
func main() {

// 创建一个gin实例
  r := gin.New()

// 使用requestid中间件，其中包含自定义生成器和自定义头键
  r.Use(
    requestid.New(
// 自定义请求ID生成器函数
      requestid.WithGenerator(func() string {
        return "test"
      }),
// 设置自定义请求头键
      requestid.WithCustomHeaderStrKey("your-customer-key"),
    ),
  )

// 示例ping请求处理
  r.GET("/ping", func(c *gin.Context) {
// 返回HTTP状态码200及当前时间戳的pong响应
    c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
  })

// 监听并在0.0.0.0:8080端口启动服务器
  r.Run(":8080")
}
```

这段代码是用Go语言编写的，主要功能是设置一个基于Gin框架的Web服务器。在服务器中使用了一个名为`requestid`的中间件，该中间件用于生成并添加请求ID到HTTP请求的头部信息。其中，通过`WithGenerator`方法设置了自定义的请求ID生成器，其返回值固定为"test"；通过`WithCustomHeaderStrKey`方法指定了请求ID将被添加到HTTP头部的键名。同时，该示例还提供了一个简单的/ping API接口，当接收到GET请求时，返回包含当前时间戳的"pong"响应。最后，服务器在0.0.0.0:8080端口运行并监听请求。

# <翻译结束>


<原文开始>
Example

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

  r := gin.New()

  r.Use(requestid.New())

  // Example ping request.
  r.GET("/ping", func(c *gin.Context) {
    c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
  })

  // Listen and Server in 0.0.0.0:8080
  r.Run(":8080")
}
```

How to get the request identifier:

```go
// Example / request.
r.GET("/", func(c *gin.Context) {
  c.String(http.StatusOK, "id:"+requestid.Get(c))
})
```

<原文结束>

# <翻译开始>
# 示例

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

// 示例ping请求处理
  r.GET("/ping", func(c *gin.Context) {
// 返回当前时间戳的pong响应
    c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
  })

// 监听并在0.0.0.0:8080端口上提供服务
  r.Run(":8080")
}
```

如何获取请求标识符：

```go
// 示例根路径（"/"）请求处理
r.GET("/", func(c *gin.Context) {
// 返回包含请求ID的消息
  c.String(http.StatusOK, "id:"+requestid.Get(c))
})
```

# <翻译结束>

