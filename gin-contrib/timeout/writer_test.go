package timeout

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
	
	"github.com/888go/gin"
	"github.com/stretchr/testify/assert"
)

func TestWriteHeader(t *testing.T) {
	code1 := 99
	errmsg1 := fmt.Sprintf("invalid http status code: %d", code1)
	code2 := 1000
	errmsg2 := fmt.Sprintf("invalid http status code: %d", code2)

	writer := Writer{}
	assert.PanicsWithValue(t, errmsg1, func() {
		writer.WriteHeader(code1)
	})
	assert.PanicsWithValue(t, errmsg2, func() {
		writer.WriteHeader(code2)
	})
}

func TestWriteHeader_SkipMinusOne(t *testing.T) {
	code := -1

	writer := Writer{}
	assert.NotPanics(t, func() {
		writer.WriteHeader(code)
		assert.False(t, writer.wroteHeaders)
	})
}

func TestWriter_Status(t *testing.T) {
	r := gin类.X创建()

	r.X中间件(New(
		WithTimeout(1*time.Second),
		WithHandler(func(c *gin类.Context) {
			c.X中间件继续()
		}),
		WithResponse(testResponse),
	))

	r.X中间件(func(c *gin类.Context) {
		c.X中间件继续()
		statusInMW := c.Writer.Status()
		c.X请求.Header.Set("X-Status-Code-MW-Set", strconv.Itoa(statusInMW))
		t.Logf("[%s] %s %s %d\n", time.Now().Format(time.RFC3339), c.X请求.Method, c.X请求.URL, statusInMW)
	})

	r.X绑定GET("/test", func(c *gin类.Context) {
		c.Writer.WriteHeader(http.StatusInternalServerError)
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, strconv.Itoa(http.StatusInternalServerError), req.Header.Get("X-Status-Code-MW-Set"))
}

// testNew is a copy of New() with a small change to the timeoutHandler() function.
// ref: https://github.com/gin-contrib/timeout/issues/31
func testNew(duration time.Duration) gin类.HandlerFunc {
	return New(
		WithTimeout(duration),
		WithHandler(func(c *gin类.Context) { c.X中间件继续() }),
		WithResponse(timeoutHandler()),
	)
}

// timeoutHandler returns a handler that returns a 504 Gateway Timeout error.
func timeoutHandler() gin类.HandlerFunc {
	gatewayTimeoutErr := struct {
		Error string `json:"error"`
	}{
		Error: "Timed out.",
	}

	return func(c *gin类.Context) {
		log.Printf("request timed out: [method=%s,path=%s]",
			c.X请求.Method, c.X请求.URL.Path)
		c.X输出JSON(http.StatusGatewayTimeout, gatewayTimeoutErr)
	}
}

// TestHTTPStatusCode tests the HTTP status code of the response.
func TestHTTPStatusCode(t *testing.T) {
	gin类.X设置运行模式(gin类.X常量_运行模式_发布)

	type testCase struct {
		Name          string
		Method        string
		Path          string
		ExpStatusCode int
		Handler       gin类.HandlerFunc
	}

	var (
		cases = []testCase{
			{
				Name:          "Plain text (200)",
				Method:        http.MethodGet,
				Path:          "/me",
				ExpStatusCode: http.StatusOK,
				Handler: func(ctx *gin类.Context) {
					ctx.X输出文本(http.StatusOK, "I'm text!")
				},
			},
			{
				Name:          "Plain text (201)",
				Method:        http.MethodGet,
				Path:          "/me",
				ExpStatusCode: http.StatusCreated,
				Handler: func(ctx *gin类.Context) {
					ctx.X输出文本(http.StatusCreated, "I'm created!")
				},
			},
			{
				Name:          "Plain text (204)",
				Method:        http.MethodGet,
				Path:          "/me",
				ExpStatusCode: http.StatusNoContent,
				Handler: func(ctx *gin类.Context) {
					ctx.X输出文本(http.StatusNoContent, "")
				},
			},
			{
				Name:          "Plain text (400)",
				Method:        http.MethodGet,
				Path:          "/me",
				ExpStatusCode: http.StatusBadRequest,
				Handler: func(ctx *gin类.Context) {
					ctx.X输出文本(http.StatusBadRequest, "")
				},
			},
			{
				Name:          "JSON (200)",
				Method:        http.MethodGet,
				Path:          "/me",
				ExpStatusCode: http.StatusOK,
				Handler: func(ctx *gin类.Context) {
					ctx.X输出JSON(http.StatusOK, gin类.H{"field": "value"})
				},
			},
			{
				Name:          "JSON (201)",
				Method:        http.MethodGet,
				Path:          "/me",
				ExpStatusCode: http.StatusCreated,
				Handler: func(ctx *gin类.Context) {
					ctx.X输出JSON(http.StatusCreated, gin类.H{"field": "value"})
				},
			},
			{
				Name:          "JSON (204)",
				Method:        http.MethodGet,
				Path:          "/me",
				ExpStatusCode: http.StatusNoContent,
				Handler: func(ctx *gin类.Context) {
					ctx.X输出JSON(http.StatusNoContent, nil)
				},
			},
			{
				Name:          "JSON (400)",
				Method:        http.MethodGet,
				Path:          "/me",
				ExpStatusCode: http.StatusBadRequest,
				Handler: func(ctx *gin类.Context) {
					ctx.X输出JSON(http.StatusBadRequest, nil)
				},
			},
			{
				Name:          "No reply",
				Method:        http.MethodGet,
				Path:          "/me",
				ExpStatusCode: http.StatusOK,
				Handler:       func(ctx *gin类.Context) {},
			},
		}

		initCase = func(c testCase) (*http.Request, *httptest.ResponseRecorder) {
			return httptest.NewRequest(c.Method, c.Path, nil), httptest.NewRecorder()
		}
	)

	for i := range cases {
		t.Run(cases[i].Name, func(tt *testing.T) {
			tt.Logf("Test case [%s]", cases[i].Name)

			router := gin类.X创建默认对象()

			router.X中间件(testNew(1 * time.Second))
			router.X绑定GET("/*root", cases[i].Handler)

			req, resp := initCase(cases[i])
			router.ServeHTTP(resp, req)

			if resp.Code != cases[i].ExpStatusCode {
				tt.Errorf("response is different from expected:\nexp: >>>%d<<<\ngot: >>>%d<<<",
					cases[i].ExpStatusCode, resp.Code)
			}
		})
	}
}
