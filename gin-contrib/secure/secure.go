package secure

import (
	"github.com/888go/gin"
)

// Config 是一个结构体，用于指定 secure 的配置选项。
type Config struct {
// AllowedHosts 是一个完全限定域名列表，其中包含允许的所有域名。
// 默认为空列表，这意味着允许任何和所有主机名。
	AllowedHosts []string
// 如果SSLRedirect设置为true，则只允许https请求。
// 默认值为false。
	SSLRedirect bool
// 如果SSLTemporaryRedirect设为true，在重定向时将使用302状态码。
// 默认值为false（即使用301状态码）。
	SSLTemporaryRedirect bool
// SSLHost 是用于将http请求重定向到https的主机名。
// 默认值为 ""，表示使用相同的主机。
	SSLHost string
// STSSeconds 是 Strict-Transport-Security 头部的最大有效期（max-age）。
// 默认值为 0，这意味着不会包含此头部。
	STSSeconds int64
// 如果将STSIncludeSubdomains设置为true，则`includeSubdomains`将被追加到Strict-Transport-Security头中。默认值为false。
	STSIncludeSubdomains bool
// 如果FrameDeny设置为true，则添加X-Frame-Options头，并将其值设为`DENY`。默认值为false。
	FrameDeny bool
// CustomFrameOptionsValue 允许自定义设置 X-Frame-Options 头部值，这将覆盖 FrameDeny 选项。
	CustomFrameOptionsValue string
// 如果ContentTypeNosniff设为true，将会添加“X-Content-Type-Options”头部，并将其值设为`nosniff`。默认值为false。
	ContentTypeNosniff bool
// 如果BrowserXssFilter为true，则添加X-XSS-Protection头，并将其值设为`1; mode=block`。默认值为false。
	BrowserXssFilter bool
// ContentSecurityPolicy 允许设置自定义的 Content-Security-Policy 头部值，默认为空字符串""。
	ContentSecurityPolicy string
// HTTP头部"Referrer-Policy"规定了在请求中随Referrer头部一起发送的referrer信息应包含哪些内容。
	ReferrerPolicy string
// 当设为true时，中间件应用的整个安全策略将被完全禁用。
	IsDevelopment bool
// 当发生错误（例如主机错误）时的处理程序。
	BadHostHandler gin.HandlerFunc
	// Prevent Internet Explorer from executing downloads in your site’s context
	IENoOpen bool
// 功能策略是一个新的头部信息，它允许网站控制哪些功能和API可以在浏览器中使用。
	FeaturePolicy string
// 如果DontRedirectIPV4Hostnames设为true，那么对IPv4地址形式的主机名的请求将不会被重定向。这是为了允许负载均衡器健康检查能够成功执行。
	DontRedirectIPV4Hostnames bool

// 如果请求是非安全的，但该字典中任意一个头部被设置为它们对应的值时，则将请求视为安全请求。
// 这在你的应用运行在一个通过http转发请求到你应用的（如Heroku）安全代理后面时非常有用。
// 以下是逐句翻译：
// 如果请求是不安全的，
// 但如果这个字典中的任何头部被设置为了它们相应的值，
// 那么就将这个请求当作安全请求来处理。
// 这在你的应用程序运行在一个会通过http将请求转发到你的应用程序的（例如Heroku这样的）安全代理之后的情况下非常有帮助。
	SSLProxyHeaders map[string]string
}

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

// ff:

// ff:

// ff:

// ff:

// ff:

// ff:

// ff:
func DefaultConfig() Config {
	return Config{
		SSLRedirect:           true,
		IsDevelopment:         false,
		STSSeconds:            315360000,
		STSIncludeSubdomains:  true,
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "default-src 'self'",
		IENoOpen:              true,
		SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
	}
}

// New 根据指定的配置创建一个安全中间件实例。
// 使用方式：router.Use(secure.N)

// ff:
// config:

// ff:
// config:

// ff:
// config:

// ff:
// config:

// ff:
// config:

// ff:
// config:

// ff:
// config:
func New(config Config) gin.HandlerFunc {
	policy := newPolicy(config)
	return func(c *gin.Context) {
		if !policy.applyToContext(c) {
			return
		}
	}
}
