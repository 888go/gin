// 版权所有 2014 Manu Martinez-Almeida。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

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
// 这里相当于返回 validate(obj)，但现在我们还不能在由 gen-proto 自动生成的结构体上添加
// `binding:""` 这个注解
	return nil
	// return validate(obj)
}
