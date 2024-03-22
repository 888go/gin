// opengintracing 包提供基于 opentracing 规范的请求追踪功能。
//
// 有关更多信息，请参阅 https://github.com/opentracing/opentracing-go
package opengintracing

import (
	"errors"
	"net/http"
	
	"github.com/888go/gin"
	"github.com/opentracing/opentracing-go"
)

const spanContextKey = "span"

// 在操作运行时可能会发生的错误。
var (
	ErrSpanNotFound = errors.New("span was not found in context")
)

// NewSpan 返回gin.HandlerFunc（中间件），该函数会启动一个新的span并将其注入到请求上下文中。
//
// 它调用ctx.Next()来测量所有后续处理器的执行时间。
func NewSpan(tracer opentracing.Tracer, operationName string, opts ...opentracing.StartSpanOption) gin类.HandlerFunc {
	return func(ctx *gin类.Context) {
		span := tracer.StartSpan(operationName, opts...)
		ctx.X设置值(spanContextKey, span)
		defer span.Finish()

		ctx.X中间件继续()
	}
}

// ParentSpanReferenceFunc 确定如何引用父级 span
//
// 参见 opentracing.SpanReferenceType
type ParentSpanReferenceFunc func(opentracing.SpanContext) opentracing.StartSpanOption

// SpanFromHeaders 返回 gin.HandlerFunc（中间件）
// 该函数从HTTP头部以TextMap格式提取父级span数据，
// 并使用ParentSpanReferenceFunc开始一个新的引用了父级span的新span。
//
// 它调用ctx.Next()来测量所有后续处理器的执行时间。
//
// 错误处理的行为由abortOnErrors选项决定。
// 如果将其设置为true，将会在出现错误时中止请求处理。
func SpanFromHeaders(tracer opentracing.Tracer, operationName string, psr ParentSpanReferenceFunc,
	abortOnErrors bool, advancedOpts ...opentracing.StartSpanOption,
) gin类.HandlerFunc {
	return func(ctx *gin类.Context) {
		spanContext, err := tracer.Extract(opentracing.TextMap, opentracing.HTTPHeadersCarrier(ctx.X请求.Header))
		if err != nil {
			if abortOnErrors {
				_ = ctx.X停止并带状态码与错误(http.StatusInternalServerError, err)
			}
			return
		}

		opts := append([]opentracing.StartSpanOption{psr(spanContext)}, advancedOpts...)

		span := tracer.StartSpan(operationName, opts...)
		ctx.X设置值(spanContextKey, span)
		defer span.Finish()

		ctx.X中间件继续()
	}
}

// SpanFromHeadersHTTPFmt 返回 gin.HandlerFunc（中间件）
// 该函数从 HTTP 头部以 HTTPHeaders 格式提取父级 span 数据，
// 并使用 ParentSpanReferenceFunc 开启一个新的引用了父级 span 的新 span。
//
// 它调用 ctx.Next() 来测量所有后续处理器的执行时间。
//
// 当出现错误时的行为由 abortOnErrors 选项决定。
// 如果该选项设置为 true，请求处理将因错误而终止。
func SpanFromHeadersHTTPFmt(tracer opentracing.Tracer, operationName string, psr ParentSpanReferenceFunc,
	abortOnErrors bool, advancedOpts ...opentracing.StartSpanOption,
) gin类.HandlerFunc {
	return func(ctx *gin类.Context) {
		spanContext, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(ctx.X请求.Header))
		if err != nil {
			if abortOnErrors {
				_ = ctx.X停止并带状态码与错误(http.StatusInternalServerError, err)
			}
			return
		}

		opts := append([]opentracing.StartSpanOption{psr(spanContext)}, advancedOpts...)

		span := tracer.StartSpan(operationName, opts...)
		ctx.X设置值(spanContextKey, span)
		defer span.Finish()

		ctx.X中间件继续()
	}
}

// SpanFromContext 返回一个gin.HandlerFunc（中间件），该中间件从请求上下文中提取父级span，并作为父级span的子级开始一个新的span。
//
// 它调用ctx.Next()来测量所有后续处理器的执行时间。
//
// 错误处理的行为由abortOnErrors选项决定。如果设置为true，将在出现错误时中止请求处理。
func SpanFromContext(tracer opentracing.Tracer, operationName string, abortOnErrors bool,
	advancedOpts ...opentracing.StartSpanOption,
) gin类.HandlerFunc {
	return func(ctx *gin类.Context) {
		var opts []opentracing.StartSpanOption
		parentSpanI, _ := ctx.X取值(spanContextKey)
		if parentSpan, typeOk := parentSpanI.(opentracing.Span); parentSpan != nil && typeOk {
			opts = append(opts, opentracing.ChildOf(parentSpan.Context()))
		} else {
			if abortOnErrors {
				_ = ctx.X停止并带状态码与错误(http.StatusInternalServerError, ErrSpanNotFound)
			}
			return
		}
		opts = append(opts, advancedOpts...)

		span := tracer.StartSpan(operationName, opts...)
		ctx.X设置值(spanContextKey, span)
		defer span.Finish()

		ctx.X中间件继续()
	}
}

// InjectToHeaders 将跨度元信息注入到请求头中。
//
// 当您想要跟踪链式请求（client->服务1->服务2）时，这可能很有用。在这种情况下，您需要保存请求头（ctx.Request.Header）并将其传递给下一级请求。
//
// 对于错误的行为由 abortOnErrors 选项决定。如果将其设置为 true，则会在出现错误时中止请求处理。
func InjectToHeaders(tracer opentracing.Tracer, abortOnErrors bool) gin类.HandlerFunc {
	return func(ctx *gin类.Context) {
		var spanContext opentracing.SpanContext
		spanI, _ := ctx.X取值(spanContextKey)
		if span, typeOk := spanI.(opentracing.Span); span != nil && typeOk {
			spanContext = span.Context()
		} else {
			if abortOnErrors {
				ctx.X停止并带状态码且返回JSON(http.StatusInternalServerError, ErrSpanNotFound)
			}
			return
		}

		_ = tracer.Inject(spanContext, opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(ctx.X请求.Header))
	}
}

// GetSpan 从上下文中提取跨度（span）。
func GetSpan(ctx *gin类.Context) (span opentracing.Span, exists bool) {
	spanI, _ := ctx.X取值(spanContextKey)
	span, ok := spanI.(opentracing.Span)
	exists = span != nil && ok
	return
}

// MustGetSpan 从上下文中提取 span（跨度）。如果未设置 span，则会触发panic。
func MustGetSpan(ctx *gin类.Context) opentracing.Span {
	return ctx.X取值PANI(spanContextKey).(opentracing.Span)
}
