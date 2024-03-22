// 版权所有 2014 Manu Martinez-Almeida。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package gin类

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"syscall"
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestPanicClean(t *testing.T) {
	buffer := new(strings.Builder)
	router := X创建()
	password := "my-super-secret-password"
	router.X中间件(RecoveryWithWriter(buffer))
	router.X绑定GET("/recovery", func(c *Context) {
		c.X停止并带状态码(http.StatusBadRequest)
		panic("Oupps, Houston, we have a problem")
	})
	// RUN
	w := PerformRequest(router, "GET", "/recovery",
		header{
			Key:   "Host",
			Value: "www.google.com",
		},
		header{
			Key:   "Authorization",
			Value: fmt.Sprintf("Bearer %s", password),
		},
		header{
			Key:   "Content-Type",
			Value: "application/json",
		},
	)
	// TEST
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// 检查缓冲区中不包含密钥
	assert.NotContains(t, buffer.String(), password)
}

// TestPanicInHandler 断言在处理器中捕获到了 panic。
func TestPanicInHandler(t *testing.T) {
	buffer := new(strings.Builder)
	router := X创建()
	router.X中间件(RecoveryWithWriter(buffer))
	router.X绑定GET("/recovery", func(_ *Context) {
		panic("Oupps, Houston, we have a problem")
	})
	// RUN
	w := PerformRequest(router, "GET", "/recovery")
	// TEST
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, buffer.String(), "panic recovered")
	assert.Contains(t, buffer.String(), "Oupps, Houston, we have a problem")
	assert.Contains(t, buffer.String(), t.Name())
	assert.NotContains(t, buffer.String(), "GET /recovery")

	// Debug模式会打印请求
	X设置运行模式(X常量_运行模式_调试)
	// RUN
	w = PerformRequest(router, "GET", "/recovery")
	// TEST
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, buffer.String(), "GET /recovery")

	X设置运行模式(X常量_运行模式_测试)
}

// TestPanicWithAbort 断言即使使用了 context.Abort，也能捕获到 panic 异常。
func TestPanicWithAbort(t *testing.T) {
	router := X创建()
	router.X中间件(RecoveryWithWriter(nil))
	router.X绑定GET("/recovery", func(c *Context) {
		c.X停止并带状态码(http.StatusBadRequest)
		panic("Oupps, Houston, we have a problem")
	})
	// RUN
	w := PerformRequest(router, "GET", "/recovery")
	// TEST
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSource(t *testing.T) {
	bs := source(nil, 0)
	assert.Equal(t, dunno, bs)

	in := [][]byte{
		[]byte("Hello world."),
		[]byte("Hi, gin.."),
	}
	bs = source(in, 10)
	assert.Equal(t, dunno, bs)

	bs = source(in, 1)
	assert.Equal(t, []byte("Hello world."), bs)
}

func TestFunction(t *testing.T) {
	bs := function(1)
	assert.Equal(t, dunno, bs)
}

// TestPanicWithBrokenPipe 断言 recovery 特别处理了
// 向已断开连接的管道写入响应的情况
func TestPanicWithBrokenPipe(t *testing.T) {
	const expectCode = 204

	expectMsgs := map[syscall.Errno]string{
		syscall.EPIPE:      "broken pipe",
		syscall.ECONNRESET: "connection reset by peer",
	}

	for errno, expectMsg := range expectMsgs {
		t.Run(expectMsg, func(t *testing.T) {
			var buf strings.Builder

			router := X创建()
			router.X中间件(RecoveryWithWriter(&buf))
			router.X绑定GET("/recovery", func(c *Context) {
				// Start writing response
				c.X设置响应协议头值("X-Test", "Value")
				c.X设置状态码(expectCode)

				// 哎呀，客户端连接已关闭
				e := &net.OpError{Err: &os.SyscallError{Err: errno}}
				panic(e)
			})
			// RUN
			w := PerformRequest(router, "GET", "/recovery")
			// TEST
			assert.Equal(t, expectCode, w.Code)
			assert.Contains(t, strings.ToLower(buf.String()), expectMsg)
		})
	}
}

func TestCustomRecoveryWithWriter(t *testing.T) {
	errBuffer := new(strings.Builder)
	buffer := new(strings.Builder)
	router := X创建()
	handleRecovery := func(c *Context, err any) {
		errBuffer.WriteString(err.(string))
		c.X停止并带状态码(http.StatusBadRequest)
	}
	router.X中间件(CustomRecoveryWithWriter(buffer, handleRecovery))
	router.X绑定GET("/recovery", func(_ *Context) {
		panic("Oupps, Houston, we have a problem")
	})
	// RUN
	w := PerformRequest(router, "GET", "/recovery")
	// TEST
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, buffer.String(), "panic recovered")
	assert.Contains(t, buffer.String(), "Oupps, Houston, we have a problem")
	assert.Contains(t, buffer.String(), t.Name())
	assert.NotContains(t, buffer.String(), "GET /recovery")

	// Debug模式会打印请求
	X设置运行模式(X常量_运行模式_调试)
	// RUN
	w = PerformRequest(router, "GET", "/recovery")
	// TEST
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, buffer.String(), "GET /recovery")

	assert.Equal(t, strings.Repeat("Oupps, Houston, we have a problem", 2), errBuffer.String())

	X设置运行模式(X常量_运行模式_测试)
}

func TestCustomRecovery(t *testing.T) {
	errBuffer := new(strings.Builder)
	buffer := new(strings.Builder)
	router := X创建()
	DefaultErrorWriter = buffer
	handleRecovery := func(c *Context, err any) {
		errBuffer.WriteString(err.(string))
		c.X停止并带状态码(http.StatusBadRequest)
	}
	router.X中间件(CustomRecovery(handleRecovery))
	router.X绑定GET("/recovery", func(_ *Context) {
		panic("Oupps, Houston, we have a problem")
	})
	// RUN
	w := PerformRequest(router, "GET", "/recovery")
	// TEST
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, buffer.String(), "panic recovered")
	assert.Contains(t, buffer.String(), "Oupps, Houston, we have a problem")
	assert.Contains(t, buffer.String(), t.Name())
	assert.NotContains(t, buffer.String(), "GET /recovery")

	// Debug模式会打印请求
	X设置运行模式(X常量_运行模式_调试)
	// RUN
	w = PerformRequest(router, "GET", "/recovery")
	// TEST
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, buffer.String(), "GET /recovery")

	assert.Equal(t, strings.Repeat("Oupps, Houston, we have a problem", 2), errBuffer.String())

	X设置运行模式(X常量_运行模式_测试)
}

func TestRecoveryWithWriterWithCustomRecovery(t *testing.T) {
	errBuffer := new(strings.Builder)
	buffer := new(strings.Builder)
	router := X创建()
	DefaultErrorWriter = buffer
	handleRecovery := func(c *Context, err any) {
		errBuffer.WriteString(err.(string))
		c.X停止并带状态码(http.StatusBadRequest)
	}
	router.X中间件(RecoveryWithWriter(DefaultErrorWriter, handleRecovery))
	router.X绑定GET("/recovery", func(_ *Context) {
		panic("Oupps, Houston, we have a problem")
	})
	// RUN
	w := PerformRequest(router, "GET", "/recovery")
	// TEST
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, buffer.String(), "panic recovered")
	assert.Contains(t, buffer.String(), "Oupps, Houston, we have a problem")
	assert.Contains(t, buffer.String(), t.Name())
	assert.NotContains(t, buffer.String(), "GET /recovery")

	// Debug模式会打印请求
	X设置运行模式(X常量_运行模式_调试)
	// RUN
	w = PerformRequest(router, "GET", "/recovery")
	// TEST
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, buffer.String(), "GET /recovery")

	assert.Equal(t, strings.Repeat("Oupps, Houston, we have a problem", 2), errBuffer.String())

	X设置运行模式(X常量_运行模式_测试)
}
