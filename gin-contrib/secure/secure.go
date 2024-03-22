package secure

import (
	"github.com/888go/gin"
)

// Config 是一个用于指定 secure 配置选项的结构体。
type Config struct {
// AllowedHosts 是一个完全合格域名列表，其中包含了允许访问的所有域名。
// 默认为空列表，这意味着允许任何和所有主机名。
	AllowedHosts []string
// 如果SSLRedirect设置为true，则只允许https请求。
// 默认值为false。
	SSLRedirect bool
// 如果SSLTemporaryRedirect为true，在重定向时将使用302状态码。默认值为false（即使用301状态码）。
	SSLTemporaryRedirect bool
// SSLHost 是用于将 HTTP 请求重定向到 HTTPS 的主机名。
// 默认值为 ""，表示使用相同的主机。
	SSLHost string
// STSSeconds 是 Strict-Transport-Security 头部的最大有效期（max-age）。
// 默认值为 0，这意味着不会包含该头部。
	STSSeconds int64
// 如果将STSIncludeSubdomains设置为true，则`includeSubdomains`将被添加到Strict-Transport-Security头中。默认值为false。
	STSIncludeSubdomains bool
// 如果FrameDeny设置为true，则会添加X-Frame-Options头，并将其值设为`DENY`。默认为false。
	FrameDeny bool
// CustomFrameOptionsValue 允许设置自定义 X-Frame-Options 头部值。这将覆盖 FrameDeny 选项。
	CustomFrameOptionsValue string
// 如果ContentTypeNosniff设为true，则会添加一个X-Content-Type-Options头，并将其值设为`nosniff`。默认值为false。
	ContentTypeNosniff bool
// 如果BrowserXssFilter设为true，则会添加“X-XSS-Protection”头部，并将其值设为`1; mode=block`。默认情况下为false。
	BrowserXssFilter bool
// ContentSecurityPolicy 允许设置自定义的“Content-Security-Policy”头部值，默认为空字符串""。
	ContentSecurityPolicy string
	// HTTP头部"Referrer-Policy"控制在请求中应随Referrer头部一起发送哪些referrer信息。
	ReferrerPolicy string
	// 当设为 true 时，中间件应用的整个安全策略将被完全禁用。
	IsDevelopment bool
	// 处理错误发生时的处理程序（例如，主机错误）。
	BadHostHandler gin类.HandlerFunc
	// Prevent Internet Explorer from executing downloads in your site’s context
	IENoOpen bool
	// 功能策略是一个新的头部信息，它允许网站控制哪些功能和API可以在浏览器中使用。
	FeaturePolicy string
// 如果DontRedirectIPV4Hostnames设为true，那么对IPv4地址形式的主机名的请求将不会被重定向。这是为了允许负载均衡器健康检查能够成功执行。
	DontRedirectIPV4Hostnames bool

// 如果请求不安全，但本字典中任意一个头部被设置为它们对应的值时，将请求视为安全请求。
// 这在你的应用运行在一个通过 http 转发请求到你应用的 secure 代理（例如 Heroku）后面时非常有用。
	SSLProxyHeaders map[string]string
}

// DefaultConfig 返回一个具有严格安全设置的 Configuration。
// ```
//   SSLRedirect:           true     // 启用 SSL 重定向
//   IsDevelopment:         false    // 设置为非开发模式
//   STSSeconds:            315360000 // 设置严格的STS（Strict-Transport-Security）头，有效期为十年
//   STSIncludeSubdomains:  true     // STS 头部包含子域名
//   FrameDeny:             true     // 禁止使用 frame 标签嵌入页面
//   ContentTypeNosniff:    true     // 阻止浏览器猜测 MIME 类型，防止内容嗅探攻击
//   BrowserXssFilter:      true     // 启用浏览器级别的 XSS 过滤器
//   ContentSecurityPolicy: "default-src 'self'" // 设置内容安全策略（CSP），默认源只能为自身站点
//   SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"}, // 当通过代理时，将 X-Forwarded-Proto 头设置为 https，确保 SSL 重定向正确执行
// ```
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

// New 函数使用指定的配置创建一个安全中间件实例。
// 示例用法：router.Use(secure.New(...))
// 其中，router 代表你的应用路由对象，secure.New(...) 用于生成并初始化安全中间件。
func New(config Config) gin类.HandlerFunc {
	policy := newPolicy(config)
	return func(c *gin类.Context) {
		if !policy.applyToContext(c) {
			return
		}
	}
}
