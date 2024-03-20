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

// ff:
// opts:
func New(opts ...Option) gin.HandlerFunc {
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

	return func(c *gin.Context) {
		// Get id from request
		rid := c.GetHeader(headerXRequestID)
		if rid == "" {
			rid = cfg.generator()
			c.Request.Header.Add(headerXRequestID, rid)
		}
		if cfg.handler != nil {
			cfg.handler(c, rid)
		}
		// 设置id以确保请求id在响应中
		c.Header(headerXRequestID, rid)
		c.Next()
	}
}

// Get 返回请求标识符

// ff:
// c:
func Get(c *gin.Context) string {
	return c.GetHeader(headerXRequestID)
}
