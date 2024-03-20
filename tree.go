// 版权所有 ? 2013 Julien Schmidt。保留所有权利。
// 本源代码的使用受 BSD 风格许可协议约束，该协议可在
// https://github.com/julienschmidt/httprouter/blob/master/LICENSE 查阅

package gin

import (
	"bytes"
	"net/url"
	"strings"
	"unicode"
	"unicode/utf8"
	
	"github.com/888go/gin/internal/bytesconv"
)

var (
	strColon = []byte(":")
	strStar  = []byte("*")
	strSlash = []byte("/")
)

// Param 是单个URL参数，包含一个键和一个值。
type Param struct {
	Key   string
	Value string
}

// Params 是一个 Param 切片，由路由器返回。
// 这个切片是有序的，第一个 URL 参数也是切片中的第一个值。
// 因此，通过索引读取值是安全的。
type Params []Param

// Get 方法返回第一个参数键与给定名称相匹配的值，并返回布尔值 true。
// 如果未找到匹配的参数，则返回一个空字符串和布尔值 false。

// ff:
// name:
func (ps Params) Get(name string) (string, bool) {
	for _, entry := range ps {
		if entry.Key == name {
			return entry.Value, true
		}
	}
	return "", false
}

// ByName 返回第一个其键与给定名称相匹配的Param的值。
// 如果未找到匹配的Param，则返回一个空字符串。

// ff:
// va:
// name:
func (ps Params) ByName(name string) (va string) {
	va, _ = ps.Get(name)
	return
}

type methodTree struct {
	method string
	root   *node
}

type methodTrees []methodTree

func (trees methodTrees) get(method string) *node {
	for _, tree := range trees {
		if tree.method == method {
			return tree.root
		}
	}
	return nil
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func longestCommonPrefix(a, b string) int {
	i := 0
	max := min(len(a), len(b))
	for i < max && a[i] == b[i] {
		i++
	}
	return i
}

// addChild 将添加一个子节点，并保持 wildcardChild 位于末尾
func (n *node) addChild(child *node) {
	if n.wildChild && len(n.children) > 0 {
		wildcardChild := n.children[len(n.children)-1]
		n.children = append(n.children[:len(n.children)-1], child, wildcardChild)
	} else {
		n.children = append(n.children, child)
	}
}

func countParams(path string) uint16 {
	var n uint16
	s := bytesconv.StringToBytes(path)
	n += uint16(bytes.Count(s, strColon))
	n += uint16(bytes.Count(s, strStar))
	return n
}

func countSections(path string) uint16 {
	s := bytesconv.StringToBytes(path)
	return uint16(bytes.Count(s, strSlash))
}

type nodeType uint8

const (
	static nodeType = iota
	root
	param
	catchAll
)

type node struct {
	path      string
	indices   string
	wildChild bool
	nType     nodeType
	priority  uint32
	children  []*node // 子节点，数组末尾最多包含 1 个 :param 样式的节点
	handlers  HandlersChain
	fullPath  string
}

// 如果有必要，递增给定子项的优先级并重新排序
func (n *node) incrementChildPrio(pos int) int {
	cs := n.children
	cs[pos].priority++
	prio := cs[pos].priority

	// 调整位置（移动到前面）
	newPos := pos
	for ; newPos > 0 && cs[newPos-1].priority < prio; newPos-- {
		// Swap node positions
		cs[newPos-1], cs[newPos] = cs[newPos], cs[newPos-1]
	}

	// 构建新的索引字符字符串
	if newPos != pos {
		n.indices = n.indices[:newPos] + // 未更改的前缀，可能为空
			n.indices[pos:pos+1] + // The index char we move
			n.indices[newPos:pos] + n.indices[pos+1:] // 在'pos'位置没有字符的Rest
	}

	return newPos
}

// addRoute 向路径中添加具有给定 handle 的节点。
// 非并发安全！
func (n *node) addRoute(path string, handlers HandlersChain) {
	fullPath := path
	n.priority++

	// Empty tree
	if len(n.path) == 0 && len(n.children) == 0 {
		n.insertChild(path, fullPath, handlers)
		n.nType = root
		return
	}

	parentFullPathIndex := 0

walk:
	for {
// 查找最长公共前缀。
// 这也意味着，公共前缀中不包含 ':' 或 '*'，
// 因为已存在的键不能包含这些字符。
		i := longestCommonPrefix(path, n.path)

		// Split edge
		if i < len(n.path) {
			child := node{
				path:      n.path[i:],
				wildChild: n.wildChild,
				nType:     static,
				indices:   n.indices,
				children:  n.children,
				handlers:  n.handlers,
				priority:  n.priority - 1,
				fullPath:  n.fullPath,
			}

			n.children = []*node{&child}
			// 使用 []byte 以正确处理 Unicode 字符转换，参考 #65 号问题
			n.indices = bytesconv.BytesToString([]byte{n.path[i]})
			n.path = path[:i]
			n.handlers = nil
			n.wildChild = false
			n.fullPath = fullPath[:parentFullPathIndex+i]
		}

		// 将新节点作为此节点的子节点
		if i < len(path) {
			path = path[i:]
			c := path[0]

			// '/' after param
			if n.nType == param && c == '/' && len(n.children) == 1 {
				parentFullPathIndex += len(n.path)
				n = n.children[0]
				n.priority++
				continue walk
			}

			// 检查是否存在下一个路径字节的子节点
			for i, max := 0, len(n.indices); i < max; i++ {
				if c == n.indices[i] {
					parentFullPathIndex += len(n.path)
					i = n.incrementChildPrio(i)
					n = n.children[i]
					continue walk
				}
			}

			// Otherwise insert it
			if c != ':' && c != '*' && n.nType != catchAll {
				// 使用 []byte 以正确处理 Unicode 字符转换，参考 #65 号问题
				n.indices += bytesconv.BytesToString([]byte{c})
				child := &node{
					fullPath: fullPath,
				}
				n.addChild(child)
				n.incrementChildPrio(len(n.indices) - 1)
				n = child
			} else if n.wildChild {
				// 插入通配符节点，需要检查是否与现有通配符冲突
				n = n.children[len(n.children)-1]
				n.priority++

				// 检查通配符是否匹配
				if len(path) >= len(n.path) && n.path == path[:len(n.path)] &&
					// 向 catchAll 添加子节点是不可能的
					n.nType != catchAll &&
					// 检查较长的通配符，例如:name和:names
					(len(n.path) >= len(path) || path[len(n.path)] == '/') {
					continue walk
				}

				// Wildcard conflict
				pathSeg := path
				if n.nType != catchAll {
					pathSeg = strings.SplitN(pathSeg, "/", 2)[0]
				}
				prefix := fullPath[:strings.Index(fullPath, pathSeg)] + n.path
				panic("'" + pathSeg +
					"' in new path '" + fullPath +
					"' conflicts with existing wildcard '" + n.path +
					"' in existing prefix '" + prefix +
					"'")
			}

			n.insertChild(path, fullPath, handlers)
			return
		}

		// 否则将处理程序添加到当前节点
		if n.handlers != nil {
			panic("handlers are already registered for path '" + fullPath + "'")
		}
		n.handlers = handlers
		n.fullPath = fullPath
		return
	}
}

// 搜索通配符段，并检查名称中是否存在无效字符。
// 如果未找到通配符，则返回索引 -1。
func findWildcard(path string) (wildcard string, i int, valid bool) {
	// Find start
	for start, c := range []byte(path) {
		// 通配符以 ':'（参数）或 '*'（捕获全部）开始
		if c != ':' && c != '*' {
			continue
		}

		// 查找结尾并检查无效字符
		valid = true
		for end, c := range []byte(path[start+1:]) {
			switch c {
			case '/':
				return path[start : start+1+end], start, valid
			case ':', '*':
				valid = false
			}
		}
		return path[start:], start, valid
	}
	return "", -1, false
}

func (n *node) insertChild(path string, fullPath string, handlers HandlersChain) {
	for {
		// 查找直到第一个通配符为止的前缀
		wildcard, i, valid := findWildcard(path)
		if i < 0 { // No wildcard found
			break
		}

		// 通配符名称中只能包含一个':'或'*'字符
		if !valid {
			panic("only one wildcard per path segment is allowed, has: '" +
				wildcard + "' in path '" + fullPath + "'")
		}

		// 检查通配符是否有名称
		if len(wildcard) < 2 {
			panic("wildcards must be named with a non-empty name in path '" + fullPath + "'")
		}

		if wildcard[0] == ':' { // param
			if i > 0 {
				// 在当前通配符前插入前缀
				n.path = path[:i]
				path = path[i:]
			}

			child := &node{
				nType:    param,
				path:     wildcard,
				fullPath: fullPath,
			}
			n.addChild(child)
			n.wildChild = true
			n = child
			n.priority++

// 如果路径没有以通配符结尾，则会有一个以'/'开头的另一个子路径
			if len(wildcard) < len(path) {
				path = path[len(wildcard):]

				child := &node{
					priority: 1,
					fullPath: fullPath,
				}
				n.addChild(child)
				n = child
				continue
			}

			// 否则我们已经完成。将句柄插入新的叶子节点
			n.handlers = handlers
			return
		}

		// catchAll
		if i+len(wildcard) != len(path) {
			panic("catch-all routes are only allowed at the end of the path in path '" + fullPath + "'")
		}

		if len(n.path) > 0 && n.path[len(n.path)-1] == '/' {
			pathSeg := strings.SplitN(n.children[0].path, "/", 2)[0]
			panic("catch-all wildcard '" + path +
				"' in new path '" + fullPath +
				"' conflicts with existing path segment '" + pathSeg +
				"' in existing prefix '" + n.path + pathSeg +
				"'")
		}

		// 当前固定宽度为1用于 '/'
		i--
		if path[i] != '/' {
			panic("no / before catch-all in path '" + fullPath + "'")
		}

		n.path = path[:i]

		// 第一个节点：匹配所有路径的“catchAll”节点，其路径为空
		child := &node{
			wildChild: true,
			nType:     catchAll,
			fullPath:  fullPath,
		}

		n.addChild(child)
		n.indices = string('/')
		n = child
		n.priority++

		// 第二个节点：存储变量的节点
		child = &node{
			path:     path[i:],
			nType:    catchAll,
			handlers: handlers,
			priority: 1,
			fullPath: fullPath,
		}
		n.children = []*node{child}

		return
	}

	// 如果没有找到通配符，直接插入路径并处理
	n.path = path
	n.handlers = handlers
	n.fullPath = fullPath
}

// nodeValue 用于存储 (*Node).getValue 方法的返回值
type nodeValue struct {
	handlers HandlersChain
	params   *Params
	tsr      bool
	fullPath string
}

type skippedNode struct {
	path        string
	node        *node
	paramsCount int16
}

// 根据给定路径（键）返回已注册的处理程序。通配符的值将保存到一个映射中。
// 如果找不到处理程序，且对于给定路径存在一个带有额外（无）尾部斜杠的处理程序，则会提出 TSR（尾部斜杠重定向）建议。
func (n *node) getValue(path string, params *Params, skippedNodes *[]skippedNode, unescape bool) (value nodeValue) {
	var globalParamsCount int16

walk: // 外层循环，用于遍历树
	for {
		prefix := n.path
		if len(path) > len(prefix) {
			if path[:len(prefix)] == prefix {
				path = path[len(prefix):]

				// 首先尝试通过匹配索引来查找所有非通配符子节点
				idxc := path[0]
				for i, c := range []byte(n.indices) {
					if c == idxc {
						// 判断n的子节点列表中最后一个子节点的路径是否以":"开头，结果与n.wildChild属性相等
						if n.wildChild {
							index := len(*skippedNodes)
							*skippedNodes = (*skippedNodes)[:index+1]
							(*skippedNodes)[index] = skippedNode{
								path: prefix + path,
								node: &node{
									path:      n.path,
									wildChild: n.wildChild,
									nType:     n.nType,
									priority:  n.priority,
									children:  n.children,
									handlers:  n.handlers,
									fullPath:  n.fullPath,
								},
								paramsCount: globalParamsCount,
							}
						}

						n = n.children[i]
						continue walk
					}
				}

				if !n.wildChild {
// 如果循环结束时路径不等于'/'且当前节点没有子节点，
// 那么当前节点需要回滚到上一个有效的已跳过节点
					if path != "/" {
						for length := len(*skippedNodes); length > 0; length-- {
							skippedNode := (*skippedNodes)[length-1]
							*skippedNodes = (*skippedNodes)[:length-1]
							if strings.HasSuffix(skippedNode.path, path) {
								path = skippedNode.path
								n = skippedNode.node
								if value.params != nil {
									*value.params = (*value.params)[:skippedNode.paramsCount]
								}
								globalParamsCount = skippedNode.paramsCount
								continue walk
							}
						}
					}

// 未找到任何内容。
// 如果该路径存在叶节点，我们可以建议重定向到不带尾部斜杠的相同 URL。
					value.tsr = path == "/" && n.handlers != nil
					return
				}

				// 处理通配符子节点，它总是在数组的末尾
				n = n.children[len(n.children)-1]
				globalParamsCount++

				switch n.nType {
				case param:
// 修复：截断参数
// 文件：tree_test.go，行号：204

					// 查找参数结束位置（可能是'/'或路径结束）
					end := 0
					for end < len(path) && path[end] != '/' {
						end++
					}

					// Save param value
					if params != nil {
						// 如果有必要，预先分配容量
						if cap(*params) < int(globalParamsCount) {
							newParams := make(Params, len(*params), globalParamsCount)
							copy(newParams, *params)
							*params = newParams
						}

						if value.params == nil {
							value.params = params
						}
						// 在预先分配的容量内扩展切片
						i := len(*value.params)
						*value.params = (*value.params)[:i+1]
						val := path[:end]
						if unescape {
							if v, err := url.QueryUnescape(val); err == nil {
								val = v
							}
						}
						(*value.params)[i] = Param{
							Key:   n.path[1:],
							Value: val,
						}
					}

					// we need to go deeper!
					if end < len(path) {
						if len(n.children) > 0 {
							path = path[end:]
							n = n.children[0]
							continue walk
						}

						// ... but we can't
						value.tsr = len(path) == end+1
						return
					}

					if value.handlers = n.handlers; value.handlers != nil {
						value.fullPath = n.fullPath
						return
					}
					if len(n.children) == 1 {
// 未找到处理程序。检查是否存在该路径加上末尾斜杠的处理程序，以便进行 TSR（可能指“透明目录重写”）推荐
						n = n.children[0]
						value.tsr = (n.path == "/" && n.handlers != nil) || (n.path == "" && n.indices == "/")
					}
					return

				case catchAll:
					// Save param value
					if params != nil {
						// 如果有必要，预先分配容量
						if cap(*params) < int(globalParamsCount) {
							newParams := make(Params, len(*params), globalParamsCount)
							copy(newParams, *params)
							*params = newParams
						}

						if value.params == nil {
							value.params = params
						}
						// 在预先分配的容量内扩展切片
						i := len(*value.params)
						*value.params = (*value.params)[:i+1]
						val := path
						if unescape {
							if v, err := url.QueryUnescape(path); err == nil {
								val = v
							}
						}
						(*value.params)[i] = Param{
							Key:   n.path[2:],
							Value: val,
						}
					}

					value.handlers = n.handlers
					value.fullPath = n.fullPath
					return

				default:
					panic("invalid node type")
				}
			}
		}

		if path == prefix {
// 如果当前路径不等于'/'，并且节点未注册处理程序，并且最近匹配的节点有一个子节点
// 那么当前节点需要回滚到上一个有效跳过的节点
			if n.handlers == nil && path != "/" {
				for length := len(*skippedNodes); length > 0; length-- {
					skippedNode := (*skippedNodes)[length-1]
					*skippedNodes = (*skippedNodes)[:length-1]
					if strings.HasSuffix(skippedNode.path, path) {
						path = skippedNode.path
						n = skippedNode.node
						if value.params != nil {
							*value.params = (*value.params)[:skippedNode.paramsCount]
						}
						globalParamsCount = skippedNode.paramsCount
						continue walk
					}
				}
				// n = latestNode.children数组的最后一个元素
			}
// 我们应该已经到达包含处理程序的节点。
// 检查该节点是否注册了处理程序。
			if value.handlers = n.handlers; value.handlers != nil {
				value.fullPath = n.fullPath
				return
			}

// 如果该路由没有处理程序，但该路由有一个通配符子路由，则必须为此路径（带有一个额外的尾部斜杠）提供一个处理程序
			if path == "/" && n.wildChild && n.nType != root {
				value.tsr = true
				return
			}

			if path == "/" && n.nType == static {
				value.tsr = true
				return
			}

// 未找到处理程序。检查是否存在针对此路径加上末尾斜杠的处理程序，以便提供末尾斜杠建议
			for i, c := range []byte(n.indices) {
				if c == '/' {
					n = n.children[i]
					value.tsr = (len(n.path) == 1 && n.handlers != nil) ||
						(n.nType == catchAll && n.children[0].handlers != nil)
					return
				}
			}

			return
		}

// 未找到任何内容。如果该路径存在叶子节点，我们可以建议重定向到同一 URL 并附加一个额外的尾部斜杠
		value.tsr = path == "/" ||
			(len(prefix) == len(path)+1 && prefix[len(path)] == '/' &&
				path == prefix[:len(prefix)-1] && n.handlers != nil)

		// 回滚到最后一个有效的跳过节点
		if !value.tsr && path != "/" {
			for length := len(*skippedNodes); length > 0; length-- {
				skippedNode := (*skippedNodes)[length-1]
				*skippedNodes = (*skippedNodes)[:length-1]
				if strings.HasSuffix(skippedNode.path, path) {
					path = skippedNode.path
					n = skippedNode.node
					if value.params != nil {
						*value.params = (*value.params)[:skippedNode.paramsCount]
					}
					globalParamsCount = skippedNode.paramsCount
					continue walk
				}
			}
		}

		return
	}
}

// 根据给定的路径进行不区分大小写的查找，并尝试找到一个处理器。
// 可选地，它也可以修正尾部的斜杠。
// 它返回大小写纠正后的路径以及一个布尔值，表示查找是否成功。
func (n *node) findCaseInsensitivePath(path string, fixTrailingSlash bool) ([]byte, bool) {
	const stackBufSize = 128

// 在常见情况下，使用栈上静态大小的缓冲区。如果路径过长，则改为在堆上分配缓冲区。
	buf := make([]byte, 0, stackBufSize)
	if length := len(path) + 1; length > stackBufSize {
		buf = make([]byte, 0, length)
	}

	ciPath := n.findCaseInsensitivePathRec(
		path,
		buf,       // 预先分配足够的内存用于新路径
		[4]byte{}, // Empty rune buffer
		fixTrailingSlash,
	)

	return ciPath, ciPath != nil
}

// 将数组中的字节向左移动n个字节
func shiftNRuneBytes(rb [4]byte, n int) [4]byte {
	switch n {
	case 0:
		return rb
	case 1:
		return [4]byte{rb[1], rb[2], rb[3], 0}
	case 2:
		return [4]byte{rb[2], rb[3]}
	case 3:
		return [4]byte{rb[3]}
	default:
		return [4]byte{}
	}
}

// 递归的大小写不敏感查找函数，由n.findCaseInsensitivePath调用
func (n *node) findCaseInsensitivePathRec(path string, ciPath []byte, rb [4]byte, fixTrailingSlash bool) []byte {
	npLen := len(n.path)

walk: // 外层循环，用于遍历树
	for len(path) >= npLen && (npLen == 0 || strings.EqualFold(path[1:npLen], n.path[1:])) {
		// 向结果添加公共前缀
		oldPath := path
		path = path[npLen:]
		ciPath = append(ciPath, n.path...)

		if len(path) == 0 {
// 我们应该已经到达包含处理程序的节点。
// 检查该节点是否注册了处理程序。
			if n.handlers != nil {
				return ciPath
			}

// 未找到处理程序。
// 尝试通过添加尾部斜杠来修复路径
			if fixTrailingSlash {
				for i, c := range []byte(n.indices) {
					if c == '/' {
						n = n.children[i]
						if (len(n.path) == 1 && n.handlers != nil) ||
							(n.nType == catchAll && n.children[0].handlers != nil) {
							return append(ciPath, '/')
						}
						return nil
					}
				}
			}
			return nil
		}

// 如果该节点没有通配符（param 或 catchAll）子节点，
// 我们可以直接查找下一个子节点并继续向下遍历树
		if !n.wildChild {
			// 跳过已处理的 rune 字节
			rb = shiftNRuneBytes(rb, npLen)

			if rb[0] != 0 {
				// Old rune not finished
				idxc := rb[0]
				for i, c := range []byte(n.indices) {
					if c == idxc {
						// 继续处理子节点
						n = n.children[i]
						npLen = len(n.path)
						continue walk
					}
				}
			} else {
				// Process a new rune
				var rv rune

// 查找rune的起始位置。
// rune字符可能包含最多4个字节，
// 因此，-4肯定属于另一个rune字符。
				var off int
				for max := min(npLen, 3); off < max; off++ {
					if i := npLen - off; utf8.RuneStart(oldPath[i]) {
						// 从缓存路径读取单个字符
						rv, _ = utf8.DecodeRuneInString(oldPath[i:])
						break
					}
				}

				// 计算当前字符的 lowercase 字节
				lo := unicode.ToLower(rv)
				utf8.EncodeRune(rb[:], lo)

				// 跳过已处理的字节
				rb = shiftNRuneBytes(rb, off)

				idxc := rb[0]
				for i, c := range []byte(n.indices) {
					// Lowercase matches
					if c == idxc {
// 必须采用递归方法，因为大小写两种字节都可能存在作为索引
						if out := n.children[i].findCaseInsensitivePathRec(
							path, ciPath, rb, fixTrailingSlash,
						); out != nil {
							return out
						}
						break
					}
				}

// 如果我们没有找到匹配项，对于大写字符执行相同的操作，
// 如果它与小写字符不同的话
				if up := unicode.ToUpper(rv); up != lo {
					utf8.EncodeRune(rb[:], up)
					rb = shiftNRuneBytes(rb, off)

					idxc := rb[0]
					for i, c := range []byte(n.indices) {
						// Uppercase matches
						if c == idxc {
							// 继续处理子节点
							n = n.children[i]
							npLen = len(n.path)
							continue walk
						}
					}
				}
			}

// 未找到任何内容。如果该路径存在叶子节点，我们可以建议重定向到去掉尾部斜杠的相同URL
			if fixTrailingSlash && path == "/" && n.handlers != nil {
				return ciPath
			}
			return nil
		}

		n = n.children[0]
		switch n.nType {
		case param:
			// 查找参数结束位置（可能是'/'或路径结束）
			end := 0
			for end < len(path) && path[end] != '/' {
				end++
			}

			// 添加参数值到不区分大小写的路径
			ciPath = append(ciPath, path[:end]...)

			// We need to go deeper!
			if end < len(path) {
				if len(n.children) > 0 {
					// 继续处理子节点
					n = n.children[0]
					npLen = len(n.path)
					path = path[end:]
					continue
				}

				// ... but we can't
				if fixTrailingSlash && len(path) == end+1 {
					return ciPath
				}
				return nil
			}

			if n.handlers != nil {
				return ciPath
			}

			if fixTrailingSlash && len(n.children) == 1 {
				// No handle found. Check if a handle for this path + a
				// trailing slash exists
				n = n.children[0]
				if n.path == "/" && n.handlers != nil {
					return append(ciPath, '/')
				}
			}

			return nil

		case catchAll:
			return append(ciPath, path...)

		default:
			panic("invalid node type")
		}
	}

// 未找到任何内容。
// 尝试通过添加或删除尾部斜杠来修复路径
	if fixTrailingSlash {
		if path == "/" {
			return ciPath
		}
		if len(path)+1 == npLen && n.path[len(path)] == '/' &&
			strings.EqualFold(path[1:], n.path[1:len(path)]) && n.handlers != nil {
			return append(ciPath, n.path...)
		}
	}
	return nil
}
