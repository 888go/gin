package timeout

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
	
	"github.com/888go/gin"
)

// Writer 是一个带有内存缓冲区的写入器
type Writer struct {
	gin.ResponseWriter
	body         *bytes.Buffer
	headers      http.Header
	mu           sync.Mutex
	timeout      bool
	wroteHeaders bool
	code         int
}

// NewWriter 将返回一个 timeout.Writer 指针

// ff:
// buf:
// w:

// ff:
// buf:
// w:

// ff:
// buf:
// w:

// ff:
// buf:
// w:

// ff:
// buf:
// w:

// ff:
// buf:
// w:
func NewWriter(w gin.ResponseWriter, buf *bytes.Buffer) *Writer {
	return &Writer{ResponseWriter: w, body: buf, headers: make(http.Header)}
}

// Write 将数据写入响应体

// ff:
// data:

// ff:
// data:

// ff:
// data:

// ff:
// data:

// ff:
// data:

// ff:
// data:
func (w *Writer) Write(data []byte) (int, error) {
	if w.timeout || w.body == nil {
		return 0, nil
	}

	w.mu.Lock()
	defer w.mu.Unlock()

	return w.body.Write(data)
}

// WriteHeader 向客户端发送带有指定状态码的HTTP响应头。
// 若响应写入器已发送过头部信息，或者发生超时，
// 此方法将不做任何操作。

// ff:
// code:

// ff:
// code:

// ff:
// code:

// ff:
// code:

// ff:
// code:

// ff:
// code:
func (w *Writer) WriteHeader(code int) {
	if w.timeout || w.wroteHeaders {
		return
	}

// gin 通过使用 -1 来跳过写入状态码
// 参考：https://github.com/gin-gonic/gin/blob/a0acf1df2814fcd828cb2d7128f2f4e2136d3fac/response_writer.go#L61
// 在 Gin 框架中，当传入的 HTTP 状态码为 -1 时，
// 表示框架将不会写出具体的状态码到响应中。
	if code == -1 {
		return
	}

	checkWriteHeaderCode(code)

	w.mu.Lock()
	defer w.mu.Unlock()

	w.writeHeader(code)
	w.ResponseWriter.WriteHeader(code)
}

func (w *Writer) writeHeader(code int) {
	w.wroteHeaders = true
	w.code = code
}

// Header 将获取响应头

// ff:

// ff:

// ff:

// ff:

// ff:

// ff:
func (w *Writer) Header() http.Header {
	return w.headers
}

// WriteString 将字符串写入响应体

// ff:
// s:

// ff:
// s:

// ff:
// s:

// ff:
// s:

// ff:
// s:

// ff:
// s:
func (w *Writer) WriteString(s string) (int, error) {
	return w.Write([]byte(s))
}

// FreeBuffer 将释放缓冲区指针

// ff:

// ff:

// ff:

// ff:

// ff:

// ff:
func (w *Writer) FreeBuffer() {
// 如果不重置body，旧的字节数据将会被放入bufPool中
	w.body.Reset()
	w.body = nil
}

// 我们必须在这里覆盖Status函数，
// 否则在其他自定义gin中间件中，gin.Context.Writer.Status()返回的HTTP状态码将始终为200。

// ff:

// ff:

// ff:

// ff:

// ff:

// ff:
func (w *Writer) Status() int {
	if w.code == 0 || w.timeout {
		return w.ResponseWriter.Status()
	}
	return w.code
}

func checkWriteHeaderCode(code int) {
	if code < 100 || code > 999 {
		panic(fmt.Sprintf("invalid http status code: %d", code))
	}
}
