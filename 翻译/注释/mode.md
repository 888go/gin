
<原文开始>
// Copyright 2014 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Manu Martinez-Almeida版权所有
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到
# <翻译结束>


<原文开始>
// EnvGinMode indicates environment name for gin mode.
<原文结束>

# <翻译开始>
// EnvGinMode为gin模式的环境名
# <翻译结束>


<原文开始>
	// DebugMode indicates gin mode is debug.
<原文结束>

# <翻译开始>
// DebugMode表示gin模式为debug
# <翻译结束>


<原文开始>
	// ReleaseMode indicates gin mode is release.
<原文结束>

# <翻译开始>
// ReleaseMode表示gin模式为release
# <翻译结束>


<原文开始>
	// TestMode indicates gin mode is test.
<原文结束>

# <翻译开始>
// TestMode表示gin模式为test
# <翻译结束>


<原文开始>
// DefaultWriter is the default io.Writer used by Gin for debug output and
// middleware output like Logger() or Recovery().
// Note that both Logger and Recovery provides custom ways to configure their
// output io.Writer.
// To support coloring in Windows use:
//
//	import "github.com/mattn/go-colorable"
//	gin.DefaultWriter = colorable.NewColorableStdout()
<原文结束>

# <翻译开始>
// defaultwwriter是默认的io
// Gin用于调试输出和中间件输出的写入器，如Logger()或Recovery()
// 请注意，Logger和Recovery都提供了自定义的方式来配置它们的输出
// 要在Windows中支持着色，请使用:import "github.com/mattn/go-colorable"杜松子酒
// defaultwwriter = colorable.NewColorableStdout()
# <翻译结束>


<原文开始>
// DefaultErrorWriter is the default io.Writer used by Gin to debug errors
<原文结束>

# <翻译开始>
// DefaultErrorWriter是默认io
// Gin用来调试错误的写入器
# <翻译结束>


<原文开始>
// SetMode sets gin mode according to input string.
<原文结束>

# <翻译开始>
// SetMode根据输入的字符串设置gin模式
# <翻译结束>


<原文开始>
// DisableBindValidation closes the default validator.
<原文结束>

# <翻译开始>
// DisableBindValidation关闭默认验证器
# <翻译结束>


<原文开始>
// EnableJsonDecoderUseNumber sets true for binding.EnableDecoderUseNumber to
// call the UseNumber method on the JSON Decoder instance.
<原文结束>

# <翻译开始>
// EnableJsonDecoderUseNumber为绑定设置为true
// EnableDecoderUseNumber以调用JSON Decoder实例上的UseNumber方法
# <翻译结束>


<原文开始>
// EnableJsonDecoderDisallowUnknownFields sets true for binding.EnableDecoderDisallowUnknownFields to
// call the DisallowUnknownFields method on the JSON Decoder instance.
<原文结束>

# <翻译开始>
// EnableJsonDecoderDisallowUnknownFields为绑定设置为true
// EnableDecoderDisallowUnknownFields调用JSON Decoder实例上的DisallowUnknownFields方法
# <翻译结束>


<原文开始>
// Mode returns current gin mode.
<原文结束>

# <翻译开始>
// Mode返回当前gin模式
# <翻译结束>

