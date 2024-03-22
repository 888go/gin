// Package ginzap provides log handling using zap package.
// Code structure based on ginrus package.
package ginzap

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
	
	"github.com/888go/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Fn func(c *gin类.Context) []zapcore.Field

// ZapLogger is the minimal logger interface compatible with zap.Logger
type ZapLogger interface {
	Info(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
}

// Config is config setting for Ginzap
type Config struct {
	TimeFormat   string
	UTC          bool
	SkipPaths    []string
	Context      Fn
	DefaultLevel zapcore.Level
}

// Ginzap returns a gin.HandlerFunc (middleware) that logs requests using uber-go/zap.
//
// Requests with errors are logged using zap.Error().
// Requests without errors are logged using zap.Info().
//
// It receives:
//  1. A time package format string (e.g. time.RFC3339).
//  2. A boolean stating whether to use UTC time zone or local.
func Ginzap(logger ZapLogger, timeFormat string, utc bool) gin类.HandlerFunc {
	return GinzapWithConfig(logger, &Config{TimeFormat: timeFormat, UTC: utc, DefaultLevel: zapcore.InfoLevel})
}

// GinzapWithConfig returns a gin.HandlerFunc using configs
func GinzapWithConfig(logger ZapLogger, conf *Config) gin类.HandlerFunc {
	skipPaths := make(map[string]bool, len(conf.SkipPaths))
	for _, path := range conf.SkipPaths {
		skipPaths[path] = true
	}

	return func(c *gin类.Context) {
		start := time.Now()
		// some evil middlewares modify this values
		path := c.X请求.URL.Path
		query := c.X请求.URL.RawQuery
		c.X中间件继续()

		if _, ok := skipPaths[path]; !ok {
			end := time.Now()
			latency := end.Sub(start)
			if conf.UTC {
				end = end.UTC()
			}

			fields := []zapcore.Field{
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.X请求.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", c.X取客户端ip()),
				zap.String("user-agent", c.X请求.UserAgent()),
				zap.Duration("latency", latency),
			}
			if conf.TimeFormat != "" {
				fields = append(fields, zap.String("time", end.Format(conf.TimeFormat)))
			}

			if conf.Context != nil {
				fields = append(fields, conf.Context(c)...)
			}

			if len(c.X错误s) > 0 {
				// Append error field if this is an erroneous request.
				for _, e := range c.X错误s.Errors() {
					logger.Error(e, fields...)
				}
			} else {
                zl, ok := logger.(*zap.Logger)
                if ok {
                    zl.Log(conf.DefaultLevel, path, fields...)
                } else {
                    logger.Error(path, fields...)
                }
			}
		}
	}
}

func defaultHandleRecovery(c *gin类.Context, err interface{}) {
	c.X停止并带状态码(http.StatusInternalServerError)
}

// RecoveryWithZap returns a gin.HandlerFunc (middleware)
// that recovers from any panics and logs requests using uber-go/zap.
// All errors are logged using zap.Error().
// stack means whether output the stack info.
// The stack info is easy to find where the error occurs but the stack info is too large.
func RecoveryWithZap(logger ZapLogger, stack bool) gin类.HandlerFunc {
	return CustomRecoveryWithZap(logger, stack, defaultHandleRecovery)
}

// CustomRecoveryWithZap returns a gin.HandlerFunc (middleware) with a custom recovery handler
// that recovers from any panics and logs requests using uber-go/zap.
// All errors are logged using zap.Error().
// stack means whether output the stack info.
// The stack info is easy to find where the error occurs but the stack info is too large.
func CustomRecoveryWithZap(logger ZapLogger, stack bool, recovery gin类.RecoveryFunc) gin类.HandlerFunc {
	return func(c *gin类.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.X请求, false)
				if brokenPipe {
					logger.Error(c.X请求.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.X错误(err.(error)) // nolint: errcheck
					c.X停止()
					return
				}

				if stack {
					logger.Error("[Recovery from panic]",
						zap.Time("time", time.Now()),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Time("time", time.Now()),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				recovery(c, err)
			}
		}()
		c.X中间件继续()
	}
}
