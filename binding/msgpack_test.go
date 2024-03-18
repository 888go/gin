// 版权所有2019 Gin Core Team
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到

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
