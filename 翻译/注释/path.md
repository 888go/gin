
<原文开始>
// Copyright 2013 Julien Schmidt. All rights reserved.
// Based on the path package, Copyright 2009 The Go Authors.
// Use of this source code is governed by a BSD-style license that can be found
// at https://github.com/julienschmidt/httprouter/blob/master/LICENSE.
<原文结束>

# <翻译开始>
// 版权所有2013朱利安施密特
// 版权所有
// 基于路径包，版权归the Go Authors所有
// 此源代码的使用受bsd风格的许可证的约束，该许可证可在https://github.com/julienschmidt/httprouter/blob/master/LICENSE上找到
# <翻译结束>


<原文开始>
// cleanPath is the URL version of path.Clean, it returns a canonical URL path
// for p, eliminating . and .. elements.
//
// The following rules are applied iteratively until no further processing can
// be done:
//  1. Replace multiple slashes with a single slash.
//  2. Eliminate each . path name element (the current directory).
//  3. Eliminate each inner .. path name element (the parent directory)
//     along with the non-.. element that precedes it.
//  4. Eliminate .. elements that begin a rooted path:
//     that is, replace "/.." by "/" at the beginning of a path.
//
// If the result of this process is an empty string, "/" is returned.
<原文结束>

# <翻译开始>
// cleanPath是path的URL版本
// 干净，它返回p的规范URL路径，消除
// 和. .元素
// 迭代地应用以下规则，直到无法进行进一步处理为止:将多个斜杠替换为单个斜杠
// 2. 消除每一个
// 路径名元素(当前目录)
// 3. 消除每个内部…路径名元素(父目录)以及非-..它前面的元素
// 4. 消除……开始根路径的元素:即替换"/.."由“/”;在一条路的起点
// 如果此过程的结果为空
# <翻译结束>


<原文开始>
	// Turn empty string into "/"
<原文结束>

# <翻译开始>
// 将空字符串转换为"/"
# <翻译结束>


<原文开始>
	// Reasonably sized buffer on stack to avoid allocations in the common case.
	// If a larger buffer is required, it gets allocated dynamically.
<原文结束>

# <翻译开始>
// 合理大小的堆栈缓冲区，以避免在通常情况下分配
// 如果需要更大的缓冲区，则动态分配
# <翻译结束>


<原文开始>
	// Invariants:
	//      reading from path; r is index of next byte to process.
	//      writing to buf; w is index of next byte to write.
<原文结束>

# <翻译开始>
// 不变量:从path读取;R是要处理的下一个字节的索引
// 给……写信;W是要写入的下一个字节的索引
# <翻译结束>


<原文开始>
	// path must start with '/'
<原文结束>

# <翻译开始>
// 路径必须以“/”开头
# <翻译结束>


<原文开始>
	// A bit more clunky without a 'lazybuf' like the path package, but the loop
	// gets completely inlined (bufApp calls).
	// loop has no expensive function calls (except 1x make)		// So in contrast to the path package this loop has no expensive function
	// calls (except make, if needed).
<原文结束>

# <翻译开始>
// 没有像path包那样的“lazybuf”会更笨拙一些，但循环会完全内联(bufApp调用)
// 循环没有昂贵的函数调用(除了1x make)所以与path包相比，这个循环没有昂贵的函数调用(除了make，如果需要的话)
# <翻译结束>


<原文开始>
			// empty path element, trailing slash is added after the end
<原文结束>

# <翻译开始>
// 空路径元素，结尾后添加斜杠
# <翻译结束>


<原文开始>
			// . element
<原文结束>

# <翻译开始>
// ． 元素
# <翻译结束>


<原文开始>
			// .. element: remove to last /
<原文结束>

# <翻译开始>
// ．． 元素:移到最后
# <翻译结束>


<原文开始>
				// can backtrack
<原文结束>

# <翻译开始>
// 可以回溯
# <翻译结束>


<原文开始>
			// Real path element.
			// Add slash if needed
<原文结束>

# <翻译开始>
// 实路径元素
// 必要时添加斜杠
# <翻译结束>


<原文开始>
			// Copy element
<原文结束>

# <翻译开始>
// 复制的元素
# <翻译结束>


<原文开始>
	// Re-append trailing slash
<原文结束>

# <翻译开始>
// 重新添加尾斜杠
# <翻译结束>


<原文开始>
	// If the original string was not modified (or only shortened at the end),
	// return the respective substring of the original string.
	// Otherwise return a new string from the buffer.
<原文结束>

# <翻译开始>
// 如果原始字符串未被修改(或仅在末尾缩短)，则返回原始字符串的相应子字符串
// 否则从缓冲区返回一个新字符串
# <翻译结束>


<原文开始>
// Internal helper to lazily create a buffer if necessary.
// Calls to this function get inlined.
<原文结束>

# <翻译开始>
// 内部帮助器在必要时惰性地创建缓冲区
// 对这个函数的调用被内联
# <翻译结束>


<原文开始>
		// No modification of the original string so far.
		// If the next character is the same as in the original string, we do
		// not yet have to allocate a buffer.
<原文结束>

# <翻译开始>
// 到目前为止没有修改原始字符串
// 如果下一个字符与原始字符串中的字符相同，则不需要分配缓冲区
# <翻译结束>


<原文开始>
		// Otherwise use either the stack buffer, if it is large enough, or
		// allocate a new buffer on the heap, and copy all previous characters.
<原文结束>

# <翻译开始>
// 否则，要么使用堆栈缓冲区(如果它足够大)，要么在堆上分配一个新的缓冲区，并复制前面的所有字符
# <翻译结束>

