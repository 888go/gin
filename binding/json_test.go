// 版权所有2019 Gin Core Team
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到

package binding

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)


// ff:
// t:
func TestJSONBindingBindBody(t *testing.T) {
	var s struct {
		Foo string `json:"foo"`
	}
	err := jsonBinding{}.BindBody([]byte(`{"foo": "FOO"}`), &s)
	require.NoError(t, err)
	assert.Equal(t, "FOO", s.Foo)
}


// ff:
// t:
func TestJSONBindingBindBodyMap(t *testing.T) {
	s := make(map[string]string)
	err := jsonBinding{}.BindBody([]byte(`{"foo": "FOO","hello":"world"}`), &s)
	require.NoError(t, err)
	assert.Len(t, s, 2)
	assert.Equal(t, "FOO", s["foo"])
	assert.Equal(t, "world", s["hello"])
}
