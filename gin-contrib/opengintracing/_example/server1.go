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
		"server01",
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
		opengintracing.SpanFromHeadersHTTPFmt(trace, "service1", fn, false),
		handler)
	r.X监听(":8001")
}

func printHeaders(message string, header http.Header) {
	fmt.Println(message)
	for k, v := range header {
		fmt.Printf("%s: %s\n", k, v)
	}
}

func handler(c *gin类.Context) {
	span, found := opengintracing.GetSpan(c)
	if found == false {
		fmt.Println("Span not found")
		c.X停止并带状态码(http.StatusInternalServerError)
		return
	}
	req, _ := http.NewRequest("POST", "http://localhost:8003", nil)

	opentracing.GlobalTracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header))

	printHeaders("Incoming headers", c.X请求.Header)
	printHeaders("Outgoing headers", req.Header)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.X停止并带状态码与错误(http.StatusInternalServerError, err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Unexpected response from service3")
	}
	c.X设置状态码(http.StatusOK)
}
