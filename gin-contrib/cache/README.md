# # 缓存 Gin 的中间件

[![构建状态](https://github.com/gin-contrib/cache/actions/workflows/testing.yml/badge.svg)](https://github.com/gin-contrib/cache/actions/workflows/testing.yml)
[![codecov](https://codecov.io/gh/gin-contrib/cache/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/cache)
[![Go 语言报告卡](https://goreportcard.com/badge/github.com/gin-contrib/cache)](https://goreportcard.com/report/github.com/gin-contrib/cache)
[![GoDoc](https://godoc.org/github.com/gin-contrib/cache?status.svg)](https://godoc.org/github.com/gin-contrib/cache)

这是一个 Gin 中间件/处理器，用于启用缓存功能。
## Usage

### # 开始使用

下载并安装：

```sh
$ go get github.com/gin-contrib/cache
```

在代码中引入：

```go
import "github.com/gin-contrib/cache"
```

#
## # 示例代码：

查看[示例](example/example.go)

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
// 初始化gin默认引擎
	r := gin.Default()

// 创建内存存储实例，缓存有效期为1秒
	store := persistence.NewInMemoryStore(time.Second)

// 定义GET请求路由"/ping"，返回当前时间戳字符串
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
	})

// 定义GET请求路由"/cache_ping"，并开启缓存功能，缓存有效期为1分钟
// 当请求此路由时，会调用匿名函数生成响应内容，并将其缓存
	r.GET("/cache_ping", cache.CachePage(store, time.Minute, func(c *gin.Context) {
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
	}))

// 在0.0.0.0:8080端口启动服务并监听请求
	r.Run(":8080")
}
```
