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
// 配置追踪
	propagator := zipkin.NewZipkinB3HTTPHeaderPropagator()
	trace, closer := jaeger.NewTracer(
		"server03",
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

// 设置路由
	r := gin.Default()
	r.POST("",
		opengintracing.SpanFromHeadersHTTPFmt(trace, "service3", fn, false),
		handler)
	r.Run(":8003")
}

func handler(c *gin.Context) {
	_, found := opengintracing.GetSpan(c)
	if found == false {
		fmt.Println("Span not found")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	fmt.Println("Incoming Headers")
	for k, v := range c.Request.Header {
		fmt.Printf("%s: %s\n", k, v)
	}
	c.Status(http.StatusOK)
}
