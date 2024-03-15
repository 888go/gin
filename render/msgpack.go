// 版权声明：2017 Manu Martinez-Almeida。保留所有权利。
// 本源代码的使用受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//go:build !nomsgpack

package render

import (
	"net/http"
	
	"github.com/ugorji/go/codec"
)

// 在此处检查接口实现以支持go构建标签nomsgpack。
// 参考：https://github.com/gin-gonic/gin/pull/1852/
// 这段Go语言代码注释翻译成中文的大致意思是：
// 为了支持go构建标签nomsgpack，我们在此处实现了相应的接口。
// 详情请参阅：https://github.com/gin-gonic/gin/pull/1852/
var (
	_ Render = MsgPack{}
)

// MsgPack 包含给定的 interface 对象。
type MsgPack struct {
	Data any
}

var msgpackContentType = []string{"application/msgpack; charset=utf-8"}

// WriteContentType (MsgPack) 写入 MsgPack 的 ContentType。
func (r MsgPack) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, msgpackContentType)
}

// Render (MsgPack) 将给定的接口对象进行编码，并使用自定义 ContentType 写入数据。
func (r MsgPack) Render(w http.ResponseWriter) error {
	return WriteMsgPack(w, r.Data)
}

// WriteMsgPack 将MsgPack ContentType写入，并编码给定的接口对象。
func WriteMsgPack(w http.ResponseWriter, obj any) error {
	writeContentType(w, msgpackContentType)
	var mh codec.MsgpackHandle
	return codec.NewEncoder(w, &mh).Encode(obj)
}
