// 版权所有 2019 Gin 核心团队。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package binding

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJSONBindingBindBody(t *testing.T) {
	var s struct {
		Foo string `json:"foo"`
	}
	err := jsonBinding{}.BindBody([]byte(`{"foo": "FOO"}`), &s)
	require.NoError(t, err)
	assert.Equal(t, "FOO", s.Foo)
}

func TestJSONBindingBindBodyMap(t *testing.T) {
	s := make(map[string]string)
	err := jsonBinding{}.BindBody([]byte(`{"foo": "FOO","hello":"world"}`), &s)
	require.NoError(t, err)
	assert.Len(t, s, 2)
	assert.Equal(t, "FOO", s["foo"])
	assert.Equal(t, "world", s["hello"])
}
