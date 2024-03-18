// 版权声明 2014 Manu Martinez-Almeida。保留所有权利。
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

// LoadHTMLGlob 是 Engine.LoadHTMLGlob 的一个包装器。

// ff:
// pattern:
func LoadHTMLGlob(pattern string) {
	engine().LoadHTMLGlob(pattern)
}

// LoadHTMLFiles 是 Engine.LoadHTMLFiles 的一个封装函数。

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

// NoRoute 为无路由情况添加处理器。默认返回404状态码。

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

// Group 创建一个新的路由分组。你应该将所有具有共同中间件或相同路径前缀的路由添加到此分组中。
// 例如，所有使用同一个授权中间件的路由可以被归为一组。

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
// 即使用特定方法（此处为"POST"）和路径（path）注册一个处理器(handle)到路由(router)

// ff:绑定POST
// handlers:处理函数
// relativePath:路由规则
func POST(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return engine().POST(relativePath, handlers...)
}

// GET 是一个快捷方式，用于 router.Handle("GET", path, handle)
// 即通过此快捷方式可以快速处理 GET 方法类型的请求，对应路径为 path，并调用 handle 处理函数

// ff:绑定GET
// handlers:处理函数
// relativePath:路由规则
func GET(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return engine().GET(relativePath, handlers...)
}

// DELETE 是一个快捷方式，用于router.Handle("DELETE", path, handle)
// 即：通过此快捷方式可以方便地为指定路由路径注册一个处理DELETE请求的方法。

// ff:绑定DELETE
// handlers:处理函数
// relativePath:路由规则
func DELETE(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return engine().DELETE(relativePath, handlers...)
}

// PATCH 是一个快捷方式，用于 router.Handle("PATCH", path, handle)
// 即使用 PATCH 方法注册路由处理器，其中：
// "PATCH" 代表 HTTP 请求方法，
// path 为待处理的请求路径，
// handle 为对应的处理函数。

// ff:绑定PATCH
// handlers:处理函数
// relativePath:路由规则
func PATCH(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return engine().PATCH(relativePath, handlers...)
}

// PUT 是一个快捷方式，用于 router.Handle("PUT", path, handle)
// 即：使用 router 处理 "PUT" 方法的请求，路径为 path，并调用 handle 函数进行处理

// ff:绑定PUT
// handlers:处理函数
// relativePath:路由规则
func PUT(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return engine().PUT(relativePath, handlers...)
}

// OPTIONS 是一个快捷方式，用于 router.Handle("OPTIONS", path, handle)
// 即：通过该选项，可以快速处理对指定路径（path）的 "OPTIONS" HTTP 方法请求，并调用相应的处理函数（handle）

// ff:绑定OPTIONS
// handlers:处理函数
// relativePath:路由规则
func OPTIONS(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return engine().OPTIONS(relativePath, handlers...)
}

// HEAD 是一个快捷方式，用于 router.Handle("HEAD", path, handle)
// 即为：router 处理 "HEAD" 方法类型的请求，路径为 path，处理函数为 handle

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

// ff:
// filepath:
// relativePath:
func StaticFile(relativePath, filepath string) gin.IRoutes {
	return engine().StaticFile(relativePath, filepath)
}

// Static 从给定的文件系统根目录提供文件服务。
// 在内部使用了 http.FileServer，因此会使用 http.NotFound 替代 Router 的 NotFound 处理程序。
// 若要使用操作系统自身的文件系统实现，
// 可以这样使用：
//
//	router.Static("/static", "/var/www")

// ff:
// root:
// relativePath:
func Static(relativePath, root string) gin.IRoutes {
	return engine().Static(relativePath, root)
}

// StaticFS 是 Engine.StaticFS 的一个包装器。

// ff:
// fs:
// relativePath:
func StaticFS(relativePath string, fs http.FileSystem) gin.IRoutes {
	return engine().StaticFS(relativePath, fs)
}

// Use 方法将一个全局中间件附加到路由。也就是说，通过Use()方法附加的中间件将会
// 包含在每一个请求的处理器链中。即便是404、405等错误状态码响应，或是静态文件请求...
// 例如，这是放置日志记录器或错误管理中间件的理想位置。

// ff:中间件
// middleware:处理函数
func Use(middlewares ...gin.HandlerFunc) gin.IRoutes {
	return engine().Use(middlewares...)
}

// Routes 返回已注册路由的切片。

// ff:
func Routes() gin.RoutesInfo {
	return engine().Routes()
}

// Run 函数连接到一个 http.Server，并开始监听和处理 HTTP 请求。
// 这是调用 http.ListenAndServe(addr, router) 的快捷方式。
// 注意：除非发生错误，否则此方法将无限期地阻塞调用它的 goroutine。

// ff:
// err:
// addr:
func Run(addr ...string) (err error) {
	return engine().Run(addr...)
}

// RunTLS 附加到一个 http.Server，并开始监听和处理 HTTPS 请求。
// 这是 http.ListenAndServeTLS(addr, certFile, keyFile, router) 的快捷方式。
// 注意：除非发生错误，否则此方法将无限期地阻塞调用的 goroutine。

// ff:
// err:
// keyFile:
// certFile:
// addr:
func RunTLS(addr, certFile, keyFile string) (err error) {
	return engine().RunTLS(addr, certFile, keyFile)
}

// RunUnix连接到一个http.Server，并开始通过指定的Unix套接字（即文件）监听和处理HTTP请求
// 注意：除非发生错误，否则此方法将无限期地阻塞调用它的goroutine。

// ff:
// err:
// file:
func RunUnix(file string) (err error) {
	return engine().RunUnix(file)
}

// RunFd 将路由器附加到 http.Server，并开始监听并透过指定的文件描述符处理 HTTP 请求。
// 注意：除非发生错误，否则该方法将无限期地阻塞调用它的 goroutine。

// ff:
// err:
// fd:
func RunFd(fd int) (err error) {
	return engine().RunFd(fd)
}
