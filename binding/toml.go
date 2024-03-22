// 版权所有 ? 2022 Gin 核心团队。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package binding

import (
	"bytes"
	"io"
	"net/http"
	
	"github.com/pelletier/go-toml/v2"
)

type tomlBinding struct{}

func (tomlBinding) Name() string {
	return "toml"
}

func (tomlBinding) Bind(req *http.Request, obj any) error {
	return decodeToml(req.Body, obj)
}

func (tomlBinding) BindBody(body []byte, obj any) error {
	return decodeToml(bytes.NewReader(body), obj)
}

func decodeToml(r io.Reader, obj any) error {
	decoder := toml.NewDecoder(r)
	if err := decoder.Decode(obj); err != nil {
		return err
	}
	return decoder.Decode(obj)
}
