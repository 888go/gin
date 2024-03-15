
<原文开始>
// RedisStore represents the cache with redis persistence
<原文结束>

# <翻译开始>
// RedisStore represents the cache with redis persistence
# <翻译结束>


<原文开始>
// NewRedisCache returns a RedisStore
// until redigo supports sharding/clustering, only one host will be in hostList
<原文结束>

# <翻译开始>
// NewRedisCache returns a RedisStore
// until redigo supports sharding/clustering, only one host will be in hostList
# <翻译结束>


<原文开始>
			// the redis protocol should probably be made sett-able
<原文结束>

# <翻译开始>
			// the redis protocol should probably be made sett-able
# <翻译结束>


<原文开始>
				// check with PING
<原文结束>

# <翻译开始>
				// check with PING
# <翻译结束>


<原文开始>
		// custom connection test method
<原文结束>

# <翻译开始>
		// custom connection test method
# <翻译结束>


<原文开始>
			// don't need check connection every time.
<原文结束>

# <翻译开始>
			// don't need check connection every time.
# <翻译结束>


<原文开始>
// NewRedisCacheWithPool returns a RedisStore using the provided pool
// until redigo supports sharding/clustering, only one host will be in hostList
<原文结束>

# <翻译开始>
// NewRedisCacheWithPool returns a RedisStore using the provided pool
// until redigo supports sharding/clustering, only one host will be in hostList
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
	// Check for existance *before* increment as per the cache contract.
	// redis will auto create the key, and we don't want that. Since we need to do increment
	// ourselves instead of natively via INCRBY (redis doesn't support wrapping), we get the value
	// and do the exists check this way to minimize calls to Redis
<原文结束>

# <翻译开始>
	// Check for existance *before* increment as per the cache contract.
	// redis will auto create the key, and we don't want that. Since we need to do increment
	// ourselves instead of natively via INCRBY (redis doesn't support wrapping), we get the value
	// and do the exists check this way to minimize calls to Redis
# <翻译结束>


<原文开始>
// Decrement (see CacheStore interface)
<原文结束>

# <翻译开始>
// Decrement (see CacheStore interface)
# <翻译结束>


<原文开始>
	// Check for existance *before* increment as per the cache contract.
	// redis will auto create the key, and we don't want that, hence the exists call
<原文结束>

# <翻译开始>
	// Check for existance *before* increment as per the cache contract.
	// redis will auto create the key, and we don't want that, hence the exists call
# <翻译结束>


<原文开始>
	// Decrement contract says you can only go to 0
	// so we go fetch the value and if the delta is greater than the amount,
	// 0 out the value
<原文结束>

# <翻译开始>
	// Decrement contract says you can only go to 0
	// so we go fetch the value and if the delta is greater than the amount,
	// 0 out the value
# <翻译结束>


<原文开始>
// Flush (see CacheStore interface)
<原文结束>

# <翻译开始>
// Flush (see CacheStore interface)
# <翻译结束>

