# # 超时

[![运行测试](https://github.com/gin-contrib/timeout/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/gin-contrib/timeout/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/timeout/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/timeout)
[![Go 代码质量报告](https://goreportcard.com/badge/github.com/gin-contrib/timeout)](https://goreportcard.com/report/github.com/gin-contrib/timeout)
[![GoDoc](https://godoc.org/github.com/gin-contrib/timeout?status.svg)](https://pkg.go.dev/github.com/gin-contrib/timeout?tab=doc)
[![加入Gitter聊天室](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin)

Timeout 对处理器进行封装，如果达到设定的超时时间，则会中止处理器的执行过程。
## # 示例

```go
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

// emptySuccessResponse 函数会在200微秒的延迟后向客户端发送一个空的成功响应（状态码为200）
func emptySuccessResponse(c *gin.Context) {
	time.Sleep(200 * time.Microsecond)
	c.String(http.StatusOK, "")
}

func main() {
// 初始化一个新的 gin 引擎
	r := gin.New()

// 配置路由，当访问根路径 "/" 时，使用 timeout 中间件处理请求
// 设置超时时间为100微秒，并在超时后调用 emptySuccessResponse 处理函数
	r.GET("/", timeout.New(
		timeout.WithTimeout(100*time.Microsecond),
		timeout.WithHandler(emptySuccessResponse),
	))

// 监听并在 0.0.0.0:8080 端口启动服务器
	if err := r.Run(":8080"); err != nil {
// 如果运行服务器时出现错误，则输出错误信息并终止程序
		log.Fatal(err)
	}
}
```

#
## # 自定义错误响应

添加新的错误响应函数：

```go
func testResponse(c *gin.Context) {
	c.String(http.StatusRequestTimeout, "测试响应")
}
```

添加 `WithResponse` 选项。

```go
r.GET("/", timeout.New(
    timeout.WithTimeout(100*time.Microsecond),
    timeout.WithHandler(emptySuccessResponse),
    timeout.WithResponse(testResponse),
))
```

#
## # 自定义中间件

```go
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

// testResponse 函数用于返回超时测试响应
func testResponse(c *gin.Context) {
	c.String(http.StatusRequestTimeout, "timeout")
}

// timeoutMiddleware 函数创建一个自定义的超时中间件
func timeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(500*time.Millisecond), // 设置请求超时时间为500毫秒
		timeout.WithHandler(func(c *gin.Context) { // 超时后执行的处理函数
			c.Next()
		}),
		timeout.WithResponse(testResponse), // 当请求超时时调用的响应函数
	)
}

func main() {
	r := gin.New() // 创建 Gin 框架实例
	r.Use(timeoutMiddleware()) // 使用自定义的超时中间件

// 定义一个响应较慢的路由
	r.GET("/slow", func(c *gin.Context) {
		time.Sleep(800 * time.Millisecond) // 模拟耗时操作，延迟800毫秒
		c.Status(http.StatusOK) // 设置HTTP状态码为200（成功）
	})

// 启动服务器监听8080端口
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err) // 如果运行时出现错误，则输出错误信息并退出程序
	}
}
```
