
<原文开始>
// RedisStore represents the cache with redis persistence
<原文结束>

# <翻译开始>
// RedisStore代表了使用Redis进行持久化的缓存
# <翻译结束>


<原文开始>
// NewRedisCache returns a RedisStore
// until redigo supports sharding/clustering, only one host will be in hostList
<原文结束>

# <翻译开始>
// NewRedisCache 返回一个 RedisStore
// 由于 redigo 目前还不支持分片/集群，因此 hostList 中目前只能包含一个主机地址
# <翻译结束>


<原文开始>
// the redis protocol should probably be made sett-able
<原文结束>

# <翻译开始>
// redis协议可能应该设置为可配置的
# <翻译结束>


<原文开始>
// custom connection test method
<原文结束>

# <翻译开始>
// 自定义连接测试方法
# <翻译结束>


<原文开始>
// don't need check connection every time.
<原文结束>

# <翻译开始>
// 不需要每次都检查连接。
# <翻译结束>


<原文开始>
// NewRedisCacheWithPool returns a RedisStore using the provided pool
// until redigo supports sharding/clustering, only one host will be in hostList
<原文结束>

# <翻译开始>
// NewRedisCacheWithPool 使用提供的连接池返回一个 RedisStore
// 在 redigo 支持分片/集群之前，hostList 中将只包含一个主机地址
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
	// Check for existance *before* increment as per the cache contract.
	// redis will auto create the key, and we don't want that. Since we need to do increment
	// ourselves instead of natively via INCRBY (redis doesn't support wrapping), we get the value
	// and do the exists check this way to minimize calls to Redis
<原文结束>

# <翻译开始>
// 根据缓存契约，在自增之前检查是否存在。
// Redis 会自动创建键，但我们不希望这样。因为我们需要自己而不是通过原生的 INCRBY（Redis 不支持自增后循环）来执行自增操作，所以我们获取值并以这种方式进行存在性检查，以尽量减少对 Redis 的调用。
# <翻译结束>


<原文开始>
// Decrement (see CacheStore interface)
<原文结束>

# <翻译开始>
// 减量（参考 CacheStore 接口）
# <翻译结束>


<原文开始>
	// Check for existance *before* increment as per the cache contract.
	// redis will auto create the key, and we don't want that, hence the exists call
<原文结束>

# <翻译开始>
// 按照缓存契约，在递增前检查是否存在。
// Redis 会自动创建键，但我们不希望这样，因此需要调用 exists 函数。
# <翻译结束>


<原文开始>
	// Decrement contract says you can only go to 0
	// so we go fetch the value and if the delta is greater than the amount,
	// 0 out the value
<原文结束>

# <翻译开始>
// Decrement contract 表示你只能减到0
// 因此，我们获取当前值，如果减少的量大于该值，
// 则将值置为0
# <翻译结束>


<原文开始>
// Flush (see CacheStore interface)
<原文结束>

# <翻译开始>
// Flush（参考 CacheStore 接口）
# <翻译结束>

