
<原文开始>
// Headers represents the header fields used to map schemes and host.
<原文结束>

# <翻译开始>
// Headers 表示用于映射方案和主机的头部字段。
# <翻译结束>


<原文开始>
// Config represents all available options for the middleware.
<原文结束>

# <翻译开始>
// Config 代表了该中间件可用的所有配置选项。
# <翻译结束>


<原文开始>
	// Scheme is the default scheme that should be used when it cannot otherwise
	// be ascertained from the incoming http.Request.
<原文结束>

# <翻译开始>
// Scheme 是默认的方案，当无法从传入的 http.Request 中确定时应使用此方案。
# <翻译结束>


<原文开始>
	// Host is the default host that should be used when it cannot otherwise
	// be ascertained from the incoming http.Request.
<原文结束>

# <翻译开始>
// Host 是默认主机，当无法从传入的 http.Request 中确定时应使用该主机。
# <翻译结束>


<原文开始>
	// Base is the base path that should be used in conjunction with proxy
	// servers that do path re-writing.
<原文结束>

# <翻译开始>
// Base 是基础路径，应与执行路径重写操作的代理服务器结合使用。
# <翻译结束>


<原文开始>
	// Header used to map schemes and host.
	// May be overriden to allow reading values from custom header fields.
<原文结束>

# <翻译开始>
// Header 用于映射方案和主机。
// 可以重写以允许从自定义头部字段读取值。
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
// Default 返回默认配置的 location 中间件。
# <翻译结束>


<原文开始>
// New returns the location middleware with user-defined custom configuration.
<原文结束>

# <翻译开始>
// New 返回一个带有用户自定义配置的 location 中间件。
# <翻译结束>


<原文开始>
// Get returns the Location information for the incoming http.Request from the
// context. If the location is not set a nil value is returned.
<原文结束>

# <翻译开始>
// Get 从 context 中获取传入 http.Request 的 Location 信息。如果未设置 location，则返回空值（nil）。
# <翻译结束>

