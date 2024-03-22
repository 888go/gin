// 版权所有 ? 2022 Gin 核心团队。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

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
