// 版权所有2022 Gin Core团队
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到

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
