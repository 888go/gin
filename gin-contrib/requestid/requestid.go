package requestid

import (
	"github.com/888go/gin"
	"github.com/google/uuid"
)

var headerXRequestID string

// Config 定义了 RequestID 中间件的配置
type config struct {
// Generator 定义了一个用于生成 ID 的函数。
// 可选。默认值：func() string {
//   return uuid.New().String()
// }
// 
// 翻译：
// Generator 描述了一个生成ID的函数。
// 非必填项。默认值为：
// func() string {
//   返回 uuid.New().String()
// }
	generator Generator
	headerKey HeaderStrKey
	handler   Handler
}

// New 初始化 RequestID 中间件。
func New(opts ...Option) gin类.HandlerFunc {
	cfg := &config{
		generator: func() string {
			return uuid.New().String()
		},
		headerKey: "X-Request-ID",
	}

	for _, opt := range opts {
		opt(cfg)
	}

	headerXRequestID = string(cfg.headerKey)

	return func(c *gin类.Context) {
		// Get id from request
		rid := c.X取请求协议头值(headerXRequestID)
		if rid == "" {
			rid = cfg.generator()
			c.X请求.Header.Add(headerXRequestID, rid)
		}
		if cfg.handler != nil {
			cfg.handler(c, rid)
		}
		// 设置id以确保请求id在响应中
		c.X设置响应协议头值(headerXRequestID, rid)
		c.X中间件继续()
	}
}

// Get 返回请求标识符
func Get(c *gin类.Context) string {
	return c.X取请求协议头值(headerXRequestID)
}
