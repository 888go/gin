
<原文开始>
// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// at https://github.com/julienschmidt/httprouter/blob/master/LICENSE
<原文结束>

# <翻译开始>
// 版权所有 ? 2013 Julien Schmidt。保留所有权利。
// 本源代码的使用受 BSD 风格许可协议约束，该协议可在
// https://github.com/julienschmidt/httprouter/blob/master/LICENSE 查阅
# <翻译结束>


<原文开始>
// Param is a single URL parameter, consisting of a key and a value.
<原文结束>

# <翻译开始>
// Param 是单个URL参数，包含一个键和一个值。
# <翻译结束>


<原文开始>
// Params is a Param-slice, as returned by the router.
// The slice is ordered, the first URL parameter is also the first slice value.
// It is therefore safe to read values by the index.
<原文结束>

# <翻译开始>
// Params 是一个 Param 切片，由路由器返回。
// 这个切片是有序的，第一个 URL 参数也是切片中的第一个值。
// 因此，通过索引读取值是安全的。
# <翻译结束>


<原文开始>
// Get returns the value of the first Param which key matches the given name and a boolean true.
// If no matching Param is found, an empty string is returned and a boolean false .
<原文结束>

# <翻译开始>
// Get 方法返回第一个参数键与给定名称相匹配的值，并返回布尔值 true。
// 如果未找到匹配的参数，则返回一个空字符串和布尔值 false。
# <翻译结束>


<原文开始>
// ByName returns the value of the first Param which key matches the given name.
// If no matching Param is found, an empty string is returned.
<原文结束>

# <翻译开始>
// ByName 返回第一个其键与给定名称相匹配的Param的值。
// 如果未找到匹配的Param，则返回一个空字符串。
# <翻译结束>


<原文开始>
// addChild will add a child node, keeping wildcardChild at the end
<原文结束>

# <翻译开始>
// addChild 将添加一个子节点，并保持 wildcardChild 位于末尾
# <翻译结束>


<原文开始>
// child nodes, at most 1 :param style node at the end of the array
<原文结束>

# <翻译开始>
// 子节点，切片末尾最多包含 1 个 :param 样式的节点
# <翻译结束>


<原文开始>
// Increments priority of the given child and reorders if necessary
<原文结束>

# <翻译开始>
// 如果有必要，递增给定子项的优先级并重新排序
# <翻译结束>


<原文开始>
// Adjust position (move to front)
<原文结束>

# <翻译开始>
// 调整位置（移动到前面）
# <翻译结束>


<原文开始>
// Build new index char string
<原文结束>

# <翻译开始>
// 构建新的索引字符字符串
# <翻译结束>


<原文开始>
// Unchanged prefix, might be empty
<原文结束>

# <翻译开始>
// 未更改的前缀，可能为空
# <翻译结束>


<原文开始>
// Rest without char at 'pos'
<原文结束>

# <翻译开始>
// 在'pos'位置没有字符的Rest
# <翻译结束>


<原文开始>
// addRoute adds a node with the given handle to the path.
// Not concurrency-safe!
<原文结束>

# <翻译开始>
// addRoute 向路径中添加具有给定 handle 的节点。
// 非并发安全！
# <翻译结束>


<原文开始>
		// Find the longest common prefix.
		// This also implies that the common prefix contains no ':' or '*'
		// since the existing key can't contain those chars.
<原文结束>

# <翻译开始>
// 查找最长公共前缀。
// 这也意味着，公共前缀中不包含 ':' 或 '*'，
// 因为已存在的键不能包含这些字符。
# <翻译结束>


<原文开始>
// []byte for proper unicode char conversion, see #65
<原文结束>

# <翻译开始>
// 使用 []byte 以正确处理 Unicode 字符转换，参考 #65 号问题
# <翻译结束>


<原文开始>
// Make new node a child of this node
<原文结束>

# <翻译开始>
// 将新节点作为此节点的子节点
# <翻译结束>


<原文开始>
// Check if a child with the next path byte exists
<原文结束>

# <翻译开始>
// 检查是否存在下一个路径字节的子节点
# <翻译结束>


<原文开始>
// inserting a wildcard node, need to check if it conflicts with the existing wildcard
<原文结束>

# <翻译开始>
// 插入通配符节点，需要检查是否与现有通配符冲突
# <翻译结束>


<原文开始>
// Check if the wildcard matches
<原文结束>

# <翻译开始>
// 检查通配符是否匹配
# <翻译结束>


<原文开始>
// Adding a child to a catchAll is not possible
<原文结束>

# <翻译开始>
// 向 catchAll 添加子节点是不可能的
# <翻译结束>


<原文开始>
// Check for longer wildcard, e.g. :name and :names
<原文结束>

# <翻译开始>
// 检查较长的通配符，例如:name和:names
# <翻译结束>


<原文开始>
// Otherwise add handle to current node
<原文结束>

# <翻译开始>
// 否则将处理程序添加到当前节点
# <翻译结束>


<原文开始>
// Search for a wildcard segment and check the name for invalid characters.
// Returns -1 as index, if no wildcard was found.
<原文结束>

# <翻译开始>
// 搜索通配符段，并检查名称中是否存在无效字符。
// 如果未找到通配符，则返回索引 -1。
# <翻译结束>


<原文开始>
// A wildcard starts with ':' (param) or '*' (catch-all)
<原文结束>

# <翻译开始>
// 通配符以 ':'（参数）或 '*'（捕获全部）开始
# <翻译结束>


<原文开始>
// Find end and check for invalid characters
<原文结束>

# <翻译开始>
// 查找结尾并检查无效字符
# <翻译结束>


<原文开始>
// Find prefix until first wildcard
<原文结束>

# <翻译开始>
// 查找直到第一个通配符为止的前缀
# <翻译结束>


<原文开始>
// The wildcard name must only contain one ':' or '*' character
<原文结束>

# <翻译开始>
// 通配符名称中只能包含一个':'或'*'字符
# <翻译结束>


<原文开始>
// check if the wildcard has a name
<原文结束>

# <翻译开始>
// 检查通配符是否有名称
# <翻译结束>


<原文开始>
// Insert prefix before the current wildcard
<原文结束>

# <翻译开始>
// 在当前通配符前插入前缀
# <翻译结束>


<原文开始>
			// if the path doesn't end with the wildcard, then there
			// will be another subpath starting with '/'
<原文结束>

# <翻译开始>
// 如果路径没有以通配符结尾，则会有一个以'/'开头的另一个子路径
# <翻译结束>


<原文开始>
// Otherwise we're done. Insert the handle in the new leaf
<原文结束>

# <翻译开始>
// 否则我们已经完成。将句柄插入新的叶子节点
# <翻译结束>


<原文开始>
// currently fixed width 1 for '/'
<原文结束>

# <翻译开始>
// 当前固定宽度为1用于 '/'
# <翻译结束>


<原文开始>
// First node: catchAll node with empty path
<原文结束>

# <翻译开始>
// 第一个节点：匹配所有路径的“catchAll”节点，其路径为空
# <翻译结束>


<原文开始>
// second node: node holding the variable
<原文结束>

# <翻译开始>
// 第二个节点：存储变量的节点
# <翻译结束>


<原文开始>
// If no wildcard was found, simply insert the path and handle
<原文结束>

# <翻译开始>
// 如果没有找到通配符，直接插入路径并处理
# <翻译结束>


<原文开始>
// nodeValue holds return values of (*Node).getValue method
<原文结束>

# <翻译开始>
// nodeValue 用于存储 (*Node).getValue 方法的返回值
# <翻译结束>


<原文开始>
// Returns the handle registered with the given path (key). The values of
// wildcards are saved to a map.
// If no handle can be found, a TSR (trailing slash redirect) recommendation is
// made if a handle exists with an extra (without the) trailing slash for the
// given path.
<原文结束>

# <翻译开始>
// 根据给定路径（键）返回已注册的处理程序。通配符的值将保存到一个映射中。
// 如果找不到处理程序，且对于给定路径存在一个带有额外（无）尾部斜杠的处理程序，则会提出 TSR（尾部斜杠重定向）建议。
# <翻译结束>


<原文开始>
// Outer loop for walking the tree
<原文结束>

# <翻译开始>
// 外层循环，用于遍历树
# <翻译结束>


<原文开始>
// Try all the non-wildcard children first by matching the indices
<原文结束>

# <翻译开始>
// 首先尝试通过匹配索引来查找所有非通配符子节点
# <翻译结束>


<原文开始>
//  strings.HasPrefix(n.children[len(n.children)-1].path, ":") == n.wildChild
<原文结束>

# <翻译开始>
// 判断n的子节点列表中最后一个子节点的路径是否以":"开头，结果与n.wildChild属性相等
# <翻译结束>


<原文开始>
					// If the path at the end of the loop is not equal to '/' and the current node has no child nodes
					// the current node needs to roll back to last valid skippedNode
<原文结束>

# <翻译开始>
// 如果循环结束时路径不等于'/'且当前节点没有子节点，
// 那么当前节点需要回滚到上一个有效的已跳过节点
# <翻译结束>


<原文开始>
					// Nothing found.
					// We can recommend to redirect to the same URL without a
					// trailing slash if a leaf exists for that path.
<原文结束>

# <翻译开始>
// 未找到任何内容。
// 如果该路径存在叶节点，我们可以建议重定向到不带尾部斜杠的相同 URL。
# <翻译结束>


<原文开始>
// Handle wildcard child, which is always at the end of the array
<原文结束>

# <翻译开始>
// 处理通配符子节点，它总是在切片的末尾
# <翻译结束>


<原文开始>
					// fix truncate the parameter
					// tree_test.go  line: 204
<原文结束>

# <翻译开始>
// 修复：截断参数
// 文件：tree_test.go，行号：204
# <翻译结束>


<原文开始>
// Find param end (either '/' or path end)
<原文结束>

# <翻译开始>
// 查找参数结束位置（可能是'/'或路径结束）
# <翻译结束>


<原文开始>
// Preallocate capacity if necessary
<原文结束>

# <翻译开始>
// 如果有必要，预先分配容量
# <翻译结束>


<原文开始>
// Expand slice within preallocated capacity
<原文结束>

# <翻译开始>
// 在预先分配的容量内扩展切片
# <翻译结束>


<原文开始>
						// No handle found. Check if a handle for this path + a
						// trailing slash exists for TSR recommendation
<原文结束>

# <翻译开始>
// 未找到处理程序。检查是否存在该路径加上末尾斜杠的处理程序，以便进行 TSR（可能指“透明目录重写”）推荐
# <翻译结束>


<原文开始>
			// If the current path does not equal '/' and the node does not have a registered handle and the most recently matched node has a child node
			// the current node needs to roll back to last valid skippedNode
<原文结束>

# <翻译开始>
// 如果当前路径不等于'/'，并且节点未注册处理程序，并且最近匹配的节点有一个子节点
// 那么当前节点需要回滚到上一个有效跳过的节点
# <翻译结束>


<原文开始>
//	n = latestNode.children[len(latestNode.children)-1]
<原文结束>

# <翻译开始>
// n = latestNode.children切片的最后一个元素
# <翻译结束>


<原文开始>
			// We should have reached the node containing the handle.
			// Check if this node has a handle registered.
<原文结束>

# <翻译开始>
// 我们应该已经到达包含处理程序的节点。
// 检查该节点是否注册了处理程序。
# <翻译结束>


<原文开始>
			// If there is no handle for this route, but this route has a
			// wildcard child, there must be a handle for this path with an
			// additional trailing slash
<原文结束>

# <翻译开始>
// 如果该路由没有处理程序，但该路由有一个通配符子路由，则必须为此路径（带有一个额外的尾部斜杠）提供一个处理程序
# <翻译结束>


<原文开始>
			// No handle found. Check if a handle for this path + a
			// trailing slash exists for trailing slash recommendation
<原文结束>

# <翻译开始>
// 未找到处理程序。检查是否存在针对此路径加上末尾斜杠的处理程序，以便提供末尾斜杠建议
# <翻译结束>


<原文开始>
		// Nothing found. We can recommend to redirect to the same URL with an
		// extra trailing slash if a leaf exists for that path
<原文结束>

# <翻译开始>
// 未找到任何内容。如果该路径存在叶子节点，我们可以建议重定向到同一 URL 并附加一个额外的尾部斜杠
# <翻译结束>


<原文开始>
// roll back to last valid skippedNode
<原文结束>

# <翻译开始>
// 回滚到最后一个有效的跳过节点
# <翻译结束>


<原文开始>
// Makes a case-insensitive lookup of the given path and tries to find a handler.
// It can optionally also fix trailing slashes.
// It returns the case-corrected path and a bool indicating whether the lookup
// was successful.
<原文结束>

# <翻译开始>
// 根据给定的路径进行不区分大小写的查找，并尝试找到一个处理器。
// 可选地，它也可以修正尾部的斜杠。
// 它返回大小写纠正后的路径以及一个布尔值，表示查找是否成功。
# <翻译结束>


<原文开始>
	// Use a static sized buffer on the stack in the common case.
	// If the path is too long, allocate a buffer on the heap instead.
<原文结束>

# <翻译开始>
// 在常见情况下，使用栈上静态大小的缓冲区。如果路径过长，则改为在堆上分配缓冲区。
# <翻译结束>


<原文开始>
// Preallocate enough memory for new path
<原文结束>

# <翻译开始>
// 预先分配足够的内存用于新路径
# <翻译结束>


<原文开始>
// Shift bytes in array by n bytes left
<原文结束>

# <翻译开始>
// 将切片中的字节向左移动n个字节
# <翻译结束>


<原文开始>
// Recursive case-insensitive lookup function used by n.findCaseInsensitivePath
<原文结束>

# <翻译开始>
// 递归的大小写不敏感查找函数，由n.findCaseInsensitivePath调用
# <翻译结束>


<原文开始>
// Add common prefix to result
<原文结束>

# <翻译开始>
// 向结果添加公共前缀
# <翻译结束>


<原文开始>
			// No handle found.
			// Try to fix the path by adding a trailing slash
<原文结束>

# <翻译开始>
// 未找到处理程序。
// 尝试通过添加尾部斜杠来修复路径
# <翻译结束>


<原文开始>
		// If this node does not have a wildcard (param or catchAll) child,
		// we can just look up the next child node and continue to walk down
		// the tree
<原文结束>

# <翻译开始>
// 如果该节点没有通配符（param 或 catchAll）子节点，
// 我们可以直接查找下一个子节点并继续向下遍历树
# <翻译结束>


<原文开始>
// Skip rune bytes already processed
<原文结束>

# <翻译开始>
// 跳过已处理的 rune 字节
# <翻译结束>


<原文开始>
// continue with child node
<原文结束>

# <翻译开始>
// 继续处理子节点
# <翻译结束>


<原文开始>
				// Find rune start.
				// Runes are up to 4 byte long,
				// -4 would definitely be another rune.
<原文结束>

# <翻译开始>
// 查找rune的起始位置。
// rune字符可能包含最多4个字节，
// 因此，-4肯定属于另一个rune字符。
# <翻译结束>


<原文开始>
// read rune from cached path
<原文结束>

# <翻译开始>
// 从缓存路径读取单个字符
# <翻译结束>


<原文开始>
// Calculate lowercase bytes of current rune
<原文结束>

# <翻译开始>
// 计算当前字符的 lowercase 字节
# <翻译结束>


<原文开始>
// Skip already processed bytes
<原文结束>

# <翻译开始>
// 跳过已处理的字节
# <翻译结束>


<原文开始>
						// must use a recursive approach since both the
						// uppercase byte and the lowercase byte might exist
						// as an index
<原文结束>

# <翻译开始>
// 必须采用递归方法，因为大小写两种字节都可能存在作为索引
# <翻译结束>


<原文开始>
				// If we found no match, the same for the uppercase rune,
				// if it differs
<原文结束>

# <翻译开始>
// 如果我们没有找到匹配项，对于大写字符执行相同的操作，
// 如果它与小写字符不同的话
# <翻译结束>


<原文开始>
// Continue with child node
<原文结束>

# <翻译开始>
// 继续处理子节点
# <翻译结束>


<原文开始>
			// Nothing found. We can recommend to redirect to the same URL
			// without a trailing slash if a leaf exists for that path
<原文结束>

# <翻译开始>
// 未找到任何内容。如果该路径存在叶子节点，我们可以建议重定向到去掉尾部斜杠的相同URL
# <翻译结束>


<原文开始>
// Add param value to case insensitive path
<原文结束>

# <翻译开始>
// 添加参数值到不区分大小写的路径
# <翻译结束>


<原文开始>
	// Nothing found.
	// Try to fix the path by adding / removing a trailing slash
<原文结束>

# <翻译开始>
// 未找到任何内容。
// 尝试通过添加或删除尾部斜杠来修复路径
# <翻译结束>

