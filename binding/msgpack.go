// 版权所有2017马努·马丁内斯-阿尔梅达
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到

//go:build !nomsgpack

package binding

import (
	"bytes"
	"io"
	"net/http"
	
	"github.com/ugorji/go/codec"
)

type msgpackBinding struct{}

func (msgpackBinding) Name() string {
	return "msgpack"
}

func (msgpackBinding) Bind(req *http.Request, obj any) error {
	return decodeMsgPack(req.Body, obj)
}

func (msgpackBinding) BindBody(body []byte, obj any) error {
	return decodeMsgPack(bytes.NewReader(body), obj)
}

func decodeMsgPack(r io.Reader, obj any) error {
	cdc := new(codec.MsgpackHandle)
	if err := codec.NewDecoder(r, cdc).Decode(&obj); err != nil {
		return err
	}
	return validate(obj)
}
