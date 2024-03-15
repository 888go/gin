
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
// NewMemcachedBinaryStoreWithConfig 根据提供的配置返回一个 MemcachedBinaryStore 实例
# <翻译结束>


<原文开始>
// Set (see CacheStore interface)
<原文结束>

# <翻译开始>
// Set (参考CacheStore接口)
// 此处的Set方法是实现CacheStore接口的一部分，用于设置缓存值。
# <翻译结束>


<原文开始>
// Add (see CacheStore interface)
<原文结束>

# <翻译开始>
// Add (参考 CacheStore 接口)
# <翻译结束>


<原文开始>
// Replace (see CacheStore interface)
<原文结束>

# <翻译开始>
// Replace (参考 CacheStore 接口)
# <翻译结束>


<原文开始>
// Get (see CacheStore interface)
<原文结束>

# <翻译开始>
// Get (参考 CacheStore 接口)
# <翻译结束>


<原文开始>
// Delete (see CacheStore interface)
<原文结束>

# <翻译开始>
// Delete (参考 CacheStore 接口)
// （该注释表明此“Delete”方法实现了 CacheStore 接口中的“Delete”方法，具体功能请参照 CacheStore 接口定义）
# <翻译结束>


<原文开始>
// Increment (see CacheStore interface)
<原文结束>

# <翻译开始>
// 自增（参见CacheStore接口）
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
// Flush （参考 CacheStore 接口）
# <翻译结束>


<原文开始>
// getExpiration converts a gin-contrib/cache expiration in the form of a
// time.Duration to a valid memcached expiration either in seconds (<30 days)
// or a Unix timestamp (>30 days)
<原文结束>

# <翻译开始>
// getExpiration 将 gin-contrib/cache 中以 time.Duration 形式表示的过期时间转换为有效的 memcached 过期时间，
// 如果时间小于30天，则转换为秒数表示；如果大于30天，则转换为 Unix 时间戳表示。
# <翻译结束>


<原文开始>
// > 30 days
<原文结束>

# <翻译开始>
// 大于30天
# <翻译结束>

