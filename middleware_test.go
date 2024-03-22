// 版权所有 2014 Manu Martinez-Almeida。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package gin类

import (
	"errors"
	"net/http"
	"strings"
	"testing"
	
	"github.com/gin-contrib/sse"
	"github.com/stretchr/testify/assert"
)

func TestMiddlewareGeneralCase(t *testing.T) {
	signature := ""
	router := X创建()
	router.X中间件(func(c *Context) {
		signature += "A"
		c.X中间件继续()
		signature += "B"
	})
	router.X中间件(func(c *Context) {
		signature += "C"
	})
	router.X绑定GET("/", func(c *Context) {
		signature += "D"
	})
	router.X绑定404(func(c *Context) {
		signature += " X "
	})
	router.X绑定405(func(c *Context) {
		signature += " XX "
	})
	// RUN
	w := PerformRequest(router, "GET", "/")

	// TEST
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "ACDB", signature)
}

func TestMiddlewareNoRoute(t *testing.T) {
	signature := ""
	router := X创建()
	router.X中间件(func(c *Context) {
		signature += "A"
		c.X中间件继续()
		signature += "B"
	})
	router.X中间件(func(c *Context) {
		signature += "C"
		c.X中间件继续()
		c.X中间件继续()
		c.X中间件继续()
		c.X中间件继续()
		signature += "D"
	})
	router.X绑定404(func(c *Context) {
		signature += "E"
		c.X中间件继续()
		signature += "F"
	}, func(c *Context) {
		signature += "G"
		c.X中间件继续()
		signature += "H"
	})
	router.X绑定405(func(c *Context) {
		signature += " X "
	})
	// RUN
	w := PerformRequest(router, "GET", "/")

	// TEST
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "ACEGHFDB", signature)
}

func TestMiddlewareNoMethodEnabled(t *testing.T) {
	signature := ""
	router := X创建()
	router.HandleMethodNotAllowed = true
	router.X中间件(func(c *Context) {
		signature += "A"
		c.X中间件继续()
		signature += "B"
	})
	router.X中间件(func(c *Context) {
		signature += "C"
		c.X中间件继续()
		signature += "D"
	})
	router.X绑定405(func(c *Context) {
		signature += "E"
		c.X中间件继续()
		signature += "F"
	}, func(c *Context) {
		signature += "G"
		c.X中间件继续()
		signature += "H"
	})
	router.X绑定404(func(c *Context) {
		signature += " X "
	})
	router.X绑定POST("/", func(c *Context) {
		signature += " XX "
	})
	// RUN
	w := PerformRequest(router, "GET", "/")

	// TEST
	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
	assert.Equal(t, "ACEGHFDB", signature)
}

func TestMiddlewareNoMethodDisabled(t *testing.T) {
	signature := ""
	router := X创建()

	// NoMethod disabled
	router.HandleMethodNotAllowed = false

	router.X中间件(func(c *Context) {
		signature += "A"
		c.X中间件继续()
		signature += "B"
	})
	router.X中间件(func(c *Context) {
		signature += "C"
		c.X中间件继续()
		signature += "D"
	})
	router.X绑定405(func(c *Context) {
		signature += "E"
		c.X中间件继续()
		signature += "F"
	}, func(c *Context) {
		signature += "G"
		c.X中间件继续()
		signature += "H"
	})
	router.X绑定404(func(c *Context) {
		signature += " X "
	})
	router.X绑定POST("/", func(c *Context) {
		signature += " XX "
	})

	// RUN
	w := PerformRequest(router, "GET", "/")

	// TEST
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "AC X DB", signature)
}

func TestMiddlewareAbort(t *testing.T) {
	signature := ""
	router := X创建()
	router.X中间件(func(c *Context) {
		signature += "A"
	})
	router.X中间件(func(c *Context) {
		signature += "C"
		c.X停止并带状态码(http.StatusUnauthorized)
		c.X中间件继续()
		signature += "D"
	})
	router.X绑定GET("/", func(c *Context) {
		signature += " X "
		c.X中间件继续()
		signature += " XX "
	})

	// RUN
	w := PerformRequest(router, "GET", "/")

	// TEST
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, "ACD", signature)
}

func TestMiddlewareAbortHandlersChainAndNext(t *testing.T) {
	signature := ""
	router := X创建()
	router.X中间件(func(c *Context) {
		signature += "A"
		c.X中间件继续()
		c.X停止并带状态码(http.StatusGone)
		signature += "B"
	})
	router.X绑定GET("/", func(c *Context) {
		signature += "C"
		c.X中间件继续()
	})
	// RUN
	w := PerformRequest(router, "GET", "/")

	// TEST
	assert.Equal(t, http.StatusGone, w.Code)
	assert.Equal(t, "ACB", signature)
}

// TestFailHandlersChain - 确保 Fail 中断会按照先进先出（FIFO）顺序使用中间件，
// 同时也会中断执行链
func TestMiddlewareFailHandlersChain(t *testing.T) {
	// SETUP
	signature := ""
	router := X创建()
	router.X中间件(func(context *Context) {
		signature += "A"
		context.X停止并带状态码与错误(http.StatusInternalServerError, errors.New("foo")) //nolint: errcheck
	})
	router.X中间件(func(context *Context) {
		signature += "B"
		context.X中间件继续()
		signature += "C"
	})
	// RUN
	w := PerformRequest(router, "GET", "/")

	// TEST
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "A", signature)
}

func TestMiddlewareWrite(t *testing.T) {
	router := X创建()
	router.X中间件(func(c *Context) {
		c.X输出文本(http.StatusBadRequest, "hola\n")
	})
	router.X中间件(func(c *Context) {
		c.X输出XML(http.StatusBadRequest, H{"foo": "bar"})
	})
	router.X中间件(func(c *Context) {
		c.X输出JSON(http.StatusBadRequest, H{"foo": "bar"})
	})
	router.X绑定GET("/", func(c *Context) {
		c.X输出JSON(http.StatusBadRequest, H{"foo": "bar"})
	}, func(c *Context) {
		c.Render底层方法(http.StatusBadRequest, sse.Event{
			Event: "test",
			Data:  "message",
		})
	})

	w := PerformRequest(router, "GET", "/")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, strings.Replace("hola\n<map><foo>bar</foo></map>{\"foo\":\"bar\"}{\"foo\":\"bar\"}event:test\ndata:message\n\n", " ", "", -1), strings.Replace(w.Body.String(), " ", "", -1))
}
