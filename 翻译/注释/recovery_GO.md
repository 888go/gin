
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
// RecoveryFunc defines the function passable to CustomRecovery.
<原文结束>

# <翻译开始>
// RecoveryFunc 定义了可以传递给 CustomRecovery 的函数。
# <翻译结束>


<原文开始>
// Recovery returns a middleware that recovers from any panics and writes a 500 if there was one.
<原文结束>

# <翻译开始>
// Recovery 返回一个中间件，该中间件可从任何 panic 中恢复，并在发生 panic 时写入一个 500 状态码。
# <翻译结束>


<原文开始>
// CustomRecovery returns a middleware that recovers from any panics and calls the provided handle func to handle it.
<原文结束>

# <翻译开始>
// CustomRecovery 返回一个中间件，该中间件可从任何 panic 中恢复，并调用提供的处理函数来处理它。
# <翻译结束>


<原文开始>
// RecoveryWithWriter returns a middleware for a given writer that recovers from any panics and writes a 500 if there was one.
<原文结束>

# <翻译开始>
// RecoveryWithWriter 返回一个中间件，针对给定的writer，在发生任何 panic 时进行恢复，并在发生 panic 时写入 500 状态码。
# <翻译结束>


<原文开始>
// CustomRecoveryWithWriter returns a middleware for a given writer that recovers from any panics and calls the provided handle func to handle it.
<原文结束>

# <翻译开始>
// CustomRecoveryWithWriter 函数为给定的 writer 返回一个中间件，该中间件可从任何 panic 中恢复，并调用提供的 handle 函数来处理它。
# <翻译结束>


<原文开始>
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
<原文结束>

# <翻译开始>
// 检查连接是否已断开，因为这并不是真正需要引发恐慌并打印堆栈跟踪的条件。
# <翻译结束>


<原文开始>
// If the connection is dead, we can't write a status to it.
<原文结束>

# <翻译开始>
// 如果连接已断开，我们无法向其写入状态。
# <翻译结束>


<原文开始>
// stack returns a nicely formatted stack frame, skipping skip frames.
<原文结束>

# <翻译开始>
// stack 返回一个格式良好的堆栈帧，跳过 skip 个帧。
# <翻译结束>


<原文开始>
	// As we loop, we open files and read them. These variables record the currently
	// loaded file.
<原文结束>

# <翻译开始>
// 在循环过程中，我们会打开文件并读取它们。这些变量记录当前加载的文件。
# <翻译结束>


<原文开始>
// Skip the expected number of frames
<原文结束>

# <翻译开始>
// 跳过预期的帧数
# <翻译结束>


<原文开始>
// Print this much at least.  If we can't find the source, it won't show.
<原文结束>

# <翻译开始>
// 至少打印这么多。如果我们找不到源，它将不会显示。
# <翻译结束>


<原文开始>
// source returns a space-trimmed slice of the n'th line.
<原文结束>

# <翻译开始>
// source 返回第n行去除两端空白字符后的切片。
# <翻译结束>


<原文开始>
// in stack trace, lines are 1-indexed but our array is 0-indexed
<原文结束>

# <翻译开始>
// 在堆栈跟踪中，行号是从1开始编号的，但我们的数组是从0开始索引的
# <翻译结束>


<原文开始>
// function returns, if possible, the name of the function containing the PC.
<原文结束>

# <翻译开始>
// 函数尝试返回包含PC的函数名称（如果可能的话）。
# <翻译结束>


<原文开始>
// timeFormat returns a customized time string for logger.
<原文结束>

# <翻译开始>
// timeFormat 返回一个自定义的时间字符串，用于日志记录。
# <翻译结束>

