
<原文开始>
	// From go/net/http/request.go:
	// For incoming requests, the Host header is promoted to the
	// Request.Host field and removed from the Header map.
<原文结束>

# <翻译开始>
// 来自 go/net/http/request.go:
// 对于传入的请求，Host 头部会被提升至 Request.Host 字段，并从 Header 映射中移除。
# <翻译结束>


<原文开始>
// no CORS request, origin == ""
<原文结束>

# <翻译开始>
// 没有CORS请求，origin（来源）为空字符串
# <翻译结束>


<原文开始>
// no CORS request, origin == host
<原文结束>

# <翻译开始>
// 无CORS请求，origin（来源）== host（主机）
# <翻译结束>


<原文开始>
// allowed CORS prefligh request
<原文结束>

# <翻译开始>
// 允许CORS预检请求
# <翻译结束>


<原文开始>
// deny CORS prefligh request
<原文结束>

# <翻译开始>
// 拒绝CORS预检请求
# <翻译结束>

