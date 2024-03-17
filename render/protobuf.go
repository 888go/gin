// 版权所有 2018 Gin 核心团队。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package render

import (
	"net/http"
	
	"google.golang.org/protobuf/proto"
)

// ProtoBuf 包含给定的接口对象。
type ProtoBuf struct {
	Data any
}

var protobufContentType = []string{"application/x-protobuf"}

// Render (ProtoBuf) 将给定的接口对象序列化，并以自定义的 ContentType 写入数据。

// ff:
// w:
func (r ProtoBuf) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)

	bytes, err := proto.Marshal(r.Data.(proto.Message))
	if err != nil {
		return err
	}

	_, err = w.Write(bytes)
	return err
}

// WriteContentType (ProtoBuf) 写入 ProtoBuf 的 ContentType。

// ff:
// w:
func (r ProtoBuf) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, protobufContentType)
}
