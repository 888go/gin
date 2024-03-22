package requestid

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	
	"github.com/888go/gin"
	"github.com/stretchr/testify/assert"
)

const testXRequestID = "test-request-id"

func emptySuccessResponse(c *gin类.Context) {
	c.X输出文本(http.StatusOK, "")
}

func Test_RequestID_CreateNew(t *testing.T) {
	r := gin类.X创建()
	r.X中间件(New())
	r.X绑定GET("/", emptySuccessResponse)

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Header().Get(headerXRequestID))
}

func Test_RequestID_PassThru(t *testing.T) {
	r := gin类.X创建()
	r.X中间件(New())
	r.X绑定GET("/", emptySuccessResponse)

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	req.Header.Set(headerXRequestID, testXRequestID)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, testXRequestID, w.Header().Get(headerXRequestID))
}

func TestRequestIDWithCustomID(t *testing.T) {
	r := gin类.X创建()
	r.X中间件(
		New(
			WithGenerator(func() string {
				return testXRequestID
			}),
		),
	)
	r.X绑定GET("/", emptySuccessResponse)

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, testXRequestID, w.Header().Get(headerXRequestID))
}

func TestRequestIDWithCustomHeaderKey(t *testing.T) {
	r := gin类.X创建()
	r.X中间件(
		New(
			WithCustomHeaderStrKey("customKey"),
		),
	)
	r.X绑定GET("/", emptySuccessResponse)

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	req.Header.Set("customKey", testXRequestID)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, testXRequestID, w.Header().Get("customKey"))
}

func TestRequestIDWithHandler(t *testing.T) {
	r := gin类.X创建()
	called := false
	r.X中间件(
		New(
			WithHandler(func(c *gin类.Context, requestID string) {
				called = true
				assert.Equal(t, testXRequestID, requestID)
			}),
		),
	)

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	req.Header.Set("X-Request-ID", testXRequestID)
	r.ServeHTTP(w, req)

	assert.True(t, called)
}

func TestRequestIDIsAttachedToRequestHeaders(t *testing.T) {
	r := gin类.X创建()

	r.X中间件(New())

	r.X绑定GET("/", func(c *gin类.Context) {
		result := c.X取请求协议头值("X-Request-ID")
		assert.NotEmpty(t, result)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	r.ServeHTTP(w, req)
}

func TestRequestIDNotNilAfterGinCopy(t *testing.T) {
	r := gin类.X创建()
	r.X中间件(New())

	r.X绑定GET("/", func(c *gin类.Context) {
		copy := c.X取副本()
		result := Get(copy)
		assert.NotEmpty(t, result)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	r.ServeHTTP(w, req)
}
