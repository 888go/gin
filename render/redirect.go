// 版权声明 2014 Manu Martinez-Almeida。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package render

import (
	"fmt"
	"net/http"
)

// Redirect 包含了 HTTP 请求引用以及重定向状态码和位置。
type Redirect struct {
	Code     int
	Request  *http.Request
	Location string
}

// Render (Redirect) 将HTTP请求重定向到新位置，并写出重定向响应。

// ff:
// w:
func (r Redirect) Render(w http.ResponseWriter) error {
	if (r.Code < http.StatusMultipleChoices || r.Code > http.StatusPermanentRedirect) && r.Code != http.StatusCreated {
		panic(fmt.Sprintf("Cannot redirect with status code %d", r.Code))
	}
	http.Redirect(w, r.Request, r.Location, r.Code)
	return nil
}

// WriteContentType (重定向) 不要写入任何 ContentType。

// ff:
// http.ResponseWriter:
func (r Redirect) WriteContentType(http.ResponseWriter) {}
