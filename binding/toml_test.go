// 版权所有2022 Gin Core团队
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到

package binding

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTOMLBindingBindBody(t *testing.T) {
	var s struct {
		Foo string `toml:"foo"`
	}
	tomlBody := `foo="FOO"`
	err := tomlBinding{}.BindBody([]byte(tomlBody), &s)
	require.NoError(t, err)
	assert.Equal(t, "FOO", s.Foo)
}
