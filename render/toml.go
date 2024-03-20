// 版权所有 ? 2022 Gin 核心团队。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package render

import (
	"net/http"
	
	"github.com/pelletier/go-toml/v2"
)

// TOML 包含给定的接口对象。
type TOML struct {
	Data any
}

var TOMLContentType = []string{"application/toml; charset=utf-8"}

// Render (TOML) 将给定的接口对象进行序列化，并使用自定义 ContentType 写入数据。
func (r TOML) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)

	bytes, err := toml.Marshal(r.Data)
	if err != nil {
		return err
	}

	_, err = w.Write(bytes)
	return err
}

// WriteContentType (TOML) 为响应写入 TOML ContentType。
func (r TOML) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, TOMLContentType)
}
