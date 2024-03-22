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

type Fn func(*gin类.Context, zerolog.Logger) zerolog.Logger

// Config 定义了 logger 中间件的配置
type config struct {
	logger Fn
	// UTC 是一个布尔值，表示是否使用 UTC 时区或本地时区。
	utc             bool
	skipPath        []string
	skipPathRegexps []*regexp.Regexp
// Output 是一个用于写入日志的writer。
// 可选配置，默认值为gin.DefaultWriter。
	output io.Writer
	// 用于状态码小于400的请求的日志级别
	defaultLevel zerolog.Level
	// 用于状态码在400至499之间的请求的日志级别
	clientErrorLevel zerolog.Level
	// 用于状态码大于等于500的请求的日志级别
	serverErrorLevel zerolog.Level
}

var isTerm bool = isatty.IsTerminal(os.Stdout.Fd())

// SetLogger 初始化日志中间件。
func SetLogger(opts ...Option) gin类.HandlerFunc {
	cfg := &config{
		defaultLevel:     zerolog.InfoLevel,
		clientErrorLevel: zerolog.WarnLevel,
		serverErrorLevel: zerolog.ErrorLevel,
		output:           gin类.DefaultWriter,
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
	return func(c *gin类.Context) {
		if cfg.logger != nil {
			l = cfg.logger(c, l)
		}

		start := time.Now()
		path := c.X请求.URL.Path
		raw := c.X请求.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}

		c.X中间件继续()
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
				Str("method", c.X请求.Method).
				Str("path", path).
				Str("ip", c.X取客户端ip()).
				Dur("latency", latency).
				Str("user_agent", c.X请求.UserAgent()).Logger()

			msg := "Request"
			if len(c.X错误s) > 0 {
				msg = c.X错误s.String()
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
func ParseLevel(levelStr string) (zerolog.Level, error) {
	return zerolog.ParseLevel(levelStr)
}
