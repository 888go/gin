// Manu Martinez-Almeida版权所有
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到

package gin

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"
	
	"github.com/stretchr/testify/assert"
)

func init() {
	SetMode(TestMode)
}


// ff:
// t:
func TestLogger(t *testing.T) {
	buffer := new(strings.Builder)
	router := New()
	router.Use(LoggerWithWriter(buffer))
	router.GET("/example", func(c *Context) {})
	router.POST("/example", func(c *Context) {})
	router.PUT("/example", func(c *Context) {})
	router.DELETE("/example", func(c *Context) {})
	router.PATCH("/example", func(c *Context) {})
	router.HEAD("/example", func(c *Context) {})
	router.OPTIONS("/example", func(c *Context) {})

	PerformRequest(router, "GET", "/example?a=100")
	assert.Contains(t, buffer.String(), "200")
	assert.Contains(t, buffer.String(), "GET")
	assert.Contains(t, buffer.String(), "/example")
	assert.Contains(t, buffer.String(), "a=100")

// 我先写了这些(扩展了上面的内容)，但后来意识到它们更像是集成测试，因为它们测试的是整个日志记录过程，而不是单个功能
// 我不确定这些应该放在哪里
	buffer.Reset()
	PerformRequest(router, "POST", "/example")
	assert.Contains(t, buffer.String(), "200")
	assert.Contains(t, buffer.String(), "POST")
	assert.Contains(t, buffer.String(), "/example")

	buffer.Reset()
	PerformRequest(router, "PUT", "/example")
	assert.Contains(t, buffer.String(), "200")
	assert.Contains(t, buffer.String(), "PUT")
	assert.Contains(t, buffer.String(), "/example")

	buffer.Reset()
	PerformRequest(router, "DELETE", "/example")
	assert.Contains(t, buffer.String(), "200")
	assert.Contains(t, buffer.String(), "DELETE")
	assert.Contains(t, buffer.String(), "/example")

	buffer.Reset()
	PerformRequest(router, "PATCH", "/example")
	assert.Contains(t, buffer.String(), "200")
	assert.Contains(t, buffer.String(), "PATCH")
	assert.Contains(t, buffer.String(), "/example")

	buffer.Reset()
	PerformRequest(router, "HEAD", "/example")
	assert.Contains(t, buffer.String(), "200")
	assert.Contains(t, buffer.String(), "HEAD")
	assert.Contains(t, buffer.String(), "/example")

	buffer.Reset()
	PerformRequest(router, "OPTIONS", "/example")
	assert.Contains(t, buffer.String(), "200")
	assert.Contains(t, buffer.String(), "OPTIONS")
	assert.Contains(t, buffer.String(), "/example")

	buffer.Reset()
	PerformRequest(router, "GET", "/notfound")
	assert.Contains(t, buffer.String(), "404")
	assert.Contains(t, buffer.String(), "GET")
	assert.Contains(t, buffer.String(), "/notfound")
}


// ff:
// t:
func TestLoggerWithConfig(t *testing.T) {
	buffer := new(strings.Builder)
	router := New()
	router.Use(LoggerWithConfig(LoggerConfig{Output: buffer}))
	router.GET("/example", func(c *Context) {})
	router.POST("/example", func(c *Context) {})
	router.PUT("/example", func(c *Context) {})
	router.DELETE("/example", func(c *Context) {})
	router.PATCH("/example", func(c *Context) {})
	router.HEAD("/example", func(c *Context) {})
	router.OPTIONS("/example", func(c *Context) {})

	PerformRequest(router, "GET", "/example?a=100")
	assert.Contains(t, buffer.String(), "200")
	assert.Contains(t, buffer.String(), "GET")
	assert.Contains(t, buffer.String(), "/example")
	assert.Contains(t, buffer.String(), "a=100")

// 我先写了这些(扩展了上面的内容)，但后来意识到它们更像是集成测试，因为它们测试的是整个日志记录过程，而不是单个功能
// 我不确定这些应该放在哪里
	buffer.Reset()
	PerformRequest(router, "POST", "/example")
	assert.Contains(t, buffer.String(), "200")
	assert.Contains(t, buffer.String(), "POST")
	assert.Contains(t, buffer.String(), "/example")

	buffer.Reset()
	PerformRequest(router, "PUT", "/example")
	assert.Contains(t, buffer.String(), "200")
	assert.Contains(t, buffer.String(), "PUT")
	assert.Contains(t, buffer.String(), "/example")

	buffer.Reset()
	PerformRequest(router, "DELETE", "/example")
	assert.Contains(t, buffer.String(), "200")
	assert.Contains(t, buffer.String(), "DELETE")
	assert.Contains(t, buffer.String(), "/example")

	buffer.Reset()
	PerformRequest(router, "PATCH", "/example")
	assert.Contains(t, buffer.String(), "200")
	assert.Contains(t, buffer.String(), "PATCH")
	assert.Contains(t, buffer.String(), "/example")

	buffer.Reset()
	PerformRequest(router, "HEAD", "/example")
	assert.Contains(t, buffer.String(), "200")
	assert.Contains(t, buffer.String(), "HEAD")
	assert.Contains(t, buffer.String(), "/example")

	buffer.Reset()
	PerformRequest(router, "OPTIONS", "/example")
	assert.Contains(t, buffer.String(), "200")
	assert.Contains(t, buffer.String(), "OPTIONS")
	assert.Contains(t, buffer.String(), "/example")

	buffer.Reset()
	PerformRequest(router, "GET", "/notfound")
	assert.Contains(t, buffer.String(), "404")
	assert.Contains(t, buffer.String(), "GET")
	assert.Contains(t, buffer.String(), "/notfound")
}


// ff:
// t:
func TestLoggerWithFormatter(t *testing.T) {
	buffer := new(strings.Builder)

	d := DefaultWriter
	DefaultWriter = buffer
	defer func() {
		DefaultWriter = d
	}()

	router := New()
	router.Use(LoggerWithFormatter(func(param LogFormatterParams) string {
		return fmt.Sprintf("[FORMATTER TEST] %v | %3d | %13v | %15s | %-7s %#v\n%s",
			param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.Method,
			param.Path,
			param.ErrorMessage,
		)
	}))
	router.GET("/example", func(c *Context) {})
	PerformRequest(router, "GET", "/example?a=100")

// 输出测试
	assert.Contains(t, buffer.String(), "[FORMATTER TEST]")
	assert.Contains(t, buffer.String(), "200")
	assert.Contains(t, buffer.String(), "GET")
	assert.Contains(t, buffer.String(), "/example")
	assert.Contains(t, buffer.String(), "a=100")
}


// ff:
// t:
func TestLoggerWithConfigFormatting(t *testing.T) {
	var gotParam LogFormatterParams
	var gotKeys map[string]any
	buffer := new(strings.Builder)

	router := New()
	router.engine.trustedCIDRs, _ = router.engine.prepareTrustedCIDRs()

	router.Use(LoggerWithConfig(LoggerConfig{
		Output: buffer,
		Formatter: func(param LogFormatterParams) string {
// 对于断言测试
			gotParam = param

			return fmt.Sprintf("[FORMATTER TEST] %v | %3d | %13v | %15s | %-7s %s\n%s",
				param.TimeStamp.Format("2006/01/02 - 15:04:05"),
				param.StatusCode,
				param.Latency,
				param.ClientIP,
				param.Method,
				param.Path,
				param.ErrorMessage,
			)
		},
	}))
	router.GET("/example", func(c *Context) {
// 设置dummy ClientIP
		c.Request.Header.Set("X-Forwarded-For", "20.20.20.20")
		gotKeys = c.Keys
		time.Sleep(time.Millisecond)
	})
	PerformRequest(router, "GET", "/example?a=100")

// 输出测试
	assert.Contains(t, buffer.String(), "[FORMATTER TEST]")
	assert.Contains(t, buffer.String(), "200")
	assert.Contains(t, buffer.String(), "GET")
	assert.Contains(t, buffer.String(), "/example")
	assert.Contains(t, buffer.String(), "a=100")

// LogFormatterParams测试
	assert.NotNil(t, gotParam.Request)
	assert.NotEmpty(t, gotParam.TimeStamp)
	assert.Equal(t, 200, gotParam.StatusCode)
	assert.NotEmpty(t, gotParam.Latency)
	assert.Equal(t, "20.20.20.20", gotParam.ClientIP)
	assert.Equal(t, "GET", gotParam.Method)
	assert.Equal(t, "/example?a=100", gotParam.Path)
	assert.Empty(t, gotParam.ErrorMessage)
	assert.Equal(t, gotKeys, gotParam.Keys)
}


// ff:
// t:
func TestDefaultLogFormatter(t *testing.T) {
	timeStamp := time.Unix(1544173902, 0).UTC()

	termFalseParam := LogFormatterParams{
		TimeStamp:    timeStamp,
		StatusCode:   200,
		Latency:      time.Second * 5,
		ClientIP:     "20.20.20.20",
		Method:       "GET",
		Path:         "/",
		ErrorMessage: "",
		isTerm:       false,
	}

	termTrueParam := LogFormatterParams{
		TimeStamp:    timeStamp,
		StatusCode:   200,
		Latency:      time.Second * 5,
		ClientIP:     "20.20.20.20",
		Method:       "GET",
		Path:         "/",
		ErrorMessage: "",
		isTerm:       true,
	}
	termTrueLongDurationParam := LogFormatterParams{
		TimeStamp:    timeStamp,
		StatusCode:   200,
		Latency:      time.Millisecond * 9876543210,
		ClientIP:     "20.20.20.20",
		Method:       "GET",
		Path:         "/",
		ErrorMessage: "",
		isTerm:       true,
	}

	termFalseLongDurationParam := LogFormatterParams{
		TimeStamp:    timeStamp,
		StatusCode:   200,
		Latency:      time.Millisecond * 9876543210,
		ClientIP:     "20.20.20.20",
		Method:       "GET",
		Path:         "/",
		ErrorMessage: "",
		isTerm:       false,
	}

	assert.Equal(t, "[GIN] 2018/12/07 - 09:11:42 | 200 |            5s |     20.20.20.20 | GET      \"/\"\n", defaultLogFormatter(termFalseParam))
	assert.Equal(t, "[GIN] 2018/12/07 - 09:11:42 | 200 |    2743h29m3s |     20.20.20.20 | GET      \"/\"\n", defaultLogFormatter(termFalseLongDurationParam))

	assert.Equal(t, "[GIN] 2018/12/07 - 09:11:42 |\x1b[97;42m 200 \x1b[0m|            5s |     20.20.20.20 |\x1b[97;44m GET     \x1b[0m \"/\"\n", defaultLogFormatter(termTrueParam))
	assert.Equal(t, "[GIN] 2018/12/07 - 09:11:42 |\x1b[97;42m 200 \x1b[0m|    2743h29m3s |     20.20.20.20 |\x1b[97;44m GET     \x1b[0m \"/\"\n", defaultLogFormatter(termTrueLongDurationParam))
}


// ff:
// t:
func TestColorForMethod(t *testing.T) {
	colorForMethod := func(method string) string {
		p := LogFormatterParams{
			Method: method,
		}
		return p.MethodColor()
	}

	assert.Equal(t, blue, colorForMethod("GET"), "get should be blue")
	assert.Equal(t, cyan, colorForMethod("POST"), "post should be cyan")
	assert.Equal(t, yellow, colorForMethod("PUT"), "put should be yellow")
	assert.Equal(t, red, colorForMethod("DELETE"), "delete should be red")
	assert.Equal(t, green, colorForMethod("PATCH"), "patch should be green")
	assert.Equal(t, magenta, colorForMethod("HEAD"), "head should be magenta")
	assert.Equal(t, white, colorForMethod("OPTIONS"), "options should be white")
	assert.Equal(t, reset, colorForMethod("TRACE"), "trace is not defined and should be the reset color")
}


// ff:
// t:
func TestColorForStatus(t *testing.T) {
	colorForStatus := func(code int) string {
		p := LogFormatterParams{
			StatusCode: code,
		}
		return p.StatusCodeColor()
	}

	assert.Equal(t, white, colorForStatus(http.StatusContinue), "1xx should be white")
	assert.Equal(t, green, colorForStatus(http.StatusOK), "2xx should be green")
	assert.Equal(t, white, colorForStatus(http.StatusMovedPermanently), "3xx should be white")
	assert.Equal(t, yellow, colorForStatus(http.StatusNotFound), "4xx should be yellow")
	assert.Equal(t, red, colorForStatus(2), "other things should be red")
}


// ff:
// t:
func TestResetColor(t *testing.T) {
	p := LogFormatterParams{}
	assert.Equal(t, string([]byte{27, 91, 48, 109}), p.ResetColor())
}


// ff:
// t:
func TestIsOutputColor(t *testing.T) {
// 用isTerm标志进行测试
	p := LogFormatterParams{
		isTerm: true,
	}

	consoleColorMode = autoColor
	assert.Equal(t, true, p.IsOutputColor())

	ForceConsoleColor()
	assert.Equal(t, true, p.IsOutputColor())

	DisableConsoleColor()
	assert.Equal(t, false, p.IsOutputColor())

	// test with isTerm flag false.
	p = LogFormatterParams{
		isTerm: false,
	}

	consoleColorMode = autoColor
	assert.Equal(t, false, p.IsOutputColor())

	ForceConsoleColor()
	assert.Equal(t, true, p.IsOutputColor())

	DisableConsoleColor()
	assert.Equal(t, false, p.IsOutputColor())

// 重置控制台颜色模式
	consoleColorMode = autoColor
}


// ff:
// t:
func TestErrorLogger(t *testing.T) {
	router := New()
	router.Use(ErrorLogger())
	router.GET("/error", func(c *Context) {
		c.Error(errors.New("this is an error")) // nolint: errcheck
// 翻译：// 不进行errcheck检查
	})
	router.GET("/abort", func(c *Context) {
		c.AbortWithError(http.StatusUnauthorized, errors.New("no authorized")) // nolint: errcheck
// 翻译：// 不进行errcheck检查
	})
	router.GET("/print", func(c *Context) {
		c.Error(errors.New("this is an error")) // nolint: errcheck
// 翻译：// 不进行errcheck检查
		c.String(http.StatusInternalServerError, "hola!")
	})

	w := PerformRequest(router, "GET", "/error")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"error\":\"this is an error\"}", w.Body.String())

	w = PerformRequest(router, "GET", "/abort")
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, "{\"error\":\"no authorized\"}", w.Body.String())

	w = PerformRequest(router, "GET", "/print")
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "hola!{\"error\":\"this is an error\"}", w.Body.String())
}


// ff:
// t:
func TestLoggerWithWriterSkippingPaths(t *testing.T) {
	buffer := new(strings.Builder)
	router := New()
	router.Use(LoggerWithWriter(buffer, "/skipped"))
	router.GET("/logged", func(c *Context) {})
	router.GET("/skipped", func(c *Context) {})

	PerformRequest(router, "GET", "/logged")
	assert.Contains(t, buffer.String(), "200")

	buffer.Reset()
	PerformRequest(router, "GET", "/skipped")
	assert.Contains(t, buffer.String(), "")
}


// ff:
// t:
func TestLoggerWithConfigSkippingPaths(t *testing.T) {
	buffer := new(strings.Builder)
	router := New()
	router.Use(LoggerWithConfig(LoggerConfig{
		Output:    buffer,
		SkipPaths: []string{"/skipped"},
	}))
	router.GET("/logged", func(c *Context) {})
	router.GET("/skipped", func(c *Context) {})

	PerformRequest(router, "GET", "/logged")
	assert.Contains(t, buffer.String(), "200")

	buffer.Reset()
	PerformRequest(router, "GET", "/skipped")
	assert.Contains(t, buffer.String(), "")
}


// ff:
// t:
func TestDisableConsoleColor(t *testing.T) {
	New()
	assert.Equal(t, autoColor, consoleColorMode)
	DisableConsoleColor()
	assert.Equal(t, disableColor, consoleColorMode)

// 重置控制台颜色模式
	consoleColorMode = autoColor
}


// ff:
// t:
func TestForceConsoleColor(t *testing.T) {
	New()
	assert.Equal(t, autoColor, consoleColorMode)
	ForceConsoleColor()
	assert.Equal(t, forceColor, consoleColorMode)

// 重置控制台颜色模式
	consoleColorMode = autoColor
}
