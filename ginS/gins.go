// Copyright 2014 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

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

// LoadHTMLGlob is a wrapper for Engine.LoadHTMLGlob.
func X加载HTML模板目录(模板目录 string) {
	engine().X加载HTML模板目录(模板目录)
}

// LoadHTMLFiles is a wrapper for Engine.LoadHTMLFiles.
func X加载HTML模板文件(模板文件s ...string) {
	engine().X加载HTML模板文件(模板文件s...)
}

// SetHTMLTemplate is a wrapper for Engine.SetHTMLTemplate.
func X设置Template模板(Template模板 *template.Template) {
	engine().X设置Template模板(Template模板)
}

// NoRoute adds handlers for NoRoute. It returns a 404 code by default.
func X绑定404(处理函数s ...gin类.HandlerFunc) {
	engine().X绑定404(处理函数s...)
}

// NoMethod is a wrapper for Engine.NoMethod.
func X绑定405(处理函数s ...gin类.HandlerFunc) {
	engine().X绑定405(处理函数s...)
}

// Group creates a new router group. You should add all the routes that have common middlewares or the same path prefix.
// For example, all the routes that use a common middleware for authorization could be grouped.
func X创建分组路由(路由规则 string, 处理函数s ...gin类.HandlerFunc) *gin类.RouterGroup {
	return engine().X创建分组路由(路由规则, 处理函数s...)
}

// Handle is a wrapper for Engine.Handle.
func X绑定(HTTP方法, 路由规则 string, 处理函数s ...gin类.HandlerFunc) gin类.IRoutes {
	return engine().X绑定(HTTP方法, 路由规则, 处理函数s...)
}

// POST is a shortcut for router.Handle("POST", path, handle)
func X绑定POST(路由规则 string, 处理函数s ...gin类.HandlerFunc) gin类.IRoutes {
	return engine().X绑定POST(路由规则, 处理函数s...)
}

// GET is a shortcut for router.Handle("GET", path, handle)
func X绑定GET(路由规则 string, 处理函数s ...gin类.HandlerFunc) gin类.IRoutes {
	return engine().X绑定GET(路由规则, 处理函数s...)
}

// DELETE is a shortcut for router.Handle("DELETE", path, handle)
func X绑定DELETE(路由规则 string, 处理函数s ...gin类.HandlerFunc) gin类.IRoutes {
	return engine().X绑定DELETE(路由规则, 处理函数s...)
}

// PATCH is a shortcut for router.Handle("PATCH", path, handle)
func X绑定PATCH(路由规则 string, 处理函数s ...gin类.HandlerFunc) gin类.IRoutes {
	return engine().X绑定PATCH(路由规则, 处理函数s...)
}

// PUT is a shortcut for router.Handle("PUT", path, handle)
func X绑定PUT(路由规则 string, 处理函数s ...gin类.HandlerFunc) gin类.IRoutes {
	return engine().X绑定PUT(路由规则, 处理函数s...)
}

// OPTIONS is a shortcut for router.Handle("OPTIONS", path, handle)
func X绑定OPTIONS(路由规则 string, 处理函数s ...gin类.HandlerFunc) gin类.IRoutes {
	return engine().X绑定OPTIONS(路由规则, 处理函数s...)
}

// HEAD is a shortcut for router.Handle("HEAD", path, handle)
func X绑定HEAD(路由规则 string, 处理函数s ...gin类.HandlerFunc) gin类.IRoutes {
	return engine().X绑定HEAD(路由规则, 处理函数s...)
}

// Any is a wrapper for Engine.Any.
func X绑定Any(路由规则 string, 处理函数s ...gin类.HandlerFunc) gin类.IRoutes {
	return engine().X绑定Any(路由规则, 处理函数s...)
}

// StaticFile is a wrapper for Engine.StaticFile.
func X绑定静态单文件(URL路径, 文件路径 string) gin类.IRoutes {
	return engine().X绑定静态单文件(URL路径, 文件路径)
}

// Static serves files from the given file system root.
// Internally a http.FileServer is used, therefore http.NotFound is used instead
// of the Router's NotFound handler.
// To use the operating system's file system implementation,
// use :
//
//	router.Static("/static", "/var/www")
func X绑定静态文件目录(URL路径前缀, 绑定目录 string) gin类.IRoutes {
	return engine().X绑定静态文件目录(URL路径前缀, 绑定目录)
}

// StaticFS is a wrapper for Engine.StaticFS.
func X绑定静态文件目录FS(URL路径前缀 string, fs http.FileSystem) gin类.IRoutes {
	return engine().X绑定静态文件目录FS(URL路径前缀, fs)
}

// Use attaches a global middleware to the router. i.e. the middlewares attached through Use() will be
// included in the handlers chain for every single request. Even 404, 405, static files...
// For example, this is the right place for a logger or error management middleware.
func X中间件(middlewares ...gin类.HandlerFunc) gin类.IRoutes {
	return engine().X中间件(middlewares...)
}

// Routes returns a slice of registered routes.
func X取路由数组() gin类.RoutesInfo {
	return engine().X取路由数组()
}

// Run attaches to a http.Server and starts listening and serving HTTP requests.
// It is a shortcut for http.ListenAndServe(addr, router)
// Note: this method will block the calling goroutine indefinitely unless an error happens.
func X监听(地址与端口 ...string) (错误 error) {
	return engine().X监听(地址与端口...)
}

// RunTLS attaches to a http.Server and starts listening and serving HTTPS requests.
// It is a shortcut for http.ListenAndServeTLS(addr, certFile, keyFile, router)
// Note: this method will block the calling goroutine indefinitely unless an error happens.
func X监听TLS(地址与端口, cert文件, key文件 string) (错误 error) {
	return engine().X监听TLS(地址与端口, cert文件, key文件)
}

// RunUnix attaches to a http.Server and starts listening and serving HTTP requests
// through the specified unix socket (i.e. a file)
// Note: this method will block the calling goroutine indefinitely unless an error happens.
func X监听Unix(文件路径 string) (错误 error) {
	return engine().X监听Unix(文件路径)
}

// RunFd attaches the router to a http.Server and starts listening and serving HTTP requests
// through the specified file descriptor.
// Note: the method will block the calling goroutine indefinitely unless on error happens.
func X监听Fd(fd int) (错误 error) {
	return engine().X监听Fd(fd)
}
