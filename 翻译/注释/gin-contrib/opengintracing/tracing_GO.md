
<原文开始>
// Package opengintracing provides requests tracing functional using opentracing specification.
//
// See https://github.com/opentracing/opentracing-go for more information
<原文结束>

# <翻译开始>
// Package opengintracing 提供了基于 opentracing 规范的请求跟踪功能。
//
// 有关更多信息，请参见 https://github.com/opentracing/opentracing-go
# <翻译结束>


<原文开始>
// Errors which may occur at operation time.
<原文结束>

# <翻译开始>
// 运行时可能会发生的错误。
# <翻译结束>


<原文开始>
// NewSpan returns gin.HandlerFunc (middleware) that starts a new span and injects it to request context.
//
// It calls ctx.Next() to measure execution time of all following handlers.
<原文结束>

# <翻译开始>
// NewSpan 返回一个gin.HandlerFunc（中间件），该中间件会启动一个新的span并将其注入到请求上下文中。
//
// 它调用ctx.Next()以测量所有后续处理器的执行时间。
# <翻译结束>


<原文开始>
// ParentSpanReferenceFunc determines how to reference parent span
//
// See opentracing.SpanReferenceType
<原文结束>

# <翻译开始>
// ParentSpanReferenceFunc 定义如何引用父级 span（跟踪范围）
//
// 请参阅 opentracing.SpanReferenceType 类型定义
# <翻译结束>


<原文开始>
// SpanFromHeaders returns gin.HandlerFunc (middleware)
// that extracts parent span data from HTTP headers in TextMap format and
// starts a new span referenced to parent with ParentSpanReferenceFunc.
//
// It calls ctx.Next() to measure execution time of all following handlers.
//
// Behaviour on errors determined by abortOnErrors option.
// If it set to true request handling will be aborted with error.
<原文结束>

# <翻译开始>
// SpanFromHeaders 返回一个gin.HandlerFunc（中间件）
// 该函数从HTTP头部以TextMap格式提取父级span数据，
// 并使用ParentSpanReferenceFunc开始一个新的引用至父级span的新span。
//
// 它调用ctx.Next()来测量所有后续处理器的执行时间。
//
// 错误处理的行为由abortOnErrors选项决定。
// 如果将其设置为true，将在发生错误时中止请求处理。
# <翻译结束>


<原文开始>
// SpanFromHeadersHTTPFmt returns gin.HandlerFunc (middleware)
// that extracts parent span data from HTTP headers in HTTPHeaders format and
// starts a new span referenced to parent with ParentSpanReferenceFunc.
//
// It calls ctx.Next() to measure execution time of all following handlers.
//
// Behaviour on errors determined by abortOnErrors option.
// If it set to true request handling will be aborted with error.
<原文结束>

# <翻译开始>
// SpanFromHeadersHTTPFmt 返回一个gin.HandlerFunc（中间件）
// 该函数从HTTP头以HTTPHeaders格式提取父级span数据，并使用ParentSpanReferenceFunc
// 开启一个新的引用了父级span的新span。
//
// 它调用ctx.Next()来测量所有后续处理器的执行时间。
//
// 错误处理的行为由abortOnErrors选项决定。
// 若将其设置为true，将在发生错误时中止请求处理并返回错误。
# <翻译结束>


<原文开始>
// SpanFromContext returns gin.HandlerFunc (middleware) that extracts parent span from request context
// and starts a new span as child of parent span.
//
// It calls ctx.Next() to measure execution time of all following handlers.
//
// Behaviour on errors determined by abortOnErrors option.
// If it set to true request handling will be aborted with error.
<原文结束>

# <翻译开始>
// SpanFromContext 返回一个gin.HandlerFunc（中间件），该中间件从请求上下文中提取父级span，并以父级span为起点开始一个新的子span。
// 它调用ctx.Next()来测量所有后续处理器的执行时间。
// 错误处理的行为由abortOnErrors选项决定。如果设置为true，则在出现错误时将中止请求处理并返回错误。
# <翻译结束>


<原文开始>
// InjectToHeaders injects span meta-information to request headers.
//
// It may be useful when you want to trace chained request (client->service 1->service 2).
// In this case you have to save request headers (ctx.Request.Header) and pass it to next level request.
//
// Behaviour on errors determined by abortOnErrors option.
// If it set to true request handling will be aborted with error.
<原文结束>

# <翻译开始>
// InjectToHeaders 将追踪的span元信息注入到请求头中。
//
// 当您想要追踪链式请求（客户端->服务1->服务2）时，这可能非常有用。在这种情况下，
// 您需要保存请求头（ctx.Request.Header）并将其传递给下一级请求。
//
// 出错时的行为由abortOnErrors选项决定。
// 若该选项设置为true，处理请求将会因错误而中止。
# <翻译结束>


<原文开始>
// GetSpan extracts span from context.
<原文结束>

# <翻译开始>
// GetSpan 从上下文中提取跨度（span）。
// 在分布式追踪系统中，"span"通常代表一次操作或请求的追踪片段。这个函数是从给定的 context 中获取并返回这个追踪片段。
# <翻译结束>


<原文开始>
// MustGetSpan extracts span from context. It panics if span was not set.
<原文结束>

# <翻译开始>
// MustGetSpan 从上下文中提取跨度（span）。如果未设置跨度，则会引发恐慌。
# <翻译结束>


<原文开始>
// ff:
// opts:
// operationName:
// tracer:
<原文结束>

# <翻译开始>
// ff:
// opts:
// operationName:
// tracer:
# <翻译结束>


<原文开始>
// ff:
// abortOnErrors:
// tracer:
<原文结束>

# <翻译开始>
// ff:
// abortOnErrors:
// tracer:
# <翻译结束>


<原文开始>
// ff:
// exists:
// span:
// ctx:
<原文结束>

# <翻译开始>
// ff:
// exists:
// span:
// ctx:
# <翻译结束>

