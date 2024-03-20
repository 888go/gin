// 版权所有 ? 2018 Gin 核心团队。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package render

import (
	"io"
	"net/http"
	"strconv"
)

// Reader 包含了 IO 读取器及其长度，以及自定义的 ContentType 和其他头部信息。
type Reader struct {
	ContentType   string
	ContentLength int64
	Reader        io.Reader
	Headers       map[string]string
}

// Render (Reader) 通过自定义的 ContentType 和头部信息写入数据。

// ff:
// err:
// w:
func (r Reader) Render(w http.ResponseWriter) (err error) {
	r.WriteContentType(w)
	if r.ContentLength >= 0 {
		if r.Headers == nil {
			r.Headers = map[string]string{}
		}
		r.Headers["Content-Length"] = strconv.FormatInt(r.ContentLength, 10)
	}
	r.writeHeaders(w, r.Headers)
	_, err = io.Copy(w, r.Reader)
	return
}

// WriteContentType (针对Reader) 写入自定义的 ContentType。

// ff:
// w:
func (r Reader) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, []string{r.ContentType})
}

// writeHeaders 写入自定义 Header。
func (r Reader) writeHeaders(w http.ResponseWriter, headers map[string]string) {
	header := w.Header()
	for k, v := range headers {
		if header.Get(k) == "" {
			header.Set(k, v)
		}
	}
}
