# # 缓存 Gin 的中间件

[![构建状态](https://github.com/gin-contrib/cache/actions/workflows/testing.yml/badge.svg)](https://github.com/gin-contrib/cache/actions/workflows/testing.yml)
[![codecov](https://codecov.io/gh/gin-contrib/cache/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/cache)
[![Go 语言报告卡](https://goreportcard.com/badge/github.com/gin-contrib/cache)](https://goreportcard.com/report/github.com/gin-contrib/cache)
[![GoDoc](https://godoc.org/github.com/gin-contrib/cache?status.svg)](https://godoc.org/github.com/gin-contrib/cache)

启用缓存功能的 Gin 中间件/处理器。
## Usage

### # 开始使用

下载并安装：

```sh
$ go get github.com/gin-contrib/cache
```

在代码中导入：

```go
import "github.com/gin-contrib/cache"
```
## # 规范示例：

查看 [example](example/example.go)

```go
package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 创建一个内存存储，缓存有效期为1秒
	store := persistence.NewInMemoryStore(time.Second)

	// 不缓存的页面路由
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// 缓存处理的页面路由，缓存有效期为1分钟
	r.GET("/cache_ping", cache.CachePage(store, time.Minute, func(c *gin.Context) {
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
	}))

	// 监听并在 0.0.0.0:8080 端口启动服务
	r.Run(":8080")
}
```
