# Gin Web Framework

<img align="right" width="159px" src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png">

[![Build Status](https://github.com/gin-gonic/gin/workflows/Run%20Tests/badge.svg?branch=master)](https://github.com/gin-gonic/gin/actions?query=branch%3Amaster)
[![codecov](https://codecov.io/gh/gin-gonic/gin/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-gonic/gin)
[![Go Report Card](https://goreportcard.com/badge/github.com/gin-gonic/gin)](https://goreportcard.com/report/github.com/gin-gonic/gin)
[![GoDoc](https://pkg.go.dev/badge/github.com/gin-gonic/gin?status.svg)](https://pkg.go.dev/github.com/gin-gonic/gin?tab=doc)
[![Sourcegraph](https://sourcegraph.com/github.com/gin-gonic/gin/-/badge.svg)](https://sourcegraph.com/github.com/gin-gonic/gin?badge)
[![Open Source Helpers](https://www.codetriage.com/gin-gonic/gin/badges/users.svg)](https://www.codetriage.com/gin-gonic/gin)
[![Release](https://img.shields.io/github/release/gin-gonic/gin.svg?style=flat-square)](https://github.com/gin-gonic/gin/releases)
[![TODOs](https://badgen.net/https/api.tickgit.com/badgen/github.com/gin-gonic/gin)](https://www.tickgit.com/browse?repo=github.com/gin-gonic/gin)

Gin is a web framework written in [Go](https://go.dev/). It features a martini-like API with performance that is up to 40 times faster thanks to [httprouter](https://github.com/julienschmidt/httprouter). If you need performance and good productivity, you will love Gin.

**The key features of Gin are:**

- Zero allocation router
- Fast
- Middleware support
- Crash-free
- JSON validation
- Routes grouping
- Error management
- Rendering built-in
- Extendable


## Getting started

### # 前提条件

- **[Go](https://go.dev/)**: 最近**三个主要版本**中的任意一个 [发布版本](https://go.dev/doc/devel/release)（我们使用这些版本进行测试）。
## # 获取 Gin

通过 [Go 模块](https://github.com/golang/go/wiki/Modules) 支持，只需在代码中添加以下导入语句：

```go
import "github.com/gin-gonic/gin"
```

然后运行 `go [build|run|test]` 会自动获取所需的依赖项。

否则，可以运行以下 Go 命令来安装 `gin` 包：

```sh
$ go get -u github.com/gin-gonic/gin
```

#
## # 运行 Gin

首先，你需要导入 Gin 包以便使用 Gin，一个最简单的示例如下 `example.go`：

```go
package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  r.Run() // 监听并服务在 0.0.0.0:8080（对于 Windows 是 "localhost:8080"）
}
```

然后使用 Go 命令来运行这个示例：

```
# # 运行example.go并在浏览器中访问0.0.0.0:8080/ping
```
$ go run example.go
```

（由于您提供的内容不完整，此处无法展示更多上下文或代码含义。不过以上是命令行提示，指示用户使用Go语言运行名为“example.go”的程序，并在浏览器中打开地址“0.0.0.0:8080/ping”来查看运行结果。）
## Learn more examples

#### # 快速入门

想要学习和实践更多示例，请阅读包含 API 示例和构建标签的[Gin 快速入门](docs/doc.md)。
## # 示例

在[Gin示例](https://github.com/gin-gonic/examples)仓库中，有一系列展示了Gin在各种应用场景下使用的现成示例。
## Documentation

See [API documentation and descriptions](https://godoc.org/github.com/gin-gonic/gin) for package.

All documentation is available on the Gin website.

- [English](https://gin-gonic.com/docs/)
- [简体中文](https://gin-gonic.com/zh-cn/docs/)
- [繁體中文](https://gin-gonic.com/zh-tw/docs/)
- [日本語](https://gin-gonic.com/ja/docs/)
- [Español](https://gin-gonic.com/es/docs/)
- [한국어](https://gin-gonic.com/ko-kr/docs/)
- [Turkish](https://gin-gonic.com/tr/docs/)
- [Persian](https://gin-gonic.com/fa/docs/)

### # 关于Gin的文章

精心整理的Gin框架优秀资源列表：

- [教程：使用Go和Gin开发RESTful API](https://go.dev/doc/tutorial/web-service-gin)
## Benchmarks

Gin uses a custom version of [HttpRouter](https://github.com/julienschmidt/httprouter), [see all benchmarks details](/BENCHMARKS.md).

| Benchmark name                 |       (1) |             (2) |          (3) |             (4) |
| ------------------------------ | ---------:| ---------------:| ------------:| ---------------:|
| BenchmarkGin_GithubAll         | **43550** | **27364 ns/op** |   **0 B/op** | **0 allocs/op** |
| BenchmarkAce_GithubAll         |     40543 |     29670 ns/op |       0 B/op |     0 allocs/op |
| BenchmarkAero_GithubAll        |     57632 |     20648 ns/op |       0 B/op |     0 allocs/op |
| BenchmarkBear_GithubAll        |      9234 |    216179 ns/op |   86448 B/op |   943 allocs/op |
| BenchmarkBeego_GithubAll       |      7407 |    243496 ns/op |   71456 B/op |   609 allocs/op |
| BenchmarkBone_GithubAll        |       420 |   2922835 ns/op |  720160 B/op |  8620 allocs/op |
| BenchmarkChi_GithubAll         |      7620 |    238331 ns/op |   87696 B/op |   609 allocs/op |
| BenchmarkDenco_GithubAll       |     18355 |     64494 ns/op |   20224 B/op |   167 allocs/op |
| BenchmarkEcho_GithubAll        |     31251 |     38479 ns/op |       0 B/op |     0 allocs/op |
| BenchmarkGocraftWeb_GithubAll  |      4117 |    300062 ns/op |  131656 B/op |  1686 allocs/op |
| BenchmarkGoji_GithubAll        |      3274 |    416158 ns/op |   56112 B/op |   334 allocs/op |
| BenchmarkGojiv2_GithubAll      |      1402 |    870518 ns/op |  352720 B/op |  4321 allocs/op |
| BenchmarkGoJsonRest_GithubAll  |      2976 |    401507 ns/op |  134371 B/op |  2737 allocs/op |
| BenchmarkGoRestful_GithubAll   |       410 |   2913158 ns/op |  910144 B/op |  2938 allocs/op |
| BenchmarkGorillaMux_GithubAll  |       346 |   3384987 ns/op |  251650 B/op |  1994 allocs/op |
| BenchmarkGowwwRouter_GithubAll |     10000 |    143025 ns/op |   72144 B/op |   501 allocs/op |
| BenchmarkHttpRouter_GithubAll  |     55938 |     21360 ns/op |       0 B/op |     0 allocs/op |
| BenchmarkHttpTreeMux_GithubAll |     10000 |    153944 ns/op |   65856 B/op |   671 allocs/op |
| BenchmarkKocha_GithubAll       |     10000 |    106315 ns/op |   23304 B/op |   843 allocs/op |
| BenchmarkLARS_GithubAll        |     47779 |     25084 ns/op |       0 B/op |     0 allocs/op |
| BenchmarkMacaron_GithubAll     |      3266 |    371907 ns/op |  149409 B/op |  1624 allocs/op |
| BenchmarkMartini_GithubAll     |       331 |   3444706 ns/op |  226551 B/op |  2325 allocs/op |
| BenchmarkPat_GithubAll         |       273 |   4381818 ns/op | 1483152 B/op | 26963 allocs/op |
| BenchmarkPossum_GithubAll      |     10000 |    164367 ns/op |   84448 B/op |   609 allocs/op |
| BenchmarkR2router_GithubAll    |     10000 |    160220 ns/op |   77328 B/op |   979 allocs/op |
| BenchmarkRivet_GithubAll       |     14625 |     82453 ns/op |   16272 B/op |   167 allocs/op |
| BenchmarkTango_GithubAll       |      6255 |    279611 ns/op |   63826 B/op |  1618 allocs/op |
| BenchmarkTigerTonic_GithubAll  |      2008 |    687874 ns/op |  193856 B/op |  4474 allocs/op |
| BenchmarkTraffic_GithubAll     |       355 |   3478508 ns/op |  820744 B/op | 14114 allocs/op |
| BenchmarkVulcan_GithubAll      |      6885 |    193333 ns/op |   19894 B/op |   609 allocs/op |

- (1): Total Repetitions achieved in constant time, higher means more confident result
- (2): Single Repetition Duration (ns/op), lower is better
- (3): Heap Memory (B/op), lower is better
- (4): Average Allocations per Repetition (allocs/op), lower is better


## # 中间件

你可以在 [gin-contrib](https://github.com/gin-contrib) 找到许多有用的 Gin 中间件。
## # 用户

使用[Gin](https://github.com/gin-gonic/gin) web框架的精彩项目列表。

- [gorush](https://github.com/appleboy/gorush): 一个用Go编写的推送通知服务器。
- [fnproject](https://github.com/fnproject/fn): 适用于容器、云平台无关的无服务器平台。
- [photoprism](https://github.com/photoprism/photoprism): 基于Go和Google TensorFlow的个人照片管理工具。
- [lura](https://github.com/luraproject/lura): 高性能API网关，支持中间件。
- [picfit](https://github.com/thoas/picfit): 一个用Go编写的图像缩放服务器。
- [dkron](https://github.com/distribworks/dkron): 分布式、容错的工作调度系统。
## # 贡献

Gin 由数百名贡献者共同完成。我们非常感谢您的帮助！

请参阅 [CONTRIBUTING](CONTRIBUTING.md) 获取提交补丁和贡献工作流程的详细信息。
