# # 位置

[![运行测试](https://github.com/gin-contrib/location/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/gin-contrib/location/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/location/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/location)
[![Go 代码质量报告](https://goreportcard.com/badge/github.com/gin-contrib/location)](https://goreportcard.com/report/github.com/gin-contrib/location)
[![GoDoc](https://godoc.org/github.com/gin-contrib/location?status.svg)](https://godoc.org/github.com/gin-contrib/location)

这个 Gin 中间件可以用于自动发现并公开服务器的主机名和协议，通过检查传入的 `http.Request` 中的信息。如果不使用此插件，替代方案则是显式地通过命令行参数或环境变量向服务器提供这些信息。
## Usage

### # 开始使用

下载并安装：

```bash
go get github.com/gin-contrib/location
```

在你的代码中导入：

```go
import "github.com/gin-contrib/location"
```

#
## # 默认设置

```go
package main

import (
  "github.com/gin-contrib/location"
  "github.com/gin-gonic/gin"
)

func main() {
// 初始化路由
  router := gin.Default()

// 配置自动检测请求的scheme（协议）和host（主机）
// - 当无法确定默认scheme时，使用http
// - 当无法确定默认host时，使用localhost:8080
  router.Use(location.Default())

// 定义GET请求根路径处理函数
  router.GET("/", func(c *gin.Context) {
// 获取当前请求的URL信息
    url := location.Get(c)

// 可以通过url对象获取如下信息：
// url.Scheme   // 请求协议（http或https等）
// url.Host     // 主机名及端口号
// url.Path     // 请求路径
  })

// 启动服务器监听路由
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
// 初始化路由
  router := gin.Default()

// 配置自动检测请求的 scheme（协议）和 host（主机），默认为 https://foo.com/base
// - 当无法确定默认 scheme 时，使用 https
// - 当无法确定默认 host 时，使用 foo.com
// - 路径包含 /base
  router.Use(location.New(location.Config{
    Scheme: "https",
    Host: "foo.com",
    Base: "/base",
    Headers: location.Headers{Scheme: "X-Forwarded-Proto", Host: "X-Forwarded-For"},
  }))

// 定义 GET 请求处理器
  router.GET("/", func(c *gin.Context) {
// 获取当前请求的 URL 信息
    url := location.Get(c)

// 使用以下属性获取 URL 的各个部分
// url.Scheme （协议）
// url.Host   （主机）
// url.Path   （路径）
  })

// 运行服务器
  router.Run()
}
```
## # 参与贡献

 Fork（分叉） -> Patch（打补丁） -> Push（推送） -> Pull Request（拉取请求）
## License

MIT
