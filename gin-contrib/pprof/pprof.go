package pprof

import (
	"net/http/pprof"
	
	"github.com/888go/gin"
)

const (
	// DefaultPrefix url prefix of pprof
	DefaultPrefix = "/debug/pprof"
)

func getPrefix(prefixOptions ...string) string {
	prefix := DefaultPrefix
	if len(prefixOptions) > 0 {
		prefix = prefixOptions[0]
	}
	return prefix
}

// Register the standard HandlerFuncs from the net/http/pprof package with
// the provided gin.Engine. prefixOptions is a optional. If not prefixOptions,
// the default path prefix is used, otherwise first prefixOptions will be path prefix.
func Register(r *gin类.Engine, prefixOptions ...string) {
	RouteRegister(&(r.RouterGroup), prefixOptions...)
}

// RouteRegister the standard HandlerFuncs from the net/http/pprof package with
// the provided gin.GrouterGroup. prefixOptions is a optional. If not prefixOptions,
// the default path prefix is used, otherwise first prefixOptions will be path prefix.
func RouteRegister(rg *gin类.RouterGroup, prefixOptions ...string) {
	prefix := getPrefix(prefixOptions...)

	prefixRouter := rg.X创建分组路由(prefix)
	{
		prefixRouter.X绑定GET("/", gin类.WrapF(pprof.Index))
		prefixRouter.X绑定GET("/cmdline", gin类.WrapF(pprof.Cmdline))
		prefixRouter.X绑定GET("/profile", gin类.WrapF(pprof.Profile))
		prefixRouter.X绑定POST("/symbol", gin类.WrapF(pprof.Symbol))
		prefixRouter.X绑定GET("/symbol", gin类.WrapF(pprof.Symbol))
		prefixRouter.X绑定GET("/trace", gin类.WrapF(pprof.Trace))
		prefixRouter.X绑定GET("/allocs", gin类.WrapH(pprof.Handler("allocs")))
		prefixRouter.X绑定GET("/block", gin类.WrapH(pprof.Handler("block")))
		prefixRouter.X绑定GET("/goroutine", gin类.WrapH(pprof.Handler("goroutine")))
		prefixRouter.X绑定GET("/heap", gin类.WrapH(pprof.Handler("heap")))
		prefixRouter.X绑定GET("/mutex", gin类.WrapH(pprof.Handler("mutex")))
		prefixRouter.X绑定GET("/threadcreate", gin类.WrapH(pprof.Handler("threadcreate")))
	}
}
