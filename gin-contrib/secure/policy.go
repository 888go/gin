package secure

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	
	"github.com/888go/gin"
)

type (
// Secure 是一个中间件，用于帮助设置一些基本的安全功能。可以通过提供一个 secure.Options 结构体来配置要启用哪些功能，并且可以覆盖一些默认值。
	policy struct {
		// 使用Options结构体来自定义Secure。
		config       Config
		fixedHeaders []header
	}

	header struct {
		key   string
		value []string
	}
)

// 使用提供的选项构建一个新的Policy实例。
func newPolicy(config Config) *policy {
	policy := &policy{}
	policy.loadConfig(config)
	return policy
}

func (p *policy) loadConfig(config Config) {
	p.config = config
	p.fixedHeaders = make([]header, 0, 5)

	// Frame Options header.
	if len(config.CustomFrameOptionsValue) > 0 {
		p.addHeader("X-Frame-Options", config.CustomFrameOptionsValue)
	} else if config.FrameDeny {
		p.addHeader("X-Frame-Options", "DENY")
	}

	// 内容类型选项头。
	if config.ContentTypeNosniff {
		p.addHeader("X-Content-Type-Options", "nosniff")
	}

	// XSS Protection header.
	if config.BrowserXssFilter {
		p.addHeader("X-Xss-Protection", "1; mode=block")
	}

	// 内容安全策略头。
	if len(config.ContentSecurityPolicy) > 0 {
		p.addHeader("Content-Security-Policy", config.ContentSecurityPolicy)
	}

	if len(config.ReferrerPolicy) > 0 {
		p.addHeader("Referrer-Policy", config.ReferrerPolicy)
	}

	// 严格传输安全（Strict Transport Security）头部。
	if config.STSSeconds != 0 {
		stsSub := ""
		if config.STSIncludeSubdomains {
			stsSub = "; includeSubdomains"
		}

// TODO：待办事项
// "max-age=%d%s" 需重构
		p.addHeader(
			"Strict-Transport-Security",
			fmt.Sprintf("max-age=%d%s", config.STSSeconds, stsSub))
	}

	// X-Download-Options 头部信息。
	if config.IENoOpen {
		p.addHeader("X-Download-Options", "noopen")
	}

	// FeaturePolicy header.
	if len(config.FeaturePolicy) > 0 {
		p.addHeader("Feature-Policy", config.FeaturePolicy)
	}
}

func (p *policy) addHeader(key string, value string) {
	p.fixedHeaders = append(p.fixedHeaders, header{
		key:   key,
		value: []string{value},
	})
}

func (p *policy) applyToContext(c *gin类.Context) bool {
	if !p.config.IsDevelopment {
		p.writeSecureHeaders(c)

		if !p.checkAllowHosts(c) {
			return false
		}
		if !p.checkSSL(c) {
			return false
		}
	}
	return true
}

func (p *policy) writeSecureHeaders(c *gin类.Context) {
	header := c.Writer.Header()
	for _, pair := range p.fixedHeaders {
		header[pair.key] = pair.value
	}
}

func (p *policy) checkAllowHosts(c *gin类.Context) bool {
	if len(p.config.AllowedHosts) == 0 {
		return true
	}

	host := c.X请求.Host
	if len(host) == 0 {
		host = c.X请求.URL.Host
	}

	for _, allowedHost := range p.config.AllowedHosts {
		if strings.EqualFold(allowedHost, host) {
			return true
		}
	}

	if p.config.BadHostHandler != nil {
		p.config.BadHostHandler(c)
	} else {
		c.X停止并带状态码(403)
	}

	return false
}

// 检查一个主机（可能带有端口号）是否为IPV4地址
func isIPV4(host string) bool {
	if index := strings.IndexByte(host, ':'); index != -1 {
		host = host[:index]
	}
	return net.ParseIP(host) != nil
}

func (p *policy) isSSLRequest(req *http.Request) bool {
	if strings.EqualFold(req.URL.Scheme, "https") || req.TLS != nil {
		return true
	}

	for h, v := range p.config.SSLProxyHeaders {
		hv, ok := req.Header[h]

		if !ok {
			continue
		}

		if strings.EqualFold(hv[0], v) {
			return true
		}
	}

	if p.config.DontRedirectIPV4Hostnames && isIPV4(req.Host) {
		return true
	}

	return false
}

func (p *policy) checkSSL(c *gin类.Context) bool {
	if !p.config.SSLRedirect {
		return true
	}

	req := c.X请求
	isSSLRequest := p.isSSLRequest(req)
	if isSSLRequest {
		return true
	}

// TODO：（待办事项）
// req.Host 与 req.URL.Host 的对比
	url := req.URL
	url.Scheme = "https"
	url.Host = req.Host

	if len(p.config.SSLHost) > 0 {
		url.Host = p.config.SSLHost
	}

	status := http.StatusMovedPermanently
	if p.config.SSLTemporaryRedirect {
		status = http.StatusTemporaryRedirect
	}
	c.X重定向(status, url.String())
	c.X停止()
	return false
}
