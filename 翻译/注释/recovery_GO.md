
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
// RecoveryFunc defines the function passable to CustomRecovery.
<原文结束>

# <翻译开始>
// recoveryfunction定义了CustomRecovery可传递的函数
# <翻译结束>


<原文开始>
// Recovery returns a middleware that recovers from any panics and writes a 500 if there was one.
<原文结束>

# <翻译开始>
// Recovery返回一个中间件，它可以从任何Panic中恢复，如果有Panic，则写入500
# <翻译结束>


<原文开始>
// CustomRecovery returns a middleware that recovers from any panics and calls the provided handle func to handle it.
<原文结束>

# <翻译开始>
// CustomRecovery返回一个中间件，它可以从任何Panic中恢复，并调用提供的handle函数来处理它
# <翻译结束>


<原文开始>
// RecoveryWithWriter returns a middleware for a given writer that recovers from any panics and writes a 500 if there was one.
<原文结束>

# <翻译开始>
// RecoveryWithWriter为给定的写入器返回一个中间件，该写入器可以从任何Panic中恢复并写入500(如果有的话)
# <翻译结束>


<原文开始>
// CustomRecoveryWithWriter returns a middleware for a given writer that recovers from any panics and calls the provided handle func to handle it.
<原文结束>

# <翻译开始>
// CustomRecoveryWithWriter为给定的编写器返回一个中间件，该编写器可以从任何Panic中恢复，并调用提供的handle函数来处理它
# <翻译结束>


<原文开始>
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
<原文结束>

# <翻译开始>
// 检查是否有断开的连接，因为它实际上并不是需要进行紧急堆栈跟踪的条件
# <翻译结束>


<原文开始>
					// If the connection is dead, we can't write a status to it.
<原文结束>

# <翻译开始>
// 如果连接已死，我们就不能向它写入状态
# <翻译结束>


<原文开始>
// stack returns a nicely formatted stack frame, skipping skip frames.
<原文结束>

# <翻译开始>
// Stack返回一个格式良好的堆栈帧，跳过跳过帧
# <翻译结束>


<原文开始>
// the returned data
<原文结束>

# <翻译开始>
// 返回的数据
# <翻译结束>


<原文开始>
	// As we loop, we open files and read them. These variables record the currently
	// loaded file.
<原文结束>

# <翻译开始>
// 在循环时，打开文件并读取它们
// 这些变量记录当前加载的文件
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
// 至少打印这么多
// 如果我们找不到源头，就不会显示出来
# <翻译结束>


<原文开始>
// source returns a space-trimmed slice of the n'th line.
<原文结束>

# <翻译开始>
// Source返回第n行经过空格处理的切片
# <翻译结束>


<原文开始>
// in stack trace, lines are 1-indexed but our array is 0-indexed
<原文结束>

# <翻译开始>
// 在堆栈跟踪中，行是1索引的，但是数组是0索引的
# <翻译结束>


<原文开始>
// function returns, if possible, the name of the function containing the PC.
<原文结束>

# <翻译开始>
// 如果可能的话，函数返回包含PC的函数名
# <翻译结束>


<原文开始>
// timeFormat returns a customized time string for logger.
<原文结束>

# <翻译开始>
// timeFormat返回记录器的自定义时间字符串
# <翻译结束>

