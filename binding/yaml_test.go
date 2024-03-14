// 版权所有2019 Gin Core Team
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到

package binding
import (
	"testing"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	)
func TestYAMLBindingBindBody(t *testing.T) {
	var s struct {
		Foo string `yaml:"foo"`
	}
	err := yamlBinding{}.BindBody([]byte("foo: FOO"), &s)
	require.NoError(t, err)
	assert.Equal(t, "FOO", s.Foo)
}
