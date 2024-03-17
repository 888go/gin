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

func emptySuccessResponse(c *gin.Context) {
	c.String(http.StatusOK, "")
}


// ff:
// t:

// ff:
// t:
func Test_RequestID_CreateNew(t *testing.T) {
	r := gin.New()
	r.Use(New())
	r.GET("/", emptySuccessResponse)

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Header().Get(headerXRequestID))
}


// ff:
// t:

// ff:
// t:
func Test_RequestID_PassThru(t *testing.T) {
	r := gin.New()
	r.Use(New())
	r.GET("/", emptySuccessResponse)

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	req.Header.Set(headerXRequestID, testXRequestID)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, testXRequestID, w.Header().Get(headerXRequestID))
}


// ff:
// t:

// ff:
// t:
func TestRequestIDWithCustomID(t *testing.T) {
	r := gin.New()
	r.Use(
		New(
			WithGenerator(func() string {
				return testXRequestID
			}),
		),
	)
	r.GET("/", emptySuccessResponse)

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, testXRequestID, w.Header().Get(headerXRequestID))
}


// ff:
// t:

// ff:
// t:
func TestRequestIDWithCustomHeaderKey(t *testing.T) {
	r := gin.New()
	r.Use(
		New(
			WithCustomHeaderStrKey("customKey"),
		),
	)
	r.GET("/", emptySuccessResponse)

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	req.Header.Set("customKey", testXRequestID)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, testXRequestID, w.Header().Get("customKey"))
}


// ff:
// t:

// ff:
// t:
func TestRequestIDWithHandler(t *testing.T) {
	r := gin.New()
	called := false
	r.Use(
		New(
			WithHandler(func(c *gin.Context, requestID string) {
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


// ff:
// t:

// ff:
// t:
func TestRequestIDIsAttachedToRequestHeaders(t *testing.T) {
	r := gin.New()

	r.Use(New())

	r.GET("/", func(c *gin.Context) {
		result := c.GetHeader("X-Request-ID")
		assert.NotEmpty(t, result)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	r.ServeHTTP(w, req)
}


// ff:
// t:

// ff:
// t:
func TestRequestIDNotNilAfterGinCopy(t *testing.T) {
	r := gin.New()
	r.Use(New())

	r.GET("/", func(c *gin.Context) {
		copy := c.Copy()
		result := Get(copy)
		assert.NotEmpty(t, result)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	r.ServeHTTP(w, req)
}
