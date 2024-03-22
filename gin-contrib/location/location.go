package location

import (
	"net/url"
	
	"github.com/888go/gin"
)

const key = "location"

// Headers 代表用于映射方案和主机的头部字段。
type Headers struct {
	Scheme string
	Host   string
}

// Config 表示该中间件的所有可用选项。
type Config struct {
// Scheme是当无法从传入的http.Request中明确获知时，应使用的默认方案。
	Scheme string

// Host 是默认主机，当无法从传入的 http.Request 中明确获取时，应使用此主机。
	Host string

// Base 是基路径，应与进行路径重写操作的代理服务器结合使用。
	Base string

// 该Header用于映射方案和主机。
// 可以被覆盖以允许从自定义头部字段读取值。
	Headers Headers
}

// DefaultConfig 返回一个映射到本机的通用默认配置。
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

// 默认返回使用默认配置的位置中间件。
func Default() gin类.HandlerFunc {
	config := DefaultConfig()
	return New(config)
}

// New 函数返回一个使用用户自定义配置的 location 中间件。
func New(config Config) gin类.HandlerFunc {
	location := newLocation(config)

	return func(c *gin类.Context) {
		location.applyToContext(c)
	}
}

// Get 从上下文获取传入 http.Request 的 Location 信息。如果未设置位置信息，则返回 nil 值。
func Get(c *gin类.Context) *url.URL {
	v, ok := c.X取值(key)

	if !ok {
		return nil
	}

	vv, ok := v.(*url.URL)

	if !ok {
		return nil
	}

	return vv
}
