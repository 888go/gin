package main

import (
	"fmt"
	"net/http"
	"regexp"
	"time"
	
	"github.com/888go/gin/gin-contrib/logger"
	"github.com/888go/gin/gin-contrib/requestid"
	"github.com/888go/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/trace"
)

var rxURL = regexp.MustCompile(`^/regexp\d*`)

func main() {
	r := gin.New()

// 添加一个日志中间件，其功能包括：
//   - 记录所有请求，就像综合访问和错误日志一样。
//   - 将日志记录到标准输出（stdout）。
// r.Use(logger.SetLogger())

// 示例：请求 pong
	r.GET("/pong", logger.SetLogger(), func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
	})

// 示例 ping 请求。
	r.GET("/ping", logger.SetLogger(
		logger.WithSkipPath([]string{"/skip"}),
		logger.WithUTC(true),
		logger.WithSkipPathRegexps(rxURL),
	), func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
	})

// 示例：跳过路径请求。
	r.GET("/skip", logger.SetLogger(
		logger.WithSkipPath([]string{"/skip"}),
	), func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
	})

// 示例：跳过路径请求。
	r.GET("/regexp1", logger.SetLogger(
		logger.WithSkipPathRegexps(rxURL),
	), func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
	})

// 示例：跳过路径请求。
	r.GET("/regexp2", logger.SetLogger(
		logger.WithSkipPathRegexps(rxURL),
	), func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
	})

// 添加自定义字段。
	r.GET("/id", requestid.New(requestid.WithGenerator(func() string {
		return "foobar"
	})), logger.SetLogger(
		logger.WithLogger(func(c *gin.Context, l zerolog.Logger) zerolog.Logger {
			if trace.SpanFromContext(c.Request.Context()).SpanContext().IsValid() {
				l = l.With().
					Str("trace_id", trace.SpanFromContext(c.Request.Context()).SpanContext().TraceID().String()).
					Str("span_id", trace.SpanFromContext(c.Request.Context()).SpanContext().SpanID().String()).
					Logger()
			}

			return l.With().
				Str("id", requestid.Get(c)).
				Str("foo", "bar").
				Str("path", c.Request.URL.Path).
				Logger()
		}),
	), func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
	})

// 示例：JSON格式的日志
	r.GET("/json", logger.SetLogger(
		logger.WithLogger(func(_ *gin.Context, l zerolog.Logger) zerolog.Logger {
			return l.Output(gin.DefaultWriter).With().Logger()
		}),
	), func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
	})

// 在0.0.0.0:8080监听并服务
	if err := r.Run(":8080"); err != nil {
		log.Fatal().Msg("can' start server with 8080 port")
	}
}
