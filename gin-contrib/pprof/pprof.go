package pprof

import (
	"net/http/pprof"
	
	"github.com/888go/gin"
)

const (
// DefaultPrefix 是 pprof 的默认 URL 前缀
	DefaultPrefix = "/debug/pprof"
)

func getPrefix(prefixOptions ...string) string {
	prefix := DefaultPrefix
	if len(prefixOptions) > 0 {
		prefix = prefixOptions[0]
	}
	return prefix
}

// 注册 net/http/pprof 包中标准的 HandlerFuncs 到提供的 gin.Engine 中。
// prefixOptions 是可选参数。如果不提供 prefixOptions，则使用默认路径前缀，
// 否则将使用第一个 prefixOptions 作为路径前缀。
// 这段代码注释是为一个 Go 函数写的，这个函数的功能是将 `net/http/pprof` 包中的性能分析处理器注册到 Gin 框架的路由引擎中，并且允许自定义路径前缀。
func Register(r *gin.Engine, prefixOptions ...string) {
	RouteRegister(&(r.RouterGroup), prefixOptions...)
}

// RouteRegister 将来自 net/http/pprof 包的标准 HandlerFuncs 注册到提供的 gin.GrouterGroup。
// prefixOptions 是可选的。如果不提供 prefixOptions，则使用默认路径前缀，否则将使用第一个 prefixOptions 作为路径前缀。
func RouteRegister(rg *gin.RouterGroup, prefixOptions ...string) {
	prefix := getPrefix(prefixOptions...)

	prefixRouter := rg.Group(prefix)
	{
		prefixRouter.GET("/", gin.WrapF(pprof.Index))
		prefixRouter.GET("/cmdline", gin.WrapF(pprof.Cmdline))
		prefixRouter.GET("/profile", gin.WrapF(pprof.Profile))
		prefixRouter.POST("/symbol", gin.WrapF(pprof.Symbol))
		prefixRouter.GET("/symbol", gin.WrapF(pprof.Symbol))
		prefixRouter.GET("/trace", gin.WrapF(pprof.Trace))
		prefixRouter.GET("/allocs", gin.WrapH(pprof.Handler("allocs")))
		prefixRouter.GET("/block", gin.WrapH(pprof.Handler("block")))
		prefixRouter.GET("/goroutine", gin.WrapH(pprof.Handler("goroutine")))
		prefixRouter.GET("/heap", gin.WrapH(pprof.Handler("heap")))
		prefixRouter.GET("/mutex", gin.WrapH(pprof.Handler("mutex")))
		prefixRouter.GET("/threadcreate", gin.WrapH(pprof.Handler("threadcreate")))
	}
}
