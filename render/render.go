// 版权声明 2014 Manu Martinez-Almeida。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package render

import (
	"net/http"
)

// Render接口需要被JSON、XML、HTML、YAML等实现。
// 这段Go语言代码的注释翻译成中文为：该Render接口是为了让JSON、XML、HTML、YAML等数据格式实现而设计的。
type Render interface {
// Render通过自定义的ContentType写入数据。
	Render(http.ResponseWriter) error
// WriteContentType 写入自定义 ContentType。
	WriteContentType(w http.ResponseWriter)
}

var (
	_ Render     = JSON{}
	_ Render     = IndentedJSON{}
	_ Render     = SecureJSON{}
	_ Render     = JsonpJSON{}
	_ Render     = XML{}
	_ Render     = String{}
	_ Render     = Redirect{}
	_ Render     = Data{}
	_ Render     = HTML{}
	_ HTMLRender = HTMLDebug{}
	_ HTMLRender = HTMLProduction{}
	_ Render     = YAML{}
	_ Render     = Reader{}
	_ Render     = AsciiJSON{}
	_ Render     = ProtoBuf{}
	_ Render     = TOML{}
)

func writeContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}
