// 版权所有 2014 Manu Martinez-Almeida。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package render

import (
	"fmt"
	"net/http"
	
	"github.com/888go/gin/internal/bytesconv"
)

// String 包含给定的接口对象切片及其格式。
type String struct {
	Format string
	Data   []any
}

var plainContentType = []string{"text/plain; charset=utf-8"}

// Render (String) 通过自定义 ContentType 写入数据。

// ff:
// w:
func (r String) Render(w http.ResponseWriter) error {
	return WriteString(w, r.Format, r.Data)
}

// WriteContentType (字符串) 写入纯文本 ContentType。

// ff:
// w:
func (r String) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, plainContentType)
}

// WriteString 根据其格式写入数据，并写入自定义 ContentType。

// ff:
// err:
// data:
// format:
// w:
func WriteString(w http.ResponseWriter, format string, data []any) (err error) {
	writeContentType(w, plainContentType)
	if len(data) > 0 {
		_, err = fmt.Fprintf(w, format, data...)
		return
	}
	_, err = w.Write(bytesconv.StringToBytes(format))
	return
}
