
<原文开始>
// MemcachedBinaryStore represents the cache with memcached persistence using
// the binary protocol
<原文结束>

# <翻译开始>
// MemcachedBinaryStore代表使用二进制协议的缓存，其持久化机制为Memcached
# <翻译结束>


<原文开始>
// NewMemcachedBinaryStore returns a MemcachedBinaryStore
<原文结束>

# <翻译开始>
// NewMemcachedBinaryStore 返回一个 MemcachedBinaryStore
# <翻译结束>


<原文开始>
// NewMemcachedBinaryStoreWithConfig returns a MemcachedBinaryStore using the provided configuration
<原文结束>

# <翻译开始>
// NewMemcachedBinaryStoreWithConfig 使用提供的配置返回一个 MemcachedBinaryStore 实例
# <翻译结束>


<原文开始>
// Set (see CacheStore interface)
<原文结束>

# <翻译开始>
// Set（参见 CacheStore 接口）
# <翻译结束>


<原文开始>
// Add (see CacheStore interface)
<原文结束>

# <翻译开始>
// Add （参见 CacheStore 接口）
# <翻译结束>


<原文开始>
// Replace (see CacheStore interface)
<原文结束>

# <翻译开始>
// Replace（参见 CacheStore 接口）
# <翻译结束>


<原文开始>
// Get (see CacheStore interface)
<原文结束>

# <翻译开始>
// Get（参见 CacheStore 接口）
# <翻译结束>


<原文开始>
// Delete (see CacheStore interface)
<原文结束>

# <翻译开始>
// Delete（参考 CacheStore 接口）
# <翻译结束>


<原文开始>
// Increment (see CacheStore interface)
<原文结束>

# <翻译开始>
// 自增（参见 CacheStore 接口）
# <翻译结束>


<原文开始>
// Decrement (see CacheStore interface)
<原文结束>

# <翻译开始>
// 减量（参考 CacheStore 接口）
# <翻译结束>


<原文开始>
// Flush (see CacheStore interface)
<原文结束>

# <翻译开始>
// Flush（参考 CacheStore 接口）
# <翻译结束>


<原文开始>
// getExpiration converts a gin-contrib/cache expiration in the form of a
// time.Duration to a valid memcached expiration either in seconds (<30 days)
// or a Unix timestamp (>30 days)
<原文结束>

# <翻译开始>
// getExpiration 将gin-contrib/cache中以time.Duration形式表示的过期时间转换为有效的memcached过期时间，
// 过期时间要么表示为秒（小于30天），要么表示为Unix时间戳（大于30天）
# <翻译结束>

