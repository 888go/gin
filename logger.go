// 版权所有 2014 Manu Martinez-Almeida。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package gin

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

// LoggerConfig 定义了 Logger 中间件的配置。
type LoggerConfig struct {
	// 可选的。默认值为gin.defaultLogFormatter
	Formatter LogFormatter

// Output 是一个用于写入日志的writer。
// 可选配置，默认值为gin.DefaultWriter。
	Output io.Writer

// SkipPaths 是一个 URL 路径数组，其中的请求日志不会被记录。
// 可选配置。
	SkipPaths []string
}

// LogFormatter 提供了传递给 LoggerWithFormatter 的格式化函数的签名
type LogFormatter func(params LogFormatterParams) string

// LogFormatterParams 是一个结构体，当需要进行日志记录时，任何格式化器都将接收到这个结构体作为参数
type LogFormatterParams struct {
	Request *http.Request

	// TimeStamp 表示服务器返回响应后的时刻。
	TimeStamp time.Time
	// StatusCode 是HTTP响应代码。
	StatusCode int
	// 延迟是服务器处理特定请求所需的时间。
	Latency time.Duration
	// ClientIP 等同于 Context 的 ClientIP 方法。
	ClientIP string
	// Method是请求中给定的HTTP方法。
	Method string
	// Path 是客户端请求的路径。
	Path string
	// ErrorMessage在处理请求时发生错误时设置。
	ErrorMessage string
	// isTerm 判断 gin 的输出描述符是否指向一个终端。
	isTerm bool
	// BodySize 是 Response Body 的大小
	BodySize int
	// Keys 是在请求的上下文中设置的键。
	Keys map[string]any
}

// StatusCodeColor 是用于将 HTTP 状态码适当地以 ANSI 颜色格式输出到终端的。
func (p *LogFormatterParams) StatusCodeColor() string {
	code := p.StatusCode

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

// MethodColor 是用于适当地将HTTP方法以颜色格式输出到终端的ANSI颜色。
func (p *LogFormatterParams) MethodColor() string {
	method := p.Method

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

// ResetColor 重置所有转义属性。
func (p *LogFormatterParams) ResetColor() string {
	return reset
}

// IsOutputColor 指示是否可以在日志中输出颜色。
func (p *LogFormatterParams) IsOutputColor() bool {
	return consoleColorMode == forceColor || (consoleColorMode == autoColor && p.isTerm)
}

// defaultLogFormatter 是 Logger 中间件默认使用的日志格式化函数。
var defaultLogFormatter = func(param LogFormatterParams) string {
	var statusColor, methodColor, resetColor string
	if param.IsOutputColor() {
		statusColor = param.StatusCodeColor()
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
	}

	if param.Latency > time.Minute {
		param.Latency = param.Latency.Truncate(time.Second)
	}
	return fmt.Sprintf("[GIN] %v |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		statusColor, param.StatusCode, resetColor,
		param.Latency,
		param.ClientIP,
		methodColor, param.Method, resetColor,
		param.Path,
		param.ErrorMessage,
	)
}

// DisableConsoleColor 禁用控制台中的颜色输出。
func DisableConsoleColor() {
	consoleColorMode = disableColor
}

// ForceConsoleColor 强制在控制台输出彩色内容
func ForceConsoleColor() {
	consoleColorMode = forceColor
}

// ErrorLogger 返回一个适用于任何错误类型的 HandlerFunc。
func ErrorLogger() HandlerFunc {
	return ErrorLoggerT(ErrorTypeAny)
}

// ErrorLoggerT为给定的错误类型返回一个HandlerFunc。
func ErrorLoggerT(typ ErrorType) HandlerFunc {
	return func(c *Context) {
		c.Next()
		errors := c.Errors.ByType(typ)
		if len(errors) > 0 {
			c.JSON(-1, errors)
		}
	}
}

// Logger 创建一个 Logger 中间件，该中间件会将日志写入 gin.DefaultWriter。
// 默认情况下，gin.DefaultWriter = os.Stdout。
func Logger() HandlerFunc {
	return LoggerWithConfig(LoggerConfig{})
}

// LoggerWithFormatter 根据指定的日志格式化函数实例化一个 Logger 中间件。
func LoggerWithFormatter(f LogFormatter) HandlerFunc {
	return LoggerWithConfig(LoggerConfig{
		Formatter: f,
	})
}

// LoggerWithWriter 通过指定的写入器缓冲区实例化一个 Logger 中间件。
// 示例：os.Stdout（标准输出），以写入模式打开的文件，套接字等...
func LoggerWithWriter(out io.Writer, notlogged ...string) HandlerFunc {
	return LoggerWithConfig(LoggerConfig{
		Output:    out,
		SkipPaths: notlogged,
	})
}

// LoggerWithConfig 通过配置实例化一个 Logger 中间件。
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
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// 当路径未被跳过时才记录日志
		if _, ok := skip[path]; !ok {
			param := LogFormatterParams{
				Request: c.Request,
				isTerm:  isTerm,
				Keys:    c.Keys,
			}

			// Stop timer
			param.TimeStamp = time.Now()
			param.Latency = param.TimeStamp.Sub(start)

			param.ClientIP = c.ClientIP()
			param.Method = c.Request.Method
			param.StatusCode = c.Writer.Status()
			param.ErrorMessage = c.Errors.ByType(ErrorTypePrivate).String()

			param.BodySize = c.Writer.Size()

			if raw != "" {
				path = path + "?" + raw
			}

			param.Path = path

			fmt.Fprint(out, formatter(param))
		}
	}
}
