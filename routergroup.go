// Manu Martinez-Almeida版权所有
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到

package gin

import (
	"net/http"
	"path"
	"regexp"
	"strings"
)

var (
	// regEnLetter匹配http方法名的英文字母
	regEnLetter = regexp.MustCompile("^[A-Z]+$")

	// RouterGroup的anyMethods:任何方法
	anyMethods = []string{
		http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch,
		http.MethodHead, http.MethodOptions, http.MethodDelete, http.MethodConnect,
		http.MethodTrace,
	}
)

// IRouter定义了所有路由器句柄接口，包括单个和组路由器
type IRouter interface {
	IRoutes
	Group(string, ...HandlerFunc) *RouterGroup
}

// irroutes定义了所有路由器句柄接口
type IRoutes interface {
	Use(...HandlerFunc) IRoutes

	Handle(string, string, ...HandlerFunc) IRoutes
	Any(string, ...HandlerFunc) IRoutes
	GET(string, ...HandlerFunc) IRoutes
	POST(string, ...HandlerFunc) IRoutes
	DELETE(string, ...HandlerFunc) IRoutes
	PATCH(string, ...HandlerFunc) IRoutes
	PUT(string, ...HandlerFunc) IRoutes
	OPTIONS(string, ...HandlerFunc) IRoutes
	HEAD(string, ...HandlerFunc) IRoutes
	Match([]string, string, ...HandlerFunc) IRoutes

	StaticFile(string, string) IRoutes
	StaticFileFS(string, string, http.FileSystem) IRoutes
	Static(string, string) IRoutes
	StaticFS(string, http.FileSystem) IRoutes
}

// RouterGroup在内部用于配置路由器，它与一个前缀和一组处理程序(中间件)相关联
type RouterGroup struct {
	Handlers HandlersChain
	basePath string
	engine   *Engine
	root     bool
}

var _ IRouter = (*RouterGroup)(nil)

// 使用将中间件添加到组中，参见GitHub中的示例代码

// ff:中间件
// middleware:处理方法
func (group *RouterGroup) Use(middleware ...HandlerFunc) IRoutes {
	group.Handlers = append(group.Handlers, middleware...)
	return group.returnObj()
}

// Group命令用于创建新的路由器组
// 您应该添加所有具有相同中间件或相同路径前缀的路由
// 例如，可以对使用公共中间件进行授权的所有路由进行分组

// ff:创建分组路由
// handlers:处理方法
// relativePath:路由规则
func (group *RouterGroup) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup {
	return &RouterGroup{
		Handlers: group.combineHandlers(handlers),
		basePath: group.calculateAbsolutePath(relativePath),
		engine:   group.engine,
	}
}

// BasePath返回路由器组的基路径
// 例如:v:= router.Group("/rest/n/v1/api")，则v. basepath()为"/rest/n/v1/api"

// ff:取路由基础路径
func (group *RouterGroup) BasePath() string {
	return group.basePath
}

func (group *RouterGroup) handle(httpMethod, relativePath string, handlers HandlersChain) IRoutes {
	absolutePath := group.calculateAbsolutePath(relativePath)
	handlers = group.combineHandlers(handlers)
	group.engine.addRoute(httpMethod, absolutePath, handlers)
	return group.returnObj()
}

// Handle用给定的路径和方法注册一个新的请求句柄和中间件
// 最后一个处理程序应该是真正的处理程序，其他的应该是中间件，可以并且应该在不同的路由之间共享
// 参见GitHub中的示例代码
// 对于GET、POST、PUT、PATCH和DELETE请求，可以使用各自的快捷函数
// 此功能用于批量加载，并允许使用不太常用的非标准化或自定义方法(例如用于与代理的内部通信)

// ff:绑定
// handlers:处理方法
// relativePath:路由规则
// httpMethod:HTTP方法
func (group *RouterGroup) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) IRoutes {
	if matched := regEnLetter.MatchString(httpMethod); !matched {
		panic("http method " + httpMethod + " is not valid")
	}
	return group.handle(httpMethod, relativePath, handlers)
}

// POST是router.Handle("POST" path, handlers)的快捷方式

// ff:绑定POST
// handlers:处理方法
// relativePath:路由规则
func (group *RouterGroup) POST(relativePath string, handlers ...HandlerFunc) IRoutes {
	return group.handle(http.MethodPost, relativePath, handlers)
}

// GET是router.Handle("GET" path, handlers)的快捷方式

// ff:绑定GET
// handlers:处理方法
// relativePath:路由规则
func (group *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) IRoutes {
	return group.handle(http.MethodGet, relativePath, handlers)
}

// DELETE是router.Handle("DELETE"， path, handlers)的快捷方式

// ff:绑定DELETE
// handlers:处理方法
// relativePath:路由规则
func (group *RouterGroup) DELETE(relativePath string, handlers ...HandlerFunc) IRoutes {
	return group.handle(http.MethodDelete, relativePath, handlers)
}

// PATCH是router.Handle("PATCH" path, handlers)的快捷方式

// ff:绑定PATCH
// handlers:处理方法
// relativePath:路由规则
func (group *RouterGroup) PATCH(relativePath string, handlers ...HandlerFunc) IRoutes {
	return group.handle(http.MethodPatch, relativePath, handlers)
}

// PUT是router.Handle("PUT" path, handlers)的快捷方式

// ff:绑定PUT
// handlers:处理方法
// relativePath:路由规则
func (group *RouterGroup) PUT(relativePath string, handlers ...HandlerFunc) IRoutes {
	return group.handle(http.MethodPut, relativePath, handlers)
}

// OPTIONS是router.Handle("OPTIONS" path, handlers)的快捷方式

// ff:绑定OPTIONS
// handlers:处理方法
// relativePath:路由规则
func (group *RouterGroup) OPTIONS(relativePath string, handlers ...HandlerFunc) IRoutes {
	return group.handle(http.MethodOptions, relativePath, handlers)
}

// HEAD是router.Handle("HEAD" path, handlers)的快捷方式

// ff:绑定HEAD
// handlers:处理方法
// relativePath:路由规则
func (group *RouterGroup) HEAD(relativePath string, handlers ...HandlerFunc) IRoutes {
	return group.handle(http.MethodHead, relativePath, handlers)
}

// Any注册了一个匹配所有HTTP方法的路由
// Get, post, put, patch, head, options, delete, connect, trace

// ff:绑定Any
// handlers:处理方法
// relativePath:路由规则
func (group *RouterGroup) Any(relativePath string, handlers ...HandlerFunc) IRoutes {
	for _, method := range anyMethods {
		group.handle(method, relativePath, handlers)
	}

	return group.returnObj()
}

// Match注册一个与你声明的指定方法匹配的路由

// ff:
// handlers:
// relativePath:
// methods:
func (group *RouterGroup) Match(methods []string, relativePath string, handlers ...HandlerFunc) IRoutes {
	for _, method := range methods {
		group.handle(method, relativePath, handlers)
	}

	return group.returnObj()
}

// 为了服务本地文件系统的单个文件，StaticFile注册单个路由
// router.StaticFile(“favicon.ico&quot“
// /资源/ favicon.ico")

// ff:
// filepath:
// relativePath:
func (group *RouterGroup) StaticFile(relativePath, filepath string) IRoutes {
	return group.staticFileHandler(relativePath, func(c *Context) {
		c.File(filepath)
	})
}

// StaticFileFS的工作原理就像' StaticFile '，但一个自定义的' http
// 可以用FileSystem代替
// router. staticfiles ("favicon.ico"， " /resources/favicon.ico"， Dir{"."， false}) Gin默认使用:Gin .Dir()

// ff:
// fs:
// filepath:
// relativePath:
func (group *RouterGroup) StaticFileFS(relativePath, filepath string, fs http.FileSystem) IRoutes {
	return group.staticFileHandler(relativePath, func(c *Context) {
		c.FileFromFS(filepath, fs)
	})
}

func (group *RouterGroup) staticFileHandler(relativePath string, handler HandlerFunc) IRoutes {
	if strings.Contains(relativePath, ":") || strings.Contains(relativePath, "*") {
		panic("URL parameters can not be used when serving a static file")
	}
	group.GET(relativePath, handler)
	group.HEAD(relativePath, handler)
	return group.returnObj()
}

// 静态从给定的文件系统根目录提供文件
// 内部http
// 使用的是FileServer，因此是http
// 使用NotFound来代替路由器的NotFound处理程序
// 要使用操作系统的文件系统实现，使用:router.Static("/static"， "/var/www")

// ff:
// root:
// relativePath:
func (group *RouterGroup) Static(relativePath, root string) IRoutes {
	return group.StaticFS(relativePath, Dir(root, false))
}

// StaticFS的工作原理就像' Static() '，但一个自定义的' http
// 可以使用FileSystem’代替
// Gin默认使用:Gin . dir ()

// ff:
// fs:
// relativePath:
func (group *RouterGroup) StaticFS(relativePath string, fs http.FileSystem) IRoutes {
	if strings.Contains(relativePath, ":") || strings.Contains(relativePath, "*") {
		panic("URL parameters can not be used when serving a static folder")
	}
	handler := group.createStaticHandler(relativePath, fs)
	urlPattern := path.Join(relativePath, "/*filepath")

	// 注册GET和HEAD处理程序
	group.GET(urlPattern, handler)
	group.HEAD(urlPattern, handler)
	return group.returnObj()
}

func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := group.calculateAbsolutePath(relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))

	return func(c *Context) {
		if _, noListing := fs.(*onlyFilesFS); noListing {
			c.Writer.WriteHeader(http.StatusNotFound)
		}

		file := c.Param("filepath")
		// 检查文件是否存在和/或我们是否有访问它的权限
		f, err := fs.Open(file)
		if err != nil {
			c.Writer.WriteHeader(http.StatusNotFound)
			c.handlers = group.engine.noRoute
			// 重置指数
			c.index = -1
			return
		}
		f.Close()

		fileServer.ServeHTTP(c.Writer, c.Request)
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
