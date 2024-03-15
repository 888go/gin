
<原文开始>
static middleware

[![Run Tests](https://github.com/gin-contrib/static/actions/workflows/go.yml/badge.svg)](https://github.com/gin-contrib/static/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/static/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/static)
[![Go Report Card](https://goreportcard.com/badge/github.com/gin-contrib/static)](https://goreportcard.com/report/github.com/gin-contrib/static)
[![GoDoc](https://godoc.org/github.com/gin-contrib/static?status.svg)](https://godoc.org/github.com/gin-contrib/static)

Static middleware


<原文结束>

# <翻译开始>
# 静态中间件

[![运行测试](https://github.com/gin-contrib/static/actions/workflows/go.yml/badge.svg)](https://github.com/gin-contrib/static/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/static/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/static)
[![Go 代码质量报告](https://goreportcard.com/badge/github.com/gin-contrib/static)](https://goreportcard.com/report/github.com/gin-contrib/static)
[![GoDoc 文档](https://godoc.org/github.com/gin-contrib/static?status.svg)](https://godoc.org/github.com/gin-contrib/static)

静态中间件

# <翻译结束>


<原文开始>
Start using it

Download and install it:

```sh
go get github.com/gin-contrib/static
```

Import it in your code:

```go
import "github.com/gin-contrib/static"
```

#
<原文结束>

# <翻译开始>
# 开始使用

下载并安装：

```sh
go get github.com/gin-contrib/static
```

在代码中导入：

```go
import "github.com/gin-contrib/static"
```

#

# <翻译结束>


<原文开始>
Canonical example

See the [example](_example)

```go
package main

import (
  "github.com/gin-contrib/static"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  // if Allow DirectoryIndex
  //r.Use(static.Serve("/", static.LocalFile("/tmp", true)))
  // set prefix
  //r.Use(static.Serve("/static", static.LocalFile("/tmp", true)))

  r.Use(static.Serve("/", static.LocalFile("/tmp", false)))
  r.GET("/ping", func(c *gin.Context) {
    c.String(200, "test")
  })
  // Listen and Server in 0.0.0.0:8080
  if err := r.Run(":8080"); err != nil {
    log.Fatal(err)
  }
}
```

<原文结束>

# <翻译开始>
# 示例

参见 [示例](_example)

```go
package main

import (
  "github.com/gin-contrib/static"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

// 如果允许目录索引
//r.Use(static.Serve("/", static.LocalFile("/tmp", true)))
// 设置前缀
//r.Use(static.Serve("/static", static.LocalFile("/tmp", true)))

  r.Use(static.Serve("/", static.LocalFile("/tmp", false)))
  r.GET("/ping", func(c *gin.Context) {
    c.String(200, "test")
  })
// 监听并服务在 0.0.0.0:8080
  if err := r.Run(":8080"); err != nil {
    log.Fatal(err)
  }
}
```

这个Go代码示例展示了如何使用`gin`框架结合`gin-contrib/static`中间件来提供静态文件服务。在这个例子中，它将本地`/tmp`目录作为静态文件根目录，并监听在`0.0.0.0:8080`端口上。同时，它还定义了一个GET路由`/ping`，当访问该路由时返回HTTP状态码200和字符串"test"。此外，代码中还包含两行被注释掉的示例代码，分别用于开启目录索引功能以及为静态文件服务设置前缀。

# <翻译结束>

