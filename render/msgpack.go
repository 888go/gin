// 版权所有 ? 2017 Manu Martinez-Almeida。保留所有权利。
// 本源代码的使用受 MIT 风格许可证协议约束，
// 该协议可在 LICENSE 文件中查阅。

//go:build !nomsgpack

package render

import (
	"net/http"
	
	"github.com/ugorji/go/codec"
)

// 在此处检查接口实现以支持go构建标签nomsgpack。
// 参考：https://github.com/gin-gonic/gin/pull/1852/
// 
// 这段注释的大致意思是，为了支持名为"nomsgpack"的Go构建标签，这里实现了某个接口。更多详情可以参考提供的GitHub链接（ gin-gonic/gin 项目的一个拉取请求）。
var (
	_ Render = MsgPack{}
)

// MsgPack 包含给定的接口对象。
type MsgPack struct {
	Data any
}

var msgpackContentType = []string{"application/msgpack; charset=utf-8"}

// WriteContentType (MsgPack) 写入 MsgPack 的 ContentType。

// ff:
// w:
func (r MsgPack) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, msgpackContentType)
}

// Render (MsgPack) 将给定的接口对象进行编码，并使用自定义 ContentType 写入数据。

// ff:
// w:
func (r MsgPack) Render(w http.ResponseWriter) error {
	return WriteMsgPack(w, r.Data)
}

// WriteMsgPack 将MsgPack ContentType写入，并对给定的接口对象进行编码。

// ff:
// obj:
// w:
func WriteMsgPack(w http.ResponseWriter, obj any) error {
	writeContentType(w, msgpackContentType)
	var mh codec.MsgpackHandle
	return codec.NewEncoder(w, &mh).Encode(obj)
}
