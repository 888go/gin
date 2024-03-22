package requestid

import (
	"github.com/888go/gin"
	"github.com/google/uuid"
)

var headerXRequestID string

// Config defines the config for RequestID middleware
type config struct {
	// Generator defines a function to generate an ID.
	// Optional. Default: func() string {
	//   return uuid.New().String()
	// }
	generator Generator
	headerKey HeaderStrKey
	handler   Handler
}

// New initializes the RequestID middleware.
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
		// Set the id to ensure that the requestid is in the response
		c.X设置响应协议头值(headerXRequestID, rid)
		c.X中间件继续()
	}
}

// Get returns the request identifier
func Get(c *gin类.Context) string {
	return c.X取请求协议头值(headerXRequestID)
}
