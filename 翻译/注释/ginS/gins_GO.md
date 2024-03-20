
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
// LoadHTMLGlob is a wrapper for Engine.LoadHTMLGlob.
<原文结束>

# <翻译开始>
// LoadHTMLGlob 是 Engine.LoadHTMLGlob 的一个包装函数。
# <翻译结束>


<原文开始>
// LoadHTMLFiles is a wrapper for Engine.LoadHTMLFiles.
<原文结束>

# <翻译开始>
// LoadHTMLFiles 是对 Engine.LoadHTMLFiles 的一个封装。
# <翻译结束>


<原文开始>
// SetHTMLTemplate is a wrapper for Engine.SetHTMLTemplate.
<原文结束>

# <翻译开始>
// SetHTMLTemplate 是 Engine.SetHTMLTemplate 的一个包装函数。
# <翻译结束>


<原文开始>
// NoRoute adds handlers for NoRoute. It returns a 404 code by default.
<原文结束>

# <翻译开始>
// NoRoute 添加处理函数，用于未找到路由的情况（NoRoute）。默认情况下返回404状态码。
# <翻译结束>


<原文开始>
// NoMethod is a wrapper for Engine.NoMethod.
<原文结束>

# <翻译开始>
// NoMethod 是 Engine.NoMethod 的一个包装器。
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
// Handle is a wrapper for Engine.Handle.
<原文结束>

# <翻译开始>
// Handle 是 Engine.Handle 的一个包装器。
# <翻译结束>


<原文开始>
// POST is a shortcut for router.Handle("POST", path, handle)
<原文结束>

# <翻译开始>
// POST 是一个快捷方式，用于 router.Handle("POST", path, handle)
# <翻译结束>


<原文开始>
// GET is a shortcut for router.Handle("GET", path, handle)
<原文结束>

# <翻译开始>
// GET 是一个快捷方式，等同于 router.Handle("GET", path, handle)
# <翻译结束>


<原文开始>
// DELETE is a shortcut for router.Handle("DELETE", path, handle)
<原文结束>

# <翻译开始>
// DELETE 是一个快捷方式，等同于 router.Handle("DELETE", path, handle)
# <翻译结束>


<原文开始>
// PATCH is a shortcut for router.Handle("PATCH", path, handle)
<原文结束>

# <翻译开始>
// PATCH 是一个快捷方式，用于 router.Handle("PATCH", path, handle)
# <翻译结束>


<原文开始>
// PUT is a shortcut for router.Handle("PUT", path, handle)
<原文结束>

# <翻译开始>
// PUT 是一个快捷方式，等同于 router.Handle("PUT", path, handle)
# <翻译结束>


<原文开始>
// OPTIONS is a shortcut for router.Handle("OPTIONS", path, handle)
<原文结束>

# <翻译开始>
// OPTIONS 是一个快捷方式，用于 router.Handle("OPTIONS", path, handle)
# <翻译结束>


<原文开始>
// HEAD is a shortcut for router.Handle("HEAD", path, handle)
<原文结束>

# <翻译开始>
// HEAD 是一个快捷方式，用于 router.Handle("HEAD", path, handle)
# <翻译结束>


<原文开始>
// Any is a wrapper for Engine.Any.
<原文结束>

# <翻译开始>
// Any 是 Engine.Any 的一个包装器。
# <翻译结束>


<原文开始>
// StaticFile is a wrapper for Engine.StaticFile.
<原文结束>

# <翻译开始>
// StaticFile 是 Engine.StaticFile 的一个包装器。
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
// 如:
// r.Static("/static", "./文件夹") //当你访问http://localhost:8080/static时，它会服务于./文件夹 目录下的文件
//
// 在内部使用了 http.FileServer，因此会使用 http.NotFound 替代 Router 的 NotFound 处理程序。
// 要使用操作系统自身的文件系统实现，请按如下方式使用：
// router.Static("/static", "/var/www")
# <翻译结束>


<原文开始>
// StaticFS is a wrapper for Engine.StaticFS.
<原文结束>

# <翻译开始>
// StaticFS 是 Engine.StaticFS 的一个包装器。
# <翻译结束>


<原文开始>
// Use attaches a global middleware to the router. i.e. the middlewares attached through Use() will be
// included in the handlers chain for every single request. Even 404, 405, static files...
// For example, this is the right place for a logger or error management middleware.
<原文结束>

# <翻译开始>
// Use 方法将全局中间件附加到路由器。即通过 Use() 附加的中间件将会
// 包含在每一个请求的处理器链中。即使是 404、405 状态码的响应，或者是静态文件的处理...
// 例如，这里适合放置日志记录器或错误管理中间件。
# <翻译结束>


<原文开始>
// Routes returns a slice of registered routes.
<原文结束>

# <翻译开始>
// Routes 返回已注册路由的切片。
# <翻译结束>


<原文开始>
// Run attaches to a http.Server and starts listening and serving HTTP requests.
// It is a shortcut for http.ListenAndServe(addr, router)
// Note: this method will block the calling goroutine indefinitely unless an error happens.
<原文结束>

# <翻译开始>
// Run 函数连接到一个 http.Server，并开始监听和处理 HTTP 请求。
// 这是 http.ListenAndServe(addr, router) 的快捷方式。
// 注意：除非发生错误，否则此方法将无限期地阻塞调用它的 goroutine。
# <翻译结束>


<原文开始>
// RunTLS attaches to a http.Server and starts listening and serving HTTPS requests.
// It is a shortcut for http.ListenAndServeTLS(addr, certFile, keyFile, router)
// Note: this method will block the calling goroutine indefinitely unless an error happens.
<原文结束>

# <翻译开始>
// RunTLS 绑定到一个 http.Server，并开始监听和处理 HTTPS 请求。
// 这是 http.ListenAndServeTLS(addr, certFile, keyFile, router) 的快捷方式。
// 注意：除非发生错误，否则此方法将无限期地阻塞调用它的 goroutine。
# <翻译结束>


<原文开始>
// RunUnix attaches to a http.Server and starts listening and serving HTTP requests
// through the specified unix socket (i.e. a file)
// Note: this method will block the calling goroutine indefinitely unless an error happens.
<原文结束>

# <翻译开始>
// RunUnix连接到一个http.Server，并开始通过指定的unix套接字（即文件）监听和处理HTTP请求。
// 注意：除非发生错误，否则此方法将无限期地阻塞调用的goroutine。
# <翻译结束>


<原文开始>
// RunFd attaches the router to a http.Server and starts listening and serving HTTP requests
// through the specified file descriptor.
// Note: the method will block the calling goroutine indefinitely unless on error happens.
<原文结束>

# <翻译开始>
// RunFd 将路由器附加到 http.Server，并开始通过指定的文件描述符监听和处理 HTTP 请求。
// 注意：除非发生错误，否则该方法将无限期地阻塞调用它的 goroutine。
# <翻译结束>

