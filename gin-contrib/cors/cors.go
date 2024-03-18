package cors

import (
	"errors"
	"strings"
	"time"
	
	"github.com/888go/gin"
)

// Config 代表了该中间件可用的所有配置选项。
type Config struct {
	AllowAllOrigins bool

// AllowOrigins 是一个允许跨域请求发起的源列表。
// 如果该列表中存在特殊的 "*" 值，则所有来源都将被允许。
// 默认值为 []
	AllowOrigins []string

// AllowOriginFunc 是一个自定义函数，用于验证请求来源。它接收一个起源（origin）作为参数，并在允许访问时返回 true，否则返回 false。如果设置了这个选项，则会忽略 AllowOrigins 的内容。
	AllowOriginFunc func(origin string) bool

// AllowMethods 是一个方法列表，用于指定客户端在跨域请求中允许使用的方法。默认值包括常见的简单方法（GET, POST, PUT, PATCH, DELETE, HEAD 和 OPTIONS）
	AllowMethods []string

// AllowPrivateNetwork 表示响应中是否应包含允许私有网络头信息
	AllowPrivateNetwork bool

// AllowHeaders 是一个包含客户端在跨域请求中允许使用的非简单头部列表。
	AllowHeaders []string

// AllowCredentials 指示请求是否可以包含用户凭据，如cookies、HTTP认证或客户端SSL证书。
	AllowCredentials bool

// ExposeHeaders 表示哪些头部信息是安全的，可以暴露给 CORS（跨源资源共享）API 规范的 API。
	ExposeHeaders []string

// MaxAge 指示预检请求结果可以被缓存的时间长度（以秒为单位）
	MaxAge time.Duration

// 允许添加如下形式的源地址：
// http://some-domain/* 表示以 http://some-domain/ 开头的所有URL
// https://api.* 表示以 https://api. 开头，且域名后跟任意字符的所有URL
// http://some.*.subdomain.com 表示以 http://some. 以及任意子域名.subdomain.com 结尾的所有URL
	AllowWildcard bool

// 允许使用流行的浏览器扩展方案
	AllowBrowserExtensions bool

// 允许使用WebSocket协议
	AllowWebSockets bool

// 允许使用 file:// 协议模式（危险！）仅在您100%确定需要时才使用它
	AllowFiles bool

// 允许为旧版浏览器/客户端传入自定义的OPTIONS响应状态码
	OptionsResponseStatusCode int
}

// AddAllowMethods 允许添加自定义方法

// ff:
// methods:

// ff:
// methods:

// ff:
// methods:

// ff:
// methods:

// ff:
// methods:

// ff:
// methods:
func (c *Config) AddAllowMethods(methods ...string) {
	c.AllowMethods = append(c.AllowMethods, methods...)
}

// AddAllowHeaders 允许添加自定义头信息

// ff:
// headers:

// ff:
// headers:

// ff:
// headers:

// ff:
// headers:

// ff:
// headers:

// ff:
// headers:
func (c *Config) AddAllowHeaders(headers ...string) {
	c.AllowHeaders = append(c.AllowHeaders, headers...)
}

// AddExposeHeaders 允许添加自定义暴露头信息

// ff:
// headers:

// ff:
// headers:

// ff:
// headers:

// ff:
// headers:

// ff:
// headers:

// ff:
// headers:
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

// Validate 是用于检查用户自定义配置的功能。

// ff:

// ff:

// ff:

// ff:

// ff:

// ff:
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

// ff:

// ff:

// ff:

// ff:

// ff:

// ff:
func DefaultConfig() Config {
	return Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}
}

// Default 返回默认配置的 location 中间件。

// ff:

// ff:

// ff:

// ff:

// ff:

// ff:
func Default() gin.HandlerFunc {
	config := DefaultConfig()
	config.AllowAllOrigins = true
	return New(config)
}

// New 返回一个带有用户自定义配置的 location 中间件。

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
	cors := newCors(config)
	return func(c *gin.Context) {
		cors.applyCors(c)
	}
}
