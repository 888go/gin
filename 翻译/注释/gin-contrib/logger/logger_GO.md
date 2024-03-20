
<原文开始>
// Config defines the config for logger middleware
<原文结束>

# <翻译开始>
// Config 定义了 logger 中间件的配置
# <翻译结束>


<原文开始>
// UTC a boolean stating whether to use UTC time zone or local.
<原文结束>

# <翻译开始>
// UTC 是一个布尔值，表示是否使用 UTC 时区或本地时区。
# <翻译结束>


<原文开始>
	// Output is a writer where logs are written.
	// Optional. Default value is gin.DefaultWriter.
<原文结束>

# <翻译开始>
// Output 是一个用于写入日志的writer。
// 可选配置，默认值为gin.DefaultWriter。
# <翻译结束>


<原文开始>
// the log level used for request with status code < 400
<原文结束>

# <翻译开始>
// 用于状态码小于400的请求的日志级别
# <翻译结束>


<原文开始>
// the log level used for request with status code between 400 and 499
<原文结束>

# <翻译开始>
// 用于状态码在400至499之间的请求的日志级别
# <翻译结束>


<原文开始>
// the log level used for request with status code >= 500
<原文结束>

# <翻译开始>
// 用于状态码大于等于500的请求的日志级别
# <翻译结束>


<原文开始>
// SetLogger initializes the logging middleware.
<原文结束>

# <翻译开始>
// SetLogger 初始化日志中间件。
# <翻译结束>


<原文开始>
// Loop through each option
<原文结束>

# <翻译开始>
// 遍历每个选项
# <翻译结束>


<原文开始>
// Call the option giving the instantiated
<原文结束>

# <翻译开始>
// 调用选项，传入已实例化的
# <翻译结束>


<原文开始>
// ParseLevel converts a level string into a zerolog Level value.
// returns an error if the input string does not match known values.
<原文结束>

# <翻译开始>
// ParseLevel将级别字符串转换为zerolog的Level值。
// 如果输入字符串与已知值不匹配，则返回错误。
# <翻译结束>

