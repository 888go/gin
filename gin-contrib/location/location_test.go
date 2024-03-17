package location

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	
	"github.com/888go/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	gin.SetMode(gin.TestMode)
}

var defaultHeaders = Headers {
	Scheme: "X-Forwarded-Proto",
	Host:   "X-Forwarded-For",
}

var tests = []struct {
	want string
	conf Config
	req  *http.Request
}{
// 默认值
	{
		want: "http://localhost:8080",
		conf: DefaultConfig(),
		req: &http.Request{
			Header: http.Header{},
			URL:    &url.URL{},
		},
	},

// URL方案
	{
		want: "https://localhost:8080",
		conf: DefaultConfig(),
		req: &http.Request{
			Header: http.Header{},
			URL: &url.URL{
				Scheme: "https",
			},
		},
	},

// x-forward headers
// （注释翻译：）// x-forward headers，这个注释表明该代码段与“x-forwarded-headers”相关，这是一个HTTP头部信息，通常用于标识请求在经过代理服务器或负载均衡器时的原始来源信息。
	{
		want: "https://bar.com/bar",
		conf: Config{"http", "foo.com", "/bar", defaultHeaders},
		req: &http.Request{
			Header: http.Header{
				"X-Forwarded-Proto": {"https"},
				"X-Forwarded-For":   {"bar.com"},
			},
			URL: &url.URL{},
		},
	},

// X-Host 头部信息
	{
		want: "http://bar.com/bar",
		conf: Config{"http", "foo.com", "/bar", defaultHeaders},
		req: &http.Request{
			Header: http.Header{
				"X-Host": {"bar.com"},
			},
			URL: &url.URL{},
		},
	},

// URL 主机
	{
		want: "http://bar.com/bar",
		conf: Config{"http", "foo.com", "/bar", defaultHeaders},
		req: &http.Request{
			Header: http.Header{},
			URL: &url.URL{
				Host: "bar.com",
			},
		},
	},

// 请求
	{
		want: "https://baz.com/bar",
		conf: Config{"http", "foo.com", "/bar", defaultHeaders},
		req: &http.Request{
			Proto:  "HTTPS://",
			Host:   "baz.com",
			Header: http.Header{},
			URL:    &url.URL{},
		},
	},

// tls // （Transport Layer Security，传输层安全协议）
	{
		want: "https://foo.com/bar",
		conf: Config{"http", "foo.com", "/bar", defaultHeaders},
		req: &http.Request{
			TLS:    &tls.ConnectionState{},
			Header: http.Header{},
			URL:    &url.URL{},
		},
	},

// X-Forwarded-Host：主机头信息
	{
		want: "http://bar.com/bar",
		conf: Config{"http", "foo.com", "/bar", Headers{
			Scheme: "X-Forwarded-Proto",
			Host:   "X-Forwarded-Host",
		}},
		req: &http.Request{
			Header: http.Header{
				"X-Forwarded-Host": {"bar.com"},
			},
			URL: &url.URL{},
		},
	},
}


// ff:
// t:

// ff:
// t:
func TestLocation(t *testing.T) {
	for _, test := range tests {
		c := new(gin.Context)
		c.Request = test.req
		loc := newLocation(test.conf)
		loc.applyToContext(c)

		got := Get(c)

		if got.String() != test.want {
			t.Errorf("wanted location %s, got %s", got.String(), test.want)
		}
	}
}

func defaultRouter() *gin.Engine {
	router := gin.New()
	router.Use(Default())

	router.GET("/", func(c *gin.Context) {
		url := Get(c)
		c.String(200, url.String())
	})

	return router
}

func performRequest(r http.Handler, method string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}


// ff:
// t:

// ff:
// t:
func TestDefault(t *testing.T) {
	router := defaultRouter()
	w := performRequest(router, "GET")

	assert.Equal(t, "http://localhost:8080", w.Body.String())
}

func customRouter(config Config) *gin.Engine {
	router := gin.New()
	router.Use(New(config))

	router.GET("/", func(c *gin.Context) {
		url := Get(c)
		c.String(200, url.String())
	})

	return router
}


// ff:
// t:

// ff:
// t:
func TestCustom(t *testing.T) {
	router := customRouter(Config{
		Scheme: "https",
		Host:   "foo.com",
		Base:   "/base",
	})
	w := performRequest(router, "GET")

	assert.Equal(t, "https://foo.com/base", w.Body.String())
}
