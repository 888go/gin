package timeout

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	
	"github.com/888go/gin"
	"github.com/stretchr/testify/assert"
)

func emptySuccessResponse(c *gin类.Context) {
	time.Sleep(200 * time.Microsecond)
	c.X输出文本(http.StatusOK, "")
}

func TestTimeout(t *testing.T) {
	r := gin类.X创建()
	r.X绑定GET("/", New(WithTimeout(50*time.Microsecond), WithHandler(emptySuccessResponse)))

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusRequestTimeout, w.Code)
	assert.Equal(t, http.StatusText(http.StatusRequestTimeout), w.Body.String())
}

func TestWithoutTimeout(t *testing.T) {
	r := gin类.X创建()
	r.X绑定GET("/", New(WithTimeout(-1*time.Microsecond), WithHandler(emptySuccessResponse)))

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "", w.Body.String())
}

func testResponse(c *gin类.Context) {
	c.X输出文本(http.StatusRequestTimeout, "test response")
}

func TestCustomResponse(t *testing.T) {
	r := gin类.X创建()
	r.X绑定GET("/", New(
		WithTimeout(100*time.Microsecond),
		WithHandler(emptySuccessResponse),
		WithResponse(testResponse),
	))

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusRequestTimeout, w.Code)
	assert.Equal(t, "test response", w.Body.String())
}

func emptySuccessResponse2(c *gin类.Context) {
	time.Sleep(50 * time.Microsecond)
	c.X输出文本(http.StatusOK, "")
}

func TestSuccess(t *testing.T) {
	r := gin类.X创建()
	r.X绑定GET("/", New(
		WithTimeout(1*time.Second),
		WithHandler(emptySuccessResponse2),
		WithResponse(testResponse),
	))

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "", w.Body.String())
}

func panicResponse(c *gin类.Context) {
	panic("test")
}

func TestPanic(t *testing.T) {
	r := gin类.X创建()
	r.X中间件(gin类.Recovery())
	r.X绑定GET("/", New(
		WithTimeout(1*time.Second),
		WithHandler(panicResponse),
	))

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "", w.Body.String())
}
