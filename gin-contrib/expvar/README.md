# # expvar

[![运行测试](https://github.com/gin-contrib/expvar/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/gin-contrib/expvar/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/expvar/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/expvar)
[![Go 代码质量报告](https://goreportcard.com/badge/github.com/gin-contrib/expvar)](https://goreportcard.com/report/github.com/gin-contrib/expvar)
[![GoDoc](https://godoc.org/github.com/gin-contrib/expvar?status.svg)](https://godoc.org/github.com/gin-contrib/expvar)

这是一个用于gin框架的expvar处理器，[expvar](https://golang.org/pkg/expvar/) 提供了一个标准化接口来公开变量。
## Usage

### # 开始使用它

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
// 初始化默认的Gin引擎
  r := gin.Default()

// 添加GET路由，用于访问/debug/vars，并使用expvar.Handler处理请求
  r.GET("/debug/vars", expvar.Handler())

// 在8080端口运行web服务器，如果运行时出现错误，则输出错误信息并终止程序
  if err := r.Run(":8080"); err != nil {
    log.Fatal(err)
  }
}
```
