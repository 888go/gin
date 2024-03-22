
<原文开始>
// Copyright 2017 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2017 Manu Martinez-Almeida。保留所有权利。
// 本源代码的使用受 MIT 风格许可证协议约束，
// 该协议可在 LICENSE 文件中查阅。
# <翻译结束>


<原文开始>
// Dir returns a http.FileSystem that can be used by http.FileServer(). It is used internally
// in router.Static().
// if listDirectory == true, then it works the same as http.Dir() otherwise it returns
// a filesystem that prevents http.FileServer() to list the directory files.
<原文结束>

# <翻译开始>
// Dir 返回一个可用于 http.FileServer() 的 http.FileSystem。它在 router.Static() 内部使用。
// 如果 listDirectory 为 true，则其行为与 http.Dir() 相同；否则，它将返回一个文件系统，
// 阻止 http.FileServer() 列出目录中的文件。
# <翻译结束>


<原文开始>
// Open conforms to http.Filesystem.
<原文结束>

# <翻译开始>
// Open 符合 http.Filesystem 接口。
//
// 注意!!! 此方法不能翻译, 因为是http包的接口实现
# <翻译结束>


<原文开始>
// Readdir overrides the http.File default implementation.
<原文结束>

# <翻译开始>
// Readdir 重写（覆盖）了 http.File 的默认实现。
# <翻译结束>


<原文开始>
// this disables directory listing
<原文结束>

# <翻译开始>
// 这将禁用目录列表
# <翻译结束>

