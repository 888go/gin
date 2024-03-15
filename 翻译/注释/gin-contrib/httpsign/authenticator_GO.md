
<原文开始>
// Authenticator is the gin authenticator middleware.
<原文结束>

# <翻译开始>
// Authenticator 是 Gin 框架的身份验证中间件。
# <翻译结束>


<原文开始>
// Option is the option to the Authenticator constructor.
<原文结束>

# <翻译开始>
// Option 是 Authenticator 构造函数的选项。
# <翻译结束>


<原文开始>
// WithValidator configures the Authenticator to use custom validator.
// The default validators are time based and digest.
<原文结束>

# <翻译开始>
// WithValidator 配置 Authenticator 以使用自定义验证器。
// 默认的验证器基于时间和摘要。
# <翻译结束>


<原文开始>
// WithRequiredHeaders is list of all requires HTTP headers that the client
// have to include in the singing string for the request to be considered valid.
// If not provided, the created Authenticator instance will use defaultRequiredHeaders variable.
<原文结束>

# <翻译开始>
// WithRequiredHeaders 是一个包含所有必需HTTP头的列表，客户端
// 必须在签名字符串中包含这些头信息，以便请求被认为是有效的。
// 如果未提供，创建的Authenticator实例将使用默认的defaultRequiredHeaders变量。
# <翻译结束>


<原文开始>
// NewAuthenticator creates a new Authenticator instance with
// given allowed permissions and required header and secret keys.
<原文结束>

# <翻译开始>
// NewAuthenticator 创建一个全新的 Authenticator 实例，其参数包括
// 允许的权限列表以及请求头中所需的密钥和秘密键。
# <翻译结束>


<原文开始>
// Authenticated returns a gin middleware which permits given permissions in parameter.
<原文结束>

# <翻译开始>
// Authenticated 返回一个 gin 中间件，该中间件允许在参数中指定的权限。
# <翻译结束>


<原文开始>
// isValidHeader check if all server required header is in header list
<原文结束>

# <翻译开始>
// isValidHeader 检查是否所有服务器要求的头部都在头部列表中
# <翻译结束>

