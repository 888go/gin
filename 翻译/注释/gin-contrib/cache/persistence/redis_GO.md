
<原文开始>
// RedisStore represents the cache with redis persistence
<原文结束>

# <翻译开始>
// RedisStore 表示使用 Redis 进行持久化的缓存
# <翻译结束>


<原文开始>
// NewRedisCache returns a RedisStore
// until redigo supports sharding/clustering, only one host will be in hostList
<原文结束>

# <翻译开始>
// NewRedisCache 返回一个 RedisStore
// 由于 redigo 库目前还不支持分片/集群，所以 hostList 中当前仅包含一个主机地址
# <翻译结束>


<原文开始>
			// the redis protocol should probably be made sett-able
<原文结束>

# <翻译开始>
// redis协议可能应该设置为可配置的
# <翻译结束>


<原文开始>
				// check with PING
<原文结束>

# <翻译开始>
// 使用PING进行检查
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
// 直到 redigo 支持分片/集群，hostList 中将仅包含一个主机
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
	// Check for existance *before* increment as per the cache contract.
	// redis will auto create the key, and we don't want that. Since we need to do increment
	// ourselves instead of natively via INCRBY (redis doesn't support wrapping), we get the value
	// and do the exists check this way to minimize calls to Redis
<原文结束>

# <翻译开始>
// 根据缓存契约，在自增操作**之前**检查键是否存在。
// Redis 会自动创建键，但我们并不希望这样。因为我们需要自己进行自增操作（由于 Redis 不支持数值溢出，不能直接使用 INCRBY 命令），所以我们先获取值，通过这种方式来检查键是否存在，以尽量减少对 Redis 的调用次数。
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
// 根据缓存契约，在自增之前检查键是否存在。
// Redis 会自动创建键，但我们不希望这样，因此需要调用 exists 方法。
# <翻译结束>


<原文开始>
	// Decrement contract says you can only go to 0
	// so we go fetch the value and if the delta is greater than the amount,
	// 0 out the value
<原文结束>

# <翻译开始>
// 减少合约规定只能减到0
// 因此我们获取当前值，如果减少量大于该值，则将值置零
# <翻译结束>


<原文开始>
// Flush (see CacheStore interface)
<原文结束>

# <翻译开始>
// Flush （参考 CacheStore 接口）
# <翻译结束>

