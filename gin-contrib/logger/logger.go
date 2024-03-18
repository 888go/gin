package logger

import (
	"io"
	"net/http"
	"os"
	"regexp"
	"time"
	
	"github.com/888go/gin"
	"github.com/mattn/go-isatty"
	"github.com/rs/zerolog"
)

type Fn func(*gin.Context, zerolog.Logger) zerolog.Logger

// Config 定义了日志中间件的配置
type config struct {
	logger Fn
// UTC 是一个布尔值，表示是否使用 UTC 时区还是本地时区。
	utc             bool
	skipPath        []string
	skipPathRegexps []*regexp.Regexp
// Output 是一个用于写入日志的 writer。
// 可选配置项，默认值为 gin.DefaultWriter。
	output io.Writer
// 用于状态码小于400的请求的日志级别
	defaultLevel zerolog.Level
// 用于状态码在400到499之间的请求的日志级别
	clientErrorLevel zerolog.Level
// 用于状态码大于等于500的请求的日志级别
	serverErrorLevel zerolog.Level
}

var isTerm bool = isatty.IsTerminal(os.Stdout.Fd())

// SetLogger 初始化日志中间件。

// ff:
// opts:

// ff:
// opts:

// ff:
// opts:

// ff:
// opts:

// ff:
// opts:

// ff:
// opts:

// ff:
// opts:
func SetLogger(opts ...Option) gin.HandlerFunc {
	cfg := &config{
		defaultLevel:     zerolog.InfoLevel,
		clientErrorLevel: zerolog.WarnLevel,
		serverErrorLevel: zerolog.ErrorLevel,
		output:           gin.DefaultWriter,
	}

// 遍历每个选项
	for _, o := range opts {
// 调用选项，传入已实例化的
		o.apply(cfg)
	}

	var skip map[string]struct{}
	if length := len(cfg.skipPath); length > 0 {
		skip = make(map[string]struct{}, length)
		for _, path := range cfg.skipPath {
			skip[path] = struct{}{}
		}
	}

	l := zerolog.New(cfg.output).
		Output(
			zerolog.ConsoleWriter{
				Out:     cfg.output,
				NoColor: !isTerm,
			},
		).
		With().
		Timestamp().
		Logger()
	return func(c *gin.Context) {
		if cfg.logger != nil {
			l = cfg.logger(c, l)
		}

		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}

		c.Next()
		track := true

		if _, ok := skip[path]; ok {
			track = false
		}

		if track && len(cfg.skipPathRegexps) > 0 {
			for _, reg := range cfg.skipPathRegexps {
				if !reg.MatchString(path) {
					continue
				}

				track = false
				break
			}
		}

		if track {
			end := time.Now()
			if cfg.utc {
				end = end.UTC()
			}
			latency := end.Sub(start)

			l = l.With().
				Int("status", c.Writer.Status()).
				Str("method", c.Request.Method).
				Str("path", path).
				Str("ip", c.ClientIP()).
				Dur("latency", latency).
				Str("user_agent", c.Request.UserAgent()).Logger()

			msg := "Request"
			if len(c.Errors) > 0 {
				msg = c.Errors.String()
			}

			switch {
			case c.Writer.Status() >= http.StatusBadRequest && c.Writer.Status() < http.StatusInternalServerError:
				{
					l.WithLevel(cfg.clientErrorLevel).
						Msg(msg)
				}
			case c.Writer.Status() >= http.StatusInternalServerError:
				{
					l.WithLevel(cfg.serverErrorLevel).
						Msg(msg)
				}
			default:
				l.WithLevel(cfg.defaultLevel).
					Msg(msg)
			}
		}
	}
}

// ParseLevel将级别字符串转换为zerolog的Level值。
// 如果输入字符串与已知值不匹配，则返回错误。

// ff:
// zerolog.Level:
// levelStr:

// ff:
// zerolog.Level:
// levelStr:

// ff:
// zerolog.Level:
// levelStr:

// ff:
// zerolog.Level:
// levelStr:

// ff:
// zerolog.Level:
// levelStr:

// ff:
// zerolog.Level:
// levelStr:

// ff:
// zerolog.Level:
// levelStr:
func ParseLevel(levelStr string) (zerolog.Level, error) {
	return zerolog.ParseLevel(levelStr)
}
