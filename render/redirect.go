// 版权所有 2014 Manu Martinez-Almeida。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package render

import (
	"fmt"
	"net/http"
)

// Redirect 包含了HTTP请求引用，以及重定向状态码和目标位置。
type Redirect struct {
	Code     int
	Request  *http.Request
	Location string
}

// Render (Redirect) 将HTTP请求重定向到新位置并写入重定向响应。
func (r Redirect) Render(w http.ResponseWriter) error {
	if (r.Code < http.StatusMultipleChoices || r.Code > http.StatusPermanentRedirect) && r.Code != http.StatusCreated {
		panic(fmt.Sprintf("Cannot redirect with status code %d", r.Code))
	}
	http.Redirect(w, r.Request, r.Location, r.Code)
	return nil
}

// WriteContentType (Redirect) 不要写入任何 ContentType。
func (r Redirect) WriteContentType(http.ResponseWriter) {}
