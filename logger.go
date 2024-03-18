// Manu Martinez-Almeida版权所有
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到

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

// LoggerConfig定义了Logger中间件的配置
type LoggerConfig struct {
// 可选的
// 默认值为gin.defaultLogFormatter
	Formatter LogFormatter

// Output是写入日志的写入器
// 可选的
// 默认值为gin. defaultwwriter
	Output io.Writer

// skipppaths是一个url路径数组，不写入日志
// 可选的
	SkipPaths []string
}

// LogFormatter给出传递给LoggerWithFormatter的formatter函数的签名
type LogFormatter func(params LogFormatterParams) string

// LogFormatterParams是任何格式化程序在需要进行日志记录时要传递的结构
type LogFormatterParams struct {
	Request *http.Request

// TimeStamp显示服务器返回响应后的时间
	TimeStamp time.Time
// StatusCode是HTTP响应码
	StatusCode int
// 延迟是服务器处理某个请求所需的时间
	Latency time.Duration
// ClientIP等于Context的ClientIP方法
	ClientIP string
// 方法是给定给请求的HTTP方法
	Method string
// Path是客户端请求的路径
	Path string
// 如果在处理请求时发生错误，则设置ErrorMessage
	ErrorMessage string
// isTerm显示gin的输出描述符是否指向终端
	isTerm bool
// BodySize是响应体的大小
	BodySize int
// 键是在请求的上下文中设置的键
	Keys map[string]any
}

// StatusCodeColor是用于将http状态码适当地记录到终端的ANSI颜色

// ff:
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

// MethodColor是用于将http方法适当地记录到终端的ANSI颜色

// ff:
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

// ResetColor重置所有转义属性

// ff:
func (p *LogFormatterParams) ResetColor() string {
	return reset
}

// IsOutputColor是否可以输出颜色到日志中

// ff:
func (p *LogFormatterParams) IsOutputColor() bool {
	return consoleColorMode == forceColor || (consoleColorMode == autoColor && p.isTerm)
}

// defaultLogFormatter是Logger中间件使用的默认日志格式函数
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

// DisableConsoleColor禁用控制台的颜色输出

// ff:关闭控制台颜色
func DisableConsoleColor() {
	consoleColorMode = disableColor
}

// ForceConsoleColor强制控制台的颜色输出

// ff:开启控制台颜色
func ForceConsoleColor() {
	consoleColorMode = forceColor
}

// ErrorLogger为任何错误类型返回一个HandlerFunc

// ff:
func ErrorLogger() HandlerFunc {
	return ErrorLoggerT(ErrorTypeAny)
}

// ErrorLoggerT返回给定错误类型的HandlerFunc

// ff:
// typ:
func ErrorLoggerT(typ ErrorType) HandlerFunc {
	return func(c *Context) {
		c.Next()
		errors := c.Errors.ByType(typ)
		if len(errors) > 0 {
			c.JSON(-1, errors)
		}
	}
}

// Logger实例化一个Logger中间件，它将把日志写入gin. defaultwwriter
// 缺省为gin
// defaultwwriter = os.Stdout

// ff:
func Logger() HandlerFunc {
	return LoggerWithConfig(LoggerConfig{})
}

// LoggerWithFormatter实例:一个具有指定日志格式功能的Logger中间件

// ff:
// f:
func LoggerWithFormatter(f LogFormatter) HandlerFunc {
	return LoggerWithConfig(LoggerConfig{
		Formatter: f,
	})
}

// LoggerWithWriter实例:一个具有指定写入器缓冲区的Logger中间件
// 例如:操作系统
// 标准输出，以写模式打开的文件，套接字…

// ff:
// notlogged:
// out:
func LoggerWithWriter(out io.Writer, notlogged ...string) HandlerFunc {
	return LoggerWithConfig(LoggerConfig{
		Output:    out,
		SkipPaths: notlogged,
	})
}

// LoggerWithConfig实例是一个带有config的Logger中间件

// ff:
// conf:
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
// 启动定时器
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

// 处理请求
		c.Next()

// 仅当路径未被跳过时记录日志
		if _, ok := skip[path]; !ok {
			param := LogFormatterParams{
				Request: c.Request,
				isTerm:  isTerm,
				Keys:    c.Keys,
			}

// 停止计时器
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
