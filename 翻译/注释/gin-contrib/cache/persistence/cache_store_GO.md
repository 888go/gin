
<原文开始>
// CacheStore is the interface of a cache backend
<原文结束>

# <翻译开始>
// CacheStore是缓存后端接口
# <翻译结束>


<原文开始>
	// Get retrieves an item from the cache. Returns the item or nil, and a bool indicating
	// whether the key was found.
<原文结束>

# <翻译开始>
// Get从缓存中检索项
// 返回项或nil，以及指示是否找到键的bool值
# <翻译结束>


<原文开始>
	// Set sets an item to the cache, replacing any existing item.
<原文结束>

# <翻译开始>
// Set将项设置到缓存，替换任何现有项
# <翻译结束>


<原文开始>
	// Add adds an item to the cache only if an item doesn't already exist for the given
	// key, or if the existing item has expired. Returns an error otherwise.
<原文结束>

# <翻译开始>
// Add仅在给定键的项不存在或现有项已过期时向缓存添加项
// 否则返回错误
# <翻译结束>


<原文开始>
	// Replace sets a new value for the cache key only if it already exists. Returns an
	// error if it does not.
<原文结束>

# <翻译开始>
// Replace仅在缓存键已经存在时才为该键设置新值
// 如果没有，则返回错误
# <翻译结束>


<原文开始>
	// Delete removes an item from the cache. Does nothing if the key is not in the cache.
<原文结束>

# <翻译开始>
// Delete从缓存中删除项
// 如果键不在缓存中，则不执行任何操作
# <翻译结束>


<原文开始>
	// Increment increments a real number, and returns error if the value is not real
<原文结束>

# <翻译开始>
// Increment对实数递增，如果值不是实数则返回error
# <翻译结束>


<原文开始>
	// Decrement decrements a real number, and returns error if the value is not real
<原文结束>

# <翻译开始>
// 递减一个实数，如果值不是实数，则返回错误
# <翻译结束>


<原文开始>
	// Flush deletes all items from the cache.
<原文结束>

# <翻译开始>
// 刷新从缓存中删除所有项
# <翻译结束>

