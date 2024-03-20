
<原文开始>
// Copyright 2014 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2014 Manu Martinez-Almeida。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// regEnLetter matches english letters for http method name
<原文结束>

# <翻译开始>
// regEnLetter 匹配用于HTTP方法名称的英文字母
# <翻译结束>


<原文开始>
// anyMethods for RouterGroup Any method
<原文结束>

# <翻译开始>
// anyMethods 用于 RouterGroup 的任意方法
# <翻译结束>


<原文开始>
// IRouter defines all router handle interface includes single and group router.
<原文结束>

# <翻译开始>
// IRouter 定义了包括单个和组合路由在内的所有路由处理器接口。
# <翻译结束>


<原文开始>
// IRoutes defines all router handle interface.
<原文结束>

# <翻译开始>
// IRoutes 定义了所有路由处理器的接口。
# <翻译结束>


<原文开始>
// RouterGroup is used internally to configure router, a RouterGroup is associated with
// a prefix and an array of handlers (middleware).
<原文结束>

# <翻译开始>
// RouterGroup 用于内部配置路由器，一个 RouterGroup 与一个前缀和一组处理器（中间件）关联。
# <翻译结束>


<原文开始>
// Use adds middleware to the group, see example code in GitHub.
<原文结束>

# <翻译开始>
// Use 向组中添加中间件，参见 GitHub 中的示例代码。
# <翻译结束>


<原文开始>
// Group creates a new router group. You should add all the routes that have common middlewares or the same path prefix.
// For example, all the routes that use a common middleware for authorization could be grouped.
<原文结束>

# <翻译开始>
// Group 创建一个新的路由分组。你应该在此添加所有具有共同中间件或相同路径前缀的路由。
// 例如，所有使用共同授权中间件的路由可以被归为一组。
# <翻译结束>


<原文开始>
// BasePath returns the base path of router group.
// For example, if v := router.Group("/rest/n/v1/api"), v.BasePath() is "/rest/n/v1/api".
<原文结束>

# <翻译开始>
// BasePath 返回路由组的基础路径。
// 例如，如果 v := router.Group("/rest/n/v1/api")，那么 v.BasePath() 的值就是 "/rest/n/v1/api"。
# <翻译结束>


<原文开始>
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
<原文结束>

# <翻译开始>
// Handle 注册一个新的请求处理程序和中间件，关联给定的路径和方法。
// 最后一个处理器应该是真正的处理器，其他的应该是可以在不同路由间共享的中间件。
// 请参阅 GitHub 上的示例代码。
//
// 对于 GET、POST、PUT、PATCH 和 DELETE 请求，可以分别使用相应的快捷函数。
//
// 该函数主要用于批量加载，并允许使用不太常用、非标准化或自定义的方法（例如，用于与代理的内部通信）。
# <翻译结束>


<原文开始>
// POST is a shortcut for router.Handle("POST", path, handlers).
<原文结束>

# <翻译开始>
// POST 是一个快捷方式，用于 router.Handle("POST", path, handlers)。
# <翻译结束>


<原文开始>
// GET is a shortcut for router.Handle("GET", path, handlers).
<原文结束>

# <翻译开始>
// GET 是一个快捷方式，用于 router.Handle("GET", path, handlers)。
# <翻译结束>


<原文开始>
// DELETE is a shortcut for router.Handle("DELETE", path, handlers).
<原文结束>

# <翻译开始>
// DELETE 是一个快捷方式，用于 router.Handle("DELETE", path, handlers)。
# <翻译结束>


<原文开始>
// PATCH is a shortcut for router.Handle("PATCH", path, handlers).
<原文结束>

# <翻译开始>
// PATCH 是一个快捷方式，用于 router.Handle("PATCH", path, handlers)。
# <翻译结束>


<原文开始>
// PUT is a shortcut for router.Handle("PUT", path, handlers).
<原文结束>

# <翻译开始>
// PUT 是一个快捷方式，用于 router.Handle("PUT", path, handlers)。
# <翻译结束>


<原文开始>
// OPTIONS is a shortcut for router.Handle("OPTIONS", path, handlers).
<原文结束>

# <翻译开始>
// OPTIONS 是一个快捷方式，用于 router.Handle("OPTIONS", path, handlers)。
# <翻译结束>


<原文开始>
// HEAD is a shortcut for router.Handle("HEAD", path, handlers).
<原文结束>

# <翻译开始>
// HEAD 是一个快捷方式，用于router.Handle("HEAD", path, handlers)。
# <翻译结束>


<原文开始>
// Any registers a route that matches all the HTTP methods.
// GET, POST, PUT, PATCH, HEAD, OPTIONS, DELETE, CONNECT, TRACE.
<原文结束>

# <翻译开始>
// Any 注册一个路由，该路由会匹配所有 HTTP 方法。
// 包括 GET, POST, PUT, PATCH, HEAD, OPTIONS, DELETE, CONNECT 和 TRACE。
# <翻译结束>


<原文开始>
// Match registers a route that matches the specified methods that you declared.
<原文结束>

# <翻译开始>
// Match 注册一条路由，该路由会匹配你声明的指定方法。
# <翻译结束>


<原文开始>
// StaticFile registers a single route in order to serve a single file of the local filesystem.
// router.StaticFile("favicon.ico", "./resources/favicon.ico")
<原文结束>

# <翻译开始>
// StaticFile 注册一个单独的路由，以便从本地文件系统服务单个文件。
// 示例：router.StaticFile("favicon.ico", "./resources/favicon.ico")
# <翻译结束>


<原文开始>
// StaticFileFS works just like `StaticFile` but a custom `http.FileSystem` can be used instead..
// router.StaticFileFS("favicon.ico", "./resources/favicon.ico", Dir{".", false})
// Gin by default uses: gin.Dir()
<原文结束>

# <翻译开始>
// StaticFileFS 的工作方式与 `StaticFile` 类似，但是可以使用自定义的 `http.FileSystem`。
// 示例：router.StaticFileFS("favicon.ico", "./resources/favicon.ico", Dir{".", false})
// Gin 框架默认使用：gin.Dir()
// 翻译：
// StaticFileFS 与 `StaticFile` 功能类似，但允许使用自定义的 `http.FileSystem` 实例替代。
// 例如：通过 router.StaticFileFS 将 favicon.ico 静态文件映射到 "./resources/favicon.ico"，并指定 Dir{".", false} 参数。
// Gin 框架默认情况下使用的静态文件系统为 gin.Dir()。
# <翻译结束>


<原文开始>
// Static serves files from the given file system root.
// Internally a http.FileServer is used, therefore http.NotFound is used instead
// of the Router's NotFound handler.
// To use the operating system's file system implementation,
// use :
//
//	router.Static("/static", "/var/www")
<原文结束>

# <翻译开始>
// Static 从给定的文件系统根目录提供文件服务。
// 在内部使用了 http.FileServer，因此会使用 http.NotFound 替代 Router 的 NotFound 处理程序。
// 要使用操作系统自身的文件系统实现，请按如下方式使用：
//
//	router.Static("/static", "/var/www")
# <翻译结束>


<原文开始>
// StaticFS works just like `Static()` but a custom `http.FileSystem` can be used instead.
// Gin by default uses: gin.Dir()
<原文结束>

# <翻译开始>
// StaticFS的功能类似于`Static()`，但是可以使用自定义的`http.FileSystem`。
// 如: router.StaticFS("/more_static", http.Dir("my_file_system")) 
// Gin默认使用: gin.Dir()
//
// 注意: StaticFS 比Static一个多了个功能，当目录下不存 index.html 文件，会直接列出该目录下的所有文件。
# <翻译结束>


<原文开始>
// Register GET and HEAD handlers
<原文结束>

# <翻译开始>
// 注册 GET 和 HEAD 请求处理器
# <翻译结束>


<原文开始>
// Check if file exists and/or if we have permission to access it
<原文结束>

# <翻译开始>
// 检查文件是否存在，以及/或我们是否有权限访问它
# <翻译结束>

