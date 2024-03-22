package pprof

import (
	"net/http/pprof"
	
	"github.com/888go/gin"
)

const (
	// DefaultPrefix 是pprof的默认URL前缀
	DefaultPrefix = "/debug/pprof"
)

func getPrefix(prefixOptions ...string) string {
	prefix := DefaultPrefix
	if len(prefixOptions) > 0 {
		prefix = prefixOptions[0]
	}
	return prefix
}

// 使用提供的gin.Engine注册net/http/pprof包中的标准HandlerFuncs。
// prefixOptions是可选的。如果不提供prefixOptions，则使用默认路径前缀，否则将使用第一个prefixOptions作为路径前缀。
func Register(r *gin类.Engine, prefixOptions ...string) {
	RouteRegister(&(r.RouterGroup), prefixOptions...)
}

// RouteRegister 将标准的来自 net/http/pprof 包中的 HandlerFuncs 与提供的 gin.GrouterGroup 进行注册。
// prefixOptions 是可选参数。如果不提供 prefixOptions，将使用默认路径前缀；否则，首个 prefixOptions 将作为路径前缀。
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
