// Package opengintracing 提供了基于 opentracing 规范的请求跟踪功能。
//
// 有关更多信息，请参见 https://github.com/opentracing/opentracing-go
package opengintracing

import (
	"errors"
	"net/http"
	
	"github.com/888go/gin"
	"github.com/opentracing/opentracing-go"
)

const spanContextKey = "span"

// 运行时可能会发生的错误。
var (
	ErrSpanNotFound = errors.New("span was not found in context")
)

// NewSpan 返回一个gin.HandlerFunc（中间件），该中间件会启动一个新的span并将其注入到请求上下文中。
//
// 它调用ctx.Next()以测量所有后续处理器的执行时间。

// ff:
// opts:
// operationName:
// tracer:

// ff:
// opts:
// operationName:
// tracer:

// ff:
// opts:
// operationName:
// tracer:

// ff:
// opts:
// operationName:
// tracer:

// ff:
// opts:
// operationName:
// tracer:
func NewSpan(tracer opentracing.Tracer, operationName string, opts ...opentracing.StartSpanOption) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		span := tracer.StartSpan(operationName, opts...)
		ctx.Set(spanContextKey, span)
		defer span.Finish()

		ctx.Next()
	}
}

// ParentSpanReferenceFunc 定义如何引用父级 span（跟踪范围）
//
// 请参阅 opentracing.SpanReferenceType 类型定义
type ParentSpanReferenceFunc func(opentracing.SpanContext) opentracing.StartSpanOption

// SpanFromHeaders 返回一个gin.HandlerFunc（中间件）
// 该函数从HTTP头部以TextMap格式提取父级span数据，
// 并使用ParentSpanReferenceFunc开始一个新的引用至父级span的新span。
//
// 它调用ctx.Next()来测量所有后续处理器的执行时间。
//
// 错误处理的行为由abortOnErrors选项决定。
// 如果将其设置为true，将在发生错误时中止请求处理。
func SpanFromHeaders(tracer opentracing.Tracer, operationName string, psr ParentSpanReferenceFunc,
	abortOnErrors bool, advancedOpts ...opentracing.StartSpanOption,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		spanContext, err := tracer.Extract(opentracing.TextMap, opentracing.HTTPHeadersCarrier(ctx.Request.Header))
		if err != nil {
			if abortOnErrors {
				_ = ctx.AbortWithError(http.StatusInternalServerError, err)
			}
			return
		}

		opts := append([]opentracing.StartSpanOption{psr(spanContext)}, advancedOpts...)

		span := tracer.StartSpan(operationName, opts...)
		ctx.Set(spanContextKey, span)
		defer span.Finish()

		ctx.Next()
	}
}

// SpanFromHeadersHTTPFmt 返回一个gin.HandlerFunc（中间件）
// 该函数从HTTP头以HTTPHeaders格式提取父级span数据，并使用ParentSpanReferenceFunc
// 开启一个新的引用了父级span的新span。
//
// 它调用ctx.Next()来测量所有后续处理器的执行时间。
//
// 错误处理的行为由abortOnErrors选项决定。
// 若将其设置为true，将在发生错误时中止请求处理并返回错误。
func SpanFromHeadersHTTPFmt(tracer opentracing.Tracer, operationName string, psr ParentSpanReferenceFunc,
	abortOnErrors bool, advancedOpts ...opentracing.StartSpanOption,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		spanContext, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(ctx.Request.Header))
		if err != nil {
			if abortOnErrors {
				_ = ctx.AbortWithError(http.StatusInternalServerError, err)
			}
			return
		}

		opts := append([]opentracing.StartSpanOption{psr(spanContext)}, advancedOpts...)

		span := tracer.StartSpan(operationName, opts...)
		ctx.Set(spanContextKey, span)
		defer span.Finish()

		ctx.Next()
	}
}

// SpanFromContext 返回一个gin.HandlerFunc（中间件），该中间件从请求上下文中提取父级span，并以父级span为起点开始一个新的子span。
// 它调用ctx.Next()来测量所有后续处理器的执行时间。
// 错误处理的行为由abortOnErrors选项决定。如果设置为true，则在出现错误时将中止请求处理并返回错误。
func SpanFromContext(tracer opentracing.Tracer, operationName string, abortOnErrors bool,
	advancedOpts ...opentracing.StartSpanOption,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var opts []opentracing.StartSpanOption
		parentSpanI, _ := ctx.Get(spanContextKey)
		if parentSpan, typeOk := parentSpanI.(opentracing.Span); parentSpan != nil && typeOk {
			opts = append(opts, opentracing.ChildOf(parentSpan.Context()))
		} else {
			if abortOnErrors {
				_ = ctx.AbortWithError(http.StatusInternalServerError, ErrSpanNotFound)
			}
			return
		}
		opts = append(opts, advancedOpts...)

		span := tracer.StartSpan(operationName, opts...)
		ctx.Set(spanContextKey, span)
		defer span.Finish()

		ctx.Next()
	}
}

// InjectToHeaders 将追踪的span元信息注入到请求头中。
//
// 当您想要追踪链式请求（客户端->服务1->服务2）时，这可能非常有用。在这种情况下，
// 您需要保存请求头（ctx.Request.Header）并将其传递给下一级请求。
//
// 出错时的行为由abortOnErrors选项决定。
// 若该选项设置为true，处理请求将会因错误而中止。

// ff:
// abortOnErrors:
// tracer:

// ff:
// abortOnErrors:
// tracer:

// ff:
// abortOnErrors:
// tracer:

// ff:
// abortOnErrors:
// tracer:

// ff:
// abortOnErrors:
// tracer:
func InjectToHeaders(tracer opentracing.Tracer, abortOnErrors bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var spanContext opentracing.SpanContext
		spanI, _ := ctx.Get(spanContextKey)
		if span, typeOk := spanI.(opentracing.Span); span != nil && typeOk {
			spanContext = span.Context()
		} else {
			if abortOnErrors {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, ErrSpanNotFound)
			}
			return
		}

		_ = tracer.Inject(spanContext, opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(ctx.Request.Header))
	}
}

// GetSpan 从上下文中提取跨度（span）。
// 在分布式追踪系统中，"span"通常代表一次操作或请求的追踪片段。这个函数是从给定的 context 中获取并返回这个追踪片段。

// ff:
// exists:
// span:
// ctx:

// ff:
// exists:
// span:
// ctx:

// ff:
// exists:
// span:
// ctx:

// ff:
// exists:
// span:
// ctx:

// ff:
// exists:
// span:
// ctx:
func GetSpan(ctx *gin.Context) (span opentracing.Span, exists bool) {
	spanI, _ := ctx.Get(spanContextKey)
	span, ok := spanI.(opentracing.Span)
	exists = span != nil && ok
	return
}

// MustGetSpan 从上下文中提取跨度（span）。如果未设置跨度，则会引发恐慌。

// ff:
// ctx:

// ff:
// ctx:

// ff:
// ctx:

// ff:
// ctx:

// ff:
// ctx:
func MustGetSpan(ctx *gin.Context) opentracing.Span {
	return ctx.MustGet(spanContextKey).(opentracing.Span)
}
