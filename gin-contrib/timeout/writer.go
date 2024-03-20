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
func NewWriter(w gin.ResponseWriter, buf *bytes.Buffer) *Writer {
	return &Writer{ResponseWriter: w, body: buf, headers: make(http.Header)}
}

// Write 将数据写入响应体
func (w *Writer) Write(data []byte) (int, error) {
	if w.timeout || w.body == nil {
		return 0, nil
	}

	w.mu.Lock()
	defer w.mu.Unlock()

	return w.body.Write(data)
}

// WriteHeader 向客户端发送带有指定状态码的 HTTP 响应头。
// 如果响应写入器已写入了头信息，或者发生超时，
// 此方法将不做任何操作。
func (w *Writer) WriteHeader(code int) {
	if w.timeout || w.wroteHeaders {
		return
	}

// gin 使用 -1 来跳过写入状态码
// 详情见 https://github.com/gin-gonic/gin/blob/a0acf1df2814fcd828cb2d7128f2f4e2136d3fac/response_writer.go#L61
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
func (w *Writer) Header() http.Header {
	return w.headers
}

// WriteString 将字符串写入响应体
func (w *Writer) WriteString(s string) (int, error) {
	return w.Write([]byte(s))
}

// FreeBuffer 会释放缓冲区指针
func (w *Writer) FreeBuffer() {
	// 如果不重置body，旧的字节数据将会被放入bufPool中
	w.body.Reset()
	w.body = nil
}

// 我们必须在这里覆盖 Status 函数，
// 否则在其他自定义 gin 中间件中，gin.Context.Writer.Status() 返回的 HTTP 状态码将始终为 200。
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
