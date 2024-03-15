
<原文开始>
// Config is a struct for specifying configuration options for the secure.
<原文结束>

# <翻译开始>
// Config 是一个结构体，用于指定 secure 的配置选项。
# <翻译结束>


<原文开始>
	// AllowedHosts is a list of fully qualified domain names that are allowed.
	//Default is empty list, which allows any and all host names.
<原文结束>

# <翻译开始>
// AllowedHosts 是一个完全限定域名列表，其中包含允许的所有域名。
// 默认为空列表，这意味着允许任何和所有主机名。
# <翻译结束>


<原文开始>
	// If SSLRedirect is set to true, then only allow https requests.
	// Default is false.
<原文结束>

# <翻译开始>
// 如果SSLRedirect设置为true，则只允许https请求。
// 默认值为false。
# <翻译结束>


<原文开始>
	// If SSLTemporaryRedirect is true, the a 302 will be used while redirecting.
	// Default is false (301).
<原文结束>

# <翻译开始>
// 如果SSLTemporaryRedirect设为true，在重定向时将使用302状态码。
// 默认值为false（即使用301状态码）。
# <翻译结束>


<原文开始>
	// SSLHost is the host name that is used to redirect http requests to https.
	// Default is "", which indicates to use the same host.
<原文结束>

# <翻译开始>
// SSLHost 是用于将http请求重定向到https的主机名。
// 默认值为 ""，表示使用相同的主机。
# <翻译结束>


<原文开始>
	// STSSeconds is the max-age of the Strict-Transport-Security header.
	// Default is 0, which would NOT include the header.
<原文结束>

# <翻译开始>
// STSSeconds 是 Strict-Transport-Security 头部的最大有效期（max-age）。
// 默认值为 0，这意味着不会包含此头部。
# <翻译结束>


<原文开始>
	// If STSIncludeSubdomains is set to true, the `includeSubdomains` will
	// be appended to the Strict-Transport-Security header. Default is false.
<原文结束>

# <翻译开始>
// 如果将STSIncludeSubdomains设置为true，则`includeSubdomains`将被追加到Strict-Transport-Security头中。默认值为false。
# <翻译结束>


<原文开始>
	// If FrameDeny is set to true, adds the X-Frame-Options header with
	// the value of `DENY`. Default is false.
<原文结束>

# <翻译开始>
// 如果FrameDeny设置为true，则添加X-Frame-Options头，并将其值设为`DENY`。默认值为false。
# <翻译结束>


<原文开始>
	// CustomFrameOptionsValue allows the X-Frame-Options header value
	// to be set with a custom value. This overrides the FrameDeny option.
<原文结束>

# <翻译开始>
// CustomFrameOptionsValue 允许自定义设置 X-Frame-Options 头部值，这将覆盖 FrameDeny 选项。
# <翻译结束>


<原文开始>
	// If ContentTypeNosniff is true, adds the X-Content-Type-Options header
	// with the value `nosniff`. Default is false.
<原文结束>

# <翻译开始>
// 如果ContentTypeNosniff设为true，将会添加“X-Content-Type-Options”头部，并将其值设为`nosniff`。默认值为false。
# <翻译结束>


<原文开始>
	// If BrowserXssFilter is true, adds the X-XSS-Protection header with
	// the value `1; mode=block`. Default is false.
<原文结束>

# <翻译开始>
// 如果BrowserXssFilter为true，则添加X-XSS-Protection头，并将其值设为`1; mode=block`。默认值为false。
# <翻译结束>


<原文开始>
	// ContentSecurityPolicy allows the Content-Security-Policy header value
	// to be set with a custom value. Default is "".
<原文结束>

# <翻译开始>
// ContentSecurityPolicy 允许设置自定义的 Content-Security-Policy 头部值，默认为空字符串""。
# <翻译结束>


<原文开始>
	// HTTP header "Referrer-Policy" governs which referrer information, sent in the Referrer header, should be included with requests made.
<原文结束>

# <翻译开始>
// HTTP头部"Referrer-Policy"规定了在请求中随Referrer头部一起发送的referrer信息应包含哪些内容。
# <翻译结束>


<原文开始>
	// When true, the whole security policy applied by the middleware is disabled completely.
<原文结束>

# <翻译开始>
// 当设为true时，中间件应用的整个安全策略将被完全禁用。
# <翻译结束>


<原文开始>
	// Handlers for when an error occurs (ie bad host).
<原文结束>

# <翻译开始>
// 当发生错误（例如主机错误）时的处理程序。
# <翻译结束>


<原文开始>
	// Feature Policy is a new header that allows a site to control which features and APIs can be used in the browser.
<原文结束>

# <翻译开始>
// 功能策略是一个新的头部信息，它允许网站控制哪些功能和API可以在浏览器中使用。
# <翻译结束>


<原文开始>
	// If DontRedirectIPV4Hostnames is true, requests to hostnames that are IPV4
	// addresses aren't redirected. This is to allow load balancer health checks
	// to succeed.
<原文结束>

# <翻译开始>
// 如果DontRedirectIPV4Hostnames设为true，那么对IPv4地址形式的主机名的请求将不会被重定向。这是为了允许负载均衡器健康检查能够成功执行。
# <翻译结束>


<原文开始>
	// If the request is insecure, treat it as secure if any of the headers in this dict are set to their corresponding value
	// This is useful when your app is running behind a secure proxy that forwards requests to your app over http (such as on Heroku).
<原文结束>

# <翻译开始>
// 如果请求是非安全的，但该字典中任意一个头部被设置为它们对应的值时，则将请求视为安全请求。
// 这在你的应用运行在一个通过http转发请求到你应用的（如Heroku）安全代理后面时非常有用。
// 以下是逐句翻译：
// 如果请求是不安全的，
// 但如果这个字典中的任何头部被设置为了它们相应的值，
// 那么就将这个请求当作安全请求来处理。
// 这在你的应用程序运行在一个会通过http将请求转发到你的应用程序的（例如Heroku这样的）安全代理之后的情况下非常有帮助。
# <翻译结束>


<原文开始>
// DefaultConfig returns a Configuration with strict security settings.
// ```
//		SSLRedirect:           true
//		IsDevelopment:         false
//		STSSeconds:            315360000
//		STSIncludeSubdomains:  true
//		FrameDeny:             true
//		ContentTypeNosniff:    true
//		BrowserXssFilter:      true
//		ContentSecurityPolicy: "default-src 'self'"
//		SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
// ```
<原文结束>

# <翻译开始>
// DefaultConfig 返回一个具有严格安全设置的 Configuration。
// ```
//		SSLRedirect:           // 是否启用 HTTPS 重定向（默认为 true）
//		IsDevelopment:         // 是否处于开发模式（默认为 false）
//		STSSeconds:            // 设置 HTTP Strict Transport Security (HSTS) 的有效期，单位为秒（默认为 315360000，即十年）
//		STSIncludeSubdomains:  // 是否在 HSTS 策略中包含子域名（默认为 true）
//		FrameDeny:             // 是否禁止页面在 frame 或 iframe 中加载（默认为 true，防止点击劫持攻击）
//		ContentTypeNosniff:    // 是否启用 X-Content-Type-Options: nosniff 防止浏览器猜测 MIME 类型（默认为 true）
//		BrowserXssFilter:      // 是否启用 X-XSS-Protection: 1; mode=block 防御跨站脚本攻击（默认为 true）
//		ContentSecurityPolicy: "default-src 'self'", // 内容安全策略，默认只允许加载同一源下的资源
//		SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"}, // 用于识别经过代理服务器的 HTTPS 请求头映射（默认将 "X-Forwarded-Proto" 设置为 "https"）
// ```
# <翻译结束>


<原文开始>
// New creates an instance of the secure middleware using the specified configuration.
// router.Use(secure.N)
<原文结束>

# <翻译开始>
// New 根据指定的配置创建一个安全中间件实例。
// 使用方式：router.Use(secure.N)
# <翻译结束>

