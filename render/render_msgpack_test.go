// 版权所有 2014 Manu Martinez-Almeida。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//go:build !nomsgpack

package render

import (
	"bytes"
	"net/http/httptest"
	"testing"
	
	"github.com/stretchr/testify/assert"
	"github.com/ugorji/go/codec"
)

// TODO：单元测试
// 测试错误

func TestRenderMsgPack(t *testing.T) {
	w := httptest.NewRecorder()
	data := map[string]any{
		"foo": "bar",
	}

	(MsgPack{data}).WriteContentType(w)
	assert.Equal(t, "application/msgpack; charset=utf-8", w.Header().Get("Content-Type"))

	err := (MsgPack{data}).Render(w)

	assert.NoError(t, err)

	h := new(codec.MsgpackHandle)
	assert.NotNil(t, h)
	buf := bytes.NewBuffer([]byte{})
	assert.NotNil(t, buf)
	err = codec.NewEncoder(buf, h).Encode(data)

	assert.NoError(t, err)
	assert.Equal(t, w.Body.String(), buf.String())
	assert.Equal(t, "application/msgpack; charset=utf-8", w.Header().Get("Content-Type"))
}
