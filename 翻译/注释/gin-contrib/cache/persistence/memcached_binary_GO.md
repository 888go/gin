
<原文开始>
// MemcachedBinaryStore represents the cache with memcached persistence using
// the binary protocol
<原文结束>

# <翻译开始>
// MemcachedBinaryStore表示使用二进制协议的memcached持久性缓存
# <翻译结束>


<原文开始>
// NewMemcachedBinaryStore returns a MemcachedBinaryStore
<原文结束>

# <翻译开始>
// NewMemcachedBinaryStore返回一个MemcachedBinaryStore
# <翻译结束>


<原文开始>
// NewMemcachedBinaryStoreWithConfig returns a MemcachedBinaryStore using the provided configuration
<原文结束>

# <翻译开始>
// NewMemcachedBinaryStoreWithConfig returns a MemcachedBinaryStore using the provided configuration
# <翻译结束>


<原文开始>
// Set (see CacheStore interface)
<原文结束>

# <翻译开始>
// Set (see CacheStore interface)
# <翻译结束>


<原文开始>
// Add (see CacheStore interface)
<原文结束>

# <翻译开始>
// Add (see CacheStore interface)
# <翻译结束>


<原文开始>
// Replace (see CacheStore interface)
<原文结束>

# <翻译开始>
// Replace (see CacheStore interface)
# <翻译结束>


<原文开始>
// Get (see CacheStore interface)
<原文结束>

# <翻译开始>
// Get (see CacheStore interface)
# <翻译结束>


<原文开始>
// Delete (see CacheStore interface)
<原文结束>

# <翻译开始>
// Delete (see CacheStore interface)
# <翻译结束>


<原文开始>
// Increment (see CacheStore interface)
<原文结束>

# <翻译开始>
// Increment (see CacheStore interface)
# <翻译结束>


<原文开始>
// Decrement (see CacheStore interface)
<原文结束>

# <翻译开始>
// Decrement (see CacheStore interface)
# <翻译结束>


<原文开始>
// Flush (see CacheStore interface)
<原文结束>

# <翻译开始>
// Flush (see CacheStore interface)
# <翻译结束>


<原文开始>
// getExpiration converts a gin-contrib/cache expiration in the form of a
// time.Duration to a valid memcached expiration either in seconds (<30 days)
// or a Unix timestamp (>30 days)
<原文结束>

# <翻译开始>
// getExpiration converts a gin-contrib/cache expiration in the form of a
// time.Duration to a valid memcached expiration either in seconds (<30 days)
// or a Unix timestamp (>30 days)
# <翻译结束>


<原文开始>
// > 30 days
<原文结束>

# <翻译开始>
// > 30 days
# <翻译结束>

