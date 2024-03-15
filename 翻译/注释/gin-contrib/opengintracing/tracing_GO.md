
<原文开始>
// Package opengintracing provides requests tracing functional using opentracing specification.
//
// See https://github.com/opentracing/opentracing-go for more information
<原文结束>

# <翻译开始>
// Package opengintracing provides requests tracing functional using opentracing specification.
//
// See https://github.com/opentracing/opentracing-go for more information
# <翻译结束>


<原文开始>
// Errors which may occur at operation time.
<原文结束>

# <翻译开始>
// Errors which may occur at operation time.
# <翻译结束>


<原文开始>
// NewSpan returns gin.HandlerFunc (middleware) that starts a new span and injects it to request context.
//
// It calls ctx.Next() to measure execution time of all following handlers.
<原文结束>

# <翻译开始>
// NewSpan returns gin.HandlerFunc (middleware) that starts a new span and injects it to request context.
//
// It calls ctx.Next() to measure execution time of all following handlers.
# <翻译结束>


<原文开始>
// ParentSpanReferenceFunc determines how to reference parent span
//
// See opentracing.SpanReferenceType
<原文结束>

# <翻译开始>
// ParentSpanReferenceFunc determines how to reference parent span
//
// See opentracing.SpanReferenceType
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
// SpanFromHeaders returns gin.HandlerFunc (middleware)
// that extracts parent span data from HTTP headers in TextMap format and
// starts a new span referenced to parent with ParentSpanReferenceFunc.
//
// It calls ctx.Next() to measure execution time of all following handlers.
//
// Behaviour on errors determined by abortOnErrors option.
// If it set to true request handling will be aborted with error.
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
// SpanFromHeadersHTTPFmt returns gin.HandlerFunc (middleware)
// that extracts parent span data from HTTP headers in HTTPHeaders format and
// starts a new span referenced to parent with ParentSpanReferenceFunc.
//
// It calls ctx.Next() to measure execution time of all following handlers.
//
// Behaviour on errors determined by abortOnErrors option.
// If it set to true request handling will be aborted with error.
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
// SpanFromContext returns gin.HandlerFunc (middleware) that extracts parent span from request context
// and starts a new span as child of parent span.
//
// It calls ctx.Next() to measure execution time of all following handlers.
//
// Behaviour on errors determined by abortOnErrors option.
// If it set to true request handling will be aborted with error.
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
// InjectToHeaders injects span meta-information to request headers.
//
// It may be useful when you want to trace chained request (client->service 1->service 2).
// In this case you have to save request headers (ctx.Request.Header) and pass it to next level request.
//
// Behaviour on errors determined by abortOnErrors option.
// If it set to true request handling will be aborted with error.
# <翻译结束>


<原文开始>
// GetSpan extracts span from context.
<原文结束>

# <翻译开始>
// GetSpan extracts span from context.
# <翻译结束>


<原文开始>
// MustGetSpan extracts span from context. It panics if span was not set.
<原文结束>

# <翻译开始>
// MustGetSpan extracts span from context. It panics if span was not set.
# <翻译结束>

