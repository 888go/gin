
<原文开始>
// Test typical cache interactions
<原文结束>

# <翻译开始>
// 测试典型的缓存交互
# <翻译结束>


<原文开始>
// Test the increment-decrement cases
<原文结束>

# <翻译开始>
// 测试递增-递减用例
# <翻译结束>


<原文开始>
	// Normal increment / decrement operation.
<原文结束>

# <翻译开始>
// 正常的递增/递减操作
# <翻译结束>


<原文开始>
	// Increment wraparound
<原文结束>

# <翻译开始>
// 增量的
# <翻译结束>


<原文开始>
	// Decrement capped at 0
<原文结束>

# <翻译开始>
// 减量上限为0
# <翻译结束>


<原文开始>
	// memcached does not support expiration times less than 1 second.
<原文结束>

# <翻译开始>
// Memcached不支持小于1秒的过期时间
# <翻译结束>


<原文开始>
	// Test Set w/ DEFAULT
<原文结束>

# <翻译开始>
// 测试集w/ DEFAULT
# <翻译结束>


<原文开始>
	// Test Set w/ short time
<原文结束>

# <翻译开始>
// 短时间测试集
# <翻译结束>


<原文开始>
	// Test Set w/ longer time.
<原文结束>

# <翻译开始>
// 测试集w/更长时间
# <翻译结束>


<原文开始>
	// Test Set w/ forever.
<原文结束>

# <翻译开始>
// 测试集w/ forever
# <翻译结束>


<原文开始>
	// Replace in an empty cache.
<原文结束>

# <翻译开始>
// 在空缓存中替换
# <翻译结束>


<原文开始>
	// Set a value of 1, and replace it with 2
<原文结束>

# <翻译开始>
// 设置值为1，并用2替换
# <翻译结束>


<原文开始>
	// Wait for it to expire and replace with 3 (unsuccessfully).
<原文结束>

# <翻译开始>
// 等待它过期并替换为3(不成功)
# <翻译结束>


<原文开始>
	// Add to an empty cache.
<原文结束>

# <翻译开始>
// 添加到空缓存中
# <翻译结束>


<原文开始>
	// Try to add again. (fail)
<原文结束>

# <翻译开始>
// 试着再加一次
// (失败)
# <翻译结束>


<原文开始>
	// Wait for it to expire, and add again.
<原文结束>

# <翻译开始>
// 等待它过期，然后再次添加
# <翻译结束>


<原文开始>
	// Get and verify the value.
<原文结束>

# <翻译开始>
// 获取并验证该值
# <翻译结束>

