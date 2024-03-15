
<原文开始>
// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// at https://github.com/julienschmidt/httprouter/blob/master/LICENSE
<原文结束>

# <翻译开始>
// 版权所有2013朱利安施密特
// 版权所有
// 此源代码的使用受bsd风格的许可证的约束，该许可证可在https://github.com/julienschmidt/httprouter/blob/master/LICENSE上找到
# <翻译结束>


<原文开始>
// Param is a single URL parameter, consisting of a key and a value.
<原文结束>

# <翻译开始>
// Param是一个URL参数，由一个键和一个值组成
# <翻译结束>


<原文开始>
// Params is a Param-slice, as returned by the router.
// The slice is ordered, the first URL parameter is also the first slice value.
// It is therefore safe to read values by the index.
<原文结束>

# <翻译开始>
// Params是一个Param-slice，由路由器返回
// 片是有序的，第一个URL参数也是第一个片值
// 因此，通过索引读取值是安全的
# <翻译结束>


<原文开始>
// Get returns the value of the first Param which key matches the given name and a boolean true.
// If no matching Param is found, an empty string is returned and a boolean false .
<原文结束>

# <翻译开始>
// Get返回与给定名称匹配的第一个Param的值和一个布尔值true
// 如果没有找到匹配的参数，则返回一个空字符串和一个布尔值false
# <翻译结束>


<原文开始>
// ByName returns the value of the first Param which key matches the given name.
// If no matching Param is found, an empty string is returned.
<原文结束>

# <翻译开始>
// ByName返回与给定名称匹配的第一个参数的值
// 如果没有找到匹配的Param，则返回一个空字符串
# <翻译结束>


<原文开始>
// addChild will add a child node, keeping wildcardChild at the end
<原文结束>

# <翻译开始>
// addChild将添加一个子节点，将wildcardChild保留在最后
# <翻译结束>


<原文开始>
// child nodes, at most 1 :param style node at the end of the array
<原文结束>

# <翻译开始>
// 子节点，最多1个:数组末尾的参数样式节点
# <翻译结束>


<原文开始>
// Increments priority of the given child and reorders if necessary
<原文结束>

# <翻译开始>
// 增加给定子元素的优先级，并在必要时重新排序
# <翻译结束>


<原文开始>
	// Adjust position (move to front)
<原文结束>

# <翻译开始>
// 调整位置(移动到前面)
# <翻译结束>


<原文开始>
		// Swap node positions
<原文结束>

# <翻译开始>
// 交换节点位置
# <翻译结束>


<原文开始>
	// Build new index char string
<原文结束>

# <翻译开始>
// 构建新的索引字符串
# <翻译结束>


<原文开始>
// Unchanged prefix, might be empty
<原文结束>

# <翻译开始>
// 未更改前缀，可能为空
# <翻译结束>


<原文开始>
// The index char we move
<原文结束>

# <翻译开始>
// 我们移动的索引字符
# <翻译结束>


<原文开始>
// addRoute adds a node with the given handle to the path.
// Not concurrency-safe!
<原文结束>

# <翻译开始>
// addRoute将具有给定句柄的节点添加到路径中
// 不是concurrency-safe !
# <翻译结束>


<原文开始>
	// Empty tree
<原文结束>

# <翻译开始>
// 空树
# <翻译结束>


<原文开始>
		// Find the longest common prefix.
		// This also implies that the common prefix contains no ':' or '*'
		// since the existing key can't contain those chars.
<原文结束>

# <翻译开始>
// 找出最长的公共前缀
// 这还意味着公共前缀不包含':'或'*'，因为现有键不能包含这些字符
# <翻译结束>


<原文开始>
		// Split edge
<原文结束>

# <翻译开始>
// 分裂的边缘
# <翻译结束>


<原文开始>
			// []byte for proper unicode char conversion, see #65
<原文结束>

# <翻译开始>
// []byte用于正确的unicode字符转换，参见#65
# <翻译结束>


<原文开始>
		// Make new node a child of this node
<原文结束>

# <翻译开始>
// 使新节点成为此节点的子节点
# <翻译结束>


<原文开始>
			// '/' after param
<原文结束>

# <翻译开始>
// 参数后的'/'
# <翻译结束>


<原文开始>
			// Check if a child with the next path byte exists
<原文结束>

# <翻译开始>
// 检查下一个路径字节的子节点是否存在
# <翻译结束>


<原文开始>
			// Otherwise insert it
<原文结束>

# <翻译开始>
// 否则插入
# <翻译结束>


<原文开始>
				// []byte for proper unicode char conversion, see #65
<原文结束>

# <翻译开始>
// []byte用于正确的unicode字符转换，参见#65
# <翻译结束>


<原文开始>
				// inserting a wildcard node, need to check if it conflicts with the existing wildcard
<原文结束>

# <翻译开始>
// 插入一个通配符节点，需要检查它是否与现有的通配符冲突
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
// 不能将子对象添加到catchAll中
# <翻译结束>


<原文开始>
					// Check for longer wildcard, e.g. :name and :names
<原文结束>

# <翻译开始>
// 检查较长的通配符，例如:name和:names
# <翻译结束>


<原文开始>
				// Wildcard conflict
<原文结束>

# <翻译开始>
// 通配符的冲突
# <翻译结束>


<原文开始>
		// Otherwise add handle to current node
<原文结束>

# <翻译开始>
// 否则向当前节点添加句柄
# <翻译结束>


<原文开始>
// Search for a wildcard segment and check the name for invalid characters.
// Returns -1 as index, if no wildcard was found.
<原文结束>

# <翻译开始>
// 搜索通配符段并检查名称是否有无效字符
// 如果没有找到通配符，则返回-1作为索引
# <翻译结束>


<原文开始>
	// Find start
<原文结束>

# <翻译开始>
// 找到开始
# <翻译结束>


<原文开始>
		// A wildcard starts with ':' (param) or '*' (catch-all)
<原文结束>

# <翻译开始>
// 通配符以':'(参数)或'*'(通配符)开头
# <翻译结束>


<原文开始>
		// Find end and check for invalid characters
<原文结束>

# <翻译开始>
// 查找结束符并检查无效字符
# <翻译结束>


<原文开始>
		// Find prefix until first wildcard
<原文结束>

# <翻译开始>
// 查找前缀直到第一个通配符
# <翻译结束>


<原文开始>
// No wildcard found
<原文结束>

# <翻译开始>
// 没有找到通配符
# <翻译结束>


<原文开始>
		// The wildcard name must only contain one ':' or '*' character
<原文结束>

# <翻译开始>
// 通配符名称只能包含一个':'或'*'字符
# <翻译结束>


<原文开始>
		// check if the wildcard has a name
<原文结束>

# <翻译开始>
// 检查通配符是否有名称
# <翻译结束>


<原文开始>
// param
<原文结束>

# <翻译开始>
// 参数
# <翻译结束>


<原文开始>
				// Insert prefix before the current wildcard
<原文结束>

# <翻译开始>
// 在当前通配符之前插入前缀
# <翻译结束>


<原文开始>
			// if the path doesn't end with the wildcard, then there
			// will be another subpath starting with '/'
<原文结束>

# <翻译开始>
// 如果路径没有以通配符结尾，那么将会有另一个子路径以'/'开头
# <翻译结束>


<原文开始>
			// Otherwise we're done. Insert the handle in the new leaf
<原文结束>

# <翻译开始>
// 否则我们就做完了
// 将手柄插入新叶中
# <翻译结束>


<原文开始>
		// catchAll
<原文结束>

# <翻译开始>
// 包罗万象的
# <翻译结束>


<原文开始>
		// currently fixed width 1 for '/'
<原文结束>

# <翻译开始>
// 目前固定宽度为1的“/”
# <翻译结束>


<原文开始>
		// First node: catchAll node with empty path
<原文结束>

# <翻译开始>
// 第一个节点:空路径的catchAll节点
# <翻译结束>


<原文开始>
		// second node: node holding the variable
<原文结束>

# <翻译开始>
// 第二个节点:保存变量的节点
# <翻译结束>


<原文开始>
	// If no wildcard was found, simply insert the path and handle
<原文结束>

# <翻译开始>
// 如果没有找到通配符，只需插入路径和句柄
# <翻译结束>


<原文开始>
// nodeValue holds return values of (*Node).getValue method
<原文结束>

# <翻译开始>
// nodeValue保存(*Node)的返回值
// getValue方法
# <翻译结束>


<原文开始>
// Returns the handle registered with the given path (key). The values of
// wildcards are saved to a map.
// If no handle can be found, a TSR (trailing slash redirect) recommendation is
// made if a handle exists with an extra (without the) trailing slash for the
// given path.
<原文结束>

# <翻译开始>
// 返回用给定路径(键)注册的句柄
// 通配符的值保存到映射中
// 如果找不到句柄，如果存在一个带有额外(不带)尾斜杠的句柄，则会提出TSR(尾斜杠重定向)建议
# <翻译结束>


<原文开始>
// Outer loop for walking the tree
<原文结束>

# <翻译开始>
// 行走树的外循环
# <翻译结束>


<原文开始>
				// Try all the non-wildcard children first by matching the indices
<原文结束>

# <翻译开始>
// 首先通过匹配索引来尝试所有非通配符子节点
# <翻译结束>


<原文开始>
						//  strings.HasPrefix(n.children[len(n.children)-1].path, ":") == n.wildChild
<原文结束>

# <翻译开始>
// strings.HasPrefix (n.children len (n.children)[1]
// path， ":") == n.wildChild
# <翻译结束>


<原文开始>
					// If the path at the end of the loop is not equal to '/' and the current node has no child nodes
					// the current node needs to roll back to last valid skippedNode
<原文结束>

# <翻译开始>
// 如果循环结束时的路径不等于'/'，并且当前节点没有子节点，则当前节点需要回滚到最后一个有效的skippedNode
# <翻译结束>


<原文开始>
					// Nothing found.
					// We can recommend to redirect to the same URL without a
					// trailing slash if a leaf exists for that path.
<原文结束>

# <翻译开始>
// 没有什么发现
// 我们可以建议重定向到相同的URL，如果该路径存在叶子，则不带尾斜杠
# <翻译结束>


<原文开始>
				// Handle wildcard child, which is always at the end of the array
<原文结束>

# <翻译开始>
// 处理通配符子，它总是在数组的末尾
# <翻译结束>


<原文开始>
					// fix truncate the parameter
					// tree_test.go  line: 204
<原文结束>

# <翻译开始>
// 修复截断参数tree_test的问题
// Go line: 204
# <翻译结束>


<原文开始>
					// Find param end (either '/' or path end)
<原文结束>

# <翻译开始>
// 查找参数结束('/'或路径结束)
# <翻译结束>


<原文开始>
					// Save param value
<原文结束>

# <翻译开始>
// 保存参数值
# <翻译结束>


<原文开始>
						// Preallocate capacity if necessary
<原文结束>

# <翻译开始>
// 必要时预分配容量
# <翻译结束>


<原文开始>
						// Expand slice within preallocated capacity
<原文结束>

# <翻译开始>
// 在预分配的容量内扩展切片
# <翻译结束>


<原文开始>
					// we need to go deeper!
<原文结束>

# <翻译开始>
// 我们得再深入一点!
# <翻译结束>


<原文开始>
						// ... but we can't
<原文结束>

# <翻译开始>
// …但我们不能
# <翻译结束>


<原文开始>
						// No handle found. Check if a handle for this path + a
						// trailing slash exists for TSR recommendation
<原文结束>

# <翻译开始>
// 没有找到手柄
// 检查是否存在此路径的句柄+尾斜杠以供TSR推荐
# <翻译结束>


<原文开始>
			// If the current path does not equal '/' and the node does not have a registered handle and the most recently matched node has a child node
			// the current node needs to roll back to last valid skippedNode
<原文结束>

# <翻译开始>
// 如果当前路径不等于'/'，并且节点没有注册句柄，并且最近匹配的节点有子节点，则当前节点需要回滚到最后一个有效的skippedNode
# <翻译结束>


<原文开始>
				//	n = latestNode.children[len(latestNode.children)-1]
<原文结束>

# <翻译开始>
// n = latestNode.children[len(latestNode.children)-1]
# <翻译结束>


<原文开始>
			// We should have reached the node containing the handle.
			// Check if this node has a handle registered.
<原文结束>

# <翻译开始>
// 我们应该已经到达包含句柄的节点
// 检查此节点是否已注册句柄
# <翻译结束>


<原文开始>
			// If there is no handle for this route, but this route has a
			// wildcard child, there must be a handle for this path with an
			// additional trailing slash
<原文结束>

# <翻译开始>
// 如果这个路由没有句柄，但是这个路由有一个通配符子节点，那么这个路径必须有一个附加斜杠的句柄
# <翻译结束>


<原文开始>
			// No handle found. Check if a handle for this path + a
			// trailing slash exists for trailing slash recommendation
<原文结束>

# <翻译开始>
// 没有找到手柄
// 检查该路径的句柄+尾斜杠是否存在以推荐尾斜杠
# <翻译结束>


<原文开始>
		// Nothing found. We can recommend to redirect to the same URL with an
		// extra trailing slash if a leaf exists for that path
<原文结束>

# <翻译开始>
// 没有什么发现
// 我们可以建议重定向到相同的URL，如果该路径存在叶子，则使用额外的斜杠
# <翻译结束>


<原文开始>
		// roll back to last valid skippedNode
<原文结束>

# <翻译开始>
// 回滚到最后一个有效的skippedNode
# <翻译结束>


<原文开始>
// Makes a case-insensitive lookup of the given path and tries to find a handler.
// It can optionally also fix trailing slashes.
// It returns the case-corrected path and a bool indicating whether the lookup
// was successful.
<原文结束>

# <翻译开始>
// 对给定路径进行不区分大小写的查找，并尝试查找处理程序
// 它还可以选择修复尾随斜杠
// 它返回经过大小写校正的路径和一个bool值，该值指示查找是否成功
# <翻译结束>


<原文开始>
	// Use a static sized buffer on the stack in the common case.
	// If the path is too long, allocate a buffer on the heap instead.
<原文结束>

# <翻译开始>
// 一般情况下，在堆栈上使用静态大小的缓冲区
// 如果路径太长，则在堆上分配一个缓冲区
# <翻译结束>


<原文开始>
// Preallocate enough memory for new path
<原文结束>

# <翻译开始>
// 为新路径预先分配足够的内存
# <翻译结束>


<原文开始>
// Empty rune buffer
<原文结束>

# <翻译开始>
// 空符文缓冲区
# <翻译结束>


<原文开始>
// Shift bytes in array by n bytes left
<原文结束>

# <翻译开始>
// 将数组中的字节向左移动n个字节
# <翻译结束>


<原文开始>
// Recursive case-insensitive lookup function used by n.findCaseInsensitivePath
<原文结束>

# <翻译开始>
// n.findCaseInsensitivePath使用的递归不区分大小写的查找函数
# <翻译结束>


<原文开始>
		// Add common prefix to result
<原文结束>

# <翻译开始>
// 在结果中添加公共前缀
# <翻译结束>


<原文开始>
			// No handle found.
			// Try to fix the path by adding a trailing slash
<原文结束>

# <翻译开始>
// 没有找到手柄
// 尝试通过添加尾斜杠来修复路径
# <翻译结束>


<原文开始>
		// If this node does not have a wildcard (param or catchAll) child,
		// we can just look up the next child node and continue to walk down
		// the tree
<原文结束>

# <翻译开始>
// 如果该节点没有通配符(param或catchAll)子节点，我们可以只查找下一个子节点并继续沿着树向下走
# <翻译结束>


<原文开始>
			// Skip rune bytes already processed
<原文结束>

# <翻译开始>
// 跳过已处理的符文字节
# <翻译结束>


<原文开始>
				// Old rune not finished
<原文结束>

# <翻译开始>
// 旧符文还没写完
# <翻译结束>


<原文开始>
						// continue with child node
<原文结束>

# <翻译开始>
// 继续子节点
# <翻译结束>


<原文开始>
				// Process a new rune
<原文结束>

# <翻译开始>
// 处理一个新的符文
# <翻译结束>


<原文开始>
				// Find rune start.
				// Runes are up to 4 byte long,
				// -4 would definitely be another rune.
<原文结束>

# <翻译开始>
// 找到符文开始
// 符文最多有4字节长，-4肯定是另一个符文
# <翻译结束>


<原文开始>
						// read rune from cached path
<原文结束>

# <翻译开始>
// 从缓存路径读取符文
# <翻译结束>


<原文开始>
				// Calculate lowercase bytes of current rune
<原文结束>

# <翻译开始>
// 计算当前符文的小写字节
# <翻译结束>


<原文开始>
				// Skip already processed bytes
<原文结束>

# <翻译开始>
// 跳过已经处理的字节
# <翻译结束>


<原文开始>
					// Lowercase matches
<原文结束>

# <翻译开始>
// 小写字母相匹配
# <翻译结束>


<原文开始>
						// must use a recursive approach since both the
						// uppercase byte and the lowercase byte might exist
						// as an index
<原文结束>

# <翻译开始>
// 必须使用递归方法，因为大写字节和小写字节都可能作为索引存在
# <翻译结束>


<原文开始>
				// If we found no match, the same for the uppercase rune,
				// if it differs
<原文结束>

# <翻译开始>
// 如果找不到匹配，大写符文也一样，如果不同
# <翻译结束>


<原文开始>
						// Uppercase matches
<原文结束>

# <翻译开始>
// 大写字母相匹配
# <翻译结束>


<原文开始>
							// Continue with child node
<原文结束>

# <翻译开始>
// 继续子节点
# <翻译结束>


<原文开始>
			// Nothing found. We can recommend to redirect to the same URL
			// without a trailing slash if a leaf exists for that path
<原文结束>

# <翻译开始>
// 没有什么发现
// 我们可以建议重定向到相同的URL，如果该路径存在叶子，则不带尾斜杠
# <翻译结束>


<原文开始>
			// Add param value to case insensitive path
<原文结束>

# <翻译开始>
// 为不区分大小写的路径添加参数值
# <翻译结束>


<原文开始>
			// We need to go deeper!
<原文结束>

# <翻译开始>
// 我们得再深入一点!
# <翻译结束>


<原文开始>
				// No handle found. Check if a handle for this path + a
				// trailing slash exists
<原文结束>

# <翻译开始>
// 没有找到手柄
// 检查是否存在此路径的句柄+尾斜杠
# <翻译结束>


<原文开始>
	// Nothing found.
	// Try to fix the path by adding / removing a trailing slash
<原文结束>

# <翻译开始>
// 没有什么发现
// 尝试通过添加/删除尾斜杠来修复路径
# <翻译结束>

