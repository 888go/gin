
<原文开始>
// Copyright 2013 Julien Schmidt. All rights reserved.
// Based on the path package, Copyright 2009 The Go Authors.
// Use of this source code is governed by a BSD-style license that can be found
// at https://github.com/julienschmidt/httprouter/blob/master/LICENSE.
<原文结束>

# <翻译开始>
// 版权声明：2013年 Julien Schmidt。保留所有权利。
// 本代码基于2009年 The Go Authors 的 path 包。
// 使用本源代码须遵循 BSD 风格的许可证，该许可证可在
// https://github.com/julienschmidt/httprouter/blob/master/LICENSE 查阅。
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
// cleanPath 是 path.Clean 函数在 URL 版本中的实现，它返回一个规范化的 URL 路径 p，消除其中的 "." 和 ".." 元素。
//
// 以下规则会迭代应用，直到无法进一步处理为止：
//  1. 将多个斜杠替换为单个斜杠。
//  2. 消除每个表示当前目录的 "." 路径名称元素。
//  3. 消除每个表示父目录的内层 ".." 路径名称元素及其前面紧随的非 ".." 元素。
//  4. 消除根路径开始处的 ".." 元素：即，在路径开头将 "/.." 替换为 "/"。
//
// 如果此过程的结果为空字符串，则返回 "/"。
# <翻译结束>


<原文开始>
// Turn empty string into "/"
<原文结束>

# <翻译开始>
// 将空字符串转换为 "/"
# <翻译结束>


<原文开始>
	// Reasonably sized buffer on stack to avoid allocations in the common case.
	// If a larger buffer is required, it gets allocated dynamically.
<原文结束>

# <翻译开始>
// 为避免在常见场景中进行内存分配，栈上预留了适度大小的缓冲区。
// 如果需要更大的缓冲区，则会动态分配。
# <翻译结束>


<原文开始>
	// Invariants:
	//      reading from path; r is index of next byte to process.
	//      writing to buf; w is index of next byte to write.
<原文结束>

# <翻译开始>
// 保持不变的条件（不变式）：
//      正从 path 进行读取操作，r 是待处理的下一个字节索引。
//      正向 buf 进行写入操作，w 是待写入的下一个字节索引。
# <翻译结束>


<原文开始>
// path must start with '/'
<原文结束>

# <翻译开始>
// 路径必须以'/'开头
# <翻译结束>


<原文开始>
	// A bit more clunky without a 'lazybuf' like the path package, but the loop
	// gets completely inlined (bufApp calls).
	// loop has no expensive function calls (except 1x make)		// So in contrast to the path package this loop has no expensive function
	// calls (except make, if needed).
<原文结束>

# <翻译开始>
// 没有像 path 包中的 'lazybuf' 那样精巧，但循环会完全内联（bufApp 调用）。
// 循环中没有昂贵的函数调用（除了最多 1 次 make）。 // 因此，与 path 包相比，这个循环中没有昂贵的函数调用（除了必要时的 make）。
# <翻译结束>


<原文开始>
// empty path element, trailing slash is added after the end
<原文结束>

# <翻译开始>
// 空路径元素，末尾的斜杠会在结尾处添加
# <翻译结束>


<原文开始>
// .. element: remove to last /
<原文结束>

# <翻译开始>
// .. element: 移除到末尾 /
# <翻译结束>


<原文开始>
			// Real path element.
			// Add slash if needed
<原文结束>

# <翻译开始>
// 真实路径元素
// 如有必要添加斜杠
# <翻译结束>


<原文开始>
// Re-append trailing slash
<原文结束>

# <翻译开始>
// 重新追加尾部斜杠
# <翻译结束>


<原文开始>
	// If the original string was not modified (or only shortened at the end),
	// return the respective substring of the original string.
	// Otherwise return a new string from the buffer.
<原文结束>

# <翻译开始>
// 如果原始字符串未被修改（或者只是在末尾缩短了），
// 则返回原始字符串的相应子串。
// 否则，从缓冲区返回一个新的字符串。
# <翻译结束>


<原文开始>
// Internal helper to lazily create a buffer if necessary.
// Calls to this function get inlined.
<原文结束>

# <翻译开始>
// 内部辅助函数，按需惰性创建缓冲区。
// 对该函数的调用会被内联处理。
# <翻译结束>


<原文开始>
		// No modification of the original string so far.
		// If the next character is the same as in the original string, we do
		// not yet have to allocate a buffer.
<原文结束>

# <翻译开始>
// 到目前为止，原始字符串尚未被修改。
// 如果下一个字符与原始字符串中的相同，我们还不需要分配缓冲区。
# <翻译结束>


<原文开始>
		// Otherwise use either the stack buffer, if it is large enough, or
		// allocate a new buffer on the heap, and copy all previous characters.
<原文结束>

# <翻译开始>
// 否则，如果堆栈缓冲区足够大，则使用堆栈缓冲区；否则，在堆上分配一个新的缓冲区，并复制之前的所有字符。
# <翻译结束>

