
<原文开始>
// Copyright 2014 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Manu Martinez-Almeida版权所有
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到
# <翻译结束>


<原文开始>
	// regEnLetter matches english letters for http method name
<原文结束>

# <翻译开始>
// regEnLetter匹配http方法名的英文字母
# <翻译结束>


<原文开始>
	// anyMethods for RouterGroup Any method
<原文结束>

# <翻译开始>
// RouterGroup的anyMethods:任何方法
# <翻译结束>


<原文开始>
// IRouter defines all router handle interface includes single and group router.
<原文结束>

# <翻译开始>
// IRouter定义了所有路由器句柄接口，包括单个和组路由器
# <翻译结束>


<原文开始>
// IRoutes defines all router handle interface.
<原文结束>

# <翻译开始>
// irroutes定义了所有路由器句柄接口
# <翻译结束>


<原文开始>
// RouterGroup is used internally to configure router, a RouterGroup is associated with
// a prefix and an array of handlers (middleware).
<原文结束>

# <翻译开始>
// RouterGroup在内部用于配置路由器，它与一个前缀和一组处理程序(中间件)相关联
# <翻译结束>


<原文开始>
// Use adds middleware to the group, see example code in GitHub.
<原文结束>

# <翻译开始>
// 使用将中间件添加到组中，参见GitHub中的示例代码
# <翻译结束>


<原文开始>
// Group creates a new router group. You should add all the routes that have common middlewares or the same path prefix.
// For example, all the routes that use a common middleware for authorization could be grouped.
<原文结束>

# <翻译开始>
// Group命令用于创建新的路由器组
// 您应该添加所有具有相同中间件或相同路径前缀的路由
// 例如，可以对使用公共中间件进行授权的所有路由进行分组
# <翻译结束>


<原文开始>
// BasePath returns the base path of router group.
// For example, if v := router.Group("/rest/n/v1/api"), v.BasePath() is "/rest/n/v1/api".
<原文结束>

# <翻译开始>
// BasePath返回路由器组的基路径
// 例如:v:= router.Group("/rest/n/v1/api")，则v. basepath()为"/rest/n/v1/api"
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
// Handle用给定的路径和方法注册一个新的请求句柄和中间件
// 最后一个处理程序应该是真正的处理程序，其他的应该是中间件，可以并且应该在不同的路由之间共享
// 参见GitHub中的示例代码
// 对于GET、POST、PUT、PATCH和DELETE请求，可以使用各自的快捷函数
// 此功能用于批量加载，并允许使用不太常用的非标准化或自定义方法(例如用于与代理的内部通信)
# <翻译结束>


<原文开始>
// POST is a shortcut for router.Handle("POST", path, handlers).
<原文结束>

# <翻译开始>
// POST是router.Handle("POST" path, handlers)的快捷方式
# <翻译结束>


<原文开始>
// GET is a shortcut for router.Handle("GET", path, handlers).
<原文结束>

# <翻译开始>
// GET是router.Handle("GET" path, handlers)的快捷方式
# <翻译结束>


<原文开始>
// DELETE is a shortcut for router.Handle("DELETE", path, handlers).
<原文结束>

# <翻译开始>
// DELETE是router.Handle("DELETE"， path, handlers)的快捷方式
# <翻译结束>


<原文开始>
// PATCH is a shortcut for router.Handle("PATCH", path, handlers).
<原文结束>

# <翻译开始>
// PATCH是router.Handle("PATCH" path, handlers)的快捷方式
# <翻译结束>


<原文开始>
// PUT is a shortcut for router.Handle("PUT", path, handlers).
<原文结束>

# <翻译开始>
// PUT是router.Handle("PUT" path, handlers)的快捷方式
# <翻译结束>


<原文开始>
// OPTIONS is a shortcut for router.Handle("OPTIONS", path, handlers).
<原文结束>

# <翻译开始>
// OPTIONS是router.Handle("OPTIONS" path, handlers)的快捷方式
# <翻译结束>


<原文开始>
// HEAD is a shortcut for router.Handle("HEAD", path, handlers).
<原文结束>

# <翻译开始>
// HEAD是router.Handle("HEAD" path, handlers)的快捷方式
# <翻译结束>


<原文开始>
// Any registers a route that matches all the HTTP methods.
// GET, POST, PUT, PATCH, HEAD, OPTIONS, DELETE, CONNECT, TRACE.
<原文结束>

# <翻译开始>
// Any注册了一个匹配所有HTTP方法的路由
// Get, post, put, patch, head, options, delete, connect, trace
# <翻译结束>


<原文开始>
// Match registers a route that matches the specified methods that you declared.
<原文结束>

# <翻译开始>
// Match注册一个与你声明的指定方法匹配的路由
# <翻译结束>


<原文开始>
// StaticFile registers a single route in order to serve a single file of the local filesystem.
// router.StaticFile("favicon.ico", "./resources/favicon.ico")
<原文结束>

# <翻译开始>
// 为了服务本地文件系统的单个文件，StaticFile注册单个路由
// router.StaticFile(“favicon.ico&quot“
// /资源/ favicon.ico")
# <翻译结束>


<原文开始>
// StaticFileFS works just like `StaticFile` but a custom `http.FileSystem` can be used instead..
// router.StaticFileFS("favicon.ico", "./resources/favicon.ico", Dir{".", false})
// Gin by default uses: gin.Dir()
<原文结束>

# <翻译开始>
// StaticFileFS的工作原理就像' StaticFile '，但一个自定义的' http
// 可以用FileSystem代替
// router. staticfiles ("favicon.ico"， " /resources/favicon.ico"， Dir{"."， false}) Gin默认使用:Gin .Dir()
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
// 静态从给定的文件系统根目录提供文件
// 内部http
// 使用的是FileServer，因此是http
// 使用NotFound来代替路由器的NotFound处理程序
// 要使用操作系统的文件系统实现，使用:router.Static("/static"， "/var/www")
# <翻译结束>


<原文开始>
// StaticFS works just like `Static()` but a custom `http.FileSystem` can be used instead.
// Gin by default uses: gin.Dir()
<原文结束>

# <翻译开始>
// StaticFS的工作原理就像' Static() '，但一个自定义的' http
// 可以使用FileSystem’代替
// Gin默认使用:Gin . dir ()
# <翻译结束>


<原文开始>
	// Register GET and HEAD handlers
<原文结束>

# <翻译开始>
// 注册GET和HEAD处理程序
# <翻译结束>


<原文开始>
		// Check if file exists and/or if we have permission to access it
<原文结束>

# <翻译开始>
// 检查文件是否存在和/或我们是否有访问它的权限
# <翻译结束>


<原文开始>
			// Reset index
<原文结束>

# <翻译开始>
// 重置指数
# <翻译结束>

