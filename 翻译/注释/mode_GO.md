
<原文开始>
// Copyright 2014 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2014 Manu Martinez-Almeida。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// EnvGinMode indicates environment name for gin mode.
<原文结束>

# <翻译开始>
// EnvGinMode 指示 Gin 模式的环境名称。
# <翻译结束>


<原文开始>
// DebugMode indicates gin mode is debug.
<原文结束>

# <翻译开始>
// DebugMode 指示 gin 模式为调试模式。
# <翻译结束>


<原文开始>
// ReleaseMode indicates gin mode is release.
<原文结束>

# <翻译开始>
// ReleaseMode 表示 gin 模式为发布模式。
# <翻译结束>


<原文开始>
// TestMode indicates gin mode is test.
<原文结束>

# <翻译开始>
// TestMode 表示 gin 模式为测试模式。
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
// DefaultWriter 是 Gin 默认使用的 io.Writer，用于调试输出以及中间件输出，如 Logger() 和 Recovery()。
// 注意，Logger 和 Recovery 都提供了自定义配置其输出 io.Writer 的方法。
// 若要在 Windows 系统中支持彩色输出，请使用：
//
//	导入 "github.com/mattn/go-colorable"
//	gin.DefaultWriter = colorable.NewColorableStdout()
# <翻译结束>


<原文开始>
// DefaultErrorWriter is the default io.Writer used by Gin to debug errors
<原文结束>

# <翻译开始>
// DefaultErrorWriter 是 Gin 默认使用的 io.Writer，用于调试错误
# <翻译结束>


<原文开始>
// SetMode sets gin mode according to input string.
<原文结束>

# <翻译开始>
// SetMode 根据输入的字符串设置 gin 模式。
# <翻译结束>


<原文开始>
// DisableBindValidation closes the default validator.
<原文结束>

# <翻译开始>
// DisableBindValidation 关闭默认的验证器。
# <翻译结束>


<原文开始>
// EnableJsonDecoderUseNumber sets true for binding.EnableDecoderUseNumber to
// call the UseNumber method on the JSON Decoder instance.
<原文结束>

# <翻译开始>
// EnableJsonDecoderUseNumber 将参数设置为 true 以启用 binding.EnableDecoderUseNumber，
// 这样就会在 JSON 解码器实例上调用 UseNumber 方法。
# <翻译结束>


<原文开始>
// EnableJsonDecoderDisallowUnknownFields sets true for binding.EnableDecoderDisallowUnknownFields to
// call the DisallowUnknownFields method on the JSON Decoder instance.
<原文结束>

# <翻译开始>
// EnableJsonDecoderDisallowUnknownFields 将 binding.EnableDecoderDisallowUnknownFields 设为 true，
// 以便在 JSON 解码器实例上调用 DisallowUnknownFields 方法。
# <翻译结束>


<原文开始>
// Mode returns current gin mode.
<原文结束>

# <翻译开始>
// Mode 返回当前 gin 模式。
# <翻译结束>

