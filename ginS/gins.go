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
var internalEngine *gin.Engine

func engine() *gin.Engine {
	once.Do(func() {
		internalEngine = gin.Default()
	})
	return internalEngine
}

// LoadHTMLGlob 是 Engine.LoadHTMLGlob 的一个包装函数。

// ff:
// pattern:
func LoadHTMLGlob(pattern string) {
	engine().LoadHTMLGlob(pattern)
}

// LoadHTMLFiles 是对 Engine.LoadHTMLFiles 的一个封装。

// ff:
// files:
func LoadHTMLFiles(files ...string) {
	engine().LoadHTMLFiles(files...)
}

// SetHTMLTemplate 是 Engine.SetHTMLTemplate 的一个包装函数。

// ff:
// templ:
func SetHTMLTemplate(templ *template.Template) {
	engine().SetHTMLTemplate(templ)
}

// NoRoute 添加处理函数，用于未找到路由的情况（NoRoute）。默认情况下返回404状态码。

// ff:
// handlers:
func NoRoute(handlers ...gin.HandlerFunc) {
	engine().NoRoute(handlers...)
}

// NoMethod 是 Engine.NoMethod 的一个包装器。

// ff:
// handlers:
func NoMethod(handlers ...gin.HandlerFunc) {
	engine().NoMethod(handlers...)
}

// Group 创建一个新的路由分组。你应该在此添加所有具有共同中间件或相同路径前缀的路由。
// 例如，所有使用共同授权中间件的路由可以被归为一组。

// ff:
// handlers:
// relativePath:
func Group(relativePath string, handlers ...gin.HandlerFunc) *gin.RouterGroup {
	return engine().Group(relativePath, handlers...)
}

// Handle 是 Engine.Handle 的一个包装器。

// ff:绑定
// handlers:处理函数
// relativePath:路由规则
// httpMethod:HTTP方法
func Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return engine().Handle(httpMethod, relativePath, handlers...)
}

// POST 是一个快捷方式，用于 router.Handle("POST", path, handle)

// ff:绑定POST
// handlers:处理函数
// relativePath:路由规则
func POST(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return engine().POST(relativePath, handlers...)
}

// GET 是一个快捷方式，等同于 router.Handle("GET", path, handle)

// ff:绑定GET
// handlers:处理函数
// relativePath:路由规则
func GET(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return engine().GET(relativePath, handlers...)
}

// DELETE 是一个快捷方式，等同于 router.Handle("DELETE", path, handle)

// ff:绑定DELETE
// handlers:处理函数
// relativePath:路由规则
func DELETE(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return engine().DELETE(relativePath, handlers...)
}

// PATCH 是一个快捷方式，用于 router.Handle("PATCH", path, handle)

// ff:绑定PATCH
// handlers:处理函数
// relativePath:路由规则
func PATCH(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return engine().PATCH(relativePath, handlers...)
}

// PUT 是一个快捷方式，等同于 router.Handle("PUT", path, handle)

// ff:绑定PUT
// handlers:处理函数
// relativePath:路由规则
func PUT(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return engine().PUT(relativePath, handlers...)
}

// OPTIONS 是一个快捷方式，用于 router.Handle("OPTIONS", path, handle)

// ff:绑定OPTIONS
// handlers:处理函数
// relativePath:路由规则
func OPTIONS(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return engine().OPTIONS(relativePath, handlers...)
}

// HEAD 是一个快捷方式，用于 router.Handle("HEAD", path, handle)

// ff:绑定HEAD
// handlers:处理函数
// relativePath:路由规则
func HEAD(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return engine().HEAD(relativePath, handlers...)
}

// Any 是 Engine.Any 的一个包装器。

// ff:绑定Any
// handlers:处理函数
// relativePath:路由规则
func Any(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return engine().Any(relativePath, handlers...)
}

// StaticFile 是 Engine.StaticFile 的一个包装器。

// ff:绑定静态单文件
// filepath:文件路径
// relativePath:URL路径
func StaticFile(relativePath, filepath string) gin.IRoutes {
	return engine().StaticFile(relativePath, filepath)
}

// Static 从给定的文件系统根目录提供文件服务。
// 在内部使用了 http.FileServer，因此会使用 http.NotFound 替代 Router 的 NotFound 处理程序。
// 要使用操作系统自身的文件系统实现，请按如下方式使用：
//
//	router.Static("/static", "/var/www")

// ff:绑定静态文件目录
// root:绑定目录
// relativePath:URL路径前缀
func Static(relativePath, root string) gin.IRoutes {
	return engine().Static(relativePath, root)
}

// StaticFS 是 Engine.StaticFS 的一个包装器。

// ff:绑定静态文件目录FS
// fs:
// relativePath:URL路径前缀
func StaticFS(relativePath string, fs http.FileSystem) gin.IRoutes {
	return engine().StaticFS(relativePath, fs)
}

// Use 方法将全局中间件附加到路由器。即通过 Use() 附加的中间件将会
// 包含在每一个请求的处理器链中。即使是 404、405 状态码的响应，或者是静态文件的处理...
// 例如，这里适合放置日志记录器或错误管理中间件。

// ff:中间件
// middlewares:
func Use(middlewares ...gin.HandlerFunc) gin.IRoutes {
	return engine().Use(middlewares...)
}

// Routes 返回已注册路由的切片。

// ff:
func Routes() gin.RoutesInfo {
	return engine().Routes()
}

// Run 函数连接到一个 http.Server，并开始监听和处理 HTTP 请求。
// 这是 http.ListenAndServe(addr, router) 的快捷方式。
// 注意：除非发生错误，否则此方法将无限期地阻塞调用它的 goroutine。

// ff:
// err:
// addr:
func Run(addr ...string) (err error) {
	return engine().Run(addr...)
}

// RunTLS 绑定到一个 http.Server，并开始监听和处理 HTTPS 请求。
// 这是 http.ListenAndServeTLS(addr, certFile, keyFile, router) 的快捷方式。
// 注意：除非发生错误，否则此方法将无限期地阻塞调用它的 goroutine。

// ff:
// err:
// keyFile:
// certFile:
// addr:
func RunTLS(addr, certFile, keyFile string) (err error) {
	return engine().RunTLS(addr, certFile, keyFile)
}

// RunUnix连接到一个http.Server，并开始通过指定的unix套接字（即文件）监听和处理HTTP请求。
// 注意：除非发生错误，否则此方法将无限期地阻塞调用的goroutine。

// ff:
// err:
// file:
func RunUnix(file string) (err error) {
	return engine().RunUnix(file)
}

// RunFd 将路由器附加到 http.Server，并开始通过指定的文件描述符监听和处理 HTTP 请求。
// 注意：除非发生错误，否则该方法将无限期地阻塞调用它的 goroutine。

// ff:
// err:
// fd:
func RunFd(fd int) (err error) {
	return engine().RunFd(fd)
}
