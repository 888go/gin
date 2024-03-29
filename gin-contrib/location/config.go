package location

import (
	"net/http"
	"net/url"
	"strings"
	
	"github.com/888go/gin"
)

type (
	location struct {
		scheme  string
		host    string
		base    string
		headers Headers
	}
)

func newLocation(config Config) *location {
	return &location{
		scheme: config.Scheme,
		host:   config.Host,
		base:   config.Base,
		headers: config.Headers,
	}
}

func (l *location) applyToContext(c *gin类.Context) {
	value := new(url.URL)
	value.Scheme = l.resolveScheme(c.X请求)
	value.Host = l.resolveHost(c.X请求)
	value.Path = l.base
	c.X设置值(key, value)
}

func (l *location) resolveScheme(r *http.Request) string {
	switch {
	case r.Header.Get(l.headers.Scheme) == "https":
		return "https"
	case r.URL.Scheme == "https":
		return "https"
	case r.TLS != nil:
		return "https"
	case strings.HasPrefix(r.Proto, "HTTPS"):
		return "https"
	default:
		return l.scheme
	}
}

func (l *location) resolveHost(r *http.Request) (host string) {
	switch {
	case r.Header.Get(l.headers.Host) != "":
		return r.Header.Get(l.headers.Host)
	case r.Header.Get("X-Host") != "":
		return r.Header.Get("X-Host")
	case r.Host != "":
		return r.Host
	case r.URL.Host != "":
		return r.URL.Host
	default:
		return l.host
	}
}
