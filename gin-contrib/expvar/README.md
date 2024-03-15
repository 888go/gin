# # expvar

[![运行测试](https://github.com/gin-contrib/expvar/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/gin-contrib/expvar/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/expvar/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/expvar)
[![Go 项目报告卡](https://goreportcard.com/badge/github.com/gin-contrib/expvar)](https://goreportcard.com/report/github.com/gin-contrib/expvar)
[![GoDoc](https://godoc.org/github.com/gin-contrib/expvar?status.svg)](https://godoc.org/github.com/gin-contrib/expvar)

这是一个用于 gin 框架的 expvar 处理器，[expvar](https://golang.org/pkg/expvar/) 提供了一个公开变量的标准接口。
## Usage

### # 开始使用

下载并安装：

```sh
go get github.com/gin-contrib/expvar
```

在代码中导入：

```go
import "github.com/gin-contrib/expvar"
```

#
## # 标准示例

```go
package main

import (
  "log"

  "github.com/gin-contrib/expvar"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  r.GET("/debug/vars", expvar.Handler())

  if err := r.Run(":8080"); err != nil {
    log.Fatal(err)
  }
}
```

翻译：

```go
package main

import (
  "log"

  "github.com/gin-contrib/expvar" // Gin框架的expvar中间件包
  "github.com/gin-gonic/gin"      // Gin Web框架
)

func main() {
  r := gin.Default() // 初始化一个默认配置的Gin引擎

  r.GET("/debug/vars", expvar.Handler()) // 配置GET请求路由，用于暴露调试变量信息

  if err := r.Run(":8080"); err != nil { // 在8080端口运行Web服务器
    log.Fatal(err) // 如果运行时出现错误，则记录并终止程序
  }
}
```
