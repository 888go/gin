
<原文开始>
Timeout

[![Run Tests](https://github.com/gin-contrib/timeout/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/gin-contrib/timeout/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/timeout/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/timeout)
[![Go Report Card](https://goreportcard.com/badge/github.com/gin-contrib/timeout)](https://goreportcard.com/report/github.com/gin-contrib/timeout)
[![GoDoc](https://godoc.org/github.com/gin-contrib/timeout?status.svg)](https://pkg.go.dev/github.com/gin-contrib/timeout?tab=doc)
[![Join the chat at https://gitter.im/gin-gonic/gin](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin)

Timeout wraps a handler and aborts the process of the handler if the timeout is reached.


<原文结束>

# <翻译开始>
# 超时

[![运行测试](https://github.com/gin-contrib/timeout/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/gin-contrib/timeout/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/timeout/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/timeout)
[![Go 报告卡](https://goreportcard.com/badge/github.com/gin-contrib/timeout)](https://goreportcard.com/report/github.com/gin-contrib/timeout)
[![GoDoc](https://godoc.org/github.com/gin-contrib/timeout?status.svg)](https://pkg.go.dev/github.com/gin-contrib/timeout?tab=doc)
[![加入聊天 https://gitter.im/gin-gonic/gin](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin)

Timeout 封装了一个处理器，如果达到超时时间，则会中止该处理器的执行过程。

# <翻译结束>


<原文开始>
Example

```go
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func emptySuccessResponse(c *gin.Context) {
	time.Sleep(200 * time.Microsecond)
	c.String(http.StatusOK, "")
}

func main() {
	r := gin.New()

	r.GET("/", timeout.New(
		timeout.WithTimeout(100*time.Microsecond),
		timeout.WithHandler(emptySuccessResponse),
	))

	// Listen and Server in 0.0.0.0:8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
```

#
<原文结束>

# <翻译开始>
# 示例

```go
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

// 空成功响应函数，模拟耗时操作后返回空字符串和200状态码
func emptySuccessResponse(c *gin.Context) {
	time.Sleep(200 * time.Microsecond)
	c.String(http.StatusOK, "")
}

func main() {
	// 创建一个 gin 新实例
	r := gin.New()

	// 配置路由，访问根路径（"/"）时使用 timeout 中间件
	// 设置超时时间为100微秒，并在超时后调用 emptySuccessResponse 函数处理请求
	r.GET("/", timeout.New(
		timeout.WithTimeout(100*time.Microsecond),
		timeout.WithHandler(emptySuccessResponse),
	))

	// 监听并服务在 0.0.0.0:8080 端口
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
```

#

# <翻译结束>


<原文开始>
custom error response

Add new error response func:

```go
func testResponse(c *gin.Context) {
	c.String(http.StatusRequestTimeout, "test response")
}
```

Add `WithResponse` option.

```go
	r.GET("/", timeout.New(
		timeout.WithTimeout(100*time.Microsecond),
		timeout.WithHandler(emptySuccessResponse),
		timeout.WithResponse(testResponse),
	))
```

#
<原文结束>

# <翻译开始>
# 自定义错误响应

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

# <翻译结束>


<原文开始>
custom middleware

```go
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func testResponse(c *gin.Context) {
	c.String(http.StatusRequestTimeout, "timeout")
}

func timeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(500*time.Millisecond),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(testResponse),
	)
}

func main() {
	r := gin.New()
	r.Use(timeoutMiddleware())
	r.GET("/slow", func(c *gin.Context) {
		time.Sleep(800 * time.Millisecond)
		c.Status(http.StatusOK)
	})
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
```

<原文结束>

# <翻译开始>
# 自定义中间件

```go
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

// testResponse 函数用于测试返回超时响应
func testResponse(c *gin.Context) {
	c.String(http.StatusRequestTimeout, "timeout")
}

// timeoutMiddleware 函数创建一个自定义的超时中间件
func timeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(500*time.Millisecond), // 设置请求超时时间为500毫秒
		timeout.WithHandler(func(c *gin.Context) {
			c.Next() // 超时后调用下一个处理器
		}),
		timeout.WithResponse(testResponse), // 当发生超时时，调用testResponse函数处理响应
	)
}

func main() {
	r := gin.New() // 创建一个新的Gin引擎实例
	r.Use(timeoutMiddleware()) // 使用自定义的超时中间件
	r.GET("/slow", func(c *gin.Context) {
		time.Sleep(800 * time.Millisecond) // 模拟耗时操作，休眠800毫秒
		c.Status(http.StatusOK) // 设置HTTP状态码为200（成功）
	})
	if err := r.Run(":8080"); err != nil { // 在8080端口运行服务器
		log.Fatal(err) // 如果运行时出现错误，则记录并终止程序
	}
}
```

# <翻译结束>

