
<原文开始>
// this is our reverse server ip address
<原文结束>

# <翻译开始>
// 这是我们的反向服务器IP地址
# <翻译结束>


<原文开始>
// maybe we can have many real server addresses and do some load balanced strategy.
<原文结束>

# <翻译开始>
// 可能我们可以拥有多个真实的服务器地址，并执行某种负载均衡策略。
# <翻译结束>


<原文开始>
// a fake function that we can do strategy here.
<原文结束>

# <翻译开始>
// 这是一个模拟函数，我们可以在其中实现策略。
# <翻译结束>


<原文开始>
// step 1: resolve proxy address, change scheme and host in requets
<原文结束>

# <翻译开始>
// 步骤1：解析代理地址，更改请求中的方案和主机
# <翻译结束>


<原文开始>
// step 2: use http.Transport to do request to real server.
<原文结束>

# <翻译开始>
// 步骤 2：使用 http.Transport 向真实服务器发起请求。
# <翻译结束>


<原文开始>
// step 3: return real server response to upstream.
<原文结束>

# <翻译开始>
// 步骤3：将真实服务器响应返回给上游。
# <翻译结束>

