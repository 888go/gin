package timeout

import (
	"net/http"
	"time"
	
	"github.com/888go/gin"
)

// Option for timeout
type Option func(*Timeout)

// WithTimeout set timeout
func WithTimeout(timeout time.Duration) Option {
	return func(t *Timeout) {
		t.timeout = timeout
	}
}

// WithHandler add gin handler
func WithHandler(h gin类.HandlerFunc) Option {
	return func(t *Timeout) {
		t.handler = h
	}
}

// WithResponse add gin handler
func WithResponse(h gin类.HandlerFunc) Option {
	return func(t *Timeout) {
		t.response = h
	}
}

func defaultResponse(c *gin类.Context) {
	c.X输出文本(http.StatusRequestTimeout, http.StatusText(http.StatusRequestTimeout))
}

// Timeout struct
type Timeout struct {
	timeout  time.Duration
	handler  gin类.HandlerFunc
	response gin类.HandlerFunc
}
