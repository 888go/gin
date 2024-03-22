package requestid

import (
	"github.com/888go/gin"
)

// 队列系统的选项
type Option func(*config)

type (
	Generator func() string
	Handler   func(c *gin类.Context, requestID string)
)

type HeaderStrKey string

// WithGenerator 设置生成器函数
func WithGenerator(g Generator) Option {
	return func(cfg *config) {
		cfg.generator = g
	}
}

// WithCustomHeaderStrKey 为请求ID设置自定义头部键
func WithCustomHeaderStrKey(s HeaderStrKey) Option {
	return func(cfg *config) {
		cfg.headerKey = s
	}
}

// WithHandler 为带有上下文的请求ID设置处理函数
func WithHandler(handler Handler) Option {
	return func(cfg *config) {
		cfg.handler = handler
	}
}
