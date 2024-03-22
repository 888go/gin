// ginzap包提供了使用zap包进行日志处理的功能。
// 代码结构基于ginrus包。
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

// ZapLogger 是与 zap.Logger 兼容的最小日志器接口
type ZapLogger interface {
	Info(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
}

// Config 是 Ginzap 的配置设置
type Config struct {
	TimeFormat   string
	UTC          bool
	SkipPaths    []string
	Context      Fn
	DefaultLevel zapcore.Level
}

// Ginzap 返回一个 gin.HandlerFunc（中间件），该中间件使用 uber-go/zap 记录请求日志。
//
// 对于包含错误的请求，使用 zap.Error() 进行记录。
// 对于没有错误的请求，使用 zap.Info() 进行记录。
//
// 它接收以下参数：
//  1. 一个 time 包的时间格式字符串（例如 time.RFC3339）。
//  2. 一个布尔值，表示是否使用 UTC 时区或本地时区。
func Ginzap(logger ZapLogger, timeFormat string, utc bool) gin类.HandlerFunc {
	return GinzapWithConfig(logger, &Config{TimeFormat: timeFormat, UTC: utc, DefaultLevel: zapcore.InfoLevel})
}

// GinzapWithConfig 返回一个使用配置的 gin.HandlerFunc
func GinzapWithConfig(logger ZapLogger, conf *Config) gin类.HandlerFunc {
	skipPaths := make(map[string]bool, len(conf.SkipPaths))
	for _, path := range conf.SkipPaths {
		skipPaths[path] = true
	}

	return func(c *gin类.Context) {
		start := time.Now()
		// 一些恶意中间件会修改这些值
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
				// 如果这是错误请求，则追加错误字段。
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

// RecoveryWithZap 返回一个gin.HandlerFunc（中间件）
// 该中间件能够从任何panic中恢复，并使用uber-go/zap记录请求信息。
// 所有错误都会通过zap.Error()进行日志记录。
// stack 参数表示是否输出堆栈信息。
// 堆栈信息有助于快速定位错误发生的位置，但其体积较大。
func RecoveryWithZap(logger ZapLogger, stack bool) gin类.HandlerFunc {
	return CustomRecoveryWithZap(logger, stack, defaultHandleRecovery)
}

// CustomRecoveryWithZap 返回一个gin.HandlerFunc（中间件），它具有自定义恢复处理器，
// 可从任何panic中恢复，并使用uber-go/zap库记录请求信息。
// 所有错误都会通过zap.Error()方法进行日志记录。
// stack 参数表示是否输出堆栈信息。
// 堆栈信息有助于快速定位错误发生位置，但其信息量较大。
func CustomRecoveryWithZap(logger ZapLogger, stack bool, recovery gin类.RecoveryFunc) gin类.HandlerFunc {
	return func(c *gin类.Context) {
		defer func() {
			if err := recover(); err != nil {
// 检查连接是否已断开，因为这并不是真正需要引发恐慌并打印堆栈跟踪的条件。
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
					// 如果连接已断开，我们无法向其写入状态。
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
