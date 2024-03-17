package cors

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
	
	"github.com/888go/gin"
	"github.com/stretchr/testify/assert"
)

func newTestRouter(config Config) *gin.Engine {
	router := gin.New()
	router.Use(New(config))
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "get")
	})
	router.POST("/", func(c *gin.Context) {
		c.String(http.StatusOK, "post")
	})
	router.PATCH("/", func(c *gin.Context) {
		c.String(http.StatusOK, "patch")
	})
	return router
}

func performRequest(r http.Handler, method, origin string) *httptest.ResponseRecorder {
	return performRequestWithHeaders(r, method, origin, http.Header{})
}

func performRequestWithHeaders(r http.Handler, method, origin string, header http.Header) *httptest.ResponseRecorder {
	req, _ := http.NewRequestWithContext(context.Background(), method, "/", nil)
// 来自 go/net/http/request.go:
// 对于传入的请求，Host 头部会被提升至 Request.Host 字段，并从 Header 映射中移除。
	req.Host = header.Get("Host")
	header.Del("Host")
	if len(origin) > 0 {
		header.Set("Origin", origin)
	}
	req.Header = header
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}


// ff:
// t:

// ff:
// t:
func TestConfigAddAllow(t *testing.T) {
	config := Config{}
	config.AddAllowMethods("POST")
	config.AddAllowMethods("GET", "PUT")
	config.AddExposeHeaders()

	config.AddAllowHeaders("Some", " cool")
	config.AddAllowHeaders("header")
	config.AddExposeHeaders()

	config.AddExposeHeaders()
	config.AddExposeHeaders("exposed", "header")
	config.AddExposeHeaders("hey")

	assert.Equal(t, config.AllowMethods, []string{"POST", "GET", "PUT"})
	assert.Equal(t, config.AllowHeaders, []string{"Some", " cool", "header"})
	assert.Equal(t, config.ExposeHeaders, []string{"exposed", "header", "hey"})
}


// ff:
// t:

// ff:
// t:
func TestBadConfig(t *testing.T) {
	assert.Panics(t, func() { New(Config{}) })
	assert.Panics(t, func() {
		New(Config{
			AllowAllOrigins: true,
			AllowOrigins:    []string{"http://google.com"},
		})
	})
	assert.Panics(t, func() {
		New(Config{
			AllowAllOrigins: true,
			AllowOriginFunc: func(origin string) bool { return false },
		})
	})
	assert.Panics(t, func() {
		New(Config{
			AllowOrigins: []string{"google.com"},
		})
	})
}


// ff:
// t:

// ff:
// t:
func TestNormalize(t *testing.T) {
	values := normalize([]string{
		"http-Access ", "Post", "POST", " poSt  ",
		"HTTP-Access", "",
	})
	assert.Equal(t, values, []string{"http-access", "post", ""})

	values = normalize(nil)
	assert.Nil(t, values)

	values = normalize([]string{})
	assert.Equal(t, values, []string{})
}


// ff:
// t:

// ff:
// t:
func TestConvert(t *testing.T) {
	methods := []string{"Get", "GET", "get"}
	headers := []string{"X-CSRF-TOKEN", "X-CSRF-Token", "x-csrf-token"}

	assert.Equal(t, []string{"GET", "GET", "GET"}, convert(methods, strings.ToUpper))
	assert.Equal(t, []string{"X-Csrf-Token", "X-Csrf-Token", "X-Csrf-Token"}, convert(headers, http.CanonicalHeaderKey))
}


// ff:
// t:

// ff:
// t:
func TestGenerateNormalHeaders_AllowAllOrigins(t *testing.T) {
	header := generateNormalHeaders(Config{
		AllowAllOrigins: false,
	})
	assert.Equal(t, header.Get("Access-Control-Allow-Origin"), "")
	assert.Equal(t, header.Get("Vary"), "Origin")
	assert.Len(t, header, 1)

	header = generateNormalHeaders(Config{
		AllowAllOrigins: true,
	})
	assert.Equal(t, header.Get("Access-Control-Allow-Origin"), "*")
	assert.Equal(t, header.Get("Vary"), "")
	assert.Len(t, header, 1)
}


// ff:
// t:

// ff:
// t:
func TestGenerateNormalHeaders_AllowCredentials(t *testing.T) {
	header := generateNormalHeaders(Config{
		AllowCredentials: true,
	})
	assert.Equal(t, header.Get("Access-Control-Allow-Credentials"), "true")
	assert.Equal(t, header.Get("Vary"), "Origin")
	assert.Len(t, header, 2)
}


// ff:
// t:

// ff:
// t:
func TestGenerateNormalHeaders_ExposedHeaders(t *testing.T) {
	header := generateNormalHeaders(Config{
		ExposeHeaders: []string{"X-user", "xPassword"},
	})
	assert.Equal(t, header.Get("Access-Control-Expose-Headers"), "X-User,Xpassword")
	assert.Equal(t, header.Get("Vary"), "Origin")
	assert.Len(t, header, 2)
}


// ff:
// t:

// ff:
// t:
func TestGeneratePreflightHeaders(t *testing.T) {
	header := generatePreflightHeaders(Config{
		AllowAllOrigins: false,
	})
	assert.Equal(t, header.Get("Access-Control-Allow-Origin"), "")
	assert.Equal(t, header.Get("Vary"), "Origin")
	assert.Len(t, header, 1)

	header = generateNormalHeaders(Config{
		AllowAllOrigins: true,
	})
	assert.Equal(t, header.Get("Access-Control-Allow-Origin"), "*")
	assert.Equal(t, header.Get("Vary"), "")
	assert.Len(t, header, 1)
}


// ff:
// t:

// ff:
// t:
func TestGeneratePreflightHeaders_AllowCredentials(t *testing.T) {
	header := generatePreflightHeaders(Config{
		AllowCredentials: true,
	})
	assert.Equal(t, header.Get("Access-Control-Allow-Credentials"), "true")
	assert.Equal(t, header.Get("Vary"), "Origin")
	assert.Len(t, header, 2)
}


// ff:
// t:

// ff:
// t:
func TestGeneratePreflightHeaders_AllowPrivateNetwork(t *testing.T) {
	header := generatePreflightHeaders(Config{
		AllowPrivateNetwork: true,
	})
	assert.Equal(t, header.Get("Access-Control-Allow-Private-Network"), "true")
	assert.Equal(t, header.Get("Vary"), "Origin")
	assert.Len(t, header, 2)
}


// ff:
// t:

// ff:
// t:
func TestGeneratePreflightHeaders_AllowMethods(t *testing.T) {
	header := generatePreflightHeaders(Config{
		AllowMethods: []string{"GET ", "post", "PUT", " put  "},
	})
	assert.Equal(t, header.Get("Access-Control-Allow-Methods"), "GET,POST,PUT")
	assert.Equal(t, header.Get("Vary"), "Origin")
	assert.Len(t, header, 2)
}


// ff:
// t:

// ff:
// t:
func TestGeneratePreflightHeaders_AllowHeaders(t *testing.T) {
	header := generatePreflightHeaders(Config{
		AllowHeaders: []string{"X-user", "Content-Type"},
	})
	assert.Equal(t, header.Get("Access-Control-Allow-Headers"), "X-User,Content-Type")
	assert.Equal(t, header.Get("Vary"), "Origin")
	assert.Len(t, header, 2)
}


// ff:
// t:

// ff:
// t:
func TestGeneratePreflightHeaders_MaxAge(t *testing.T) {
	header := generatePreflightHeaders(Config{
		MaxAge: 12 * time.Hour,
	})
	assert.Equal(t, header.Get("Access-Control-Max-Age"), "43200") // 12 * 60 * 60
// 这段代码是 Go 语言中的一个数值表达式，计算的是 12 小时乘以每小时的分钟数（60）再乘以每分钟的秒数（60），所以这个表达式的结果代表了 12 小时的总秒数。
	assert.Equal(t, header.Get("Vary"), "Origin")
	assert.Len(t, header, 2)
}


// ff:
// t:

// ff:
// t:
func TestValidateOrigin(t *testing.T) {
	cors := newCors(Config{
		AllowAllOrigins: true,
	})
	assert.True(t, cors.validateOrigin("http://google.com"))
	assert.True(t, cors.validateOrigin("https://google.com"))
	assert.True(t, cors.validateOrigin("example.com"))
	assert.True(t, cors.validateOrigin("chrome-extension://random-extension-id"))

	cors = newCors(Config{
		AllowOrigins: []string{"https://google.com", "https://github.com"},
		AllowOriginFunc: func(origin string) bool {
			return (origin == "http://news.ycombinator.com")
		},
		AllowBrowserExtensions: true,
	})
	assert.False(t, cors.validateOrigin("http://google.com"))
	assert.True(t, cors.validateOrigin("https://google.com"))
	assert.True(t, cors.validateOrigin("https://github.com"))
	assert.True(t, cors.validateOrigin("http://news.ycombinator.com"))
	assert.False(t, cors.validateOrigin("http://example.com"))
	assert.False(t, cors.validateOrigin("google.com"))
	assert.False(t, cors.validateOrigin("chrome-extension://random-extension-id"))

	cors = newCors(Config{
		AllowOrigins: []string{"https://google.com", "https://github.com"},
	})
	assert.False(t, cors.validateOrigin("chrome-extension://random-extension-id"))
	assert.False(t, cors.validateOrigin("file://some-dangerous-file.js"))
	assert.False(t, cors.validateOrigin("wss://socket-connection"))

	cors = newCors(Config{
		AllowOrigins: []string{
			"chrome-extension://*",
			"safari-extension://my-extension-*-app",
			"*.some-domain.com",
		},
		AllowBrowserExtensions: true,
		AllowWildcard:          true,
	})
	assert.True(t, cors.validateOrigin("chrome-extension://random-extension-id"))
	assert.True(t, cors.validateOrigin("chrome-extension://another-one"))
	assert.True(t, cors.validateOrigin("safari-extension://my-extension-one-app"))
	assert.True(t, cors.validateOrigin("safari-extension://my-extension-two-app"))
	assert.False(t, cors.validateOrigin("moz-extension://ext-id-we-not-allow"))
	assert.True(t, cors.validateOrigin("http://api.some-domain.com"))
	assert.False(t, cors.validateOrigin("http://api.another-domain.com"))

	cors = newCors(Config{
		AllowOrigins:    []string{"file://safe-file.js", "wss://some-session-layer-connection"},
		AllowFiles:      true,
		AllowWebSockets: true,
	})
	assert.True(t, cors.validateOrigin("file://safe-file.js"))
	assert.False(t, cors.validateOrigin("file://some-dangerous-file.js"))
	assert.True(t, cors.validateOrigin("wss://some-session-layer-connection"))
	assert.False(t, cors.validateOrigin("ws://not-what-we-expected"))

	cors = newCors(Config{
		AllowOrigins: []string{"*"},
	})
	assert.True(t, cors.validateOrigin("http://google.com"))
	assert.True(t, cors.validateOrigin("https://google.com"))
	assert.True(t, cors.validateOrigin("example.com"))
	assert.True(t, cors.validateOrigin("chrome-extension://random-extension-id"))
}


// ff:
// t:

// ff:
// t:
func TestPassesAllowOrigins(t *testing.T) {
	router := newTestRouter(Config{
		AllowOrigins:     []string{"http://google.com"},
		AllowMethods:     []string{" GeT ", "get", "post", "PUT  ", "Head", "POST"},
		AllowHeaders:     []string{"Content-type", "timeStamp "},
		ExposeHeaders:    []string{"Data", "x-User"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://github.com"
		},
	})

// 没有CORS请求，origin为空字符串
	w := performRequest(router, "GET", "")
	assert.Equal(t, "get", w.Body.String())
	assert.Empty(t, w.Header().Get("Access-Control-Allow-Origin"))
	assert.Empty(t, w.Header().Get("Access-Control-Allow-Credentials"))
	assert.Empty(t, w.Header().Get("Access-Control-Expose-Headers"))

// 没有CORS请求，origin（来源）等于host（主机）
	h := http.Header{}
	h.Set("Host", "facebook.com")
	w = performRequestWithHeaders(router, "GET", "http://facebook.com", h)
	assert.Equal(t, "get", w.Body.String())
	assert.Empty(t, w.Header().Get("Access-Control-Allow-Origin"))
	assert.Empty(t, w.Header().Get("Access-Control-Allow-Credentials"))
	assert.Empty(t, w.Header().Get("Access-Control-Expose-Headers"))

// 允许跨源资源共享（CORS）请求
	w = performRequest(router, "GET", "http://google.com")
	assert.Equal(t, "get", w.Body.String())
	assert.Equal(t, "http://google.com", w.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "", w.Header().Get("Access-Control-Allow-Credentials"))
	assert.Equal(t, "Data,X-User", w.Header().Get("Access-Control-Expose-Headers"))

	w = performRequest(router, "GET", "http://github.com")
	assert.Equal(t, "get", w.Body.String())
	assert.Equal(t, "http://github.com", w.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "", w.Header().Get("Access-Control-Allow-Credentials"))
	assert.Equal(t, "Data,X-User", w.Header().Get("Access-Control-Expose-Headers"))

// 拒绝CORS请求
	w = performRequest(router, "GET", "https://google.com")
	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.Empty(t, w.Header().Get("Access-Control-Allow-Origin"))
	assert.Empty(t, w.Header().Get("Access-Control-Allow-Credentials"))
	assert.Empty(t, w.Header().Get("Access-Control-Expose-Headers"))

// 允许CORS预检请求
	w = performRequest(router, "OPTIONS", "http://github.com")
	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Equal(t, "http://github.com", w.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "", w.Header().Get("Access-Control-Allow-Credentials"))
	assert.Equal(t, "GET,POST,PUT,HEAD", w.Header().Get("Access-Control-Allow-Methods"))
	assert.Equal(t, "Content-Type,Timestamp", w.Header().Get("Access-Control-Allow-Headers"))
	assert.Equal(t, "43200", w.Header().Get("Access-Control-Max-Age"))

// 拒绝CORS预检请求
	w = performRequest(router, "OPTIONS", "http://example.com")
	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.Empty(t, w.Header().Get("Access-Control-Allow-Origin"))
	assert.Empty(t, w.Header().Get("Access-Control-Allow-Credentials"))
	assert.Empty(t, w.Header().Get("Access-Control-Allow-Methods"))
	assert.Empty(t, w.Header().Get("Access-Control-Allow-Headers"))
	assert.Empty(t, w.Header().Get("Access-Control-Max-Age"))
}


// ff:
// t:

// ff:
// t:
func TestPassesAllowAllOrigins(t *testing.T) {
	router := newTestRouter(Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{" Patch ", "get", "post", "POST"},
		AllowHeaders:     []string{"Content-type", "  testheader "},
		ExposeHeaders:    []string{"Data2", "x-User2"},
		AllowCredentials: false,
		MaxAge:           10 * time.Hour,
	})

// 没有CORS请求，origin为空字符串
	w := performRequest(router, "GET", "")
	assert.Equal(t, "get", w.Body.String())
	assert.Empty(t, w.Header().Get("Access-Control-Allow-Origin"))
	assert.Empty(t, w.Header().Get("Access-Control-Allow-Credentials"))
	assert.Empty(t, w.Header().Get("Access-Control-Expose-Headers"))
	assert.Empty(t, w.Header().Get("Access-Control-Allow-Credentials"))

// 允许跨源资源共享（CORS）请求
	w = performRequest(router, "POST", "example.com")
	assert.Equal(t, "post", w.Body.String())
	assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "Data2,X-User2", w.Header().Get("Access-Control-Expose-Headers"))
	assert.Empty(t, w.Header().Get("Access-Control-Allow-Credentials"))
	assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))

// 允许CORS预检请求
	w = performRequest(router, "OPTIONS", "https://facebook.com")
	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "PATCH,GET,POST", w.Header().Get("Access-Control-Allow-Methods"))
	assert.Equal(t, "Content-Type,Testheader", w.Header().Get("Access-Control-Allow-Headers"))
	assert.Equal(t, "36000", w.Header().Get("Access-Control-Max-Age"))
	assert.Empty(t, w.Header().Get("Access-Control-Allow-Credentials"))
}


// ff:
// t:

// ff:
// t:
func TestWildcard(t *testing.T) {
	router := newTestRouter(Config{
		AllowOrigins:  []string{"https://*.github.com", "https://api.*", "http://*", "https://facebook.com", "*.golang.org"},
		AllowMethods:  []string{"GET"},
		AllowWildcard: true,
	})

	w := performRequest(router, "GET", "https://gist.github.com")
	assert.Equal(t, 200, w.Code)

	w = performRequest(router, "GET", "https://api.github.com/v1/users")
	assert.Equal(t, 200, w.Code)

	w = performRequest(router, "GET", "https://giphy.com/")
	assert.Equal(t, 403, w.Code)

	w = performRequest(router, "GET", "http://hard-to-find-http-example.com")
	assert.Equal(t, 200, w.Code)

	w = performRequest(router, "GET", "https://facebook.com")
	assert.Equal(t, 200, w.Code)

	w = performRequest(router, "GET", "https://something.golang.org")
	assert.Equal(t, 200, w.Code)

	w = performRequest(router, "GET", "https://something.go.org")
	assert.Equal(t, 403, w.Code)

	router = newTestRouter(Config{
		AllowOrigins: []string{"https://github.com", "https://facebook.com"},
		AllowMethods: []string{"GET"},
	})

	w = performRequest(router, "GET", "https://gist.github.com")
	assert.Equal(t, 403, w.Code)

	w = performRequest(router, "GET", "https://github.com")
	assert.Equal(t, 200, w.Code)
}
