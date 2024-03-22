package main

import (
	"fmt"
	"net/http"
	
	"github.com/888go/gin/gin-contrib/opengintracing"
	"github.com/888go/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/zipkin"
)

func main() {
	// Configure tracing
	propagator := zipkin.NewZipkinB3HTTPHeaderPropagator()
	trace, closer := jaeger.NewTracer(
		"server02",
		jaeger.NewConstSampler(true),
		jaeger.NewNullReporter(),
		jaeger.TracerOptions.Injector(opentracing.HTTPHeaders, propagator),
		jaeger.TracerOptions.Extractor(opentracing.HTTPHeaders, propagator),
		jaeger.TracerOptions.ZipkinSharedRPCSpan(true),
	)
	defer closer.Close()
	opentracing.SetGlobalTracer(trace)
	var fn opengintracing.ParentSpanReferenceFunc
	fn = func(sc opentracing.SpanContext) opentracing.StartSpanOption {
		return opentracing.ChildOf(sc)
	}

	// Set up routes
	r := gin类.X创建默认对象()
	r.X绑定POST("",
		opengintracing.SpanFromHeadersHTTPFmt(trace, "service2", fn, false),
		handler)
	r.X监听(":8002")
}

func handler(c *gin类.Context) {
	_, found := opengintracing.GetSpan(c)
	if found == false {
		fmt.Println("Span not found")
		c.X停止并带状态码(http.StatusInternalServerError)
		return
	}
	fmt.Println("Incoming Headers")
	for k, v := range c.X请求.Header {
		fmt.Printf("%s: %s\n", k, v)
	}
	c.X设置状态码(http.StatusOK)
}
