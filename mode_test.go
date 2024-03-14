// Manu Martinez-Almeida版权所有
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到

package gin
import (
	"flag"
	"os"
	"testing"
	
	"e.coding.net/gogit/go/gin/binding"
	"github.com/stretchr/testify/assert"
	)
func init() {
	os.Setenv(EnvGinMode, TestMode)
}

func TestSetMode(t *testing.T) {
	assert.Equal(t, testCode, ginMode)
	assert.Equal(t, TestMode, Mode())
	os.Unsetenv(EnvGinMode)

	SetMode("")
	assert.Equal(t, testCode, ginMode)
	assert.Equal(t, TestMode, Mode())

	tmp := flag.CommandLine
	flag.CommandLine = flag.NewFlagSet("", flag.ContinueOnError)
	SetMode("")
	assert.Equal(t, debugCode, ginMode)
	assert.Equal(t, DebugMode, Mode())
	flag.CommandLine = tmp

	SetMode(DebugMode)
	assert.Equal(t, debugCode, ginMode)
	assert.Equal(t, DebugMode, Mode())

	SetMode(ReleaseMode)
	assert.Equal(t, releaseCode, ginMode)
	assert.Equal(t, ReleaseMode, Mode())

	SetMode(TestMode)
	assert.Equal(t, testCode, ginMode)
	assert.Equal(t, TestMode, Mode())

	assert.Panics(t, func() { SetMode("unknown") })
}

func TestDisableBindValidation(t *testing.T) {
	v := binding.Validator
	assert.NotNil(t, binding.Validator)
	DisableBindValidation()
	assert.Nil(t, binding.Validator)
	binding.Validator = v
}

func TestEnableJsonDecoderUseNumber(t *testing.T) {
	assert.False(t, binding.EnableDecoderUseNumber)
	EnableJsonDecoderUseNumber()
	assert.True(t, binding.EnableDecoderUseNumber)
}

func TestEnableJsonDecoderDisallowUnknownFields(t *testing.T) {
	assert.False(t, binding.EnableDecoderDisallowUnknownFields)
	EnableJsonDecoderDisallowUnknownFields()
	assert.True(t, binding.EnableDecoderDisallowUnknownFields)
}
