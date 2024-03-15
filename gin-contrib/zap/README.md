# # zap

[![运行测试](https://github.com/gin-contrib/zap/actions/workflows/go.yml/badge.svg)](https://github.com/gin-contrib/zap/actions/workflows/go.yml)
[![Go 项目报告卡](https://goreportcard.com/badge/github.com/gin-contrib/zap)](https://goreportcard.com/report/github.com/gin-contrib/zap)
[![GoDoc 文档](https://godoc.org/github.com/gin-contrib/zap?status.svg)](https://godoc.org/github.com/gin-contrib/zap)
[![加入聊天室](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin)

通过 [zap](https://github.com/uber-go/zap) 提供的替代日志方案。感谢 [@yezooz](https://github.com/yezooz) 的 [Pull Request](https://github.com/gin-gonic/contrib/pull/129)。
## Usage

### # 开始使用

下载并安装：

```sh
go get github.com/gin-contrib/zap
```

在代码中导入：

```go
import "github.com/gin-contrib/zap"
```
## # 示例

参见 [example](_example/example01/main.go)。

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

  // 创建一个zap日志器实例
  logger, _ := zap.NewProduction()

  // 添加ginzap中间件，其功能包括：
  //   - 记录所有请求，类似于综合访问和错误日志。
  //   - 将日志输出到stdout。
  //   - 使用UTC时间格式的RFC3339格式。
  r.Use(ginzap.Ginzap(logger, time.RFC3339, true))

  // 记录所有panic信息到错误日志
  //   - stack 参数表示是否输出堆栈信息。
  r.Use(ginzap.RecoveryWithZap(logger, true))

  // 示例ping请求
  r.GET("/ping", func(c *gin.Context) {
    c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
  })

  // 当发生panic时的示例
  r.GET("/panic", func(c *gin.Context) {
    panic("发生了意外错误！")
  })

  // 监听并在0.0.0.0:8080端口上启动服务
  r.Run(":8080")
}
```
## # 当您想要对特定路径跳过日志记录时，请使用 GinzapWithConfig。

```go
r.Use(GinzapWithConfig(utcLogger, &Config{
  TimeFormat: time.RFC3339,
  UTC: true,
  SkipPaths: []string{"/no_log"},
}))
```

翻译：
若要针对特定路径禁用日志记录，请使用 GinzapWithConfig。

```go
r.Use(GinzapWithConfig(utcLogger, &Config{
  时间格式: time.RFC3339,
  是否使用UTC: true,
  跳过路径: []string{"/no_log"},
}))
```
## Custom Zap fields

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
