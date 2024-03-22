// Copyright 2014 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gin类

import (
	"flag"
	"os"
	"testing"
	
	"github.com/888go/gin/binding"
	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv(EnvGinMode, X常量_运行模式_测试)
}

func TestSetMode(t *testing.T) {
	assert.Equal(t, testCode, ginMode)
	assert.Equal(t, X常量_运行模式_测试, X取运行模式())
	os.Unsetenv(EnvGinMode)

	X设置运行模式("")
	assert.Equal(t, testCode, ginMode)
	assert.Equal(t, X常量_运行模式_测试, X取运行模式())

	tmp := flag.CommandLine
	flag.CommandLine = flag.NewFlagSet("", flag.ContinueOnError)
	X设置运行模式("")
	assert.Equal(t, debugCode, ginMode)
	assert.Equal(t, X常量_运行模式_调试, X取运行模式())
	flag.CommandLine = tmp

	X设置运行模式(X常量_运行模式_调试)
	assert.Equal(t, debugCode, ginMode)
	assert.Equal(t, X常量_运行模式_调试, X取运行模式())

	X设置运行模式(X常量_运行模式_发布)
	assert.Equal(t, releaseCode, ginMode)
	assert.Equal(t, X常量_运行模式_发布, X取运行模式())

	X设置运行模式(X常量_运行模式_测试)
	assert.Equal(t, testCode, ginMode)
	assert.Equal(t, X常量_运行模式_测试, X取运行模式())

	assert.Panics(t, func() { X设置运行模式("unknown") })
}

func TestDisableBindValidation(t *testing.T) {
	v := binding.Validator
	assert.NotNil(t, binding.Validator)
	X关闭Validator验证器()
	assert.Nil(t, binding.Validator)
	binding.Validator = v
}

func TestEnableJsonDecoderUseNumber(t *testing.T) {
	assert.False(t, binding.EnableDecoderUseNumber)
	X启用Json解码器使用Number()
	assert.True(t, binding.EnableDecoderUseNumber)
}

func TestEnableJsonDecoderDisallowUnknownFields(t *testing.T) {
	assert.False(t, binding.EnableDecoderDisallowUnknownFields)
	X启用json解码器禁止未知字段()
	assert.True(t, binding.EnableDecoderDisallowUnknownFields)
}
