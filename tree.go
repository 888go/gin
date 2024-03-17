// 版权所有2013朱利安施密特
// 版权所有
// 此源代码的使用受bsd风格的许可证的约束，该许可证可在https://github.com/julienschmidt/httprouter/blob/master/LICENSE上找到

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

// Param是一个URL参数，由一个键和一个值组成
type Param struct {
	Key   string
	Value string
}

// Params是一个Param-slice，由路由器返回
// 片是有序的，第一个URL参数也是第一个片值
// 因此，通过索引读取值是安全的
type Params []Param

// Get返回与给定名称匹配的第一个Param的值和一个布尔值true
// 如果没有找到匹配的参数，则返回一个空字符串和一个布尔值false

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

// ByName返回与给定名称匹配的第一个参数的值
// 如果没有找到匹配的Param，则返回一个空字符串

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

// addChild将添加一个子节点，将wildcardChild保留在最后
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
	children  []*node // 子节点，最多1个:数组末尾的参数样式节点
	handlers  HandlersChain
	fullPath  string
}

// 增加给定子元素的优先级，并在必要时重新排序
func (n *node) incrementChildPrio(pos int) int {
	cs := n.children
	cs[pos].priority++
	prio := cs[pos].priority

// 调整位置(移动到前面)
	newPos := pos
	for ; newPos > 0 && cs[newPos-1].priority < prio; newPos-- {
// 交换节点位置
		cs[newPos-1], cs[newPos] = cs[newPos], cs[newPos-1]
	}

// 构建新的索引字符串
	if newPos != pos {
		n.indices = n.indices[:newPos] + // 未更改前缀，可能为空
			n.indices[pos:pos+1] + // 我们移动的索引字符
			n.indices[newPos:pos] + n.indices[pos+1:] // 在'pos'位置移除字符后的剩余部分
	}

	return newPos
}

// addRoute将具有给定句柄的节点添加到路径中
// 不是concurrency-safe !
func (n *node) addRoute(path string, handlers HandlersChain) {
	fullPath := path
	n.priority++

// 空树
	if len(n.path) == 0 && len(n.children) == 0 {
		n.insertChild(path, fullPath, handlers)
		n.nType = root
		return
	}

	parentFullPathIndex := 0

walk:
	for {
// 找出最长的公共前缀
// 这还意味着公共前缀不包含':'或'*'，因为现有键不能包含这些字符
		i := longestCommonPrefix(path, n.path)

// 分裂的边缘
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
// []byte用于正确的unicode字符转换，参见#65
			n.indices = bytesconv.BytesToString([]byte{n.path[i]})
			n.path = path[:i]
			n.handlers = nil
			n.wildChild = false
			n.fullPath = fullPath[:parentFullPathIndex+i]
		}

// 使新节点成为此节点的子节点
		if i < len(path) {
			path = path[i:]
			c := path[0]

// 参数后的'/'
			if n.nType == param && c == '/' && len(n.children) == 1 {
				parentFullPathIndex += len(n.path)
				n = n.children[0]
				n.priority++
				continue walk
			}

// 检查下一个路径字节的子节点是否存在
			for i, max := 0, len(n.indices); i < max; i++ {
				if c == n.indices[i] {
					parentFullPathIndex += len(n.path)
					i = n.incrementChildPrio(i)
					n = n.children[i]
					continue walk
				}
			}

// 否则插入
			if c != ':' && c != '*' && n.nType != catchAll {
	// []byte用于正确的unicode字符转换，参见#65
				n.indices += bytesconv.BytesToString([]byte{c})
				child := &node{
					fullPath: fullPath,
				}
				n.addChild(child)
				n.incrementChildPrio(len(n.indices) - 1)
				n = child
			} else if n.wildChild {
// 插入一个通配符节点，需要检查它是否与现有的通配符冲突
				n = n.children[len(n.children)-1]
				n.priority++

// 检查通配符是否匹配
				if len(path) >= len(n.path) && n.path == path[:len(n.path)] &&
// 不能将子对象添加到catchAll中
					n.nType != catchAll &&
// 检查较长的通配符，例如:name和:names
					(len(n.path) >= len(path) || path[len(n.path)] == '/') {
					continue walk
				}

// 通配符的冲突
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

// 否则向当前节点添加句柄
		if n.handlers != nil {
			panic("handlers are already registered for path '" + fullPath + "'")
		}
		n.handlers = handlers
		n.fullPath = fullPath
		return
	}
}

// 搜索通配符段并检查名称是否有无效字符
// 如果没有找到通配符，则返回-1作为索引
func findWildcard(path string) (wildcard string, i int, valid bool) {
// 找到开始
	for start, c := range []byte(path) {
// 通配符以':'(参数)或'*'(通配符)开头
		if c != ':' && c != '*' {
			continue
		}

// 查找结束符并检查无效字符
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
// 查找前缀直到第一个通配符
		wildcard, i, valid := findWildcard(path)
		if i < 0 { // 没有找到通配符
			break
		}

// 通配符名称只能包含一个':'或'*'字符
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
// 在当前通配符之前插入前缀
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

// 如果路径没有以通配符结尾，那么将会有另一个子路径以'/'开头
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

// 否则我们就做完了
// 将手柄插入新叶中
			n.handlers = handlers
			return
		}

// 包罗万象的
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

// 目前固定宽度为1的“/”
		i--
		if path[i] != '/' {
			panic("no / before catch-all in path '" + fullPath + "'")
		}

		n.path = path[:i]

// 第一个节点:空路径的catchAll节点
		child := &node{
			wildChild: true,
			nType:     catchAll,
			fullPath:  fullPath,
		}

		n.addChild(child)
		n.indices = string('/')
		n = child
		n.priority++

// 第二个节点:保存变量的节点
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

// 如果没有找到通配符，只需插入路径和句柄
	n.path = path
	n.handlers = handlers
	n.fullPath = fullPath
}

// nodeValue保存(*Node)的返回值
// getValue方法
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

// 返回用给定路径(键)注册的句柄
// 通配符的值保存到映射中
// 如果找不到句柄，如果存在一个带有额外(不带)尾斜杠的句柄，则会提出TSR(尾斜杠重定向)建议
func (n *node) getValue(path string, params *Params, skippedNodes *[]skippedNode, unescape bool) (value nodeValue) {
	var globalParamsCount int16

walk: // 行走树的外循环
	for {
		prefix := n.path
		if len(path) > len(prefix) {
			if path[:len(prefix)] == prefix {
				path = path[len(prefix):]

// 首先通过匹配索引来尝试所有非通配符子节点
				idxc := path[0]
				for i, c := range []byte(n.indices) {
					if c == idxc {
// strings.HasPrefix (n.children len (n.children)[1]
// path， ":") == n.wildChild
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
// 如果循环结束时的路径不等于'/'，并且当前节点没有子节点，则当前节点需要回滚到最后一个有效的skippedNode
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

// 没有什么发现
// 我们可以建议重定向到相同的URL，如果该路径存在叶子，则不带尾斜杠
					value.tsr = path == "/" && n.handlers != nil
					return
				}

// 处理通配符子，它总是在数组的末尾
				n = n.children[len(n.children)-1]
				globalParamsCount++

				switch n.nType {
				case param:
// 修复截断参数tree_test的问题
// Go line: 204

// 查找参数结束('/'或路径结束)
					end := 0
					for end < len(path) && path[end] != '/' {
						end++
					}

// 保存参数值
					if params != nil {
// 必要时预分配容量
						if cap(*params) < int(globalParamsCount) {
							newParams := make(Params, len(*params), globalParamsCount)
							copy(newParams, *params)
							*params = newParams
						}

						if value.params == nil {
							value.params = params
						}
// 在预分配的容量内扩展切片
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

// 我们得再深入一点!
					if end < len(path) {
						if len(n.children) > 0 {
							path = path[end:]
							n = n.children[0]
							continue walk
						}

// …但我们不能
						value.tsr = len(path) == end+1
						return
					}

					if value.handlers = n.handlers; value.handlers != nil {
						value.fullPath = n.fullPath
						return
					}
					if len(n.children) == 1 {
// 没有找到手柄
// 检查是否存在此路径的句柄+尾斜杠以供TSR推荐
						n = n.children[0]
						value.tsr = (n.path == "/" && n.handlers != nil) || (n.path == "" && n.indices == "/")
					}
					return

				case catchAll:
// 保存参数值
					if params != nil {
// 必要时预分配容量
						if cap(*params) < int(globalParamsCount) {
							newParams := make(Params, len(*params), globalParamsCount)
							copy(newParams, *params)
							*params = newParams
						}

						if value.params == nil {
							value.params = params
						}
// 在预分配的容量内扩展切片
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
// 如果当前路径不等于'/'，并且节点没有注册句柄，并且最近匹配的节点有子节点，则当前节点需要回滚到最后一个有效的skippedNode
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
// n = 最新节点的子节点列表中的最后一个子节点，即latestNode.children数组的最后一个元素
			}
// 我们应该已经到达包含句柄的节点
// 检查此节点是否已注册句柄
			if value.handlers = n.handlers; value.handlers != nil {
				value.fullPath = n.fullPath
				return
			}

// 如果这个路由没有句柄，但是这个路由有一个通配符子节点，那么这个路径必须有一个附加斜杠的句柄
			if path == "/" && n.wildChild && n.nType != root {
				value.tsr = true
				return
			}

			if path == "/" && n.nType == static {
				value.tsr = true
				return
			}

// 没有找到手柄
// 检查该路径的句柄+尾斜杠是否存在以推荐尾斜杠
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

// 没有什么发现
// 我们可以建议重定向到相同的URL，如果该路径存在叶子，则使用额外的斜杠
		value.tsr = path == "/" ||
			(len(prefix) == len(path)+1 && prefix[len(path)] == '/' &&
				path == prefix[:len(prefix)-1] && n.handlers != nil)

// 回滚到最后一个有效的skippedNode
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

// 对给定路径进行不区分大小写的查找，并尝试查找处理程序
// 它还可以选择修复尾随斜杠
// 它返回经过大小写校正的路径和一个bool值，该值指示查找是否成功
func (n *node) findCaseInsensitivePath(path string, fixTrailingSlash bool) ([]byte, bool) {
	const stackBufSize = 128

// 一般情况下，在堆栈上使用静态大小的缓冲区
// 如果路径太长，则在堆上分配一个缓冲区
	buf := make([]byte, 0, stackBufSize)
	if length := len(path) + 1; length > stackBufSize {
		buf = make([]byte, 0, length)
	}

	ciPath := n.findCaseInsensitivePathRec(
		path,
		buf,       // 为新路径预先分配足够的内存
		[4]byte{}, // 空符文缓冲区
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

// n.findCaseInsensitivePath使用的递归不区分大小写的查找函数
func (n *node) findCaseInsensitivePathRec(path string, ciPath []byte, rb [4]byte, fixTrailingSlash bool) []byte {
	npLen := len(n.path)

walk: // 行走树的外循环
	for len(path) >= npLen && (npLen == 0 || strings.EqualFold(path[1:npLen], n.path[1:])) {
// 在结果中添加公共前缀
		oldPath := path
		path = path[npLen:]
		ciPath = append(ciPath, n.path...)

		if len(path) == 0 {
// 我们应该已经到达包含句柄的节点
// 检查此节点是否已注册句柄
			if n.handlers != nil {
				return ciPath
			}

// 没有找到手柄
// 尝试通过添加尾斜杠来修复路径
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

// 如果该节点没有通配符(param或catchAll)子节点，我们可以只查找下一个子节点并继续沿着树向下走
		if !n.wildChild {
// 跳过已处理的符文字节
			rb = shiftNRuneBytes(rb, npLen)

			if rb[0] != 0 {
// 旧符文还没写完
				idxc := rb[0]
				for i, c := range []byte(n.indices) {
					if c == idxc {
// 继续子节点
						n = n.children[i]
						npLen = len(n.path)
						continue walk
					}
				}
			} else {
// 处理一个新的符文
				var rv rune

// 找到符文开始
// 符文最多有4字节长，-4肯定是另一个符文
				var off int
				for max := min(npLen, 3); off < max; off++ {
					if i := npLen - off; utf8.RuneStart(oldPath[i]) {
// 从缓存路径读取符文
						rv, _ = utf8.DecodeRuneInString(oldPath[i:])
						break
					}
				}

// 计算当前符文的小写字节
				lo := unicode.ToLower(rv)
				utf8.EncodeRune(rb[:], lo)

// 跳过已经处理的字节
				rb = shiftNRuneBytes(rb, off)

				idxc := rb[0]
				for i, c := range []byte(n.indices) {
// 小写字母相匹配
					if c == idxc {
// 必须使用递归方法，因为大写字节和小写字节都可能作为索引存在
						if out := n.children[i].findCaseInsensitivePathRec(
							path, ciPath, rb, fixTrailingSlash,
						); out != nil {
							return out
						}
						break
					}
				}

// 如果找不到匹配，大写符文也一样，如果不同
				if up := unicode.ToUpper(rv); up != lo {
					utf8.EncodeRune(rb[:], up)
					rb = shiftNRuneBytes(rb, off)

					idxc := rb[0]
					for i, c := range []byte(n.indices) {
// 大写字母相匹配
						if c == idxc {
	// 继续子节点
							n = n.children[i]
							npLen = len(n.path)
							continue walk
						}
					}
				}
			}

			// Nothing found. We can recommend to redirect to the same URL
			// without a trailing slash if a leaf exists for that path
			if fixTrailingSlash && path == "/" && n.handlers != nil {
				return ciPath
			}
			return nil
		}

		n = n.children[0]
		switch n.nType {
		case param:
			// Find param end (either '/' or path end)
			end := 0
			for end < len(path) && path[end] != '/' {
				end++
			}

// 为不区分大小写的路径添加参数值
			ciPath = append(ciPath, path[:end]...)

			// We need to go deeper!
			if end < len(path) {
				if len(n.children) > 0 {
					// Continue with child node
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

// 没有什么发现
// 尝试通过添加/删除尾斜杠来修复路径
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
