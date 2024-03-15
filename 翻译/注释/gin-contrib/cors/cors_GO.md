
<原文开始>
// Config represents all available options for the middleware.
<原文结束>

# <翻译开始>
// Config 代表了该中间件可用的所有配置选项。
# <翻译结束>


<原文开始>
	// AllowOrigins is a list of origins a cross-domain request can be executed from.
	// If the special "*" value is present in the list, all origins will be allowed.
	// Default value is []
<原文结束>

# <翻译开始>
// AllowOrigins 是一个允许跨域请求发起的源列表。
// 如果该列表中存在特殊的 "*" 值，则所有来源都将被允许。
// 默认值为 []
# <翻译结束>


<原文开始>
	// AllowOriginFunc is a custom function to validate the origin. It takes the origin
	// as an argument and returns true if allowed or false otherwise. If this option is
	// set, the content of AllowOrigins is ignored.
<原文结束>

# <翻译开始>
// AllowOriginFunc 是一个自定义函数，用于验证请求来源。它接收一个起源（origin）作为参数，并在允许访问时返回 true，否则返回 false。如果设置了这个选项，则会忽略 AllowOrigins 的内容。
# <翻译结束>


<原文开始>
	// AllowMethods is a list of methods the client is allowed to use with
	// cross-domain requests. Default value is simple methods (GET, POST, PUT, PATCH, DELETE, HEAD, and OPTIONS)
<原文结束>

# <翻译开始>
// AllowMethods 是一个方法列表，用于指定客户端在跨域请求中允许使用的方法。默认值包括常见的简单方法（GET, POST, PUT, PATCH, DELETE, HEAD 和 OPTIONS）
# <翻译结束>


<原文开始>
	// AllowPrivateNetwork indicates whether the response should include allow private network header
<原文结束>

# <翻译开始>
// AllowPrivateNetwork 表示响应中是否应包含允许私有网络头信息
# <翻译结束>


<原文开始>
	// AllowHeaders is list of non simple headers the client is allowed to use with
	// cross-domain requests.
<原文结束>

# <翻译开始>
// AllowHeaders 是一个包含客户端在跨域请求中允许使用的非简单头部列表。
# <翻译结束>


<原文开始>
	// AllowCredentials indicates whether the request can include user credentials like
	// cookies, HTTP authentication or client side SSL certificates.
<原文结束>

# <翻译开始>
// AllowCredentials 指示请求是否可以包含用户凭据，如cookies、HTTP认证或客户端SSL证书。
# <翻译结束>


<原文开始>
	// ExposeHeaders indicates which headers are safe to expose to the API of a CORS
	// API specification
<原文结束>

# <翻译开始>
// ExposeHeaders 表示哪些头部信息是安全的，可以暴露给 CORS（跨源资源共享）API 规范的 API。
# <翻译结束>


<原文开始>
	// MaxAge indicates how long (with second-precision) the results of a preflight request
	// can be cached
<原文结束>

# <翻译开始>
// MaxAge 指示预检请求结果可以被缓存的时间长度（以秒为单位）
# <翻译结束>


<原文开始>
	// Allows to add origins like http://some-domain/*, https://api.* or http://some.*.subdomain.com
<原文结束>

# <翻译开始>
// 允许添加如下形式的源地址：
// http://some-domain/* 表示以 http://some-domain/ 开头的所有URL
// https://api.* 表示以 https://api. 开头，且域名后跟任意字符的所有URL
// http://some.*.subdomain.com 表示以 http://some. 以及任意子域名.subdomain.com 结尾的所有URL
# <翻译结束>


<原文开始>
	// Allows usage of popular browser extensions schemas
<原文结束>

# <翻译开始>
// 允许使用流行的浏览器扩展方案
# <翻译结束>


<原文开始>
	// Allows usage of WebSocket protocol
<原文结束>

# <翻译开始>
// 允许使用WebSocket协议
# <翻译结束>


<原文开始>
	// Allows usage of file:// schema (dangerous!) use it only when you 100% sure it's needed
<原文结束>

# <翻译开始>
// 允许使用 file:// 协议模式（危险！）仅在您100%确定需要时才使用它
# <翻译结束>


<原文开始>
	// Allows to pass custom OPTIONS response status code for old browsers / clients
<原文结束>

# <翻译开始>
// 允许为旧版浏览器/客户端传入自定义的OPTIONS响应状态码
# <翻译结束>


<原文开始>
// AddAllowMethods is allowed to add custom methods
<原文结束>

# <翻译开始>
// AddAllowMethods 允许添加自定义方法
# <翻译结束>


<原文开始>
// AddAllowHeaders is allowed to add custom headers
<原文结束>

# <翻译开始>
// AddAllowHeaders 允许添加自定义头信息
# <翻译结束>


<原文开始>
// AddExposeHeaders is allowed to add custom expose headers
<原文结束>

# <翻译开始>
// AddExposeHeaders 允许添加自定义暴露头信息
# <翻译结束>


<原文开始>
// Validate is check configuration of user defined.
<原文结束>

# <翻译开始>
// Validate 是用于检查用户自定义配置的功能。
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

