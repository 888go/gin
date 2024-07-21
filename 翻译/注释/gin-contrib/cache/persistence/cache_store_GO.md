
<原文开始>
// CacheStore is the interface of a cache backend
<原文结束>

# <翻译开始>
// CacheStore 是缓存后端的接口
# <翻译结束>


<原文开始>
	// Get retrieves an item from the cache. Returns the item or nil, and a bool indicating
	// whether the key was found.
<原文结束>

# <翻译开始>
	// Get 从缓存中检索一个项目。返回该项目或nil，以及一个布尔值，表示是否找到了该键。
# <翻译结束>


<原文开始>
// Set sets an item to the cache, replacing any existing item.
<原文结束>

# <翻译开始>
// Set 将一个项设置到缓存中，替换任何已存在的项。
# <翻译结束>


<原文开始>
	// Add adds an item to the cache only if an item doesn't already exist for the given
	// key, or if the existing item has expired. Returns an error otherwise.
<原文结束>

# <翻译开始>
	// Add 将一个项添加到缓存中，但只有在给定键下尚未存在项，或者已存在的项已过期时才会添加。否则返回错误。
# <翻译结束>


<原文开始>
	// Replace sets a new value for the cache key only if it already exists. Returns an
	// error if it does not.
<原文结束>

# <翻译开始>
	// Replace仅当缓存键已存在时设置新的值。如果不存在，则返回错误。
# <翻译结束>


<原文开始>
// Delete removes an item from the cache. Does nothing if the key is not in the cache.
<原文结束>

# <翻译开始>
// Delete 从缓存中移除一个项目。如果键不在缓存中，则不执行任何操作。
# <翻译结束>


<原文开始>
// Increment increments a real number, and returns error if the value is not real
<原文结束>

# <翻译开始>
// Increment 函数对一个实数进行增加操作，并在值不是实数时返回错误
# <翻译结束>


<原文开始>
// Decrement decrements a real number, and returns error if the value is not real
<原文结束>

# <翻译开始>
// Decrement 函数对一个实数进行减一操作，如果该值不是实数，则返回错误
# <翻译结束>


<原文开始>
// Flush deletes all items from the cache.
<原文结束>

# <翻译开始>
// Flush 清空缓存中的所有项。
# <翻译结束>

