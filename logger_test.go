// Copyright 2014 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gin类

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
	X设置运行模式(X常量_运行模式_测试)
}

func TestLogger(t *testing.T) {
	buffer := new(strings.Builder)
	router := X创建()
	router.X中间件(LoggerWithWriter(buffer))
	router.X绑定GET("/example", func(c *Context) {})
	router.X绑定POST("/example", func(c *Context) {})
	router.X绑定PUT("/example", func(c *Context) {})
	router.X绑定DELETE("/example", func(c *Context) {})
	router.X绑定PATCH("/example", func(c *Context) {})
	router.X绑定HEAD("/example", func(c *Context) {})
	router.X绑定OPTIONS("/example", func(c *Context) {})

	PerformRequest(router, "GET", "/example?a=100")
	assert.Contains(t, buffer.String(), "200")
	assert.Contains(t, buffer.String(), "GET")
	assert.Contains(t, buffer.String(), "/example")
	assert.Contains(t, buffer.String(), "a=100")

	// I wrote these first (extending the above) but then realized they are more
	// like integration tests because they test the whole logging process rather
	// than individual functions.  Im not sure where these should go.
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

func TestLoggerWithConfig(t *testing.T) {
	buffer := new(strings.Builder)
	router := X创建()
	router.X中间件(LoggerWithConfig(LoggerConfig{Output: buffer}))
	router.X绑定GET("/example", func(c *Context) {})
	router.X绑定POST("/example", func(c *Context) {})
	router.X绑定PUT("/example", func(c *Context) {})
	router.X绑定DELETE("/example", func(c *Context) {})
	router.X绑定PATCH("/example", func(c *Context) {})
	router.X绑定HEAD("/example", func(c *Context) {})
	router.X绑定OPTIONS("/example", func(c *Context) {})

	PerformRequest(router, "GET", "/example?a=100")
	assert.Contains(t, buffer.String(), "200")
	assert.Contains(t, buffer.String(), "GET")
	assert.Contains(t, buffer.String(), "/example")
	assert.Contains(t, buffer.String(), "a=100")

	// I wrote these first (extending the above) but then realized they are more
	// like integration tests because they test the whole logging process rather
	// than individual functions.  Im not sure where these should go.
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

func TestLoggerWithFormatter(t *testing.T) {
	buffer := new(strings.Builder)

	d := DefaultWriter
	DefaultWriter = buffer
	defer func() {
		DefaultWriter = d
	}()

	router := X创建()
	router.X中间件(X中间件函数_自定义日志格式(func(param LogFormatterParams) string {
		return fmt.Sprintf("[FORMATTER TEST] %v | %3d | %13v | %15s | %-7s %#v\n%s",
			param.X响应时间.Format("2006/01/02 - 15:04:05"),
			param.X状态码,
			param.X时长,
			param.X客户端IP,
			param.X方法,
			param.X路径,
			param.X错误信息,
		)
	}))
	router.X绑定GET("/example", func(c *Context) {})
	PerformRequest(router, "GET", "/example?a=100")

	// output test
	assert.Contains(t, buffer.String(), "[FORMATTER TEST]")
	assert.Contains(t, buffer.String(), "200")
	assert.Contains(t, buffer.String(), "GET")
	assert.Contains(t, buffer.String(), "/example")
	assert.Contains(t, buffer.String(), "a=100")
}

func TestLoggerWithConfigFormatting(t *testing.T) {
	var gotParam LogFormatterParams
	var gotKeys map[string]any
	buffer := new(strings.Builder)

	router := X创建()
	router.engine.trustedCIDRs, _ = router.engine.prepareTrustedCIDRs()

	router.X中间件(LoggerWithConfig(LoggerConfig{
		Output: buffer,
		Formatter: func(param LogFormatterParams) string {
			// for assert test
			gotParam = param

			return fmt.Sprintf("[FORMATTER TEST] %v | %3d | %13v | %15s | %-7s %s\n%s",
				param.X响应时间.Format("2006/01/02 - 15:04:05"),
				param.X状态码,
				param.X时长,
				param.X客户端IP,
				param.X方法,
				param.X路径,
				param.X错误信息,
			)
		},
	}))
	router.X绑定GET("/example", func(c *Context) {
		// set dummy ClientIP
		c.X请求.Header.Set("X-Forwarded-For", "20.20.20.20")
		gotKeys = c.X上下文设置值Map
		time.Sleep(time.Millisecond)
	})
	PerformRequest(router, "GET", "/example?a=100")

	// output test
	assert.Contains(t, buffer.String(), "[FORMATTER TEST]")
	assert.Contains(t, buffer.String(), "200")
	assert.Contains(t, buffer.String(), "GET")
	assert.Contains(t, buffer.String(), "/example")
	assert.Contains(t, buffer.String(), "a=100")

	// LogFormatterParams test
	assert.NotNil(t, gotParam.X请求)
	assert.NotEmpty(t, gotParam.X响应时间)
	assert.Equal(t, 200, gotParam.X状态码)
	assert.NotEmpty(t, gotParam.X时长)
	assert.Equal(t, "20.20.20.20", gotParam.X客户端IP)
	assert.Equal(t, "GET", gotParam.X方法)
	assert.Equal(t, "/example?a=100", gotParam.X路径)
	assert.Empty(t, gotParam.X错误信息)
	assert.Equal(t, gotKeys, gotParam.X上下文设置值map)
}

func TestDefaultLogFormatter(t *testing.T) {
	timeStamp := time.Unix(1544173902, 0).UTC()

	termFalseParam := LogFormatterParams{
		X响应时间:    timeStamp,
		X状态码:   200,
		X时长:      time.Second * 5,
		X客户端IP:     "20.20.20.20",
		X方法:       "GET",
		X路径:         "/",
		X错误信息: "",
		X是否输出到终端:       false,
	}

	termTrueParam := LogFormatterParams{
		X响应时间:    timeStamp,
		X状态码:   200,
		X时长:      time.Second * 5,
		X客户端IP:     "20.20.20.20",
		X方法:       "GET",
		X路径:         "/",
		X错误信息: "",
		X是否输出到终端:       true,
	}
	termTrueLongDurationParam := LogFormatterParams{
		X响应时间:    timeStamp,
		X状态码:   200,
		X时长:      time.Millisecond * 9876543210,
		X客户端IP:     "20.20.20.20",
		X方法:       "GET",
		X路径:         "/",
		X错误信息: "",
		X是否输出到终端:       true,
	}

	termFalseLongDurationParam := LogFormatterParams{
		X响应时间:    timeStamp,
		X状态码:   200,
		X时长:      time.Millisecond * 9876543210,
		X客户端IP:     "20.20.20.20",
		X方法:       "GET",
		X路径:         "/",
		X错误信息: "",
		X是否输出到终端:       false,
	}

	assert.Equal(t, "[GIN] 2018/12/07 - 09:11:42 | 200 |            5s |     20.20.20.20 | GET      \"/\"\n", defaultLogFormatter(termFalseParam))
	assert.Equal(t, "[GIN] 2018/12/07 - 09:11:42 | 200 |    2743h29m3s |     20.20.20.20 | GET      \"/\"\n", defaultLogFormatter(termFalseLongDurationParam))

	assert.Equal(t, "[GIN] 2018/12/07 - 09:11:42 |\x1b[97;42m 200 \x1b[0m|            5s |     20.20.20.20 |\x1b[97;44m GET     \x1b[0m \"/\"\n", defaultLogFormatter(termTrueParam))
	assert.Equal(t, "[GIN] 2018/12/07 - 09:11:42 |\x1b[97;42m 200 \x1b[0m|    2743h29m3s |     20.20.20.20 |\x1b[97;44m GET     \x1b[0m \"/\"\n", defaultLogFormatter(termTrueLongDurationParam))
}

func TestColorForMethod(t *testing.T) {
	colorForMethod := func(method string) string {
		p := LogFormatterParams{
			X方法: method,
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

func TestColorForStatus(t *testing.T) {
	colorForStatus := func(code int) string {
		p := LogFormatterParams{
			X状态码: code,
		}
		return p.StatusCodeColor()
	}

	assert.Equal(t, white, colorForStatus(http.StatusContinue), "1xx should be white")
	assert.Equal(t, green, colorForStatus(http.StatusOK), "2xx should be green")
	assert.Equal(t, white, colorForStatus(http.StatusMovedPermanently), "3xx should be white")
	assert.Equal(t, yellow, colorForStatus(http.StatusNotFound), "4xx should be yellow")
	assert.Equal(t, red, colorForStatus(2), "other things should be red")
}

func TestResetColor(t *testing.T) {
	p := LogFormatterParams{}
	assert.Equal(t, string([]byte{27, 91, 48, 109}), p.ResetColor())
}

func TestIsOutputColor(t *testing.T) {
	// test with isTerm flag true.
	p := LogFormatterParams{
		X是否输出到终端: true,
	}

	consoleColorMode = autoColor
	assert.Equal(t, true, p.IsOutputColor())

	X开启控制台颜色()
	assert.Equal(t, true, p.IsOutputColor())

	X关闭控制台颜色()
	assert.Equal(t, false, p.IsOutputColor())

	// test with isTerm flag false.
	p = LogFormatterParams{
		X是否输出到终端: false,
	}

	consoleColorMode = autoColor
	assert.Equal(t, false, p.IsOutputColor())

	X开启控制台颜色()
	assert.Equal(t, true, p.IsOutputColor())

	X关闭控制台颜色()
	assert.Equal(t, false, p.IsOutputColor())

	// reset console color mode.
	consoleColorMode = autoColor
}

func TestErrorLogger(t *testing.T) {
	router := X创建()
	router.X中间件(ErrorLogger())
	router.X绑定GET("/error", func(c *Context) {
		c.X错误(errors.New("this is an error")) //nolint: errcheck
	})
	router.X绑定GET("/abort", func(c *Context) {
		c.X停止并带状态码与错误(http.StatusUnauthorized, errors.New("no authorized")) //nolint: errcheck
	})
	router.X绑定GET("/print", func(c *Context) {
		c.X错误(errors.New("this is an error")) //nolint: errcheck
		c.X输出文本(http.StatusInternalServerError, "hola!")
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

func TestLoggerWithWriterSkippingPaths(t *testing.T) {
	buffer := new(strings.Builder)
	router := X创建()
	router.X中间件(LoggerWithWriter(buffer, "/skipped"))
	router.X绑定GET("/logged", func(c *Context) {})
	router.X绑定GET("/skipped", func(c *Context) {})

	PerformRequest(router, "GET", "/logged")
	assert.Contains(t, buffer.String(), "200")

	buffer.Reset()
	PerformRequest(router, "GET", "/skipped")
	assert.Contains(t, buffer.String(), "")
}

func TestLoggerWithConfigSkippingPaths(t *testing.T) {
	buffer := new(strings.Builder)
	router := X创建()
	router.X中间件(LoggerWithConfig(LoggerConfig{
		Output:    buffer,
		SkipPaths: []string{"/skipped"},
	}))
	router.X绑定GET("/logged", func(c *Context) {})
	router.X绑定GET("/skipped", func(c *Context) {})

	PerformRequest(router, "GET", "/logged")
	assert.Contains(t, buffer.String(), "200")

	buffer.Reset()
	PerformRequest(router, "GET", "/skipped")
	assert.Contains(t, buffer.String(), "")
}

func TestDisableConsoleColor(t *testing.T) {
	X创建()
	assert.Equal(t, autoColor, consoleColorMode)
	X关闭控制台颜色()
	assert.Equal(t, disableColor, consoleColorMode)

	// reset console color mode.
	consoleColorMode = autoColor
}

func TestForceConsoleColor(t *testing.T) {
	X创建()
	assert.Equal(t, autoColor, consoleColorMode)
	X开启控制台颜色()
	assert.Equal(t, forceColor, consoleColorMode)

	// reset console color mode.
	consoleColorMode = autoColor
}
