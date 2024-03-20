// 版权声明：2013年 Julien Schmidt。保留所有权利。
// 本代码基于2009年 The Go Authors 的 path 包。
// 使用本源代码须遵循 BSD 风格的许可证，该许可证可在
// https://github.com/julienschmidt/httprouter/blob/master/LICENSE 查阅。

package gin

// cleanPath 是 path.Clean 函数在 URL 版本中的实现，它返回一个规范化的 URL 路径 p，消除其中的 "." 和 ".." 元素。
//
// 以下规则会迭代应用，直到无法进一步处理为止：
//  1. 将多个斜杠替换为单个斜杠。
//  2. 消除每个表示当前目录的 "." 路径名称元素。
//  3. 消除每个表示父目录的内层 ".." 路径名称元素及其前面紧随的非 ".." 元素。
//  4. 消除根路径开始处的 ".." 元素：即，在路径开头将 "/.." 替换为 "/"。
//
// 如果此过程的结果为空字符串，则返回 "/"。
func cleanPath(p string) string {
	const stackBufSize = 128
	// 将空字符串转换为 "/"
	if p == "" {
		return "/"
	}

// 为避免在常见场景中进行内存分配，栈上预留了适度大小的缓冲区。
// 如果需要更大的缓冲区，则会动态分配。
	buf := make([]byte, 0, stackBufSize)

	n := len(p)

// 保持不变的条件（不变式）：
//      正从 path 进行读取操作，r 是待处理的下一个字节索引。
//      正向 buf 进行写入操作，w 是待写入的下一个字节索引。

	// 路径必须以'/'开头
	r := 1
	w := 1

	if p[0] != '/' {
		r = 0

		if n+1 > stackBufSize {
			buf = make([]byte, n+1)
		} else {
			buf = buf[:n+1]
		}
		buf[0] = '/'
	}

	trailing := n > 1 && p[n-1] == '/'

// 没有像 path 包中的 'lazybuf' 那样精巧，但循环会完全内联（bufApp 调用）。
// 循环中没有昂贵的函数调用（除了最多 1 次 make）。 // 因此，与 path 包相比，这个循环中没有昂贵的函数调用（除了必要时的 make）。

	for r < n {
		switch {
		case p[r] == '/':
			// 空路径元素，末尾的斜杠会在结尾处添加
			r++

		case p[r] == '.' && r+1 == n:
			trailing = true
			r++

		case p[r] == '.' && p[r+1] == '/':
			// . element
			r += 2

		case p[r] == '.' && p[r+1] == '.' && (r+2 == n || p[r+2] == '/'):
			// .. element: 移除到末尾 /
			r += 3

			if w > 1 {
				// can backtrack
				w--

				if len(buf) == 0 {
					for w > 1 && p[w] != '/' {
						w--
					}
				} else {
					for w > 1 && buf[w] != '/' {
						w--
					}
				}
			}

		default:
// 真实路径元素
// 如有必要添加斜杠
			if w > 1 {
				bufApp(&buf, p, w, '/')
				w++
			}

			// Copy element
			for r < n && p[r] != '/' {
				bufApp(&buf, p, w, p[r])
				w++
				r++
			}
		}
	}

	// 重新追加尾部斜杠
	if trailing && w > 1 {
		bufApp(&buf, p, w, '/')
		w++
	}

// 如果原始字符串未被修改（或者只是在末尾缩短了），
// 则返回原始字符串的相应子串。
// 否则，从缓冲区返回一个新的字符串。
	if len(buf) == 0 {
		return p[:w]
	}
	return string(buf[:w])
}

// 内部辅助函数，按需惰性创建缓冲区。
// 对该函数的调用会被内联处理。
func bufApp(buf *[]byte, s string, w int, c byte) {
	b := *buf
	if len(b) == 0 {
// 到目前为止，原始字符串尚未被修改。
// 如果下一个字符与原始字符串中的相同，我们还不需要分配缓冲区。
		if s[w] == c {
			return
		}

// 否则，如果堆栈缓冲区足够大，则使用堆栈缓冲区；否则，在堆上分配一个新的缓冲区，并复制之前的所有字符。
		length := len(s)
		if length > cap(b) {
			*buf = make([]byte, length)
		} else {
			*buf = (*buf)[:length]
		}
		b = *buf

		copy(b, s[:w])
	}
	b[w] = c
}
