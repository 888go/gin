package location

import (
	"net/url"
	
	"github.com/888go/gin"
)

const key = "location"

// Headers 表示用于映射方案和主机的头部字段。
type Headers struct {
	Scheme string
	Host   string
}

// Config 代表了该中间件可用的所有配置选项。
type Config struct {
// Scheme 是默认的方案，当无法从传入的 http.Request 中确定时应使用此方案。
	Scheme string

// Host 是默认主机，当无法从传入的 http.Request 中确定时应使用该主机。
	Host string

// Base 是基础路径，应与执行路径重写操作的代理服务器结合使用。
	Base string

// Header 用于映射方案和主机。
// 可以重写以允许从自定义头部字段读取值。
	Headers Headers
}

// DefaultConfig 返回一个映射到本机的通用默认配置。

// ff:

// ff:

// ff:

// ff:

// ff:
func DefaultConfig() Config {
	return Config{
		Host:   "localhost:8080",
		Scheme: "http",
		Headers: Headers {
			Scheme: "X-Forwarded-Proto",
			Host:   "X-Forwarded-For",
		},
	}
}

// Default 返回默认配置的 location 中间件。

// ff:

// ff:

// ff:

// ff:

// ff:
func Default() gin.HandlerFunc {
	config := DefaultConfig()
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
func New(config Config) gin.HandlerFunc {
	location := newLocation(config)

	return func(c *gin.Context) {
		location.applyToContext(c)
	}
}

// Get 从 context 中获取传入 http.Request 的 Location 信息。如果未设置 location，则返回空值（nil）。

// ff:
// c:

// ff:
// c:

// ff:
// c:

// ff:
// c:

// ff:
// c:
func Get(c *gin.Context) *url.URL {
	v, ok := c.Get(key)

	if !ok {
		return nil
	}

	vv, ok := v.(*url.URL)

	if !ok {
		return nil
	}

	return vv
}
