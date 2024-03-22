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
	r := gin类.X创建()

	// Add a logger middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	// r.Use(logger.SetLogger())

	// Example pong request.
	r.X绑定GET("/pong", logger.SetLogger(), func(c *gin类.Context) {
		c.X输出文本(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// Example ping request.
	r.X绑定GET("/ping", logger.SetLogger(
		logger.WithSkipPath([]string{"/skip"}),
		logger.WithUTC(true),
		logger.WithSkipPathRegexps(rxURL),
	), func(c *gin类.Context) {
		c.X输出文本(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// Example skip path request.
	r.X绑定GET("/skip", logger.SetLogger(
		logger.WithSkipPath([]string{"/skip"}),
	), func(c *gin类.Context) {
		c.X输出文本(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// Example skip path request.
	r.X绑定GET("/regexp1", logger.SetLogger(
		logger.WithSkipPathRegexps(rxURL),
	), func(c *gin类.Context) {
		c.X输出文本(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// Example skip path request.
	r.X绑定GET("/regexp2", logger.SetLogger(
		logger.WithSkipPathRegexps(rxURL),
	), func(c *gin类.Context) {
		c.X输出文本(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// add custom fields.
	r.X绑定GET("/id", requestid.New(requestid.WithGenerator(func() string {
		return "foobar"
	})), logger.SetLogger(
		logger.WithLogger(func(c *gin类.Context, l zerolog.Logger) zerolog.Logger {
			if trace.SpanFromContext(c.X请求.Context()).SpanContext().IsValid() {
				l = l.With().
					Str("trace_id", trace.SpanFromContext(c.X请求.Context()).SpanContext().TraceID().String()).
					Str("span_id", trace.SpanFromContext(c.X请求.Context()).SpanContext().SpanID().String()).
					Logger()
			}

			return l.With().
				Str("id", requestid.Get(c)).
				Str("foo", "bar").
				Str("path", c.X请求.URL.Path).
				Logger()
		}),
	), func(c *gin类.Context) {
		c.X输出文本(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// Example of JSON format log
	r.X绑定GET("/json", logger.SetLogger(
		logger.WithLogger(func(_ *gin类.Context, l zerolog.Logger) zerolog.Logger {
			return l.Output(gin类.DefaultWriter).With().Logger()
		}),
	), func(c *gin类.Context) {
		c.X输出文本(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// Listen and Server in 0.0.0.0:8080
	if err := r.X监听(":8080"); err != nil {
		log.Fatal().Msg("can' start server with 8080 port")
	}
}
