// Copyright 2014 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gin类

import (
	"flag"
	"io"
	"os"
	
	"github.com/888go/gin/binding"
)

// EnvGinMode indicates environment name for gin mode.
const EnvGinMode = "GIN_MODE"

const (
	// DebugMode indicates gin mode is debug.
	X常量_运行模式_调试 = "debug"
	// ReleaseMode indicates gin mode is release.
	X常量_运行模式_发布 = "release"
	// TestMode indicates gin mode is test.
	X常量_运行模式_测试 = "test"
)

const (
	debugCode = iota
	releaseCode
	testCode
)

// DefaultWriter is the default io.Writer used by Gin for debug output and
// middleware output like Logger() or Recovery().
// Note that both Logger and Recovery provides custom ways to configure their
// output io.Writer.
// To support coloring in Windows use:
//
//	import "github.com/mattn/go-colorable"
//	gin.DefaultWriter = colorable.NewColorableStdout()
var DefaultWriter io.Writer = os.Stdout

// DefaultErrorWriter is the default io.Writer used by Gin to debug errors
var DefaultErrorWriter io.Writer = os.Stderr

var (
	ginMode  = debugCode
	modeName = X常量_运行模式_调试
)

func init() {
	mode := os.Getenv(EnvGinMode)
	X设置运行模式(mode)
}

// SetMode sets gin mode according to input string.
func X设置运行模式(常量_运行模式 string) {
	if 常量_运行模式 == "" {
		if flag.Lookup("test.v") != nil {
			常量_运行模式 = X常量_运行模式_测试
		} else {
			常量_运行模式 = X常量_运行模式_调试
		}
	}

	switch 常量_运行模式 {
	case X常量_运行模式_调试:
		ginMode = debugCode
	case X常量_运行模式_发布:
		ginMode = releaseCode
	case X常量_运行模式_测试:
		ginMode = testCode
	default:
		panic("gin mode unknown: " + 常量_运行模式 + " (available mode: debug release test)")
	}

	modeName = 常量_运行模式
}

// DisableBindValidation closes the default validator.
func X关闭Validator验证器() {
	binding.Validator = nil
}

// EnableJsonDecoderUseNumber sets true for binding.EnableDecoderUseNumber to
// call the UseNumber method on the JSON Decoder instance.
func X启用Json解码器使用Number() {
	binding.EnableDecoderUseNumber = true
}

// EnableJsonDecoderDisallowUnknownFields sets true for binding.EnableDecoderDisallowUnknownFields to
// call the DisallowUnknownFields method on the JSON Decoder instance.
func X启用json解码器禁止未知字段() {
	binding.EnableDecoderDisallowUnknownFields = true
}

// Mode returns current gin mode.
func X取运行模式() string {
	return modeName
}
