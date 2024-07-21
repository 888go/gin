
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
// HandlerFunc defines the handler used by gin middleware as return value.
<原文结束>

# <翻译开始>
// HandlerFunc 定义了 Gin 中间件用作返回值的处理器。
# <翻译结束>


<原文开始>
// HandlersChain defines a HandlerFunc slice.
<原文结束>

# <翻译开始>
// HandlersChain 定义了一个 HandlerFunc 切片。
# <翻译结束>


<原文开始>
// Last returns the last handler in the chain. i.e. the last handler is the main one.
<原文结束>

# <翻译开始>
// Last 返回链中的最后一个处理器。即，最后一个处理器是主处理器。
# <翻译结束>


<原文开始>
// RouteInfo represents a request route's specification which contains method and path and its handler.
<原文结束>

# <翻译开始>
// RouteInfo 表示请求路由的规范，其中包含方法、路径及其处理程序。
# <翻译结束>


<原文开始>
// RoutesInfo defines a RouteInfo slice.
<原文结束>

# <翻译开始>
// RoutesInfo 定义了一个 RouteInfo 切片。
# <翻译结束>


<原文开始>
	// PlatformGoogleAppEngine when running on Google App Engine. Trust X-Appengine-Remote-Addr
	// for determining the client's IP
<原文结束>

# <翻译开始>
	// PlatformGoogleAppEngine：当在 Google App Engine 上运行时。信任 X-Appengine-Remote-Addr 头部来确定客户端的 IP 地址
# <翻译结束>


<原文开始>
	// PlatformCloudflare when using Cloudflare's CDN. Trust CF-Connecting-IP for determining
	// the client's IP
<原文结束>

# <翻译开始>
	// PlatformCloudflare 当使用Cloudflare的CDN时。信任CF-Connecting-IP来确定
	// 客户端的IP地址
# <翻译结束>


<原文开始>
// Engine is the framework's instance, it contains the muxer, middleware and configuration settings.
// Create an instance of Engine, by using New() or Default()
<原文结束>

# <翻译开始>
// Engine 是框架的实例，它包含了多路复用器（muxer）、中间件和配置设置。
// 通过使用 New() 或 Default() 创建 Engine 的一个实例。
# <翻译结束>


<原文开始>
	// RedirectTrailingSlash enables automatic redirection if the current route can't be matched but a
	// handler for the path with (without) the trailing slash exists.
	// For example if /foo/ is requested but a route only exists for /foo, the
	// client is redirected to /foo with http status code 301 for GET requests
	// and 307 for all other request methods.
<原文结束>

# <翻译开始>
	// RedirectTrailingSlash 功能会自动重定向，当当前路由无法匹配，但存在一个与请求路径（有或无尾部斜杠）相匹配的处理程序时。
	// 例如，如果请求了 /foo/，但仅存在 /foo 的路由，则客户端将被重定向到 /foo，并根据请求方法的不同返回不同的HTTP状态码：对于GET请求返回301，对于所有其他请求方法返回307。
# <翻译结束>


<原文开始>
	// RedirectFixedPath if enabled, the router tries to fix the current request path, if no
	// handle is registered for it.
	// First superfluous path elements like ../ or // are removed.
	// Afterwards the router does a case-insensitive lookup of the cleaned path.
	// If a handle can be found for this route, the router makes a redirection
	// to the corrected path with status code 301 for GET requests and 307 for
	// all other request methods.
	// For example /FOO and /..//Foo could be redirected to /foo.
	// RedirectTrailingSlash is independent of this option.
<原文结束>

# <翻译开始>
	// RedirectFixedPath：如果启用，当没有为当前请求路径注册处理程序时，路由器尝试修复该路径。
	// 首先移除诸如 ../ 或 	// 等多余的路径元素。
	// 然后，路由器对清理后的路径进行不区分大小写的查找。
	// 如果能找到与此路由匹配的处理程序，路由器将根据请求方法进行重定向：
	// 对于 GET 请求，状态码为 301；对于所有其他请求方法，状态码为 307。
	// 例如，/FOO 和 /..	//Foo 可能会被重定向到 /foo。
	// 该选项与 RedirectTrailingSlash 选项独立。
# <翻译结束>


<原文开始>
	// HandleMethodNotAllowed if enabled, the router checks if another method is allowed for the
	// current route, if the current request can not be routed.
	// If this is the case, the request is answered with 'Method Not Allowed'
	// and HTTP status code 405.
	// If no other Method is allowed, the request is delegated to the NotFound
	// handler.
<原文结束>

# <翻译开始>
	// 如果启用HandleMethodNotAllowed，当当前请求无法被路由时，路由器会检查当前路由是否允许其他方法。
	// 如果存在其他允许的方法，请求将得到响应'方法不允许'（Method Not Allowed）以及HTTP状态码405。
	// 若没有其他方法被允许，则该请求会被转发至NotFound处理器进行处理。
# <翻译结束>


<原文开始>
	// ForwardedByClientIP if enabled, client IP will be parsed from the request's headers that
	// match those stored at `(*gin.Engine).RemoteIPHeaders`. If no IP was
	// fetched, it falls back to the IP obtained from
	// `(*gin.Context).Request.RemoteAddr`.
<原文结束>

# <翻译开始>
	// ForwardedByClientIP：如果启用，将会从请求头中解析客户端IP地址，这些请求头与存储在 `(*gin.Engine).RemoteIPHeaders` 中的相匹配。如果没有获取到IP地址，则会回退到通过 `(*gin.Context).Request.RemoteAddr` 获取的IP地址。
# <翻译结束>


<原文开始>
	// AppEngine was deprecated.
	// Deprecated: USE `TrustedPlatform` WITH VALUE `gin.PlatformGoogleAppEngine` INSTEAD
	// #726 #755 If enabled, it will trust some headers starting with
	// 'X-AppEngine...' for better integration with that PaaS.
<原文结束>

# <翻译开始>
	// AppEngine 已被弃用。
	// 废弃: 请改用 `TrustedPlatform`，并设置其值为 `gin.PlatformGoogleAppEngine`
	// #726 #755 如果启用，将会信任以 'X-AppEngine...' 开头的一些头部信息，
	// 以便更好地与该 PaaS（平台即服务）进行集成。
# <翻译结束>


<原文开始>
// UseRawPath if enabled, the url.RawPath will be used to find parameters.
<原文结束>

# <翻译开始>
// 如果启用UseRawPath，将使用url.RawPath来查找参数。
# <翻译结束>


<原文开始>
	// UnescapePathValues if true, the path value will be unescaped.
	// If UseRawPath is false (by default), the UnescapePathValues effectively is true,
	// as url.Path gonna be used, which is already unescaped.
<原文结束>

# <翻译开始>
	// UnescapePathValues 如果设为 true，路径值将被解码。
	// 若 UseRawPath 为 false（默认情况），则 UnescapePathValues 实际上等同于 true，
	// 因为此时会使用已经解码过的 url.Path。
# <翻译结束>


<原文开始>
	// RemoveExtraSlash a parameter can be parsed from the URL even with extra slashes.
	// See the PR #1817 and issue #1644
<原文结束>

# <翻译开始>
	// RemoveExtraSlash：即使存在额外的斜杠，参数也可以从URL中解析出来。
	// 参见PR #1817和问题#1644
# <翻译结束>


<原文开始>
	// RemoteIPHeaders list of headers used to obtain the client IP when
	// `(*gin.Engine).ForwardedByClientIP` is `true` and
	// `(*gin.Context).Request.RemoteAddr` is matched by at least one of the
	// network origins of list defined by `(*gin.Engine).SetTrustedProxies()`.
<原文结束>

# <翻译开始>
	// RemoteIPHeaders 是一个头部列表，当 `(*gin.Engine).ForwardedByClientIP` 设置为 `true` 时，
	// 如果 `(*gin.Context).Request.RemoteAddr` 与通过 `(*gin.Engine).SetTrustedProxies()` 方法定义的网络源列表中的至少一个匹配，
	// 则会使用这些头部来获取客户端 IP 地址。
# <翻译结束>


<原文开始>
	// TrustedPlatform if set to a constant of value gin.Platform*, trusts the headers set by
	// that platform, for example to determine the client IP
<原文结束>

# <翻译开始>
	// TrustedPlatform 如果设置为gin.Platform*类型的常量值，表示信任该平台设置的头部信息，
	// 例如用于确定客户端IP地址
# <翻译结束>


<原文开始>
	// MaxMultipartMemory value of 'maxMemory' param that is given to http.Request's ParseMultipartForm
	// method call.
<原文结束>

# <翻译开始>
	// MaxMultipartMemory 是提供给 http.Request 的 ParseMultipartForm 方法调用时的 'maxMemory' 参数的值。
# <翻译结束>


<原文开始>
// UseH2C enable h2c support.
<原文结束>

# <翻译开始>
// UseH2C 启用 h2c 支持。
# <翻译结束>


<原文开始>
// ContextWithFallback enable fallback Context.Deadline(), Context.Done(), Context.Err() and Context.Value() when Context.Request.Context() is not nil.
<原文结束>

# <翻译开始>
// ContextWithFallback 用于当 Context.Request.Context() 不为空时，启用备用的 Context.Deadline()，Context.Done()，Context.Err() 和 Context.Value() 方法。
# <翻译结束>


<原文开始>
// New returns a new blank Engine instance without any middleware attached.
// By default, the configuration is:
// - RedirectTrailingSlash:  true
// - RedirectFixedPath:      false
// - HandleMethodNotAllowed: false
// - ForwardedByClientIP:    true
// - UseRawPath:             false
// - UnescapePathValues:     true
<原文结束>

# <翻译开始>
// New 函数返回一个全新的、未附加任何中间件的空白 Engine 实例。
// 默认配置为：
// - RedirectTrailingSlash:  true （自动重定向末尾的斜杠）
// - RedirectFixedPath:      false （不进行固定路径重定向）
// - HandleMethodNotAllowed: false （不处理不允许的方法）
// - ForwardedByClientIP:    true （通过客户端 IP 转发请求头）
// - UseRawPath:             false （不使用原始路径，即不做 URL 解码）
// - UnescapePathValues:     true （对路径中的参数值进行解码）
# <翻译结束>


<原文开始>
// Default returns an Engine instance with the Logger and Recovery middleware already attached.
<原文结束>

# <翻译开始>
// Default 返回一个已附加了 Logger 和 Recovery 中间件的 Engine 实例。
# <翻译结束>


<原文开始>
// Delims sets template left and right delims and returns an Engine instance.
<原文结束>

# <翻译开始>
// Delims 设置模板左右分隔符，并返回一个 Engine 实例。
# <翻译结束>


<原文开始>
// SecureJsonPrefix sets the secureJSONPrefix used in Context.SecureJSON.
<原文结束>

# <翻译开始>
// SecureJsonPrefix 设置在 Context.SecureJSON 中使用的 secureJSONPrefix。
# <翻译结束>


<原文开始>
// LoadHTMLGlob loads HTML files identified by glob pattern
// and associates the result with HTML renderer.
<原文结束>

# <翻译开始>
// LoadHTMLGlob 通过 glob 模式加载 HTML 文件，并将结果与 HTML 渲染器关联。
# <翻译结束>


<原文开始>
// LoadHTMLFiles loads a slice of HTML files
// and associates the result with HTML renderer.
<原文结束>

# <翻译开始>
// LoadHTMLFiles 加载一组 HTML 文件
// 并将结果与 HTML 渲染器关联。
# <翻译结束>


<原文开始>
// SetHTMLTemplate associate a template with HTML renderer.
<原文结束>

# <翻译开始>
// SetHTMLTemplate 将一个模板与HTML渲染器关联。
# <翻译结束>


<原文开始>
// SetFuncMap sets the FuncMap used for template.FuncMap.
<原文结束>

# <翻译开始>
// SetFuncMap 用于设置用于 template.FuncMap 的 FuncMap。
# <翻译结束>


<原文开始>
// NoRoute adds handlers for NoRoute. It returns a 404 code by default.
<原文结束>

# <翻译开始>
// NoRoute 添加处理函数，用于未找到路由的情况（NoRoute）。默认情况下返回404状态码。
# <翻译结束>


<原文开始>
// NoMethod sets the handlers called when Engine.HandleMethodNotAllowed = true.
<原文结束>

# <翻译开始>
// NoMethod 设置在 Engine.HandleMethodNotAllowed = true 时调用的处理器。
# <翻译结束>


<原文开始>
// Use attaches a global middleware to the router. i.e. the middleware attached through Use() will be
// included in the handlers chain for every single request. Even 404, 405, static files...
// For example, this is the right place for a logger or error management middleware.
<原文结束>

# <翻译开始>
// Use 方法将一个全局中间件附加到路由。也就是说，通过Use()方法附加的中间件将会
// 被包含在每一个请求的处理器链中。即便是404、405错误、静态文件等请求...
// 例如，这里适合放置日志记录器或错误管理中间件。
# <翻译结束>


<原文开始>
// Routes returns a slice of registered routes, including some useful information, such as:
// the http method, path and the handler name.
<原文结束>

# <翻译开始>
// Routes 返回已注册路由的切片，其中包括一些有用的信息，比如：
// HTTP 方法、路径以及处理器名称。
# <翻译结束>


<原文开始>
// Run attaches the router to a http.Server and starts listening and serving HTTP requests.
// It is a shortcut for http.ListenAndServe(addr, router)
// Note: this method will block the calling goroutine indefinitely unless an error happens.
<原文结束>

# <翻译开始>
// Run 将路由器附加到 http.Server，并开始监听和处理 HTTP 请求。
// 这是 http.ListenAndServe(addr, router) 的一个快捷方式。
// 注意：除非发生错误，否则此方法将无限期地阻塞调用的 goroutine。
# <翻译结束>


<原文开始>
// SetTrustedProxies set a list of network origins (IPv4 addresses,
// IPv4 CIDRs, IPv6 addresses or IPv6 CIDRs) from which to trust
// request's headers that contain alternative client IP when
// `(*gin.Engine).ForwardedByClientIP` is `true`. `TrustedProxies`
// feature is enabled by default, and it also trusts all proxies
// by default. If you want to disable this feature, use
// Engine.SetTrustedProxies(nil), then Context.ClientIP() will
// return the remote address directly.
<原文结束>

# <翻译开始>
// SetTrustedProxies 设置一个网络源列表（IPv4地址、IPv4 CIDR、IPv6地址或IPv6 CIDR），从这些源中信任请求头中包含的替代客户端IP。当`(*gin.Engine).ForwardedByClientIP`为`true`时生效。`TrustedProxies`特性默认启用，并且默认情况下信任所有代理。如果你想禁用此功能，使用Engine.SetTrustedProxies(nil)，那么Context.ClientIP()将直接返回远程地址。
# <翻译结束>


<原文开始>
// isUnsafeTrustedProxies checks if Engine.trustedCIDRs contains all IPs, it's not safe if it has (returns true)
<原文结束>

# <翻译开始>
// isUnsafeTrustedProxies 检查 Engine.trustedCIDRs 是否包含全部IP，如果包含（返回 true），则表示不安全
# <翻译结束>


<原文开始>
// parseTrustedProxies parse Engine.trustedProxies to Engine.trustedCIDRs
<原文结束>

# <翻译开始>
// parseTrustedProxies 将 Engine.trustedProxies 解析为 Engine.trustedCIDRs
# <翻译结束>


<原文开始>
// isTrustedProxy will check whether the IP address is included in the trusted list according to Engine.trustedCIDRs
<原文结束>

# <翻译开始>
// isTrustedProxy 将根据 Engine.trustedCIDRs 检查 IP 地址是否在信任列表中
# <翻译结束>


<原文开始>
// validateHeader will parse X-Forwarded-For header and return the trusted client IP address
<原文结束>

# <翻译开始>
// validateHeader 将解析 X-Forwarded-For 头部，并返回可信的客户端 IP 地址
# <翻译结束>


<原文开始>
		// X-Forwarded-For is appended by proxy
		// Check IPs in reverse order and stop when find untrusted proxy
<原文结束>

# <翻译开始>
		// X-Forwarded-For 由代理服务器追加
		// 按照逆序检查 IP 地址，并在找到不可信的代理时停止
# <翻译结束>


<原文开始>
// parseIP parse a string representation of an IP and returns a net.IP with the
// minimum byte representation or nil if input is invalid.
<原文结束>

# <翻译开始>
// parseIP 将IP地址的字符串表示形式解析为 net.IP 类型，并返回一个字节表示形式最小的 IP，如果输入无效，则返回 nil。
# <翻译结束>


<原文开始>
// return ip in a 4-byte representation
<原文结束>

# <翻译开始>
// 返回一个4字节表示的IP地址
# <翻译结束>


<原文开始>
// return ip in a 16-byte representation or nil
<原文结束>

# <翻译开始>
// 返回一个16字节表示形式的IP地址，或返回nil
# <翻译结束>


<原文开始>
// RunTLS attaches the router to a http.Server and starts listening and serving HTTPS (secure) requests.
// It is a shortcut for http.ListenAndServeTLS(addr, certFile, keyFile, router)
// Note: this method will block the calling goroutine indefinitely unless an error happens.
<原文结束>

# <翻译开始>
// RunTLS 将路由器附加到 http.Server，并开始监听和处理 HTTPS（安全）请求。
// 这是 http.ListenAndServeTLS(addr, certFile, keyFile, router) 的快捷方式。
// 注意：除非发生错误，否则此方法将无限期地阻塞调用的goroutine。
# <翻译结束>


<原文开始>
// RunUnix attaches the router to a http.Server and starts listening and serving HTTP requests
// through the specified unix socket (i.e. a file).
// Note: this method will block the calling goroutine indefinitely unless an error happens.
<原文结束>

# <翻译开始>
// RunUnix将路由器连接到http.Server，并开始通过指定的UNIX套接字（即文件）监听和处理HTTP请求。
// 注意：除非发生错误，否则此方法将无限期地阻塞调用的goroutine。
# <翻译结束>


<原文开始>
// RunFd attaches the router to a http.Server and starts listening and serving HTTP requests
// through the specified file descriptor.
// Note: this method will block the calling goroutine indefinitely unless an error happens.
<原文结束>

# <翻译开始>
// RunFd 将路由器连接到 http.Server，并开始通过指定的文件描述符监听和处理 HTTP 请求。
// 注意：除非发生错误，否则此方法将无限期地阻塞调用它的 goroutine。
# <翻译结束>


<原文开始>
// RunListener attaches the router to a http.Server and starts listening and serving HTTP requests
// through the specified net.Listener
<原文结束>

# <翻译开始>
// RunListener 将路由器附加到 http.Server，并开始通过指定的 net.Listener 监听和处理 HTTP 请求
# <翻译结束>


<原文开始>
// ServeHTTP conforms to the http.Handler interface.
<原文结束>

# <翻译开始>
// ServeHTTP 符合 http.Handler 接口。
//
// 注意!!! 此方法不能翻译, 因为是http包的接口实现
# <翻译结束>


<原文开始>
// HandleContext re-enters a context that has been rewritten.
// This can be done by setting c.Request.URL.Path to your new target.
// Disclaimer: You can loop yourself to deal with this, use wisely.
<原文结束>

# <翻译开始>
// HandleContext 该方法会重新载入一个被重写的context(可以通过c.Request.URL.Path来实现).
//
// 注意:该方法可能造成context的循环使用(会绕死你,谨慎使用)
# <翻译结束>


<原文开始>
// Find root of the tree for the given HTTP method
<原文结束>

# <翻译开始>
// 为给定的HTTP方法查找树的根节点
# <翻译结束>


<原文开始>
// Permanent redirect, request with GET method
<原文结束>

# <翻译开始>
// 永久重定向，使用GET方法请求
# <翻译结束>

