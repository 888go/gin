
<原文开始>
zap

[![Run Tests](https://github.com/gin-contrib/zap/actions/workflows/go.yml/badge.svg)](https://github.com/gin-contrib/zap/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/gin-contrib/zap)](https://goreportcard.com/report/github.com/gin-contrib/zap)
[![GoDoc](https://godoc.org/github.com/gin-contrib/zap?status.svg)](https://godoc.org/github.com/gin-contrib/zap)
[![Join the chat at https://gitter.im/gin-gonic/gin](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin)

Alternative logging through [zap](https://github.com/uber-go/zap). Thanks for [Pull Request](https://github.com/gin-gonic/contrib/pull/129) from [@yezooz](https://github.com/yezooz)


<原文结束>

# <翻译开始>
# zap

[![运行测试](https://github.com/gin-contrib/zap/actions/workflows/go.yml/badge.svg)](https://github.com/gin-contrib/zap/actions/workflows/go.yml)
[![Go 项目报告卡](https://goreportcard.com/badge/github.com/gin-contrib/zap)](https://goreportcard.com/report/github.com/gin-contrib/zap)
[![GoDoc 文档](https://godoc.org/github.com/gin-contrib/zap?status.svg)](https://godoc.org/github.com/gin-contrib/zap)
[![加入聊天室](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin)

通过[zap](https://github.com/uber-go/zap)实现替代日志记录功能。感谢[@yezooz](https://github.com/yezooz)提交的[Pull Request](https://github.com/gin-gonic/contrib/pull/129)。

# <翻译结束>


<原文开始>
Start using it

Download and install it:

```sh
go get github.com/gin-contrib/zap
```

Import it in your code:

```go
import "github.com/gin-contrib/zap"
```


<原文结束>

# <翻译开始>
# 开始使用

下载并安装：

```sh
go get github.com/gin-contrib/zap
```

在代码中引入：

```go
import "github.com/gin-contrib/zap"
```

# <翻译结束>


<原文开始>
Example

See the [example](_example/example01/main.go).

```go
package main

import (
  "fmt"
  "time"

  ginzap "github.com/gin-contrib/zap"
  "github.com/gin-gonic/gin"
  "go.uber.org/zap"
)

func main() {
  r := gin.New()

  logger, _ := zap.NewProduction()

  // Add a ginzap middleware, which:
  //   - Logs all requests, like a combined access and error log.
  //   - Logs to stdout.
  //   - RFC3339 with UTC time format.
  r.Use(ginzap.Ginzap(logger, time.RFC3339, true))

  // Logs all panic to error log
  //   - stack means whether output the stack info.
  r.Use(ginzap.RecoveryWithZap(logger, true))

  // Example ping request.
  r.GET("/ping", func(c *gin.Context) {
    c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
  })

  // Example when panic happen.
  r.GET("/panic", func(c *gin.Context) {
    panic("An unexpected error happen!")
  })

  // Listen and Server in 0.0.0.0:8080
  r.Run(":8080")
}
```


<原文结束>

# <翻译开始>
# 示例

查看 [example](_example/example01/main.go)。

```go
package main

import (
  "fmt"
  "time"

  "github.com/gin-contrib/zap" // Gin框架中集成zap日志库
  "github.com/gin-gonic/gin"    // Gin Web框架
  "go.uber.org/zap"             // Zap日志库
)

func main() {
  r := gin.New() // 创建Gin引擎实例

  logger, _ := zap.NewProduction() // 初始化zap日志生产环境配置

// 添加ginzap中间件，其功能包括：
//   - 记录所有请求信息，如同综合访问和错误日志。
//   - 将日志输出到标准输出（stdout）。
//   - 使用RFC3339格式并以UTC时间显示。
  r.Use(ginzap.Ginzap(logger, time.RFC3339, true))

// 记录所有的panic异常到错误日志
//   - stack参数表示是否输出堆栈信息。
  r.Use(ginzap.RecoveryWithZap(logger, true))

// 示例ping请求处理
  r.GET("/ping", func(c *gin.Context) {
    c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
  })

// 当panic异常发生时的示例
  r.GET("/panic", func(c *gin.Context) {
    panic("发生了一个意外错误！")
  })

// 在0.0.0.0:8080监听并启动服务
  r.Run(":8080")
}
```

# <翻译结束>


<原文开始>
Skip logging

When you want to skip logging for specific path,
please use GinzapWithConfig

```go

r.Use(GinzapWithConfig(utcLogger, &Config{
  TimeFormat: time.RFC3339,
  UTC: true,
  SkipPaths: []string{"/no_log"},
}))
```


<原文结束>

# <翻译开始>
# 跳过日志记录

当你想要对特定路径跳过日志记录时，请使用 GinzapWithConfig 方法。

```go
r.Use(GinzapWithConfig(utcLogger, &Config{
  TimeFormat: time.RFC3339,
  UTC: true,
  SkipPaths: []string{"/no_log"},
}))
```

翻译：

当需要为特定路径禁用日志记录功能时，可以采用 GinzapWithConfig 函数。

```go
r.Use(GinzapWithConfig(utcLogger, &Config{
  时间格式: time.RFC3339,
  是否使用UTC时间: true,
  忽略路径: []string{"/no_log"},
}))
```

# <翻译结束>


<原文开始>
Custom Zap fields

example for custom log request body, response request ID or log [Open Telemetry](https://opentelemetry.io/) TraceID.

```go
func main() {
  r := gin.New()

  logger, _ := zap.NewProduction()

  r.Use(ginzap.GinzapWithConfig(logger, &ginzap.Config{
    UTC:        true,
    TimeFormat: time.RFC3339,
    Context: ginzap.Fn(func(c *gin.Context) []zapcore.Field {
      fields := []zapcore.Field{}
      // log request ID
      if requestID := c.Writer.Header().Get("X-Request-Id"); requestID != "" {
        fields = append(fields, zap.String("request_id", requestID))
      }

      // log trace and span ID
      if trace.SpanFromContext(c.Request.Context()).SpanContext().IsValid() {
        fields = append(fields, zap.String("trace_id", trace.SpanFromContext(c.Request.Context()).SpanContext().TraceID().String()))
        fields = append(fields, zap.String("span_id", trace.SpanFromContext(c.Request.Context()).SpanContext().SpanID().String()))
      }

      // log request body
      var body []byte
      var buf bytes.Buffer
      tee := io.TeeReader(c.Request.Body, &buf)
      body, _ = io.ReadAll(tee)
      c.Request.Body = io.NopCloser(&buf)
      fields = append(fields, zap.String("body", string(body)))

      return fields
    }),
  }))

  // Example ping request.
  r.GET("/ping", func(c *gin.Context) {
    c.Writer.Header().Add("X-Request-Id", "1234-5678-9012")
    c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
  })

  r.POST("/ping", func(c *gin.Context) {
    c.Writer.Header().Add("X-Request-Id", "9012-5678-1234")
    c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
  })

  // Listen and Server in 0.0.0.0:8080
  r.Run(":8080")
}
```

<原文结束>

# <翻译开始>
# 自定义Zap字段

这是一个示例，用于自定义日志记录请求体、响应请求ID或日志[Open Telemetry](https://opentelemetry.io/) TraceID。

```go
func main() {
// 创建一个新的Gin实例
  r := gin.New()

// 初始化生产环境的zap日志器
  logger, _ := zap.NewProduction()

// 使用Ginzap中间件并配置自定义上下文
  r.Use(ginzap.GinzapWithConfig(logger, &ginzap.Config{
    UTC:        true,          // 使用UTC时间
    TimeFormat: time.RFC3339,  // 时间格式为RFC3339
    Context: ginzap.Fn(func(c *gin.Context) []zapcore.Field {
      fields := []zapcore.Field{} // 初始化字段列表

// 记录请求ID
      if requestID := c.Writer.Header().Get("X-Request-Id"); requestID != "" {
        fields = append(fields, zap.String("request_id", requestID))
      }

// 记录Trace ID和Span ID（Open Telemetry）
      if span := trace.SpanFromContext(c.Request.Context()); span.SpanContext().IsValid() {
        fields = append(fields, zap.String("trace_id", span.SpanContext().TraceID().String()))
        fields = append(fields, zap.String("span_id", span.SpanContext().SpanID().String()))
      }

// 记录请求体
      var body []byte
      buf := bytes.Buffer{}
      teeReader := io.TeeReader(c.Request.Body, &buf)
      body, _ = io.ReadAll(teeReader)
      c.Request.Body = io.NopCloser(&buf)
      fields = append(fields, zap.String("body", string(body)))

      return fields
    }),
  }))

// 示例ping请求
  r.GET("/ping", func(c *gin.Context) {
    c.Writer.Header().Add("X-Request-Id", "1234-5678-9012")
    c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
  })

  r.POST("/ping", func(c *gin.Context) {
    c.Writer.Header().Add("X-Request-Id", "9012-5678-1234")
    c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
  })

// 监听并服务在0.0.0.0:8080端口
  r.Run(":8080")
}
```

这段代码展示了如何使用Gin框架和Zap日志库来实现一个HTTP服务器，并在日志中记录自定义信息，包括请求ID、Open Telemetry追踪ID和Span ID以及请求体。同时，它还提供了一个简单的GET和POST接口处理函数。

# <翻译结束>

