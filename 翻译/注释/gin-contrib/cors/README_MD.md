
<原文开始>
CORS gin's middleware

[![Run Tests](https://github.com/gin-contrib/cors/actions/workflows/go.yml/badge.svg)](https://github.com/gin-contrib/cors/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/cors/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/cors)
[![Go Report Card](https://goreportcard.com/badge/github.com/gin-contrib/cors)](https://goreportcard.com/report/github.com/gin-contrib/cors)
[![GoDoc](https://godoc.org/github.com/gin-contrib/cors?status.svg)](https://godoc.org/github.com/gin-contrib/cors)

Gin middleware/handler to enable CORS support.


<原文结束>

# <翻译开始>
# CORS/gin 中间件

[![运行测试](https://github.com/gin-contrib/cors/actions/workflows/go.yml/badge.svg)](https://github.com/gin-contrib/cors/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/cors/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/cors)
[![Go 项目报告卡](https://goreportcard.com/badge/github.com/gin-contrib/cors)](https://goreportcard.com/report/github.com/gin-contrib/cors)
[![GoDoc](https://godoc.org/github.com/gin-contrib/cors?status.svg)](https://godoc.org/github.com/gin-contrib/cors)

这是一个 Gin 框架的中间件/处理器，用于启用 CORS 支持。

# <翻译结束>


<原文开始>
Start using it

Download and install it:

```sh
go get github.com/gin-contrib/cors
```

Import it in your code:

```go
import "github.com/gin-contrib/cors"
```

#
<原文结束>

# <翻译开始>
# 开始使用

下载并安装：

```sh
go get github.com/gin-contrib/cors
```

在代码中引入它：

```go
import "github.com/gin-contrib/cors"
```

#

# <翻译结束>


<原文开始>
Canonical example

```go
package main

import (
  "time"

  "github.com/gin-contrib/cors"
  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  // CORS for https://foo.com and https://github.com origins, allowing:
  // - PUT and PATCH methods
  // - Origin header
  // - Credentials share
  // - Preflight requests cached for 12 hours
  router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"https://foo.com"},
    AllowMethods:     []string{"PUT", "PATCH"},
    AllowHeaders:     []string{"Origin"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
    AllowOriginFunc: func(origin string) bool {
      return origin == "https://github.com"
    },
    MaxAge: 12 * time.Hour,
  }))
  router.Run()
}
```

#
<原文结束>

# <翻译开始>
# 规范示例

```go
package main

import (
  "time"

  "github.com/gin-contrib/cors"
  "github.com/gin-gonic/gin"
)

func main() {
  // 初始化路由
  router := gin.Default()

  // 为 https:  //foo.com 和 https:  //github.com 来源设置 CORS，允许：
  // - PUT 和 PATCH 方法
  // - Origin 请求头
  // - 凭证共享
  // - 预检请求缓存12小时
  router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"https:  //foo.com"},
    AllowMethods:     []string{"PUT", "PATCH"},
    AllowHeaders:     []string{"Origin"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
    AllowOriginFunc: func(origin string) bool {
      return origin == "https:  //github.com"
    },
    MaxAge: 12 * time.Hour,
  }))

  // 运行服务器
  router.Run()
}
```

#

# <翻译结束>


<原文开始>
Using DefaultConfig as start point

```go
func main() {
  router := gin.Default()
  // - No origin allowed by default
  // - GET,POST, PUT, HEAD methods
  // - Credentials share disabled
  // - Preflight requests cached for 12 hours
  config := cors.DefaultConfig()
  config.AllowOrigins = []string{"http://google.com"}
  // config.AllowOrigins = []string{"http://google.com", "http://facebook.com"}
  // config.AllowAllOrigins = true

  router.Use(cors.New(config))
  router.Run()
}
```

Note: while Default() allows all origins, DefaultConfig() does not and you will still have to use AllowAllOrigins.

#
<原文结束>

# <翻译开始>
# 使用DefaultConfig作为起点

```go
func main() {
  // 初始化默认的gin路由器
  router := gin.Default()

  // - 默认情况下不允许任何来源
  // - 允许GET、POST、PUT、HEAD方法
  // - 禁用凭据共享
  // - 预检请求缓存时间为12小时
  config := cors.DefaultConfig()

  // 设置允许特定来源
  config.AllowOrigins = []string{"http:  //google.com"}
  // 可以设置多个允许来源
  // config.AllowOrigins = []string{"http:  //google.com", "http:  //facebook.com"}

  // 若要允许所有来源，需要设置AllowAllOrigins为true
  // config.AllowAllOrigins = true

  // 使用自定义配置的CORS中间件
  router.Use(cors.New(config))

  // 运行服务器
  router.Run()
}

注意：虽然Default()会允许所有来源，但DefaultConfig()默认并不允许，仍需手动设置AllowAllOrigins为true来允许所有来源。

# <翻译结束>


<原文开始>
Default() allows all origins

```go
func main() {
  router := gin.Default()
  // same as
  // config := cors.DefaultConfig()
  // config.AllowAllOrigins = true
  // router.Use(cors.New(config))
  router.Use(cors.Default())
  router.Run()
}
```

Using all origins disables the ability for Gin to set cookies for clients. When dealing with credentials, don't allow all origins.

<原文结束>

# <翻译开始>
# `Default()` 函数允许所有来源

```go
func main() {
  // 创建一个 Gin 路由器，并设置默认跨域策略，允许所有来源访问
  router := gin.Default()
  // 相当于以下配置：
  // config := cors.DefaultConfig()
  // config.AllowAllOrigins = true
  // router.Use(cors.New(config))

  // 使用默认的跨域中间件
  router.Use(cors.Default())

  // 运行服务器
  router.Run()
}
```

需要注意的是，允许所有来源将导致 Gin 无法为客户端设置 cookies。在处理包含凭据的情况时，不应允许所有来源。

# <翻译结束>

