package requestid

import (
	"github.com/888go/gin"
	"github.com/google/uuid"
)

var headerXRequestID string

// Config 定义了 RequestID 中间件的配置
type config struct {
// Generator 定义了一个用于生成 ID 的函数。
// 可选，默认值：func() string {
//   return uuid.New().String()
// }
// （译文）：// Generator 用于定义一个生成 ID 的函数。
// 该参数可选，默认实现为：
// func() string {
//   // 使用uuid包生成新的UUID并返回其字符串形式
//   return uuid.New().String()
// }
	generator Generator
	headerKey HeaderStrKey
	handler   Handler
}

// New 初始化 RequestID 中间件。

// ff:
// opts:

// ff:
// opts:

// ff:
// opts:

// ff:
// opts:

// ff:
// opts:

// ff:
// opts:

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
// 从请求中获取id
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

// ff:
// c:
func Get(c *gin.Context) string {
	return c.GetHeader(headerXRequestID)
}
