
<原文开始>
// Package ginzap provides log handling using zap package.
// Code structure based on ginrus package.
<原文结束>

# <翻译开始>
// ginzap包提供了使用zap包进行日志处理的功能。
// 代码结构基于ginrus包。
# <翻译结束>


<原文开始>
// ZapLogger is the minimal logger interface compatible with zap.Logger
<原文结束>

# <翻译开始>
// ZapLogger 是与 zap.Logger 兼容的最小日志器接口
# <翻译结束>


<原文开始>
// Config is config setting for Ginzap
<原文结束>

# <翻译开始>
// Config 是 Ginzap 的配置设置
# <翻译结束>


<原文开始>
// Ginzap returns a gin.HandlerFunc (middleware) that logs requests using uber-go/zap.
//
// Requests with errors are logged using zap.Error().
// Requests without errors are logged using zap.Info().
//
// It receives:
//  1. A time package format string (e.g. time.RFC3339).
//  2. A boolean stating whether to use UTC time zone or local.
<原文结束>

# <翻译开始>
// Ginzap 返回一个 gin.HandlerFunc（中间件），该中间件使用 uber-go/zap 记录请求日志。
//
// 对于包含错误的请求，使用 zap.Error() 进行记录。
// 对于没有错误的请求，使用 zap.Info() 进行记录。
//
// 它接收以下参数：
//  1. 一个 time 包的时间格式字符串（例如 time.RFC3339）。
//  2. 一个布尔值，表示是否使用 UTC 时区或本地时区。
# <翻译结束>


<原文开始>
// GinzapWithConfig returns a gin.HandlerFunc using configs
<原文结束>

# <翻译开始>
// GinzapWithConfig 返回一个使用配置的 gin.HandlerFunc
# <翻译结束>


<原文开始>
// some evil middlewares modify this values
<原文结束>

# <翻译开始>
// 一些恶意中间件会修改这些值
# <翻译结束>


<原文开始>
// Append error field if this is an erroneous request.
<原文结束>

# <翻译开始>
// 如果这是错误请求，则追加错误字段。
# <翻译结束>


<原文开始>
// RecoveryWithZap returns a gin.HandlerFunc (middleware)
// that recovers from any panics and logs requests using uber-go/zap.
// All errors are logged using zap.Error().
// stack means whether output the stack info.
// The stack info is easy to find where the error occurs but the stack info is too large.
<原文结束>

# <翻译开始>
// RecoveryWithZap 返回一个gin.HandlerFunc（中间件）
// 该中间件能够从任何panic中恢复，并使用uber-go/zap记录请求信息。
// 所有错误都会通过zap.Error()进行日志记录。
// stack 参数表示是否输出堆栈信息。
// 堆栈信息有助于快速定位错误发生的位置，但其体积较大。
# <翻译结束>


<原文开始>
// CustomRecoveryWithZap returns a gin.HandlerFunc (middleware) with a custom recovery handler
// that recovers from any panics and logs requests using uber-go/zap.
// All errors are logged using zap.Error().
// stack means whether output the stack info.
// The stack info is easy to find where the error occurs but the stack info is too large.
<原文结束>

# <翻译开始>
// CustomRecoveryWithZap 返回一个gin.HandlerFunc（中间件），它具有自定义恢复处理器，
// 可从任何panic中恢复，并使用uber-go/zap库记录请求信息。
// 所有错误都会通过zap.Error()方法进行日志记录。
// stack 参数表示是否输出堆栈信息。
// 堆栈信息有助于快速定位错误发生位置，但其信息量较大。
# <翻译结束>


<原文开始>
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
<原文结束>

# <翻译开始>
// 检查连接是否已断开，因为这并不是真正需要引发恐慌并打印堆栈跟踪的条件。
# <翻译结束>


<原文开始>
// If the connection is dead, we can't write a status to it.
<原文结束>

# <翻译开始>
// 如果连接已断开，我们无法向其写入状态。
# <翻译结束>

