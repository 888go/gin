
<原文开始>
// Copyright 2014 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权声明 2014 Manu Martinez-Almeida。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// Delims represents a set of Left and Right delimiters for HTML template rendering.
<原文结束>

# <翻译开始>
// Delims 表示用于 HTML 模板渲染的一组左（Left）和右（Right）定界符。
# <翻译结束>


<原文开始>
	// Left delimiter, defaults to {{.
<原文结束>

# <翻译开始>
// 左侧分隔符，默认为 {{.
# <翻译结束>


<原文开始>
	// Right delimiter, defaults to }}.
<原文结束>

# <翻译开始>
// 右侧分隔符，默认为 }}。
# <翻译结束>


<原文开始>
// HTMLRender interface is to be implemented by HTMLProduction and HTMLDebug.
<原文结束>

# <翻译开始>
// HTMLRender 接口需要由 HTMLProduction 和 HTMLDebug 实现。
# <翻译结束>


<原文开始>
	// Instance returns an HTML instance.
<原文结束>

# <翻译开始>
// Instance 返回一个HTML实例。
# <翻译结束>


<原文开始>
// HTMLProduction contains template reference and its delims.
<原文结束>

# <翻译开始>
// HTMLProduction 包含模板引用及其分隔符。
# <翻译结束>


<原文开始>
// HTMLDebug contains template delims and pattern and function with file list.
<原文结束>

# <翻译开始>
// HTMLDebug 包含模板分隔符、模式以及带有文件列表的函数，主要用于调试HTML。
# <翻译结束>


<原文开始>
// HTML contains template reference and its name with given interface object.
<原文结束>

# <翻译开始>
// HTML 包含模板引用及其名称，以及给定的接口对象。
# <翻译结束>


<原文开始>
// Instance (HTMLProduction) returns an HTML instance which it realizes Render interface.
<原文结束>

# <翻译开始>
// Instance (HTMLProduction) 返回一个实现了Render接口的HTML实例。
# <翻译结束>


<原文开始>
// Instance (HTMLDebug) returns an HTML instance which it realizes Render interface.
<原文结束>

# <翻译开始>
// Instance (HTMLDebug) 返回一个实现了Render接口的HTML实例。
# <翻译结束>


<原文开始>
// Render (HTML) executes template and writes its result with custom ContentType for response.
<原文结束>

# <翻译开始>
// Render (HTML) 执行模板并使用自定义 ContentType 将其结果写入响应。
# <翻译结束>


<原文开始>
// WriteContentType (HTML) writes HTML ContentType.
<原文结束>

# <翻译开始>
// WriteContentType (HTML) 写入 HTML ContentType。
# <翻译结束>

