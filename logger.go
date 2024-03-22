// Copyright 2014 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gin类

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	
	"github.com/mattn/go-isatty"
)

type consoleColorModeValue int

const (
	autoColor consoleColorModeValue = iota
	disableColor
	forceColor
)

const (
	green   = "\033[97;42m"
	white   = "\033[90;47m"
	yellow  = "\033[90;43m"
	red     = "\033[97;41m"
	blue    = "\033[97;44m"
	magenta = "\033[97;45m"
	cyan    = "\033[97;46m"
	reset   = "\033[0m"
)

var consoleColorMode = autoColor

// LoggerConfig defines the config for Logger middleware.
type LoggerConfig struct {
	// Optional. Default value is gin.defaultLogFormatter
	Formatter LogFormatter

	// Output is a writer where logs are written.
	// Optional. Default value is gin.DefaultWriter.
	Output io.Writer

	// SkipPaths is an url path array which logs are not written.
	// Optional.
	SkipPaths []string
}

// LogFormatter gives the signature of the formatter function passed to LoggerWithFormatter
type LogFormatter func(params LogFormatterParams) string

// LogFormatterParams is the structure any formatter will be handed when time to log comes
type LogFormatterParams struct {
	X请求 *http.Request

	// TimeStamp shows the time after the server returns a response.
	X响应时间 time.Time
	// StatusCode is HTTP response code.
	X状态码 int
	// Latency is how much time the server cost to process a certain request.
	X时长 time.Duration
	// ClientIP equals Context's ClientIP method.
	X客户端IP string
	// Method is the HTTP method given to the request.
	X方法 string
	// Path is a path the client requests.
	X路径 string
	// ErrorMessage is set if error has occurred in processing the request.
	X错误信息 string
	// isTerm shows whether gin's output descriptor refers to a terminal.
	X是否输出到终端 bool
	// BodySize is the size of the Response Body
	X响应体大小 int
	// Keys are the keys set on the request's context.
	X上下文设置值map map[string]any
}

// StatusCodeColor is the ANSI color for appropriately logging http status code to a terminal.
func (p *LogFormatterParams) StatusCodeColor() string {
	code := p.X状态码

	switch {
	case code >= http.StatusContinue && code < http.StatusOK:
		return white
	case code >= http.StatusOK && code < http.StatusMultipleChoices:
		return green
	case code >= http.StatusMultipleChoices && code < http.StatusBadRequest:
		return white
	case code >= http.StatusBadRequest && code < http.StatusInternalServerError:
		return yellow
	default:
		return red
	}
}

// MethodColor is the ANSI color for appropriately logging http method to a terminal.
func (p *LogFormatterParams) MethodColor() string {
	method := p.X方法

	switch method {
	case http.MethodGet:
		return blue
	case http.MethodPost:
		return cyan
	case http.MethodPut:
		return yellow
	case http.MethodDelete:
		return red
	case http.MethodPatch:
		return green
	case http.MethodHead:
		return magenta
	case http.MethodOptions:
		return white
	default:
		return reset
	}
}

// ResetColor resets all escape attributes.
func (p *LogFormatterParams) ResetColor() string {
	return reset
}

// IsOutputColor indicates whether can colors be outputted to the log.
func (p *LogFormatterParams) IsOutputColor() bool {
	return consoleColorMode == forceColor || (consoleColorMode == autoColor && p.X是否输出到终端)
}

// defaultLogFormatter is the default log format function Logger middleware uses.
var defaultLogFormatter = func(param LogFormatterParams) string {
	var statusColor, methodColor, resetColor string
	if param.IsOutputColor() {
		statusColor = param.StatusCodeColor()
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
	}

	if param.X时长 > time.Minute {
		param.X时长 = param.X时长.Truncate(time.Second)
	}
	return fmt.Sprintf("[GIN] %v |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
		param.X响应时间.Format("2006/01/02 - 15:04:05"),
		statusColor, param.X状态码, resetColor,
		param.X时长,
		param.X客户端IP,
		methodColor, param.X方法, resetColor,
		param.X路径,
		param.X错误信息,
	)
}

// DisableConsoleColor disables color output in the console.
func X关闭控制台颜色() {
	consoleColorMode = disableColor
}

// ForceConsoleColor force color output in the console.
func X开启控制台颜色() {
	consoleColorMode = forceColor
}

// ErrorLogger returns a HandlerFunc for any error type.
func ErrorLogger() HandlerFunc {
	return ErrorLoggerT(ErrorTypeAny)
}

// ErrorLoggerT returns a HandlerFunc for a given error type.
func ErrorLoggerT(typ ErrorType) HandlerFunc {
	return func(c *Context) {
		c.X中间件继续()
		errors := c.X错误s.ByType(typ)
		if len(errors) > 0 {
			c.X输出JSON(-1, errors)
		}
	}
}

// Logger instances a Logger middleware that will write the logs to gin.DefaultWriter.
// By default, gin.DefaultWriter = os.Stdout.
func Logger() HandlerFunc {
	return LoggerWithConfig(LoggerConfig{})
}

// LoggerWithFormatter instance a Logger middleware with the specified log format function.
func X中间件函数_自定义日志格式(格式化函数 LogFormatter) HandlerFunc {
	return LoggerWithConfig(LoggerConfig{
		Formatter: 格式化函数,
	})
}

// LoggerWithWriter instance a Logger middleware with the specified writer buffer.
// Example: os.Stdout, a file opened in write mode, a socket...
func LoggerWithWriter(out io.Writer, notlogged ...string) HandlerFunc {
	return LoggerWithConfig(LoggerConfig{
		Output:    out,
		SkipPaths: notlogged,
	})
}

// LoggerWithConfig instance a Logger middleware with config.
func LoggerWithConfig(conf LoggerConfig) HandlerFunc {
	formatter := conf.Formatter
	if formatter == nil {
		formatter = defaultLogFormatter
	}

	out := conf.Output
	if out == nil {
		out = DefaultWriter
	}

	notlogged := conf.SkipPaths

	isTerm := true

	if w, ok := out.(*os.File); !ok || os.Getenv("TERM") == "dumb" ||
		(!isatty.IsTerminal(w.Fd()) && !isatty.IsCygwinTerminal(w.Fd())) {
		isTerm = false
	}

	var skip map[string]struct{}

	if length := len(notlogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range notlogged {
			skip[path] = struct{}{}
		}
	}

	return func(c *Context) {
		// Start timer
		start := time.Now()
		path := c.X请求.URL.Path
		raw := c.X请求.URL.RawQuery

		// Process request
		c.X中间件继续()

		// Log only when path is not being skipped
		if _, ok := skip[path]; !ok {
			param := LogFormatterParams{
				X请求: c.X请求,
				X是否输出到终端:  isTerm,
				X上下文设置值map:    c.X上下文设置值Map,
			}

			// Stop timer
			param.X响应时间 = time.Now()
			param.X时长 = param.X响应时间.Sub(start)

			param.X客户端IP = c.X取客户端ip()
			param.X方法 = c.X请求.Method
			param.X状态码 = c.Writer.Status()
			param.X错误信息 = c.X错误s.ByType(ErrorTypePrivate).String()

			param.X响应体大小 = c.Writer.Size()

			if raw != "" {
				path = path + "?" + raw
			}

			param.X路径 = path

			fmt.Fprint(out, formatter(param))
		}
	}
}
