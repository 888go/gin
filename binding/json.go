// 版权所有 2014 Manu Martinez-Almeida。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package binding

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	
	"github.com/888go/gin/internal/json"
)

// EnableDecoderUseNumber 用于调用 JSON 解码器实例上的 UseNumber 方法。启用 UseNumber 后，解码器将在反序列化数字时将其解析为 Number 类型而不是 float64 类型并存储到 any 类型变量中。
var EnableDecoderUseNumber = false

// EnableDecoderDisallowUnknownFields 用于调用 JSON 解码器实例上的 DisallowUnknownFields 方法。该方法启用后，当目标是一个结构体且输入包含与目标中非忽略的导出字段不匹配的对象键时，解码器会返回一个错误。
var EnableDecoderDisallowUnknownFields = false

type jsonBinding struct{}

func (jsonBinding) Name() string {
	return "json"
}

func (jsonBinding) Bind(req *http.Request, obj any) error {
	if req == nil || req.Body == nil {
		return errors.New("invalid request")
	}
	return decodeJSON(req.Body, obj)
}

func (jsonBinding) BindBody(body []byte, obj any) error {
	return decodeJSON(bytes.NewReader(body), obj)
}

func decodeJSON(r io.Reader, obj any) error {
	decoder := json.NewDecoder(r)
	if EnableDecoderUseNumber {
		decoder.UseNumber()
	}
	if EnableDecoderDisallowUnknownFields {
		decoder.DisallowUnknownFields()
	}
	if err := decoder.Decode(obj); err != nil {
		return err
	}
	return validate(obj)
}
