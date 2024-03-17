package secure

import (
	"net/http"
	"net/http/httptest"
	"testing"
	
	"github.com/888go/gin"
	"github.com/stretchr/testify/assert"
)

const (
	testResponse = "bar"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func newServer(options Config) *gin.Engine {
	router := gin.New()
	router.Use(New(options))
	router.GET("/foo", func(c *gin.Context) {
		c.String(200, testResponse)
	})
	return router
}

func performRequest(router *gin.Engine, path string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", path, nil)
	router.ServeHTTP(w, req)
	return w
}


// ff:
// t:

// ff:
// t:
func TestNoConfig(t *testing.T) {
	router := newServer(Config{
// 故意留空。
	})

	w := performRequest(router, "http://example.com/foo")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "bar", w.Body.String())
}


// ff:
// t:

// ff:
// t:
func TestDefaultConfig(t *testing.T) {
	router := newServer(DefaultConfig())

	w := performRequest(router, "https://www.example.com/foo")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "bar", w.Body.String())

	w = performRequest(router, "http://www.example.com/foo")

	assert.Equal(t, http.StatusMovedPermanently, w.Code)
	assert.Equal(t, "https://www.example.com/foo", w.Header().Get("Location"))
}


// ff:
// t:

// ff:
// t:
func TestNoAllowHosts(t *testing.T) {
	router := newServer(Config{
		AllowedHosts: []string{},
	})

	w := performRequest(router, "http://www.example.com/foo")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "bar", w.Body.String())
}


// ff:
// t:

// ff:
// t:
func TestGoodSingleAllowHosts(t *testing.T) {
	router := newServer(Config{
		AllowedHosts: []string{"www.example.com"},
	})

	w := performRequest(router, "http://www.example.com/foo")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "bar", w.Body.String())
}


// ff:
// t:

// ff:
// t:
func TestBadSingleAllowHosts(t *testing.T) {
	router := newServer(Config{
		AllowedHosts: []string{"sub.example.com"},
	})

	w := performRequest(router, "http://www.example.com/foo")

	assert.Equal(t, http.StatusForbidden, w.Code)
}


// ff:
// t:

// ff:
// t:
func TestGoodMultipleAllowHosts(t *testing.T) {
	router := newServer(Config{
		AllowedHosts: []string{"www.example.com", "sub.example.com"},
	})

	w := performRequest(router, "http://sub.example.com/foo")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "bar", w.Body.String())
}


// ff:
// t:

// ff:
// t:
func TestBadMultipleAllowHosts(t *testing.T) {
	router := newServer(Config{
		AllowedHosts: []string{"www.example.com", "sub.example.com"},
	})

	w := performRequest(router, "http://www3.example.com/foo")

	assert.Equal(t, http.StatusForbidden, w.Code)
}

// ff:
// t:

// ff:
// t:
func TestAllowHostsInDevMode(t *testing.T) {
	router := newServer(Config{
		AllowedHosts:  []string{"www.example.com", "sub.example.com"},
		IsDevelopment: true,
	})

	w := performRequest(router, "http://www3.example.com/foo")

	assert.Equal(t, http.StatusOK, w.Code)
}


// ff:
// t:

// ff:
// t:
func TestBadHostHandler(t *testing.T) {

	badHandler := func(c *gin.Context) {
		c.String(http.StatusInternalServerError, "BadHost")
		c.Abort()
	}

	router := newServer(Config{
		AllowedHosts:   []string{"www.example.com", "sub.example.com"},
		BadHostHandler: badHandler,
	})

	w := performRequest(router, "http://www3.example.com/foo")

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "BadHost", w.Body.String())
}


// ff:
// t:

// ff:
// t:
func TestSSL(t *testing.T) {
	router := newServer(Config{
		SSLRedirect: true,
	})

	w := performRequest(router, "https://www.example.com/foo")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "bar", w.Body.String())
}


// ff:
// t:

// ff:
// t:
func TestSSLInDevMode(t *testing.T) {
	router := newServer(Config{
		SSLRedirect:   true,
		IsDevelopment: true,
	})

	w := performRequest(router, "http://www.example.com/foo")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "bar", w.Body.String())
}


// ff:
// t:

// ff:
// t:
func TestBasicSSL(t *testing.T) {
	router := newServer(Config{
		SSLRedirect: true,
	})

	w := performRequest(router, "http://www.example.com/foo")

	assert.Equal(t, http.StatusMovedPermanently, w.Code)
	assert.Equal(t, "https://www.example.com/foo", w.Header().Get("Location"))
}


// ff:
// t:

// ff:
// t:
func TestDontRedirectIPV4Hostnames(t *testing.T) {
	router := newServer(Config{
		SSLRedirect: true,
		DontRedirectIPV4Hostnames: true,
	})

	w1 := performRequest(router, "http://www.example.com/foo")
	assert.Equal(t, http.StatusMovedPermanently, w1.Code)

	w2 := performRequest(router, "http://127.0.0.1/foo")
	assert.Equal(t, http.StatusOK, w2.Code)
}


// ff:
// t:

// ff:
// t:
func TestBasicSSLWithHost(t *testing.T) {
	router := newServer(Config{
		SSLRedirect: true,
		SSLHost:     "secure.example.com",
	})

	w := performRequest(router, "http://www.example.com/foo")

	assert.Equal(t, http.StatusMovedPermanently, w.Code)
	assert.Equal(t, "https://secure.example.com/foo", w.Header().Get("Location"))
}


// ff:
// t:

// ff:
// t:
func TestBadProxySSL(t *testing.T) {
	router := newServer(Config{
		SSLRedirect: true,
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/foo", nil)
	req.Host = "www.example.com"
	req.URL.Scheme = "http"
	req.Header.Add("X-Forwarded-Proto", "https")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusMovedPermanently, w.Code)
	assert.Equal(t, "https://www.example.com/foo", w.Header().Get("Location"))
}


// ff:
// t:

// ff:
// t:
func TestProxySSLWithHeaderOption(t *testing.T) {
	router := newServer(Config{
		SSLRedirect:     true,
		SSLProxyHeaders: map[string]string{"X-Arbitrary-Header": "arbitrary-value"},
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/foo", nil)
	req.Host = "www.example.com"
	req.URL.Scheme = "http"
	req.Header.Add("X-Arbitrary-Header", "arbitrary-value")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}


// ff:
// t:

// ff:
// t:
func TestProxySSLWithWrongHeaderValue(t *testing.T) {
	router := newServer(Config{
		SSLRedirect:     true,
		SSLProxyHeaders: map[string]string{"X-Arbitrary-Header": "arbitrary-value"},
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/foo", nil)
	req.Host = "www.example.com"
	req.URL.Scheme = "http"
	req.Header.Add("X-Arbitrary-Header", "wrong-value")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusMovedPermanently, w.Code)
	assert.Equal(t, "https://www.example.com/foo", w.Header().Get("Location"))
}


// ff:
// t:

// ff:
// t:
func TestStsHeader(t *testing.T) {
	router := newServer(Config{
		STSSeconds: 315360000,
	})

	w := performRequest(router, "/foo")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "max-age=315360000", w.Header().Get("Strict-Transport-Security"))
}


// ff:
// t:

// ff:
// t:
func TestStsHeaderInDevMode(t *testing.T) {
	router := newServer(Config{
		STSSeconds:    315360000,
		IsDevelopment: true,
	})

	w := performRequest(router, "/foo")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "", w.Header().Get("Strict-Transport-Security"))
}


// ff:
// t:

// ff:
// t:
func TestStsHeaderWithSubdomain(t *testing.T) {
	router := newServer(Config{
		STSSeconds:           315360000,
		STSIncludeSubdomains: true,
	})

	w := performRequest(router, "/foo")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "max-age=315360000; includeSubdomains", w.Header().Get("Strict-Transport-Security"))
}


// ff:
// t:

// ff:
// t:
func TestFrameDeny(t *testing.T) {
	router := newServer(Config{
		FrameDeny: true,
	})

	w := performRequest(router, "/foo")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "DENY", w.Header().Get("X-Frame-Options"))
}


// ff:
// t:

// ff:
// t:
func TestCustomFrameValue(t *testing.T) {
	router := newServer(Config{
		CustomFrameOptionsValue: "SAMEORIGIN",
	})

	w := performRequest(router, "/foo")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "SAMEORIGIN", w.Header().Get("X-Frame-Options"))
}


// ff:
// t:

// ff:
// t:
func TestCustomFrameValueWithDeny(t *testing.T) {
	router := newServer(Config{
		FrameDeny:               true,
		CustomFrameOptionsValue: "SAMEORIGIN",
	})

	w := performRequest(router, "/foo")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "SAMEORIGIN", w.Header().Get("X-Frame-Options"))
}


// ff:
// t:

// ff:
// t:
func TestContentNosniff(t *testing.T) {
	router := newServer(Config{
		ContentTypeNosniff: true,
	})

	w := performRequest(router, "/foo")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "nosniff", w.Header().Get("X-Content-Type-Options"))
}


// ff:
// t:

// ff:
// t:
func TestXSSProtection(t *testing.T) {
	router := newServer(Config{
		BrowserXssFilter: true,
	})

	w := performRequest(router, "/foo")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "1; mode=block", w.Header().Get("X-XSS-Protection"))
}


// ff:
// t:

// ff:
// t:
func TestReferrerPolicy(t *testing.T) {
	router := newServer(Config{
		ReferrerPolicy: "strict-origin-when-cross-origin",
	})

	w := performRequest(router, "/foo")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "strict-origin-when-cross-origin", w.Header().Get("Referrer-Policy"))
}


// ff:
// t:

// ff:
// t:
func TestFeaturePolicy(t *testing.T) {
	router := newServer(Config{
		FeaturePolicy: "vibrate 'none';",
	})

	w := performRequest(router, "/foo")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "vibrate 'none';", w.Header().Get("Feature-Policy"))
}


// ff:
// t:

// ff:
// t:
func TestCsp(t *testing.T) {
	router := newServer(Config{
		ContentSecurityPolicy: "default-src 'self'",
	})

	w := performRequest(router, "/foo")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "default-src 'self'", w.Header().Get("Content-Security-Policy"))
}


// ff:
// t:

// ff:
// t:
func TestInlineSecure(t *testing.T) {
	router := newServer(Config{
		FrameDeny: true,
	})

	w := performRequest(router, "/foo")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "DENY", w.Header().Get("X-Frame-Options"))
}


// ff:
// t:

// ff:
// t:
func TestIsIpv4Host(t *testing.T) {
	assert.Equal(t, isIPV4("127.0.0.1"), true)
	assert.Equal(t, isIPV4("127.0.0.1:8080"), true)
	assert.Equal(t, isIPV4("localhost"), false)
	assert.Equal(t, isIPV4("localhost:8080"), false)
	assert.Equal(t, isIPV4("example.com"), false)
	assert.Equal(t, isIPV4("example.com:8080"), false)
}
