
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
# 示例代码

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
// 监听并在 0.0.0.0:8080 端口上服务
  if err := r.Run(":8080"); err != nil {
    log.Fatal(err)
  }
}
```

这段代码是使用Gin框架的一个示例，主要功能是在本地服务器上提供静态文件服务，并设置了一个简单的HTTP GET接口。具体来说：

1. 导入所需包`gin-contrib/static`和`gin-gonic/gin`。
2. 初始化一个默认的Gin引擎`r := gin.Default()`。
3. 使用`static.Serve`中间件来服务本地文件系统中的静态文件，本例中为`/tmp`目录，且不开启目录索引功能。
4. 定义一个GET请求处理函数，当访问`/ping`路径时返回HTTP状态码200及字符串"test"。
5. 最后，运行服务器监听在0.0.0.0:8080端口，如果运行过程中出现错误，则记录致命错误并退出程序。

# <翻译结束>

