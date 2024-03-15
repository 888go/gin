
<原文开始>
// Copyright 2017 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有2017马努·马丁内斯-阿尔梅达
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到
# <翻译结束>


<原文开始>
// Dir returns a http.FileSystem that can be used by http.FileServer(). It is used internally
// in router.Static().
// if listDirectory == true, then it works the same as http.Dir() otherwise it returns
// a filesystem that prevents http.FileServer() to list the directory files.
<原文结束>

# <翻译开始>
// Dir返回一个http
// http.FileServer()可以使用的文件系统
// 它在router.Static()内部使用
// 如果listDirectory == true，那么它的工作方式与http.Dir()相同，否则它返回一个文件系统，阻止http.FileServer()列出目录文件
# <翻译结束>


<原文开始>
// Open conforms to http.Filesystem.
<原文结束>

# <翻译开始>
// Open符合http.Filesystem
# <翻译结束>


<原文开始>
// Readdir overrides the http.File default implementation.
<原文结束>

# <翻译开始>
// Readdir覆盖http
// 文件默认实现
# <翻译结束>


<原文开始>
	// this disables directory listing
<原文结束>

# <翻译开始>
// 这将禁用目录列表
# <翻译结束>

