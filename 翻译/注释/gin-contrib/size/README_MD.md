
<原文开始>
size

[![Run Tests](https://github.com/gin-contrib/size/actions/workflows/go.yml/badge.svg)](https://github.com/gin-contrib/size/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/size/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/size)
[![Go Report Card](https://goreportcard.com/badge/github.com/gin-contrib/size)](https://goreportcard.com/report/github.com/gin-contrib/size)
[![GoDoc](https://godoc.org/github.com/gin-contrib/size?status.svg)](https://godoc.org/github.com/gin-contrib/size)
[![Join the chat at https://gitter.im/gin-gonic/gin](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin)

Limit size of POST requests for Gin framework


<原文结束>

# <翻译开始>
# 大小

[![运行测试](https://github.com/gin-contrib/size/actions/workflows/go.yml/badge.svg)](https://github.com/gin-contrib/size/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/size/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/size)
[![Go 代码质量报告](https://goreportcard.com/badge/github.com/gin-contrib/size)](https://goreportcard.com/report/github.com/gin-contrib/size)
[![GoDoc](https://godoc.org/github.com/gin-contrib/size?status.svg)](https://godoc.org/github.com/gin-contrib/size)
[![加入Gitter聊天](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin)

限制 Gin 框架中 POST 请求的大小

# <翻译结束>


<原文开始>
Example

```go
package main

import (
  "net/http"

  limits "github.com/gin-contrib/size"
  "github.com/gin-gonic/gin"
)

func handler(ctx *gin.Context) {
  val := ctx.PostForm("b")
  if len(ctx.Errors) > 0 {
    return
  }
  ctx.String(http.StatusOK, "got %s\n", val)
}

func main() {
  r := gin.Default()
  r.Use(limits.RequestSizeLimiter(10))
  r.POST("/", handler)
  if err := r.Run(":8080"); err != nil {
    log.Fatal(err)
  }
}
```

<原文结束>

# <翻译开始>
# 示例

```go
package main

import (
  "net/http"
  "log"

  "github.com/gin-contrib/size/limits"
  "github.com/gin-gonic/gin"
)

func handler(ctx *gin.Context) {
  val := ctx.PostForm("b")
  if len(ctx.Errors) > 0 {
    return
  }
  ctx.String(http.StatusOK, "接收到 %s\n", val)
}

func main() {
  r := gin.Default()
  r.Use(limits.RequestSizeLimiter(10)) // 设置请求大小限制为10字节
  r.POST("/", handler)
  if err := r.Run(":8080"); err != nil {
    log.Fatal(err) // 如果运行服务器时出现错误，则输出错误并终止程序
  }
}
```

# <翻译结束>

