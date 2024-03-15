
<原文开始>
// Copyright 2014 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Manu Martinez-Almeida版权所有
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到
# <翻译结束>


<原文开始>
// LoggerConfig defines the config for Logger middleware.
<原文结束>

# <翻译开始>
// LoggerConfig定义了Logger中间件的配置
# <翻译结束>


<原文开始>
	// Optional. Default value is gin.defaultLogFormatter
<原文结束>

# <翻译开始>
// 可选的
// 默认值为gin.defaultLogFormatter
# <翻译结束>


<原文开始>
	// Output is a writer where logs are written.
	// Optional. Default value is gin.DefaultWriter.
<原文结束>

# <翻译开始>
// Output是写入日志的写入器
// 可选的
// 默认值为gin. defaultwwriter
# <翻译结束>


<原文开始>
	// SkipPaths is an url path array which logs are not written.
	// Optional.
<原文结束>

# <翻译开始>
// skipppaths是一个url路径数组，不写入日志
// 可选的
# <翻译结束>


<原文开始>
// LogFormatter gives the signature of the formatter function passed to LoggerWithFormatter
<原文结束>

# <翻译开始>
// LogFormatter给出传递给LoggerWithFormatter的formatter函数的签名
# <翻译结束>


<原文开始>
// LogFormatterParams is the structure any formatter will be handed when time to log comes
<原文结束>

# <翻译开始>
// LogFormatterParams是任何格式化程序在需要进行日志记录时要传递的结构
# <翻译结束>


<原文开始>
	// TimeStamp shows the time after the server returns a response.
<原文结束>

# <翻译开始>
// TimeStamp显示服务器返回响应后的时间
# <翻译结束>


<原文开始>
	// StatusCode is HTTP response code.
<原文结束>

# <翻译开始>
// StatusCode是HTTP响应码
# <翻译结束>


<原文开始>
	// Latency is how much time the server cost to process a certain request.
<原文结束>

# <翻译开始>
// 延迟是服务器处理某个请求所需的时间
# <翻译结束>


<原文开始>
	// ClientIP equals Context's ClientIP method.
<原文结束>

# <翻译开始>
// ClientIP等于Context的ClientIP方法
# <翻译结束>


<原文开始>
	// Method is the HTTP method given to the request.
<原文结束>

# <翻译开始>
// 方法是给定给请求的HTTP方法
# <翻译结束>


<原文开始>
	// Path is a path the client requests.
<原文结束>

# <翻译开始>
// Path是客户端请求的路径
# <翻译结束>


<原文开始>
	// ErrorMessage is set if error has occurred in processing the request.
<原文结束>

# <翻译开始>
// 如果在处理请求时发生错误，则设置ErrorMessage
# <翻译结束>


<原文开始>
	// isTerm shows whether gin's output descriptor refers to a terminal.
<原文结束>

# <翻译开始>
// isTerm显示gin的输出描述符是否指向终端
# <翻译结束>


<原文开始>
	// BodySize is the size of the Response Body
<原文结束>

# <翻译开始>
// BodySize是响应体的大小
# <翻译结束>


<原文开始>
	// Keys are the keys set on the request's context.
<原文结束>

# <翻译开始>
// 键是在请求的上下文中设置的键
# <翻译结束>


<原文开始>
// StatusCodeColor is the ANSI color for appropriately logging http status code to a terminal.
<原文结束>

# <翻译开始>
// StatusCodeColor是用于将http状态码适当地记录到终端的ANSI颜色
# <翻译结束>


<原文开始>
// MethodColor is the ANSI color for appropriately logging http method to a terminal.
<原文结束>

# <翻译开始>
// MethodColor是用于将http方法适当地记录到终端的ANSI颜色
# <翻译结束>


<原文开始>
// ResetColor resets all escape attributes.
<原文结束>

# <翻译开始>
// ResetColor重置所有转义属性
# <翻译结束>


<原文开始>
// IsOutputColor indicates whether can colors be outputted to the log.
<原文结束>

# <翻译开始>
// IsOutputColor是否可以输出颜色到日志中
# <翻译结束>


<原文开始>
// defaultLogFormatter is the default log format function Logger middleware uses.
<原文结束>

# <翻译开始>
// defaultLogFormatter是Logger中间件使用的默认日志格式函数
# <翻译结束>


<原文开始>
// DisableConsoleColor disables color output in the console.
<原文结束>

# <翻译开始>
// DisableConsoleColor禁用控制台的颜色输出
# <翻译结束>


<原文开始>
// ForceConsoleColor force color output in the console.
<原文结束>

# <翻译开始>
// ForceConsoleColor强制控制台的颜色输出
# <翻译结束>


<原文开始>
// ErrorLogger returns a HandlerFunc for any error type.
<原文结束>

# <翻译开始>
// ErrorLogger为任何错误类型返回一个HandlerFunc
# <翻译结束>


<原文开始>
// ErrorLoggerT returns a HandlerFunc for a given error type.
<原文结束>

# <翻译开始>
// ErrorLoggerT返回给定错误类型的HandlerFunc
# <翻译结束>


<原文开始>
// Logger instances a Logger middleware that will write the logs to gin.DefaultWriter.
// By default, gin.DefaultWriter = os.Stdout.
<原文结束>

# <翻译开始>
// Logger实例化一个Logger中间件，它将把日志写入gin. defaultwwriter
// 缺省为gin
// defaultwwriter = os.Stdout
# <翻译结束>


<原文开始>
// LoggerWithFormatter instance a Logger middleware with the specified log format function.
<原文结束>

# <翻译开始>
// LoggerWithFormatter实例:一个具有指定日志格式功能的Logger中间件
# <翻译结束>


<原文开始>
// LoggerWithWriter instance a Logger middleware with the specified writer buffer.
// Example: os.Stdout, a file opened in write mode, a socket...
<原文结束>

# <翻译开始>
// LoggerWithWriter实例:一个具有指定写入器缓冲区的Logger中间件
// 例如:操作系统
// 标准输出，以写模式打开的文件，套接字…
# <翻译结束>


<原文开始>
// LoggerWithConfig instance a Logger middleware with config.
<原文结束>

# <翻译开始>
// LoggerWithConfig实例是一个带有config的Logger中间件
# <翻译结束>


<原文开始>
		// Start timer
<原文结束>

# <翻译开始>
// 启动定时器
# <翻译结束>


<原文开始>
		// Process request
<原文结束>

# <翻译开始>
// 处理请求
# <翻译结束>


<原文开始>
		// Log only when path is not being skipped
<原文结束>

# <翻译开始>
// 仅当路径未被跳过时记录日志
# <翻译结束>


<原文开始>
			// Stop timer
<原文结束>

# <翻译开始>
// 停止计时器
# <翻译结束>

