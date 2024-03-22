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

// Config defines the config for logger middleware
type config struct {
	logger Fn
	// UTC a boolean stating whether to use UTC time zone or local.
	utc             bool
	skipPath        []string
	skipPathRegexps []*regexp.Regexp
	// Output is a writer where logs are written.
	// Optional. Default value is gin.DefaultWriter.
	output io.Writer
	// the log level used for request with status code < 400
	defaultLevel zerolog.Level
	// the log level used for request with status code between 400 and 499
	clientErrorLevel zerolog.Level
	// the log level used for request with status code >= 500
	serverErrorLevel zerolog.Level
}

var isTerm bool = isatty.IsTerminal(os.Stdout.Fd())

// SetLogger initializes the logging middleware.
func SetLogger(opts ...Option) gin类.HandlerFunc {
	cfg := &config{
		defaultLevel:     zerolog.InfoLevel,
		clientErrorLevel: zerolog.WarnLevel,
		serverErrorLevel: zerolog.ErrorLevel,
		output:           gin类.DefaultWriter,
	}

	// Loop through each option
	for _, o := range opts {
		// Call the option giving the instantiated
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

// ParseLevel converts a level string into a zerolog Level value.
// returns an error if the input string does not match known values.
func ParseLevel(levelStr string) (zerolog.Level, error) {
	return zerolog.ParseLevel(levelStr)
}
