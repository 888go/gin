package rollbar

import (
	"errors"
	"net/http/httptest"
	"testing"
	
	"github.com/888go/gin"
	"github.com/stretchr/testify/assert"
)

func TestRecovery(t *testing.T) {
	gin类.X设置运行模式(gin类.X常量_运行模式_测试)
	router := gin类.X创建()

	router.X中间件(Recovery(false))

	router.X绑定GET("/", func(c *gin类.Context) {
		baseError := errors.New("test error")
		err := &gin类.Error{
			Err:  baseError,
			Type: gin类.ErrorTypePublic,
		}
		_ = err.SetMeta("some data")
		_ = c.X错误(err)

		panic("occurs panic")
	})

	w := performRequest("GET", "/", router)
	assert.Equal(t, 500, w.Code)
}

func performRequest(method, target string, router *gin类.Engine) *httptest.ResponseRecorder {
	r := httptest.NewRequest(method, target, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	return w
}
