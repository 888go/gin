
<原文开始>
// Package opengintracing provides requests tracing functional using opentracing specification.
//
// See https://github.com/opentracing/opentracing-go for more information
<原文结束>

# <翻译开始>
// opengintracing 包提供基于 opentracing 规范的请求追踪功能。
//
// 有关更多信息，请参阅 https://github.com/opentracing/opentracing-go
# <翻译结束>


<原文开始>
// Errors which may occur at operation time.
<原文结束>

# <翻译开始>
// 在操作运行时可能会发生的错误。
# <翻译结束>


<原文开始>
// NewSpan returns gin.HandlerFunc (middleware) that starts a new span and injects it to request context.
//
// It calls ctx.Next() to measure execution time of all following handlers.
<原文结束>

# <翻译开始>
// NewSpan 返回gin.HandlerFunc（中间件），该函数会启动一个新的span并将其注入到请求上下文中。
//
// 它调用ctx.Next()来测量所有后续处理器的执行时间。
# <翻译结束>


<原文开始>
// ParentSpanReferenceFunc determines how to reference parent span
//
// See opentracing.SpanReferenceType
<原文结束>

# <翻译开始>
// ParentSpanReferenceFunc 确定如何引用父级 span
//
// 参见 opentracing.SpanReferenceType
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
// SpanFromHeaders 返回 gin.HandlerFunc（中间件）
// 该函数从HTTP头部以TextMap格式提取父级span数据，
// 并使用ParentSpanReferenceFunc开始一个新的引用了父级span的新span。
//
// 它调用ctx.Next()来测量所有后续处理器的执行时间。
//
// 错误处理的行为由abortOnErrors选项决定。
// 如果将其设置为true，将会在出现错误时中止请求处理。
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
// SpanFromHeadersHTTPFmt 返回 gin.HandlerFunc（中间件）
// 该函数从 HTTP 头部以 HTTPHeaders 格式提取父级 span 数据，
// 并使用 ParentSpanReferenceFunc 开启一个新的引用了父级 span 的新 span。
//
// 它调用 ctx.Next() 来测量所有后续处理器的执行时间。
//
// 当出现错误时的行为由 abortOnErrors 选项决定。
// 如果该选项设置为 true，请求处理将因错误而终止。
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
// SpanFromContext 返回一个gin.HandlerFunc（中间件），该中间件从请求上下文中提取父级span，并作为父级span的子级开始一个新的span。
//
// 它调用ctx.Next()来测量所有后续处理器的执行时间。
//
// 错误处理的行为由abortOnErrors选项决定。如果设置为true，将在出现错误时中止请求处理。
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
// InjectToHeaders 将跨度元信息注入到请求头中。
//
// 当您想要跟踪链式请求（client->服务1->服务2）时，这可能很有用。在这种情况下，您需要保存请求头（ctx.Request.Header）并将其传递给下一级请求。
//
// 对于错误的行为由 abortOnErrors 选项决定。如果将其设置为 true，则会在出现错误时中止请求处理。
# <翻译结束>


<原文开始>
// GetSpan extracts span from context.
<原文结束>

# <翻译开始>
// GetSpan 从上下文中提取跨度（span）。
# <翻译结束>


<原文开始>
// MustGetSpan extracts span from context. It panics if span was not set.
<原文结束>

# <翻译开始>
// MustGetSpan 从上下文中提取 span（跨度）。如果未设置 span，则会触发panic。
# <翻译结束>

