// 版权所有 2014 Manu Martinez-Almeida。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package gin类

import (
	"net/http"
	"path"
	"regexp"
	"strings"
)

var (
	// regEnLetter 匹配用于HTTP方法名称的英文字母
	regEnLetter = regexp.MustCompile("^[A-Z]+$")

	// anyMethods 用于 RouterGroup 的任意方法
	anyMethods = []string{
		http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch,
		http.MethodHead, http.MethodOptions, http.MethodDelete, http.MethodConnect,
		http.MethodTrace,
	}
)

// IRouter 定义了包括单个和组合路由在内的所有路由处理器接口。
type IRouter interface {
	IRoutes
	X创建分组路由(string, ...HandlerFunc) *RouterGroup
}

// IRoutes 定义了所有路由处理器的接口。
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

// RouterGroup 用于内部配置路由器，一个 RouterGroup 与一个前缀和一组处理器（中间件）关联。
type RouterGroup struct {
	Handlers HandlersChain
	basePath string
	engine   *Engine
	root     bool
}

var _ IRouter = (*RouterGroup)(nil)

// Use 向组中添加中间件，参见 GitHub 中的示例代码。
func (group *RouterGroup) X中间件(处理函数 ...HandlerFunc) IRoutes {
	group.Handlers = append(group.Handlers, 处理函数...)
	return group.returnObj()
}

// Group 创建一个新的路由分组。你应该在此添加所有具有共同中间件或相同路径前缀的路由。
// 例如，所有使用共同授权中间件的路由可以被归为一组。
func (group *RouterGroup) X创建分组路由(路由规则 string, 处理函数 ...HandlerFunc) *RouterGroup {
	return &RouterGroup{
		Handlers: group.combineHandlers(处理函数),
		basePath: group.calculateAbsolutePath(路由规则),
		engine:   group.engine,
	}
}

// BasePath 返回路由组的基础路径。
// 例如，如果 v := router.Group("/rest/n/v1/api")，那么 v.BasePath() 的值就是 "/rest/n/v1/api"。
func (group *RouterGroup) X取路由基础路径() string {
	return group.basePath
}

func (group *RouterGroup) handle(httpMethod, relativePath string, handlers HandlersChain) IRoutes {
	absolutePath := group.calculateAbsolutePath(relativePath)
	handlers = group.combineHandlers(handlers)
	group.engine.addRoute(httpMethod, absolutePath, handlers)
	return group.returnObj()
}

// Handle 注册一个新的请求处理程序和中间件，关联给定的路径和方法。
// 最后一个处理器应该是真正的处理器，其他的应该是可以在不同路由间共享的中间件。
// 请参阅 GitHub 上的示例代码。
//
// 对于 GET、POST、PUT、PATCH 和 DELETE 请求，可以分别使用相应的快捷函数。
//
// 该函数主要用于批量加载，并允许使用不太常用、非标准化或自定义的方法（例如，用于与代理的内部通信）。
func (group *RouterGroup) X绑定(HTTP方法, 路由规则 string, 处理函数 ...HandlerFunc) IRoutes {
	if matched := regEnLetter.MatchString(HTTP方法); !matched {
		panic("http method " + HTTP方法 + " is not valid")
	}
	return group.handle(HTTP方法, 路由规则, 处理函数)
}

// POST 是一个快捷方式，用于 router.Handle("POST", path, handlers)。
func (group *RouterGroup) X绑定POST(路由规则 string, 处理函数 ...HandlerFunc) IRoutes {
	return group.handle(http.MethodPost, 路由规则, 处理函数)
}

// GET 是一个快捷方式，用于 router.Handle("GET", path, handlers)。
func (group *RouterGroup) X绑定GET(路由规则 string, 处理函数 ...HandlerFunc) IRoutes {
	return group.handle(http.MethodGet, 路由规则, 处理函数)
}

// DELETE 是一个快捷方式，用于 router.Handle("DELETE", path, handlers)。
func (group *RouterGroup) X绑定DELETE(路由规则 string, 处理函数 ...HandlerFunc) IRoutes {
	return group.handle(http.MethodDelete, 路由规则, 处理函数)
}

// PATCH 是一个快捷方式，用于 router.Handle("PATCH", path, handlers)。
func (group *RouterGroup) X绑定PATCH(路由规则 string, 处理函数 ...HandlerFunc) IRoutes {
	return group.handle(http.MethodPatch, 路由规则, 处理函数)
}

// PUT 是一个快捷方式，用于 router.Handle("PUT", path, handlers)。
func (group *RouterGroup) X绑定PUT(路由规则 string, 处理函数 ...HandlerFunc) IRoutes {
	return group.handle(http.MethodPut, 路由规则, 处理函数)
}

// OPTIONS 是一个快捷方式，用于 router.Handle("OPTIONS", path, handlers)。
func (group *RouterGroup) X绑定OPTIONS(路由规则 string, 处理函数 ...HandlerFunc) IRoutes {
	return group.handle(http.MethodOptions, 路由规则, 处理函数)
}

// HEAD 是一个快捷方式，用于router.Handle("HEAD", path, handlers)。
func (group *RouterGroup) X绑定HEAD(路由规则 string, 处理函数 ...HandlerFunc) IRoutes {
	return group.handle(http.MethodHead, 路由规则, 处理函数)
}

// Any 注册一个路由，该路由会匹配所有 HTTP 方法。
// 包括 GET, POST, PUT, PATCH, HEAD, OPTIONS, DELETE, CONNECT 和 TRACE。
func (group *RouterGroup) X绑定Any(路由规则 string, 处理函数 ...HandlerFunc) IRoutes {
	for _, method := range anyMethods {
		group.handle(method, 路由规则, 处理函数)
	}

	return group.returnObj()
}

// Match 注册一条路由，该路由会匹配你声明的指定方法。
func (group *RouterGroup) Match(methods []string, relativePath string, handlers ...HandlerFunc) IRoutes {
	for _, method := range methods {
		group.handle(method, relativePath, handlers)
	}

	return group.returnObj()
}

// StaticFile 注册一个单独的路由，以便从本地文件系统服务单个文件。
// 示例：router.StaticFile("favicon.ico", "./resources/favicon.ico")
func (group *RouterGroup) X绑定静态单文件(URL路径, 文件路径 string) IRoutes {
	return group.staticFileHandler(URL路径, func(c *Context) {
		c.X下载文件(文件路径)
	})
}

// StaticFileFS 的工作方式与 `StaticFile` 类似，但是可以使用自定义的 `http.FileSystem`。
// 示例：router.StaticFileFS("favicon.ico", "./resources/favicon.ico", Dir{".", false})
// Gin 框架默认使用：gin.Dir()
// 翻译：
// StaticFileFS 与 `StaticFile` 功能类似，但允许使用自定义的 `http.FileSystem` 实例替代。
// 例如：通过 router.StaticFileFS 将 favicon.ico 静态文件映射到 "./resources/favicon.ico"，并指定 Dir{".", false} 参数。
// Gin 框架默认情况下使用的静态文件系统为 gin.Dir()。
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

// Static 从给定的文件系统根目录提供文件服务。
// 在内部使用了 http.FileServer，因此会使用 http.NotFound 替代 Router 的 NotFound 处理程序。
// 要使用操作系统自身的文件系统实现，请按如下方式使用：
//
//	router.Static("/static", "/var/www")
func (group *RouterGroup) X绑定静态文件目录(URL路径前缀, 绑定目录 string) IRoutes {
	return group.X绑定静态文件目录FS(URL路径前缀, Dir(绑定目录, false))
}

// StaticFS的功能类似于`Static()`，但是可以使用自定义的`http.FileSystem`。
// 如: router.StaticFS("/more_static", http.Dir("my_file_system")) 
// Gin默认使用: gin.Dir()
//
// 注意: StaticFS 比Static一个多了个功能，当目录下不存 index.html 文件，会直接列出该目录下的所有文件。
func (group *RouterGroup) X绑定静态文件目录FS(URL路径前缀 string, fs http.FileSystem) IRoutes {
	if strings.Contains(URL路径前缀, ":") || strings.Contains(URL路径前缀, "*") {
		panic("URL parameters can not be used when serving a static folder")
	}
	handler := group.createStaticHandler(URL路径前缀, fs)
	urlPattern := path.Join(URL路径前缀, "/*filepath")

	// 注册 GET 和 HEAD 请求处理器
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
		// 检查文件是否存在，以及/或我们是否有权限访问它
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
