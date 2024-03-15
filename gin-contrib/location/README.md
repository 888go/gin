# # 位置

[![运行测试](https://github.com/gin-contrib/location/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/gin-contrib/location/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/location/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/location)
[![Go 代码质量报告](https://goreportcard.com/badge/github.com/gin-contrib/location)](https://goreportcard.com/report/github.com/gin-contrib/location)
[![GoDoc](https://godoc.org/github.com/gin-contrib/location?status.svg)](https://godoc.org/github.com/gin-contrib/location)

这个 Gin 中间件可以用来自动查找并公开服务器的主机名和协议，通过检查传入的 `http.Request` 中的信息。如果不使用这个插件，则需要显式地通过命令行参数或环境变量向服务器提供此类信息。
## Usage

### # 开始使用

下载并安装：

```bash
go get github.com/gin-contrib/location
```

在代码中导入：

```go
import "github.com/gin-contrib/location"
```

#
## # 默认

```go
package main

import (
  "github.com/gin-contrib/location"
  "github.com/gin-gonic/gin"
)

func main() {
  // 初始化路由
  router := gin.Default()

  // 配置自动检测 scheme（协议）和 host（主机）
  // - 当无法确定默认 scheme 时，使用 http
  // - 当无法确定默认 host 时，使用 localhost:8080
  router.Use(location.Default())

  // 定义 GET 请求处理函数
  router.GET("/", func(c *gin.Context) {
    // 获取请求的 URL 信息
    url := location.Get(c)

    // 可以通过以下方式获取 URL 的各个部分：
    // url.Scheme   // 协议（http、https 等）
    // url.Host     // 主机（包括域名和端口）
    // url.Path     // 路径

  })

  // 运行服务器
  router.Run()
}
```
## # 自定义

```go
package main

import (
  "github.com/gin-contrib/location"
  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()

// 配置自动检测 scheme（协议）和 host（主机），并设置回退为 https://foo.com/base
// - 当无法确定默认协议时使用 https
// - 当无法确定默认主机时使用 foo.com
// - 将 /base 作为路径包含在内
  router.Use(location.New(location.Config{
    Scheme: "https",
    Host: "foo.com",
    Base: "/base",
    Headers: location.Headers{Scheme: "X-Forwarded-Proto", Host: "X-Forwarded-For"},
  }))

  router.GET("/", func(c *gin.Context) {
    url := location.Get(c)

// url.Scheme 获取协议
// url.Host 获取主机
// url.Path 获取路径
  })

  router.Run()
}
```
## # 贡献

 Fork（派生） -> Patch（打补丁） -> Push（推送） -> Pull Request（拉取请求）
## License

MIT
