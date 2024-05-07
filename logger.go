// 版权所有 2014 Manu Martinez-Almeida。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

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

// LoggerConfig 定义了 Logger 中间件的配置。
type LoggerConfig struct {
	// 可选的。默认值为gin.defaultLogFormatter
	Formatter LogFormatter

	// Output 是一个用于写入日志的writer。
	// 可选配置，默认值为gin.DefaultWriter。
	Output io.Writer

	// SkipPaths 是一个 URL 路径切片，其中的请求日志不会被记录。
	// 可选配置。
	SkipPaths []string
}

// LogFormatter 提供了传递给 LoggerWithFormatter 的格式化函数的签名
type LogFormatter func(params LogFormatterParams) string

// LogFormatterParams 是一个结构体，当需要进行日志记录时，任何格式化器都将接收到这个结构体作为参数
type LogFormatterParams struct {
	X请求 *http.Request

	// TimeStamp 表示服务器返回响应后的时刻。
	X响应时间 time.Time
	// StatusCode 是HTTP响应代码。
	X状态码 int
	// 延迟是服务器处理特定请求所需的时间。
	X时长 time.Duration
	// ClientIP 等同于 Context 的 ClientIP 方法。
	X客户端IP string
	// Method是请求中给定的HTTP方法。
	X方法 string
	// Path 是客户端请求的路径。
	X路径 string
	// ErrorMessage在处理请求时发生错误时设置。
	X错误信息 string
	// isTerm 判断 gin 的输出描述符是否指向一个终端。
	X是否输出到终端 bool
	// BodySize 是 Response Body 的大小
	X响应体大小 int
	// Keys 是在请求的上下文中设置的键。
	X上下文设置值map map[string]any
}

// StatusCodeColor 是用于将 HTTP 状态码适当地以 ANSI 颜色格式输出到终端的。
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

// MethodColor 是用于适当地将HTTP方法以颜色格式输出到终端的ANSI颜色。
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

// ResetColor 重置所有转义属性。
func (p *LogFormatterParams) ResetColor() string {
	return reset
}

// IsOutputColor 指示是否可以在日志中输出颜色。
func (p *LogFormatterParams) IsOutputColor() bool {
	return consoleColorMode == forceColor || (consoleColorMode == autoColor && p.X是否输出到终端)
}

// defaultLogFormatter 是 Logger 中间件默认使用的日志格式化函数。
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

// DisableConsoleColor 禁用控制台中的颜色输出。
func X关闭控制台颜色() {
	consoleColorMode = disableColor
}

// ForceConsoleColor 强制在控制台输出彩色内容
func X开启控制台颜色() {
	consoleColorMode = forceColor
}

// ErrorLogger 返回一个适用于任何错误类型的 HandlerFunc。
func ErrorLogger() HandlerFunc {
	return ErrorLoggerT(ErrorTypeAny)
}

// ErrorLoggerT为给定的错误类型返回一个HandlerFunc。
func ErrorLoggerT(typ ErrorType) HandlerFunc {
	return func(c *Context) {
		c.X中间件继续()
		errors := c.X错误s.ByType(typ)
		if len(errors) > 0 {
			c.X输出JSON(-1, errors)
		}
	}
}

// Logger 创建一个 Logger 中间件，该中间件会将日志写入 gin.DefaultWriter。
// 默认情况下，gin.DefaultWriter = os.Stdout。
func Logger() HandlerFunc {
	return LoggerWithConfig(LoggerConfig{})
}

// LoggerWithFormatter 根据指定的日志格式化函数实例化一个 Logger 中间件。
func X中间件函数_自定义日志格式(格式化函数 LogFormatter) HandlerFunc {
	return LoggerWithConfig(LoggerConfig{
		Formatter: 格式化函数,
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
		path := c.X请求.URL.Path
		raw := c.X请求.URL.RawQuery

		// Process request
		c.X中间件继续()

		// 当路径未被跳过时才记录日志
		if _, ok := skip[path]; !ok {
			param := LogFormatterParams{
				X请求:        c.X请求,
				X是否输出到终端:   isTerm,
				X上下文设置值map: c.X上下文设置值Map,
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
