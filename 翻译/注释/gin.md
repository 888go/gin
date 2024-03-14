
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
// 32 MB
<原文结束>

# <翻译开始>
// 32 MB
# <翻译结束>


<原文开始>
// 0.0.0.0/0 (IPv4)
<原文结束>

# <翻译开始>
// 0.0.0.0/0 (IPv4)
# <翻译结束>


<原文开始>
// ::/0 (IPv6)
<原文结束>

# <翻译开始>
// /:: 0 (IPv6)
# <翻译结束>


<原文开始>
// HandlerFunc defines the handler used by gin middleware as return value.
<原文结束>

# <翻译开始>
// HandlerFunc定义了gin中间件使用的处理程序作为返回值
# <翻译结束>


<原文开始>
// HandlersChain defines a HandlerFunc slice.
<原文结束>

# <翻译开始>
// HandlersChain定义了一个handlerfuncc片
# <翻译结束>


<原文开始>
// Last returns the last handler in the chain. i.e. the last handler is the main one.
<原文结束>

# <翻译开始>
// Last返回链中的最后一个处理程序
// 也就是说，最后一个处理器是主处理器
# <翻译结束>


<原文开始>
// RouteInfo represents a request route's specification which contains method and path and its handler.
<原文结束>

# <翻译开始>
// RouteInfo表示一个请求路由的规范，它包含方法、路径和它的处理器
# <翻译结束>


<原文开始>
// RoutesInfo defines a RouteInfo slice.
<原文结束>

# <翻译开始>
// RoutesInfo定义了一个RouteInfo切片
# <翻译结束>


<原文开始>
// Trusted platforms
<原文结束>

# <翻译开始>
// 信任的平台
# <翻译结束>


<原文开始>
	// PlatformGoogleAppEngine when running on Google App Engine. Trust X-Appengine-Remote-Addr
	// for determining the client's IP
<原文结束>

# <翻译开始>
// 在Google应用引擎上运行时的平台googleappengine
// 信任X-Appengine-Remote-Addr来确定客户端的IP
# <翻译结束>


<原文开始>
	// PlatformCloudflare when using Cloudflare's CDN. Trust CF-Connecting-IP for determining
	// the client's IP
<原文结束>

# <翻译开始>
// 使用Cloudflare的CDN时的平台Cloudflare
// Trust CF-Connecting-IP用于确定客户端的IP
# <翻译结束>


<原文开始>
// Engine is the framework's instance, it contains the muxer, middleware and configuration settings.
// Create an instance of Engine, by using New() or Default()
<原文结束>

# <翻译开始>
// 引擎是框架的实例，它包含了复用器、中间件和配置设置
// 使用New()或Default()创建Engine实例
# <翻译结束>


<原文开始>
	// RedirectTrailingSlash enables automatic redirection if the current route can't be matched but a
	// handler for the path with (without) the trailing slash exists.
	// For example if /foo/ is requested but a route only exists for /foo, the
	// client is redirected to /foo with http status code 301 for GET requests
	// and 307 for all other request methods.
<原文结束>

# <翻译开始>
// RedirectTrailingSlash在当前路由不能匹配的情况下启用自动重定向，但是存在一个带有(不带有)尾斜杠的路径处理程序
// 例如，如果请求/foo/，但只存在/foo的路由，则客户端被重定向到/foo, GET请求的http状态码为301，所有其他请求方法的http状态码为307
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
// RedirectFixedPath如果启用，如果没有为它注册句柄，路由器会尝试修复当前的请求路径
// 首先是多余的路径元素，比如…/或被移除
// 之后，路由器会对清理后的路径进行不区分大小写的查找
// 如果能找到该路由的句柄，路由器就会重定向到正确的路径，GET请求的状态码为301，其他所有请求方法的状态码为307
// 例如/FOO和/..Foo可以重定向到/ Foo
// RedirectTrailingSlash与此选项无关
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
// handlemethodnotalallowed如果使能，如果当前请求不能被路由，则路由器检查当前路由是否允许另一个方法
// 如果是这种情况，请求将返回“方法不允许”和HTTP状态码405
// 如果不允许使用其他方法，则将请求委托给NotFound处理程序
# <翻译结束>


<原文开始>
	// ForwardedByClientIP if enabled, client IP will be parsed from the request's headers that
	// match those stored at `(*gin.Engine).RemoteIPHeaders`. If no IP was
	// fetched, it falls back to the IP obtained from
	// `(*gin.Context).Request.RemoteAddr`.
<原文结束>

# <翻译开始>
// 如果启用了ForwardedByClientIP，客户端IP将从与存储在' (*gin.Engine). remoteipheaders '匹配的请求头中解析
// 如果没有获取到IP，则返回到从' (*gin.Context). request . remoteaddr '获取的IP
# <翻译结束>


<原文开始>
	// AppEngine was deprecated.
	// Deprecated: USE `TrustedPlatform` WITH VALUE `gin.PlatformGoogleAppEngine` INSTEAD
	// #726 #755 If enabled, it will trust some headers starting with
	// 'X-AppEngine...' for better integration with that PaaS.
<原文结束>

# <翻译开始>
// AppEngine已弃用
// 已弃用:使用' TrustedPlatform ' WITH VALUE ' gin
// 如果启用，它将信任一些以“X-AppEngine…”开头的标头
// 以便与该PaaS更好地集成
# <翻译结束>


<原文开始>
	// UseRawPath if enabled, the url.RawPath will be used to find parameters.
<原文结束>

# <翻译开始>
// UseRawPath如果启用，则为url
// RawPath将用于查找参数
# <翻译结束>


<原文开始>
	// UnescapePathValues if true, the path value will be unescaped.
	// If UseRawPath is false (by default), the UnescapePathValues effectively is true,
	// as url.Path gonna be used, which is already unescaped.
<原文结束>

# <翻译开始>
// UnescapePathValues如果为true，则不转义路径值
// 如果UseRawPath为false(默认情况下)，UnescapePathValues有效地为true，如url
// 路径将被使用，它已经是未转义的
# <翻译结束>


<原文开始>
	// RemoveExtraSlash a parameter can be parsed from the URL even with extra slashes.
	// See the PR #1817 and issue #1644
<原文结束>

# <翻译开始>
// 即使使用额外的斜杠，也可以从URL解析RemoveExtraSlash参数
// 见PR #1817和issue #1644
# <翻译结束>


<原文开始>
	// RemoteIPHeaders list of headers used to obtain the client IP when
	// `(*gin.Engine).ForwardedByClientIP` is `true` and
	// `(*gin.Context).Request.RemoteAddr` is matched by at least one of the
	// network origins of list defined by `(*gin.Engine).SetTrustedProxies()`.
<原文结束>

# <翻译开始>
// RemoteIPHeaders获取客户端IP时使用的报头列表(*gin.Engine)
// ForwardedByClientIP '是' true '和' (*gin.Context). request
// RemoteAddr '被' (*gin.Engine). settrustedproxies() '定义的列表的至少一个网络源匹配
# <翻译结束>


<原文开始>
	// TrustedPlatform if set to a constant of value gin.Platform*, trusts the headers set by
	// that platform, for example to determine the client IP
<原文结束>

# <翻译开始>
// TrustedPlatform设置为一个值为gin的常量
// 例如，平台*信任由该平台设置的报头来确定客户端IP
# <翻译结束>


<原文开始>
	// MaxMultipartMemory value of 'maxMemory' param that is given to http.Request's ParseMultipartForm
	// method call.
<原文结束>

# <翻译开始>
// 给http的“maxMemory”参数的MaxMultipartMemory值请求的parsemmultipartform方法调用
# <翻译结束>


<原文开始>
	// UseH2C enable h2c support.
<原文结束>

# <翻译开始>
// 启用h2c支持
# <翻译结束>


<原文开始>
	// ContextWithFallback enable fallback Context.Deadline(), Context.Done(), Context.Err() and Context.Value() when Context.Request.Context() is not nil.
<原文结束>

# <翻译开始>
// 当Context.Request.Context()不是nil时，启用回退Context.Deadline()、Context.Done()、Context.Err()和Context.Value()
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
// New返回一个新的空白Engine实例，没有附加任何中间件
// 默认配置为:—RedirectTrailingSlash: true—RedirectFixedPath: false—handlemethodnotalallowed: false—ForwardedByClientIP: true—UseRawPath: false—UnescapePathValues: true
# <翻译结束>


<原文开始>
// Default returns an Engine instance with the Logger and Recovery middleware already attached.
<原文结束>

# <翻译开始>
// Default返回一个Engine实例，其中已经附加了Logger和Recovery中间件
# <翻译结束>


<原文开始>
// Delims sets template left and right delims and returns an Engine instance.
<原文结束>

# <翻译开始>
// Delims设置模板的左和右分隔符并返回Engine实例
# <翻译结束>


<原文开始>
// SecureJsonPrefix sets the secureJSONPrefix used in Context.SecureJSON.
<原文结束>

# <翻译开始>
// SecureJsonPrefix设置Context.SecureJSON中使用的SecureJsonPrefix
# <翻译结束>


<原文开始>
// LoadHTMLGlob loads HTML files identified by glob pattern
// and associates the result with HTML renderer.
<原文结束>

# <翻译开始>
// LoadHTMLGlob加载由glob模式标识的HTML文件，并将结果与HTML渲染器相关联
# <翻译结束>


<原文开始>
// LoadHTMLFiles loads a slice of HTML files
// and associates the result with HTML renderer.
<原文结束>

# <翻译开始>
// LoadHTMLFiles加载一段HTML文件，并将结果与HTML渲染器相关联
# <翻译结束>


<原文开始>
// SetHTMLTemplate associate a template with HTML renderer.
<原文结束>

# <翻译开始>
// SetHTMLTemplate将模板与HTML渲染器关联
# <翻译结束>


<原文开始>
// SetFuncMap sets the FuncMap used for template.FuncMap.
<原文结束>

# <翻译开始>
// SetFuncMap设置用于template.FuncMap的FuncMap
# <翻译结束>


<原文开始>
// NoRoute adds handlers for NoRoute. It returns a 404 code by default.
<原文结束>

# <翻译开始>
// NoRoute为NoRoute添加处理程序
// 默认情况下，它返回404代码
# <翻译结束>


<原文开始>
// NoMethod sets the handlers called when Engine.HandleMethodNotAllowed = true.
<原文结束>

# <翻译开始>
// NoMethod设置引擎时调用的处理程序
// handlemethodnotalallowed = true
# <翻译结束>


<原文开始>
// Use attaches a global middleware to the router. i.e. the middleware attached through Use() will be
// included in the handlers chain for every single request. Even 404, 405, static files...
// For example, this is the right place for a logger or error management middleware.
<原文结束>

# <翻译开始>
// Use将全局中间件附加到路由器上
// 也就是说，通过Use()附加的中间件将被包含在每个请求的处理程序链中
// 甚至404、405、静态文件……例如，这是日志记录器或错误管理中间件的正确位置
# <翻译结束>


<原文开始>
// Routes returns a slice of registered routes, including some useful information, such as:
// the http method, path and the handler name.
<原文结束>

# <翻译开始>
// Routes返回已注册路由的切片，其中包括一些有用的信息，例如:http方法、路径和处理程序名称
# <翻译结束>


<原文开始>
// Run attaches the router to a http.Server and starts listening and serving HTTP requests.
// It is a shortcut for http.ListenAndServe(addr, router)
// Note: this method will block the calling goroutine indefinitely unless an error happens.
<原文结束>

# <翻译开始>
// Run将路由器附加到http上
// 服务器并开始监听和服务HTTP请求
// 它是http的快捷方式
// 注意:除非发生错误，否则此方法将无限期地阻塞调用例程
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
// SetTrustedProxies设置了一个网络起源列表(IPv4地址，IPv4 cidr, IPv6地址或IPv6 cidr)，从其中信任请求的头包含替代客户端IP时' (* gin.com engine)
// ForwardedByClientIP '为' true '
// ' TrustedProxies '功能是默认启用的，它也默认信任所有代理
// 如果您想禁用此功能，请使用Engine.SetTrustedProxies(nil)，然后Context.ClientIP()将直接返回远程地址
# <翻译结束>


<原文开始>
// isUnsafeTrustedProxies checks if Engine.trustedCIDRs contains all IPs, it's not safe if it has (returns true)
<原文结束>

# <翻译开始>
// isUnsafeTrustedProxies检查引擎
// trustedCIDRs包含了所有的ip地址，如果有，则不安全(返回true)
# <翻译结束>


<原文开始>
// parseTrustedProxies parse Engine.trustedProxies to Engine.trustedCIDRs
<原文结束>

# <翻译开始>
// parseTrustedProxies解析引擎
// trustedproxy to engine . trustedidrs
# <翻译结束>


<原文开始>
// isTrustedProxy will check whether the IP address is included in the trusted list according to Engine.trustedCIDRs
<原文结束>

# <翻译开始>
// isTrustedProxy会根据Engine.trustedCIDRs检查IP地址是否在可信列表中
# <翻译结束>


<原文开始>
// validateHeader will parse X-Forwarded-For header and return the trusted client IP address
<原文结束>

# <翻译开始>
// validateHeader将解析X-Forwarded-For报头并返回受信任的客户端IP地址
# <翻译结束>


<原文开始>
		// X-Forwarded-For is appended by proxy
		// Check IPs in reverse order and stop when find untrusted proxy
<原文结束>

# <翻译开始>
// 以相反的顺序检查ip，当发现不受信任的代理时停止
# <翻译结束>


<原文开始>
// parseIP parse a string representation of an IP and returns a net.IP with the
// minimum byte representation or nil if input is invalid.
<原文结束>

# <翻译开始>
// 解析IP的字符串表示形式并返回一个net
// 具有最小字节表示的IP，如果输入无效则为nil
# <翻译结束>


<原文开始>
		// return ip in a 4-byte representation
<原文结束>

# <翻译开始>
// 返回4字节表示的IP
# <翻译结束>


<原文开始>
	// return ip in a 16-byte representation or nil
<原文结束>

# <翻译开始>
// 返回16字节表示形式的IP或nil
# <翻译结束>


<原文开始>
// RunTLS attaches the router to a http.Server and starts listening and serving HTTPS (secure) requests.
// It is a shortcut for http.ListenAndServeTLS(addr, certFile, keyFile, router)
// Note: this method will block the calling goroutine indefinitely unless an error happens.
<原文结束>

# <翻译开始>
// RunTLS将路由器附加到http
// 服务器并开始监听和服务HTTPS(安全)请求
// 它是http的快捷方式
// 注意:除非发生错误，否则此方法将无限期地阻塞调用例程
# <翻译结束>


<原文开始>
// RunUnix attaches the router to a http.Server and starts listening and serving HTTP requests
// through the specified unix socket (i.e. a file).
// Note: this method will block the calling goroutine indefinitely unless an error happens.
<原文结束>

# <翻译开始>
// RunUnix将路由器附加到http
// 服务器并通过指定的unix套接字(即文件)开始侦听和服务HTTP请求
// 注意:除非发生错误，否则此方法将无限期地阻塞调用例程
# <翻译结束>


<原文开始>
// RunFd attaches the router to a http.Server and starts listening and serving HTTP requests
// through the specified file descriptor.
// Note: this method will block the calling goroutine indefinitely unless an error happens.
<原文结束>

# <翻译开始>
// RunFd将路由器附加到http
// 服务器并通过指定的文件描述符开始侦听和服务HTTP请求
// 注意:除非发生错误，否则此方法将无限期地阻塞调用例程
# <翻译结束>


<原文开始>
// RunListener attaches the router to a http.Server and starts listening and serving HTTP requests
// through the specified net.Listener
<原文结束>

# <翻译开始>
// RunListener将路由器附加到http
// 服务器并开始通过指定的网络侦听和服务HTTP请求
// 侦听器
# <翻译结束>


<原文开始>
// ServeHTTP conforms to the http.Handler interface.
<原文结束>

# <翻译开始>
// ServeHTTP符合http
// 处理程序接口
# <翻译结束>


<原文开始>
// HandleContext re-enters a context that has been rewritten.
// This can be done by setting c.Request.URL.Path to your new target.
// Disclaimer: You can loop yourself to deal with this, use wisely.
<原文结束>

# <翻译开始>
// HandleContext重新进入一个已经重写的上下文
// 这可以通过将c.Request.URL.Path设置为新目标来实现
// 免责声明:你可以循环自己来处理这个问题，明智地使用
# <翻译结束>


<原文开始>
	// Find root of the tree for the given HTTP method
<原文结束>

# <翻译开始>
// 查找给定HTTP方法的树的根
# <翻译结束>


<原文开始>
		// Find route in tree
<原文结束>

# <翻译开始>
// 在树中查找路由
# <翻译结束>


<原文开始>
// Permanent redirect, request with GET method
<原文结束>

# <翻译开始>
// 永久重定向，请求使用GET方法
# <翻译结束>

