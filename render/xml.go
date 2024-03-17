// 版权声明 2014 Manu Martinez-Almeida。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package render

import (
	"encoding/xml"
	"net/http"
)

// XML 包含给定的接口对象。
type XML struct {
	Data any
}

var xmlContentType = []string{"application/xml; charset=utf-8"}

// Render (XML) 将给定的接口对象进行编码，并使用自定义 ContentType 写入数据。

// ff:
// w:
func (r XML) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)
	return xml.NewEncoder(w).Encode(r.Data)
}

// WriteContentType (XML) 为响应写入 XML ContentType。

// ff:
// w:
func (r XML) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, xmlContentType)
}
