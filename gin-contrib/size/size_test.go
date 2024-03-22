package limits

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	
	"github.com/888go/gin"
)

func TestRequestSizeLimiterOK(t *testing.T) {
	router := gin类.X创建()
	router.X中间件(RequestSizeLimiter(10))
	router.X绑定POST("/test_ok", func(c *gin类.Context) {
		_, _ = ioutil.ReadAll(c.X请求.Body)
		if len(c.X错误s) > 0 {
			return
		}
		c.X请求.Body.Close()
		c.X输出文本(http.StatusOK, "OK")
	})
	resp := performRequest(http.MethodPost, "/test_ok", "big=abc", router)

	if resp.Code != http.StatusOK {
		t.Fatalf("error posting - http status %v", resp.Code)
	}
}

func TestRequestSizeLimiterOver(t *testing.T) {
	router := gin类.X创建()
	router.X中间件(RequestSizeLimiter(10))
	router.X绑定POST("/test_large", func(c *gin类.Context) {
		_, _ = ioutil.ReadAll(c.X请求.Body)
		if len(c.X错误s) > 0 {
			return
		}
		c.X请求.Body.Close()
		c.X输出文本(http.StatusOK, "OK")
	})
	resp := performRequest(http.MethodPost, "/test_large", "big=abcdefghijklmnop", router)

	if resp.Code != http.StatusRequestEntityTooLarge {
		t.Fatalf("error posting - http status %v", resp.Code)
	}
}

func performRequest(method, target, body string, router *gin类.Engine) *httptest.ResponseRecorder {
	var buf *bytes.Buffer
	if body != "" {
		buf = new(bytes.Buffer)
		buf.WriteString(body)
	}
	r := httptest.NewRequest(method, target, buf)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	return w
}
