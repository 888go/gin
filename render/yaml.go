// 版权所有 2014 Manu Martinez-Almeida。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package render

import (
	"net/http"
	
	"gopkg.in/yaml.v3"
)

// YAML 包含给定的接口对象。
type YAML struct {
	Data any
}

var yamlContentType = []string{"application/x-yaml; charset=utf-8"}

// Render (YAML) 将给定的接口对象进行序列化（marshals），并使用自定义 ContentType 写入数据。
func (r YAML) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)

	bytes, err := yaml.Marshal(r.Data)
	if err != nil {
		return err
	}

	_, err = w.Write(bytes)
	return err
}

// WriteContentType (YAML) 为响应写入 YAML ContentType。
func (r YAML) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, yamlContentType)
}
