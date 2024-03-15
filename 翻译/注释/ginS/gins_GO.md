
<原文开始>
// Copyright 2014 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2014 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// LoadHTMLGlob is a wrapper for Engine.LoadHTMLGlob.
<原文结束>

# <翻译开始>
// LoadHTMLGlob is a wrapper for Engine.LoadHTMLGlob.
# <翻译结束>


<原文开始>
// LoadHTMLFiles is a wrapper for Engine.LoadHTMLFiles.
<原文结束>

# <翻译开始>
// LoadHTMLFiles is a wrapper for Engine.LoadHTMLFiles.
# <翻译结束>


<原文开始>
// SetHTMLTemplate is a wrapper for Engine.SetHTMLTemplate.
<原文结束>

# <翻译开始>
// SetHTMLTemplate is a wrapper for Engine.SetHTMLTemplate.
# <翻译结束>


<原文开始>
// NoRoute adds handlers for NoRoute. It returns a 404 code by default.
<原文结束>

# <翻译开始>
// NoRoute adds handlers for NoRoute. It returns a 404 code by default.
# <翻译结束>


<原文开始>
// NoMethod is a wrapper for Engine.NoMethod.
<原文结束>

# <翻译开始>
// NoMethod is a wrapper for Engine.NoMethod.
# <翻译结束>


<原文开始>
// Group creates a new router group. You should add all the routes that have common middlewares or the same path prefix.
// For example, all the routes that use a common middleware for authorization could be grouped.
<原文结束>

# <翻译开始>
// Group creates a new router group. You should add all the routes that have common middlewares or the same path prefix.
// For example, all the routes that use a common middleware for authorization could be grouped.
# <翻译结束>


<原文开始>
// Handle is a wrapper for Engine.Handle.
<原文结束>

# <翻译开始>
// Handle is a wrapper for Engine.Handle.
# <翻译结束>


<原文开始>
// POST is a shortcut for router.Handle("POST", path, handle)
<原文结束>

# <翻译开始>
// POST is a shortcut for router.Handle("POST", path, handle)
# <翻译结束>


<原文开始>
// GET is a shortcut for router.Handle("GET", path, handle)
<原文结束>

# <翻译开始>
// GET is a shortcut for router.Handle("GET", path, handle)
# <翻译结束>


<原文开始>
// DELETE is a shortcut for router.Handle("DELETE", path, handle)
<原文结束>

# <翻译开始>
// DELETE is a shortcut for router.Handle("DELETE", path, handle)
# <翻译结束>


<原文开始>
// PATCH is a shortcut for router.Handle("PATCH", path, handle)
<原文结束>

# <翻译开始>
// PATCH is a shortcut for router.Handle("PATCH", path, handle)
# <翻译结束>


<原文开始>
// PUT is a shortcut for router.Handle("PUT", path, handle)
<原文结束>

# <翻译开始>
// PUT is a shortcut for router.Handle("PUT", path, handle)
# <翻译结束>


<原文开始>
// OPTIONS is a shortcut for router.Handle("OPTIONS", path, handle)
<原文结束>

# <翻译开始>
// OPTIONS is a shortcut for router.Handle("OPTIONS", path, handle)
# <翻译结束>


<原文开始>
// HEAD is a shortcut for router.Handle("HEAD", path, handle)
<原文结束>

# <翻译开始>
// HEAD is a shortcut for router.Handle("HEAD", path, handle)
# <翻译结束>


<原文开始>
// Any is a wrapper for Engine.Any.
<原文结束>

# <翻译开始>
// Any is a wrapper for Engine.Any.
# <翻译结束>


<原文开始>
// StaticFile is a wrapper for Engine.StaticFile.
<原文结束>

# <翻译开始>
// StaticFile is a wrapper for Engine.StaticFile.
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
// Static serves files from the given file system root.
// Internally a http.FileServer is used, therefore http.NotFound is used instead
// of the Router's NotFound handler.
// To use the operating system's file system implementation,
// use :
//
//	router.Static("/static", "/var/www")
# <翻译结束>


<原文开始>
// StaticFS is a wrapper for Engine.StaticFS.
<原文结束>

# <翻译开始>
// StaticFS is a wrapper for Engine.StaticFS.
# <翻译结束>


<原文开始>
// Use attaches a global middleware to the router. i.e. the middlewares attached through Use() will be
// included in the handlers chain for every single request. Even 404, 405, static files...
// For example, this is the right place for a logger or error management middleware.
<原文结束>

# <翻译开始>
// Use attaches a global middleware to the router. i.e. the middlewares attached through Use() will be
// included in the handlers chain for every single request. Even 404, 405, static files...
// For example, this is the right place for a logger or error management middleware.
# <翻译结束>


<原文开始>
// Routes returns a slice of registered routes.
<原文结束>

# <翻译开始>
// Routes returns a slice of registered routes.
# <翻译结束>


<原文开始>
// Run attaches to a http.Server and starts listening and serving HTTP requests.
// It is a shortcut for http.ListenAndServe(addr, router)
// Note: this method will block the calling goroutine indefinitely unless an error happens.
<原文结束>

# <翻译开始>
// Run attaches to a http.Server and starts listening and serving HTTP requests.
// It is a shortcut for http.ListenAndServe(addr, router)
// Note: this method will block the calling goroutine indefinitely unless an error happens.
# <翻译结束>


<原文开始>
// RunTLS attaches to a http.Server and starts listening and serving HTTPS requests.
// It is a shortcut for http.ListenAndServeTLS(addr, certFile, keyFile, router)
// Note: this method will block the calling goroutine indefinitely unless an error happens.
<原文结束>

# <翻译开始>
// RunTLS attaches to a http.Server and starts listening and serving HTTPS requests.
// It is a shortcut for http.ListenAndServeTLS(addr, certFile, keyFile, router)
// Note: this method will block the calling goroutine indefinitely unless an error happens.
# <翻译结束>


<原文开始>
// RunUnix attaches to a http.Server and starts listening and serving HTTP requests
// through the specified unix socket (i.e. a file)
// Note: this method will block the calling goroutine indefinitely unless an error happens.
<原文结束>

# <翻译开始>
// RunUnix attaches to a http.Server and starts listening and serving HTTP requests
// through the specified unix socket (i.e. a file)
// Note: this method will block the calling goroutine indefinitely unless an error happens.
# <翻译结束>


<原文开始>
// RunFd attaches the router to a http.Server and starts listening and serving HTTP requests
// through the specified file descriptor.
// Note: the method will block the calling goroutine indefinitely unless on error happens.
<原文结束>

# <翻译开始>
// RunFd attaches the router to a http.Server and starts listening and serving HTTP requests
// through the specified file descriptor.
// Note: the method will block the calling goroutine indefinitely unless on error happens.
# <翻译结束>

