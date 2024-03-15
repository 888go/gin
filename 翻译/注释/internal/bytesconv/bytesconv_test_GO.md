
<原文开始>
// Copyright 2020 Gin Core Team. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2020 Gin Core Team。保留所有权利。
// 本源代码的使用受 MIT 风格许可证协议约束，
// 该协议可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// go test -v
<原文结束>

# <翻译开始>
// 使用以下命令运行测试并显示详细信息：go test -v
# <翻译结束>


<原文开始>
// 6 bits to represent a letter index
<原文结束>

# <翻译开始>
// 用6位来表示一个字母索引
# <翻译结束>


<原文开始>
// All 1-bits, as many as letterIdxBits
<原文结束>

# <翻译开始>
// 生成包含1-bits的切片，数量与letterIdxBits相同
# <翻译结束>


<原文开始>
// # of letter indices fitting in 63 bits
<原文结束>

# <翻译开始>
// 符合63位大小的字母索引数量
# <翻译结束>


<原文开始>
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
<原文结束>

# <翻译开始>
// A src.Int63() 生成63个随机位，足以生成letterIdxMax个字符！
# <翻译结束>


<原文开始>
// go test -v -run=none -bench=^BenchmarkBytesConv -benchmem=true
<原文结束>

# <翻译开始>
// 运行Go测试，详细输出模式(-v)，不执行任何指定的测试函数(-run=none)，仅运行以BenchmarkBytesConv开头的基准测试(-bench=^BenchmarkBytesConv)，并开启内存使用统计(-benchmem=true)
# <翻译结束>

