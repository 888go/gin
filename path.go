// 版权所有2013朱利安施密特
// 版权所有
// 基于路径包，版权归the Go Authors所有
// 此源代码的使用受bsd风格的许可证的约束，该许可证可在https://github.com/julienschmidt/httprouter/blob/master/LICENSE上找到

package gin

// cleanPath是path的URL版本
// 干净，它返回p的规范URL路径，消除
// 和. .元素
// 迭代地应用以下规则，直到无法进行进一步处理为止:将多个斜杠替换为单个斜杠
// 2. 消除每一个
// 路径名元素(当前目录)
// 3. 消除每个内部…路径名元素(父目录)以及非-..它前面的元素
// 4. 消除……开始根路径的元素:即替换"/.."由“/”;在一条路的起点
// 如果此过程的结果为空
func cleanPath(p string) string {
	const stackBufSize = 128
// 将空字符串转换为"/"
	if p == "" {
		return "/"
	}

// 合理大小的堆栈缓冲区，以避免在通常情况下分配
// 如果需要更大的缓冲区，则动态分配
	buf := make([]byte, 0, stackBufSize)

	n := len(p)

// 不变量:从path读取;R是要处理的下一个字节的索引
// 给……写信;W是要写入的下一个字节的索引

// 路径必须以“/”开头
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

// 没有像path包那样的“lazybuf”会更笨拙一些，但循环会完全内联(bufApp调用)
// 循环没有昂贵的函数调用(除了1x make)所以与path包相比，这个循环没有昂贵的函数调用(除了make，如果需要的话)

	for r < n {
		switch {
		case p[r] == '/':
// 空路径元素，结尾后添加斜杠
			r++

		case p[r] == '.' && r+1 == n:
			trailing = true
			r++

		case p[r] == '.' && p[r+1] == '/':
// ． 元素
			r += 2

		case p[r] == '.' && p[r+1] == '.' && (r+2 == n || p[r+2] == '/'):
// ．． 元素:移到最后
			r += 3

			if w > 1 {
// 可以回溯
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
// 实路径元素
// 必要时添加斜杠
			if w > 1 {
				bufApp(&buf, p, w, '/')
				w++
			}

// 复制的元素
			for r < n && p[r] != '/' {
				bufApp(&buf, p, w, p[r])
				w++
				r++
			}
		}
	}

// 重新添加尾斜杠
	if trailing && w > 1 {
		bufApp(&buf, p, w, '/')
		w++
	}

// 如果原始字符串未被修改(或仅在末尾缩短)，则返回原始字符串的相应子字符串
// 否则从缓冲区返回一个新字符串
	if len(buf) == 0 {
		return p[:w]
	}
	return string(buf[:w])
}

// 内部帮助器在必要时惰性地创建缓冲区
// 对这个函数的调用被内联
func bufApp(buf *[]byte, s string, w int, c byte) {
	b := *buf
	if len(b) == 0 {
// 到目前为止没有修改原始字符串
// 如果下一个字符与原始字符串中的字符相同，则不需要分配缓冲区
		if s[w] == c {
			return
		}

// 否则，要么使用堆栈缓冲区(如果它足够大)，要么在堆上分配一个新的缓冲区，并复制前面的所有字符
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
