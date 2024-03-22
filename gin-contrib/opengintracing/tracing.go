// Package opengintracing provides requests tracing functional using opentracing specification.
//
// See https://github.com/opentracing/opentracing-go for more information
package opengintracing

import (
	"errors"
	"net/http"
	
	"github.com/888go/gin"
	"github.com/opentracing/opentracing-go"
)

const spanContextKey = "span"

// Errors which may occur at operation time.
var (
	ErrSpanNotFound = errors.New("span was not found in context")
)

// NewSpan returns gin.HandlerFunc (middleware) that starts a new span and injects it to request context.
//
// It calls ctx.Next() to measure execution time of all following handlers.
func NewSpan(tracer opentracing.Tracer, operationName string, opts ...opentracing.StartSpanOption) gin类.HandlerFunc {
	return func(ctx *gin类.Context) {
		span := tracer.StartSpan(operationName, opts...)
		ctx.X设置值(spanContextKey, span)
		defer span.Finish()

		ctx.X中间件继续()
	}
}

// ParentSpanReferenceFunc determines how to reference parent span
//
// See opentracing.SpanReferenceType
type ParentSpanReferenceFunc func(opentracing.SpanContext) opentracing.StartSpanOption

// SpanFromHeaders returns gin.HandlerFunc (middleware)
// that extracts parent span data from HTTP headers in TextMap format and
// starts a new span referenced to parent with ParentSpanReferenceFunc.
//
// It calls ctx.Next() to measure execution time of all following handlers.
//
// Behaviour on errors determined by abortOnErrors option.
// If it set to true request handling will be aborted with error.
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

// SpanFromHeadersHTTPFmt returns gin.HandlerFunc (middleware)
// that extracts parent span data from HTTP headers in HTTPHeaders format and
// starts a new span referenced to parent with ParentSpanReferenceFunc.
//
// It calls ctx.Next() to measure execution time of all following handlers.
//
// Behaviour on errors determined by abortOnErrors option.
// If it set to true request handling will be aborted with error.
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

// SpanFromContext returns gin.HandlerFunc (middleware) that extracts parent span from request context
// and starts a new span as child of parent span.
//
// It calls ctx.Next() to measure execution time of all following handlers.
//
// Behaviour on errors determined by abortOnErrors option.
// If it set to true request handling will be aborted with error.
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

// InjectToHeaders injects span meta-information to request headers.
//
// It may be useful when you want to trace chained request (client->service 1->service 2).
// In this case you have to save request headers (ctx.Request.Header) and pass it to next level request.
//
// Behaviour on errors determined by abortOnErrors option.
// If it set to true request handling will be aborted with error.
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

// GetSpan extracts span from context.
func GetSpan(ctx *gin类.Context) (span opentracing.Span, exists bool) {
	spanI, _ := ctx.X取值(spanContextKey)
	span, ok := spanI.(opentracing.Span)
	exists = span != nil && ok
	return
}

// MustGetSpan extracts span from context. It panics if span was not set.
func MustGetSpan(ctx *gin类.Context) opentracing.Span {
	return ctx.X取值PANI(spanContextKey).(opentracing.Span)
}
