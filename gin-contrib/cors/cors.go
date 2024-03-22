package cors

import (
	"errors"
	"strings"
	"time"
	
	"github.com/888go/gin"
)

// Config 表示该中间件的所有可用选项。
type Config struct {
	AllowAllOrigins bool

// AllowOrigins 是一个跨域请求允许执行的源列表。
// 如果列表中存在特殊的 "*" 值，则允许所有来源。
// 默认值为 []
	AllowOrigins []string

// AllowOriginFunc 是一个自定义函数，用于验证源。它接受一个 origin 参数，并在允许访问时返回 true，否则返回 false。如果设置了此选项，则忽略 AllowOrigins 的内容。
	AllowOriginFunc func(origin string) bool

// AllowMethods 是一个方法列表，列出了客户端在跨域请求中允许使用的HTTP方法。默认值包括简单的方法（GET, POST, PUT, PATCH, DELETE, HEAD 和 OPTIONS）。
	AllowMethods []string

	// AllowPrivateNetwork 指示响应中是否应包含允许私有网络头
	AllowPrivateNetwork bool

// AllowHeaders 是一个包含客户端在跨域请求中允许使用的非简单头信息的列表。
	AllowHeaders []string

// AllowCredentials 指示请求是否可以包含用户凭据，如cookies、HTTP认证或客户端SSL证书。
	AllowCredentials bool

// ExposeHeaders 表示哪些头部信息在 CORS（跨源资源共享）API 规范中是安全的，并可以暴露给 API。
	ExposeHeaders []string

// MaxAge 指示预检请求结果可以被缓存的时间长度（精确到秒）
	MaxAge time.Duration

	// 允许添加诸如 http://some-domain/*、https://api.* 或 http://some.*.subdomain.com 等来源
	AllowWildcard bool

	// 允许使用流行的浏览器扩展程序架构
	AllowBrowserExtensions bool

	// 允许使用WebSocket协议
	AllowWebSockets bool

	// 允许使用 file:// 方案（危险！）仅在您 100% 确定需要时才使用它
	AllowFiles bool

	// 允许为旧版浏览器/客户端传入自定义的OPTIONS响应状态码
	OptionsResponseStatusCode int
}

// AddAllowMethods 允许添加自定义方法
func (c *Config) AddAllowMethods(methods ...string) {
	c.AllowMethods = append(c.AllowMethods, methods...)
}

// AddAllowHeaders 允许添加自定义头
func (c *Config) AddAllowHeaders(headers ...string) {
	c.AllowHeaders = append(c.AllowHeaders, headers...)
}

// AddExposeHeaders 允许添加自定义暴露头
func (c *Config) AddExposeHeaders(headers ...string) {
	c.ExposeHeaders = append(c.ExposeHeaders, headers...)
}

func (c Config) getAllowedSchemas() []string {
	allowedSchemas := DefaultSchemas
	if c.AllowBrowserExtensions {
		allowedSchemas = append(allowedSchemas, ExtensionSchemas...)
	}
	if c.AllowWebSockets {
		allowedSchemas = append(allowedSchemas, WebSocketSchemas...)
	}
	if c.AllowFiles {
		allowedSchemas = append(allowedSchemas, FileSchemas...)
	}
	return allowedSchemas
}

func (c Config) validateAllowedSchemas(origin string) bool {
	allowedSchemas := c.getAllowedSchemas()
	for _, schema := range allowedSchemas {
		if strings.HasPrefix(origin, schema) {
			return true
		}
	}
	return false
}

// Validate 是检查用户自定义配置的功能。
func (c Config) Validate() error {
	if c.AllowAllOrigins && (c.AllowOriginFunc != nil || len(c.AllowOrigins) > 0) {
		return errors.New("conflict settings: all origins are allowed. AllowOriginFunc or AllowOrigins is not needed")
	}
	if !c.AllowAllOrigins && c.AllowOriginFunc == nil && len(c.AllowOrigins) == 0 {
		return errors.New("conflict settings: all origins disabled")
	}
	for _, origin := range c.AllowOrigins {
		if !strings.Contains(origin, "*") && !c.validateAllowedSchemas(origin) {
			return errors.New("bad origin: origins must contain '*' or include " + strings.Join(c.getAllowedSchemas(), ","))
		}
	}
	return nil
}

func (c Config) parseWildcardRules() [][]string {
	var wRules [][]string

	if !c.AllowWildcard {
		return wRules
	}

	for _, o := range c.AllowOrigins {
		if !strings.Contains(o, "*") {
			continue
		}

		if c := strings.Count(o, "*"); c > 1 {
			panic(errors.New("only one * is allowed").Error())
		}

		i := strings.Index(o, "*")
		if i == 0 {
			wRules = append(wRules, []string{"*", o[1:]})
			continue
		}
		if i == (len(o) - 1) {
			wRules = append(wRules, []string{o[:i-1], "*"})
			continue
		}

		wRules = append(wRules, []string{o[:i], o[i+1:]})
	}

	return wRules
}

// DefaultConfig 返回一个映射到本机的通用默认配置。
func DefaultConfig() Config {
	return Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}
}

// 默认返回使用默认配置的位置中间件。
func Default() gin类.HandlerFunc {
	config := DefaultConfig()
	config.AllowAllOrigins = true
	return New(config)
}

// New 函数返回一个使用用户自定义配置的 location 中间件。
func New(config Config) gin类.HandlerFunc {
	cors := newCors(config)
	return func(c *gin类.Context) {
		cors.applyCors(c)
	}
}
