package main

import (
	"bytes"
	"fmt"
	"io"
	"time"
	
	ginzap "github.com/888go/gin/gin-contrib/zap"
	
	"github.com/888go/gin"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	r := gin类.X创建()

	logger, _ := zap.NewProduction()

	r.X中间件(ginzap.GinzapWithConfig(logger, &ginzap.Config{
		UTC:        true,
		TimeFormat: time.RFC3339,
		Context: ginzap.Fn(func(c *gin类.Context) []zapcore.Field {
			fields := []zapcore.Field{}
			// log request ID
			if requestID := c.Writer.Header().Get("X-Request-Id"); requestID != "" {
				fields = append(fields, zap.String("request_id", requestID))
			}

			// log trace and span ID
			if trace.SpanFromContext(c.X请求.Context()).SpanContext().IsValid() {
				fields = append(fields, zap.String("trace_id", trace.SpanFromContext(c.X请求.Context()).SpanContext().TraceID().String()))
				fields = append(fields, zap.String("span_id", trace.SpanFromContext(c.X请求.Context()).SpanContext().SpanID().String()))
			}

			// log request body
			var body []byte
			var buf bytes.Buffer
			tee := io.TeeReader(c.X请求.Body, &buf)
			body, _ = io.ReadAll(tee)
			c.X请求.Body = io.NopCloser(&buf)
			fields = append(fields, zap.String("body", string(body)))

			return fields
		}),
	}))

	// Example ping request.
	r.X绑定GET("/ping", func(c *gin类.Context) {
		c.Writer.Header().Add("X-Request-Id", "1234-5678-9012")
		c.X输出文本(200, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	r.X绑定POST("/ping", func(c *gin类.Context) {
		c.Writer.Header().Add("X-Request-Id", "9012-5678-1234")
		c.X输出文本(200, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// Listen and Server in 0.0.0.0:8080
	r.X监听(":8080")
}
