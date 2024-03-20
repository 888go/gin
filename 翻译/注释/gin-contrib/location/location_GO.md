
<原文开始>
// Headers represents the header fields used to map schemes and host.
<原文结束>

# <翻译开始>
// Headers 代表用于映射方案和主机的头部字段。
# <翻译结束>


<原文开始>
// Config represents all available options for the middleware.
<原文结束>

# <翻译开始>
// Config 表示该中间件的所有可用选项。
# <翻译结束>


<原文开始>
	// Scheme is the default scheme that should be used when it cannot otherwise
	// be ascertained from the incoming http.Request.
<原文结束>

# <翻译开始>
// Scheme是当无法从传入的http.Request中明确获知时，应使用的默认方案。
# <翻译结束>


<原文开始>
	// Host is the default host that should be used when it cannot otherwise
	// be ascertained from the incoming http.Request.
<原文结束>

# <翻译开始>
// Host 是默认主机，当无法从传入的 http.Request 中明确获取时，应使用此主机。
# <翻译结束>


<原文开始>
	// Base is the base path that should be used in conjunction with proxy
	// servers that do path re-writing.
<原文结束>

# <翻译开始>
// Base 是基路径，应与进行路径重写操作的代理服务器结合使用。
# <翻译结束>


<原文开始>
	// Header used to map schemes and host.
	// May be overriden to allow reading values from custom header fields.
<原文结束>

# <翻译开始>
// 该Header用于映射方案和主机。
// 可以被覆盖以允许从自定义头部字段读取值。
# <翻译结束>


<原文开始>
// DefaultConfig returns a generic default configuration mapped to localhost.
<原文结束>

# <翻译开始>
// DefaultConfig 返回一个映射到本机的通用默认配置。
# <翻译结束>


<原文开始>
// Default returns the location middleware with default configuration.
<原文结束>

# <翻译开始>
// 默认返回使用默认配置的位置中间件。
# <翻译结束>


<原文开始>
// New returns the location middleware with user-defined custom configuration.
<原文结束>

# <翻译开始>
// New 函数返回一个使用用户自定义配置的 location 中间件。
# <翻译结束>


<原文开始>
// Get returns the Location information for the incoming http.Request from the
// context. If the location is not set a nil value is returned.
<原文结束>

# <翻译开始>
// Get 从上下文获取传入 http.Request 的 Location 信息。如果未设置位置信息，则返回 nil 值。
# <翻译结束>

