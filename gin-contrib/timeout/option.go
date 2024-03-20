package timeout

import (
	"net/http"
	"time"
	
	"github.com/888go/gin"
)

// Option for timeout
type Option func(*Timeout)

// WithTimeout 设置超时

// ff:
// timeout:
func WithTimeout(timeout time.Duration) Option {
	return func(t *Timeout) {
		t.timeout = timeout
	}
}

// WithHandler 添加 Gin 处理器

// ff:
// h:
func WithHandler(h gin.HandlerFunc) Option {
	return func(t *Timeout) {
		t.handler = h
	}
}

// WithResponse 添加 Gin 处理器

// ff:
// h:
func WithResponse(h gin.HandlerFunc) Option {
	return func(t *Timeout) {
		t.response = h
	}
}

func defaultResponse(c *gin.Context) {
	c.String(http.StatusRequestTimeout, http.StatusText(http.StatusRequestTimeout))
}

// Timeout struct
type Timeout struct {
	timeout  time.Duration
	handler  gin.HandlerFunc
	response gin.HandlerFunc
}
