
<原文开始>
// Copyright 2020 Gin Core Team. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2020 Gin 核心团队。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// 6 bits to represent a letter index
<原文结束>

# <翻译开始>
// 使用6位表示字母索引
# <翻译结束>


<原文开始>
// All 1-bits, as many as letterIdxBits
<原文结束>

# <翻译开始>
// 所有为1的位，数量与letterIdxBits相同
# <翻译结束>


<原文开始>
// # of letter indices fitting in 63 bits
<原文结束>

# <翻译开始>
// 符合63位的字母索引数量
# <翻译结束>


<原文开始>
// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
<原文结束>

# <翻译开始>
// A src.Int63() 生成63个随机位，足够用于letterIdxMax个字符！
# <翻译结束>


<原文开始>
// go test -v -run=none -bench=^BenchmarkBytesConv -benchmem=true
<原文结束>

# <翻译开始>
// 运行命令：go test -v（详细模式）-run=none（不运行任何正常测试用例）-bench=^BenchmarkBytesConv（仅运行名称以"BenchmarkBytesConv"开头的基准测试）-benchmem=true（在基准测试中包含内存分配统计信息）
# <翻译结束>

