// Manu Martinez-Almeida版权所有
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到

package binding
import (
	"bytes"
	"errors"
	"io"
	"net/http"
	
	"e.coding.net/gogit/go/gin/internal/json"
	)
// EnableDecoderUseNumber用于调用JSON Decoder实例上的UseNumber方法
// UseNumber导致解码器将一个数字反编组为一个any，作为一个number而不是一个float64
var EnableDecoderUseNumber = false

// EnableDecoderDisallowUnknownFields用于调用JSON Decoder实例上的DisallowUnknownFields方法
// DisallowUnknownFields导致解码器返回一个错误，当目标是一个结构，并且输入包含的对象键与目标中任何非忽略的导出字段不匹配时
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
