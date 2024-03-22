// Copyright 2014 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gin类

import (
	"net/http"
	"path"
	"regexp"
	"strings"
)

var (
	// regEnLetter matches english letters for http method name
	regEnLetter = regexp.MustCompile("^[A-Z]+$")

	// anyMethods for RouterGroup Any method
	anyMethods = []string{
		http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch,
		http.MethodHead, http.MethodOptions, http.MethodDelete, http.MethodConnect,
		http.MethodTrace,
	}
)

// IRouter defines all router handle interface includes single and group router.
type IRouter interface {
	IRoutes
	X创建分组路由(string, ...HandlerFunc) *RouterGroup
}

// IRoutes defines all router handle interface.
type IRoutes interface {
	X中间件(...HandlerFunc) IRoutes

	X绑定(string, string, ...HandlerFunc) IRoutes
	X绑定Any(string, ...HandlerFunc) IRoutes
	X绑定GET(string, ...HandlerFunc) IRoutes
	X绑定POST(string, ...HandlerFunc) IRoutes
	X绑定DELETE(string, ...HandlerFunc) IRoutes
	X绑定PATCH(string, ...HandlerFunc) IRoutes
	X绑定PUT(string, ...HandlerFunc) IRoutes
	X绑定OPTIONS(string, ...HandlerFunc) IRoutes
	X绑定HEAD(string, ...HandlerFunc) IRoutes
	Match([]string, string, ...HandlerFunc) IRoutes

	X绑定静态单文件(string, string) IRoutes
	X绑定静态单文件FS(string, string, http.FileSystem) IRoutes
	X绑定静态文件目录(string, string) IRoutes
	X绑定静态文件目录FS(string, http.FileSystem) IRoutes
}

// RouterGroup is used internally to configure router, a RouterGroup is associated with
// a prefix and an array of handlers (middleware).
type RouterGroup struct {
	Handlers HandlersChain
	basePath string
	engine   *Engine
	root     bool
}

var _ IRouter = (*RouterGroup)(nil)

// Use adds middleware to the group, see example code in GitHub.
func (group *RouterGroup) X中间件(处理函数 ...HandlerFunc) IRoutes {
	group.Handlers = append(group.Handlers, 处理函数...)
	return group.returnObj()
}

// Group creates a new router group. You should add all the routes that have common middlewares or the same path prefix.
// For example, all the routes that use a common middleware for authorization could be grouped.
func (group *RouterGroup) X创建分组路由(路由规则 string, 处理函数 ...HandlerFunc) *RouterGroup {
	return &RouterGroup{
		Handlers: group.combineHandlers(处理函数),
		basePath: group.calculateAbsolutePath(路由规则),
		engine:   group.engine,
	}
}

// BasePath returns the base path of router group.
// For example, if v := router.Group("/rest/n/v1/api"), v.BasePath() is "/rest/n/v1/api".
func (group *RouterGroup) X取路由基础路径() string {
	return group.basePath
}

func (group *RouterGroup) handle(httpMethod, relativePath string, handlers HandlersChain) IRoutes {
	absolutePath := group.calculateAbsolutePath(relativePath)
	handlers = group.combineHandlers(handlers)
	group.engine.addRoute(httpMethod, absolutePath, handlers)
	return group.returnObj()
}

// Handle registers a new request handle and middleware with the given path and method.
// The last handler should be the real handler, the other ones should be middleware that can and should be shared among different routes.
// See the example code in GitHub.
//
// For GET, POST, PUT, PATCH and DELETE requests the respective shortcut
// functions can be used.
//
// This function is intended for bulk loading and to allow the usage of less
// frequently used, non-standardized or custom methods (e.g. for internal
// communication with a proxy).
func (group *RouterGroup) X绑定(HTTP方法, 路由规则 string, 处理函数 ...HandlerFunc) IRoutes {
	if matched := regEnLetter.MatchString(HTTP方法); !matched {
		panic("http method " + HTTP方法 + " is not valid")
	}
	return group.handle(HTTP方法, 路由规则, 处理函数)
}

// POST is a shortcut for router.Handle("POST", path, handlers).
func (group *RouterGroup) X绑定POST(路由规则 string, 处理函数 ...HandlerFunc) IRoutes {
	return group.handle(http.MethodPost, 路由规则, 处理函数)
}

// GET is a shortcut for router.Handle("GET", path, handlers).
func (group *RouterGroup) X绑定GET(路由规则 string, 处理函数 ...HandlerFunc) IRoutes {
	return group.handle(http.MethodGet, 路由规则, 处理函数)
}

// DELETE is a shortcut for router.Handle("DELETE", path, handlers).
func (group *RouterGroup) X绑定DELETE(路由规则 string, 处理函数 ...HandlerFunc) IRoutes {
	return group.handle(http.MethodDelete, 路由规则, 处理函数)
}

// PATCH is a shortcut for router.Handle("PATCH", path, handlers).
func (group *RouterGroup) X绑定PATCH(路由规则 string, 处理函数 ...HandlerFunc) IRoutes {
	return group.handle(http.MethodPatch, 路由规则, 处理函数)
}

// PUT is a shortcut for router.Handle("PUT", path, handlers).
func (group *RouterGroup) X绑定PUT(路由规则 string, 处理函数 ...HandlerFunc) IRoutes {
	return group.handle(http.MethodPut, 路由规则, 处理函数)
}

// OPTIONS is a shortcut for router.Handle("OPTIONS", path, handlers).
func (group *RouterGroup) X绑定OPTIONS(路由规则 string, 处理函数 ...HandlerFunc) IRoutes {
	return group.handle(http.MethodOptions, 路由规则, 处理函数)
}

// HEAD is a shortcut for router.Handle("HEAD", path, handlers).
func (group *RouterGroup) X绑定HEAD(路由规则 string, 处理函数 ...HandlerFunc) IRoutes {
	return group.handle(http.MethodHead, 路由规则, 处理函数)
}

// Any registers a route that matches all the HTTP methods.
// GET, POST, PUT, PATCH, HEAD, OPTIONS, DELETE, CONNECT, TRACE.
func (group *RouterGroup) X绑定Any(路由规则 string, 处理函数 ...HandlerFunc) IRoutes {
	for _, method := range anyMethods {
		group.handle(method, 路由规则, 处理函数)
	}

	return group.returnObj()
}

// Match registers a route that matches the specified methods that you declared.
func (group *RouterGroup) Match(methods []string, relativePath string, handlers ...HandlerFunc) IRoutes {
	for _, method := range methods {
		group.handle(method, relativePath, handlers)
	}

	return group.returnObj()
}

// StaticFile registers a single route in order to serve a single file of the local filesystem.
// router.StaticFile("favicon.ico", "./resources/favicon.ico")
func (group *RouterGroup) X绑定静态单文件(URL路径, 文件路径 string) IRoutes {
	return group.staticFileHandler(URL路径, func(c *Context) {
		c.X下载文件(文件路径)
	})
}

// StaticFileFS works just like `StaticFile` but a custom `http.FileSystem` can be used instead..
// router.StaticFileFS("favicon.ico", "./resources/favicon.ico", Dir{".", false})
// Gin by default uses: gin.Dir()
func (group *RouterGroup) X绑定静态单文件FS(URL路径, 文件路径 string, fs http.FileSystem) IRoutes {
	return group.staticFileHandler(URL路径, func(c *Context) {
		c.X下载文件FS(文件路径, fs)
	})
}

func (group *RouterGroup) staticFileHandler(relativePath string, handler HandlerFunc) IRoutes {
	if strings.Contains(relativePath, ":") || strings.Contains(relativePath, "*") {
		panic("URL parameters can not be used when serving a static file")
	}
	group.X绑定GET(relativePath, handler)
	group.X绑定HEAD(relativePath, handler)
	return group.returnObj()
}

// Static serves files from the given file system root.
// Internally a http.FileServer is used, therefore http.NotFound is used instead
// of the Router's NotFound handler.
// To use the operating system's file system implementation,
// use :
//
//	router.Static("/static", "/var/www")
func (group *RouterGroup) X绑定静态文件目录(URL路径前缀, 绑定目录 string) IRoutes {
	return group.X绑定静态文件目录FS(URL路径前缀, Dir(绑定目录, false))
}

// StaticFS works just like `Static()` but a custom `http.FileSystem` can be used instead.
// Gin by default uses: gin.Dir()
func (group *RouterGroup) X绑定静态文件目录FS(URL路径前缀 string, fs http.FileSystem) IRoutes {
	if strings.Contains(URL路径前缀, ":") || strings.Contains(URL路径前缀, "*") {
		panic("URL parameters can not be used when serving a static folder")
	}
	handler := group.createStaticHandler(URL路径前缀, fs)
	urlPattern := path.Join(URL路径前缀, "/*filepath")

	// Register GET and HEAD handlers
	group.X绑定GET(urlPattern, handler)
	group.X绑定HEAD(urlPattern, handler)
	return group.returnObj()
}

func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := group.calculateAbsolutePath(relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))

	return func(c *Context) {
		if _, noListing := fs.(*onlyFilesFS); noListing {
			c.Writer.WriteHeader(http.StatusNotFound)
		}

		file := c.X取API参数值("filepath")
		// Check if file exists and/or if we have permission to access it
		f, err := fs.Open(file)
		if err != nil {
			c.Writer.WriteHeader(http.StatusNotFound)
			c.handlers = group.engine.noRoute
			// Reset index
			c.index = -1
			return
		}
		f.Close()

		fileServer.ServeHTTP(c.Writer, c.X请求)
	}
}

func (group *RouterGroup) combineHandlers(handlers HandlersChain) HandlersChain {
	finalSize := len(group.Handlers) + len(handlers)
	assert1(finalSize < int(abortIndex), "too many handlers")
	mergedHandlers := make(HandlersChain, finalSize)
	copy(mergedHandlers, group.Handlers)
	copy(mergedHandlers[len(group.Handlers):], handlers)
	return mergedHandlers
}

func (group *RouterGroup) calculateAbsolutePath(relativePath string) string {
	return joinPaths(group.basePath, relativePath)
}

func (group *RouterGroup) returnObj() IRoutes {
	if group.root {
		return group.engine
	}
	return group
}
