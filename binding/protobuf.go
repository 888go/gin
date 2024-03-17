// Manu Martinez-Almeida版权所有
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到

package binding

import (
	"errors"
	"io"
	"net/http"
	
	"google.golang.org/protobuf/proto"
)

type protobufBinding struct{}

func (protobufBinding) Name() string {
	return "protobuf"
}


// ff:
// obj:
// req:
func (b protobufBinding) Bind(req *http.Request, obj any) error {
	buf, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}
	return b.BindBody(buf, obj)
}

func (protobufBinding) BindBody(body []byte, obj any) error {
	msg, ok := obj.(proto.Message)
	if !ok {
		return errors.New("obj is not ProtoMessage")
	}
	if err := proto.Unmarshal(body, msg); err != nil {
		return err
	}
// 这里返回validate(obj)也是一样的，但是到目前为止，我们还不能将' binding:"" '添加到由gen-proto自动生成的结构中
	return nil
// 返回验证(obj)
}
