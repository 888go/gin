
<原文开始>
// Copyright 2014 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2014 Manu Martinez-Almeida。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// LoggerConfig defines the config for Logger middleware.
<原文结束>

# <翻译开始>
// LoggerConfig 定义了 Logger 中间件的配置。
# <翻译结束>


<原文开始>
// Optional. Default value is gin.defaultLogFormatter
<原文结束>

# <翻译开始>
// 可选的。默认值为gin.defaultLogFormatter
# <翻译结束>


<原文开始>
	// Output is a writer where logs are written.
	// Optional. Default value is gin.DefaultWriter.
<原文结束>

# <翻译开始>
// Output 是一个用于写入日志的writer。
// 可选配置，默认值为gin.DefaultWriter。
# <翻译结束>


<原文开始>
	// SkipPaths is an url path array which logs are not written.
	// Optional.
<原文结束>

# <翻译开始>
// SkipPaths 是一个 URL 路径切片，其中的请求日志不会被记录。
// 可选配置。
# <翻译结束>


<原文开始>
// LogFormatter gives the signature of the formatter function passed to LoggerWithFormatter
<原文结束>

# <翻译开始>
// LogFormatter 提供了传递给 LoggerWithFormatter 的格式化函数的签名
# <翻译结束>


<原文开始>
// LogFormatterParams is the structure any formatter will be handed when time to log comes
<原文结束>

# <翻译开始>
// LogFormatterParams 是一个结构体，当需要进行日志记录时，任何格式化器都将接收到这个结构体作为参数
# <翻译结束>


<原文开始>
// TimeStamp shows the time after the server returns a response.
<原文结束>

# <翻译开始>
// TimeStamp 表示服务器返回响应后的时刻。
# <翻译结束>


<原文开始>
// StatusCode is HTTP response code.
<原文结束>

# <翻译开始>
// StatusCode 是HTTP响应代码。
# <翻译结束>


<原文开始>
// Latency is how much time the server cost to process a certain request.
<原文结束>

# <翻译开始>
// 延迟是服务器处理特定请求所需的时间。
# <翻译结束>


<原文开始>
// ClientIP equals Context's ClientIP method.
<原文结束>

# <翻译开始>
// ClientIP 等同于 Context 的 ClientIP 方法。
# <翻译结束>


<原文开始>
// Method is the HTTP method given to the request.
<原文结束>

# <翻译开始>
// Method是请求中给定的HTTP方法。
# <翻译结束>


<原文开始>
// Path is a path the client requests.
<原文结束>

# <翻译开始>
// Path 是客户端请求的路径。
# <翻译结束>


<原文开始>
// ErrorMessage is set if error has occurred in processing the request.
<原文结束>

# <翻译开始>
// ErrorMessage在处理请求时发生错误时设置。
# <翻译结束>


<原文开始>
// isTerm shows whether gin's output descriptor refers to a terminal.
<原文结束>

# <翻译开始>
// isTerm 判断 gin 的输出描述符是否指向一个终端。
# <翻译结束>


<原文开始>
// BodySize is the size of the Response Body
<原文结束>

# <翻译开始>
// BodySize 是 Response Body 的大小
# <翻译结束>


<原文开始>
// Keys are the keys set on the request's context.
<原文结束>

# <翻译开始>
// Keys 是在请求的上下文中设置的键。
# <翻译结束>


<原文开始>
// StatusCodeColor is the ANSI color for appropriately logging http status code to a terminal.
<原文结束>

# <翻译开始>
// StatusCodeColor 是用于将 HTTP 状态码适当地以 ANSI 颜色格式输出到终端的。
# <翻译结束>


<原文开始>
// MethodColor is the ANSI color for appropriately logging http method to a terminal.
<原文结束>

# <翻译开始>
// MethodColor 是用于适当地将HTTP方法以颜色格式输出到终端的ANSI颜色。
# <翻译结束>


<原文开始>
// ResetColor resets all escape attributes.
<原文结束>

# <翻译开始>
// ResetColor 重置所有转义属性。
# <翻译结束>


<原文开始>
// IsOutputColor indicates whether can colors be outputted to the log.
<原文结束>

# <翻译开始>
// IsOutputColor 指示是否可以在日志中输出颜色。
# <翻译结束>


<原文开始>
// defaultLogFormatter is the default log format function Logger middleware uses.
<原文结束>

# <翻译开始>
// defaultLogFormatter 是 Logger 中间件默认使用的日志格式化函数。
# <翻译结束>


<原文开始>
// DisableConsoleColor disables color output in the console.
<原文结束>

# <翻译开始>
// DisableConsoleColor 禁用控制台中的颜色输出。
# <翻译结束>


<原文开始>
// ForceConsoleColor force color output in the console.
<原文结束>

# <翻译开始>
// ForceConsoleColor 强制在控制台输出彩色内容
# <翻译结束>


<原文开始>
// ErrorLogger returns a HandlerFunc for any error type.
<原文结束>

# <翻译开始>
// ErrorLogger 返回一个适用于任何错误类型的 HandlerFunc。
# <翻译结束>


<原文开始>
// ErrorLoggerT returns a HandlerFunc for a given error type.
<原文结束>

# <翻译开始>
// ErrorLoggerT为给定的错误类型返回一个HandlerFunc。
# <翻译结束>


<原文开始>
// Logger instances a Logger middleware that will write the logs to gin.DefaultWriter.
// By default, gin.DefaultWriter = os.Stdout.
<原文结束>

# <翻译开始>
// Logger 创建一个 Logger 中间件，该中间件会将日志写入 gin.DefaultWriter。
// 默认情况下，gin.DefaultWriter = os.Stdout。
# <翻译结束>


<原文开始>
// LoggerWithFormatter instance a Logger middleware with the specified log format function.
<原文结束>

# <翻译开始>
// LoggerWithFormatter 根据指定的日志格式化函数实例化一个 Logger 中间件。
# <翻译结束>


<原文开始>
// LoggerWithWriter instance a Logger middleware with the specified writer buffer.
// Example: os.Stdout, a file opened in write mode, a socket...
<原文结束>

# <翻译开始>
// LoggerWithWriter 通过指定的写入器缓冲区实例化一个 Logger 中间件。
// 示例：os.Stdout（标准输出），以写入模式打开的文件，套接字等...
# <翻译结束>


<原文开始>
// LoggerWithConfig instance a Logger middleware with config.
<原文结束>

# <翻译开始>
// LoggerWithConfig 通过配置实例化一个 Logger 中间件。
# <翻译结束>


<原文开始>
// Log only when path is not being skipped
<原文结束>

# <翻译开始>
// 当路径未被跳过时才记录日志
# <翻译结束>

