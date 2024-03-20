// 版权所有 2014 Manu Martinez-Almeida。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package gin

import (
	"flag"
	"os"
	"testing"
	
	"github.com/888go/gin/binding"
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
