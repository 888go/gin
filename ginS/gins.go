// 版权所有 2014 Manu Martinez-Almeida。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package ginS

import (
	"html/template"
	"net/http"
	"sync"
	
	"github.com/888go/gin"
)

var once sync.Once
var internalEngine *gin类.Engine

func engine() *gin类.Engine {
	once.Do(func() {
		internalEngine = gin类.X创建默认对象()
	})
	return internalEngine
}

// LoadHTMLGlob 是 Engine.LoadHTMLGlob 的一个包装函数。
func X加载HTML模板目录(模板目录 string) {
	engine().X加载HTML模板目录(模板目录)
}

// LoadHTMLFiles 是对 Engine.LoadHTMLFiles 的一个封装。
func X加载HTML模板文件(模板文件s ...string) {
	engine().X加载HTML模板文件(模板文件s...)
}

// SetHTMLTemplate 是 Engine.SetHTMLTemplate 的一个包装函数。
func X设置Template模板(Template模板 *template.Template) {
	engine().X设置Template模板(Template模板)
}

// NoRoute 添加处理函数，用于未找到路由的情况（NoRoute）。默认情况下返回404状态码。
func X绑定404(处理函数s ...gin类.HandlerFunc) {
	engine().X绑定404(处理函数s...)
}

// NoMethod 是 Engine.NoMethod 的一个包装器。
func X绑定405(处理函数s ...gin类.HandlerFunc) {
	engine().X绑定405(处理函数s...)
}

// Group 创建一个新的路由分组。你应该在此添加所有具有共同中间件或相同路径前缀的路由。
// 例如，所有使用共同授权中间件的路由可以被归为一组。
func X创建分组路由(路由规则 string, 处理函数s ...gin类.HandlerFunc) *gin类.RouterGroup {
	return engine().X创建分组路由(路由规则, 处理函数s...)
}

// Handle 是 Engine.Handle 的一个包装器。
func X绑定(HTTP方法, 路由规则 string, 处理函数s ...gin类.HandlerFunc) gin类.IRoutes {
	return engine().X绑定(HTTP方法, 路由规则, 处理函数s...)
}

// POST 是一个快捷方式，用于 router.Handle("POST", path, handle)
func X绑定POST(路由规则 string, 处理函数s ...gin类.HandlerFunc) gin类.IRoutes {
	return engine().X绑定POST(路由规则, 处理函数s...)
}

// GET 是一个快捷方式，等同于 router.Handle("GET", path, handle)
func X绑定GET(路由规则 string, 处理函数s ...gin类.HandlerFunc) gin类.IRoutes {
	return engine().X绑定GET(路由规则, 处理函数s...)
}

// DELETE 是一个快捷方式，等同于 router.Handle("DELETE", path, handle)
func X绑定DELETE(路由规则 string, 处理函数s ...gin类.HandlerFunc) gin类.IRoutes {
	return engine().X绑定DELETE(路由规则, 处理函数s...)
}

// PATCH 是一个快捷方式，用于 router.Handle("PATCH", path, handle)
func X绑定PATCH(路由规则 string, 处理函数s ...gin类.HandlerFunc) gin类.IRoutes {
	return engine().X绑定PATCH(路由规则, 处理函数s...)
}

// PUT 是一个快捷方式，等同于 router.Handle("PUT", path, handle)
func X绑定PUT(路由规则 string, 处理函数s ...gin类.HandlerFunc) gin类.IRoutes {
	return engine().X绑定PUT(路由规则, 处理函数s...)
}

// OPTIONS 是一个快捷方式，用于 router.Handle("OPTIONS", path, handle)
func X绑定OPTIONS(路由规则 string, 处理函数s ...gin类.HandlerFunc) gin类.IRoutes {
	return engine().X绑定OPTIONS(路由规则, 处理函数s...)
}

// HEAD 是一个快捷方式，用于 router.Handle("HEAD", path, handle)
func X绑定HEAD(路由规则 string, 处理函数s ...gin类.HandlerFunc) gin类.IRoutes {
	return engine().X绑定HEAD(路由规则, 处理函数s...)
}

// Any 是 Engine.Any 的一个包装器。
func X绑定Any(路由规则 string, 处理函数s ...gin类.HandlerFunc) gin类.IRoutes {
	return engine().X绑定Any(路由规则, 处理函数s...)
}

// StaticFile 是 Engine.StaticFile 的一个包装器。
func X绑定静态单文件(URL路径, 文件路径 string) gin类.IRoutes {
	return engine().X绑定静态单文件(URL路径, 文件路径)
}

// Static 从给定的文件系统根目录提供文件服务。
// 如:
// r.Static("/static", "./文件夹") //当你访问http://localhost:8080/static时，它会服务于./文件夹 目录下的文件
//
// 在内部使用了 http.FileServer，因此会使用 http.NotFound 替代 Router 的 NotFound 处理程序。
// 要使用操作系统自身的文件系统实现，请按如下方式使用：
// router.Static("/static", "/var/www")
func X绑定静态文件目录(URL路径前缀, 绑定目录 string) gin类.IRoutes {
	return engine().X绑定静态文件目录(URL路径前缀, 绑定目录)
}

// StaticFS 是 Engine.StaticFS 的一个包装器。
func X绑定静态文件目录FS(URL路径前缀 string, fs http.FileSystem) gin类.IRoutes {
	return engine().X绑定静态文件目录FS(URL路径前缀, fs)
}

// Use 方法将全局中间件附加到路由器。即通过 Use() 附加的中间件将会
// 包含在每一个请求的处理器链中。即使是 404、405 状态码的响应，或者是静态文件的处理...
// 例如，这里适合放置日志记录器或错误管理中间件。
func X中间件(middlewares ...gin类.HandlerFunc) gin类.IRoutes {
	return engine().X中间件(middlewares...)
}

// Routes 返回已注册路由的切片。
func X取路由数组() gin类.RoutesInfo {
	return engine().X取路由数组()
}

// Run 函数连接到一个 http.Server，并开始监听和处理 HTTP 请求。
// 这是 http.ListenAndServe(addr, router) 的快捷方式。
// 注意：除非发生错误，否则此方法将无限期地阻塞调用它的 goroutine。
func X监听(地址与端口 ...string) (错误 error) {
	return engine().X监听(地址与端口...)
}

// RunTLS 绑定到一个 http.Server，并开始监听和处理 HTTPS 请求。
// 这是 http.ListenAndServeTLS(addr, certFile, keyFile, router) 的快捷方式。
// 注意：除非发生错误，否则此方法将无限期地阻塞调用它的 goroutine。
func X监听TLS(地址与端口, cert文件, key文件 string) (错误 error) {
	return engine().X监听TLS(地址与端口, cert文件, key文件)
}

// RunUnix连接到一个http.Server，并开始通过指定的unix套接字（即文件）监听和处理HTTP请求。
// 注意：除非发生错误，否则此方法将无限期地阻塞调用的goroutine。
func X监听Unix(文件路径 string) (错误 error) {
	return engine().X监听Unix(文件路径)
}

// RunFd 将路由器附加到 http.Server，并开始通过指定的文件描述符监听和处理 HTTP 请求。
// 注意：除非发生错误，否则该方法将无限期地阻塞调用它的 goroutine。
func X监听Fd(fd int) (错误 error) {
	return engine().X监听Fd(fd)
}
