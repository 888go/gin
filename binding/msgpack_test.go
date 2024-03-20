// 版权所有 2019 Gin 核心团队。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//go:build !nomsgpack

package binding

import (
	"bytes"
	"testing"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/ugorji/go/codec"
)

func TestMsgpackBindingBindBody(t *testing.T) {
	type teststruct struct {
		Foo string `msgpack:"foo"`
	}
	var s teststruct
	err := msgpackBinding{}.BindBody(msgpackBody(t, teststruct{"FOO"}), &s)
	require.NoError(t, err)
	assert.Equal(t, "FOO", s.Foo)
}

func msgpackBody(t *testing.T, obj any) []byte {
	var bs bytes.Buffer
	h := &codec.MsgpackHandle{}
	err := codec.NewEncoder(&bs, h).Encode(obj)
	require.NoError(t, err)
	return bs.Bytes()
}
