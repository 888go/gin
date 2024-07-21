
<原文开始>
// It keeps a list of clients those are currently attached
// and broadcasting events to those clients.
<原文结束>

# <翻译开始>
// 它维护一个当前已连接的客户端列表，并向这些客户端广播事件。
# <翻译结束>


<原文开始>
// Events are pushed to this channel by the main events-gathering routine
<原文结束>

# <翻译开始>
// 主要事件收集程序会将事件推送到此通道
# <翻译结束>


<原文开始>
// Closed client connections
<原文结束>

# <翻译开始>
// 已关闭的客户端连接
# <翻译结束>


<原文开始>
// Total client connections
<原文结束>

# <翻译开始>
// 总客户端连接数
# <翻译结束>


<原文开始>
// New event messages are broadcast to all registered client connection channels
<原文结束>

# <翻译开始>
// 新的事件消息将被广播到所有已注册的客户端连接通道
# <翻译结束>


<原文开始>
// Initialize new streaming server
<原文结束>

# <翻译开始>
// 初始化新的流媒体服务器
# <翻译结束>


<原文开始>
// We are streaming current time to clients in the interval 10 seconds
<原文结束>

# <翻译开始>
// 我们以10秒间隔向客户端流式传输当前时间
# <翻译结束>


<原文开始>
// Send current time to clients message channel
<原文结束>

# <翻译开始>
// 将当前时间发送到客户端消息通道
# <翻译结束>


<原文开始>
// username : admin, password : admin123
<原文结束>

# <翻译开始>
// 用户名：admin，密码：admin123
# <翻译结束>


<原文开始>
	// Authorized client can stream the event
	// Add event-streaming headers
<原文结束>

# <翻译开始>
	// 授权的客户端可以流式接收该事件
	// 添加事件流传输所需的头部信息
# <翻译结束>


<原文开始>
// Stream message to client from message channel
<原文结束>

# <翻译开始>
// 从消息通道向客户端流式传输消息
# <翻译结束>


<原文开始>
// Initialize event and Start procnteessing requests
<原文结束>

# <翻译开始>
// 初始化事件并开始处理请求
# <翻译结束>


<原文开始>
// It Listens all incoming requests from clients.
// Handles addition and removal of clients and broadcast messages to clients.
<原文结束>

# <翻译开始>
// 它监听来自客户端的所有入站请求。
// 处理客户端的添加和移除，并向客户端广播消息。
# <翻译结束>


<原文开始>
// Add new available client
<原文结束>

# <翻译开始>
// 添加新的可用客户端
# <翻译结束>


<原文开始>
// Broadcast message to client
<原文结束>

# <翻译开始>
// 向客户端广播消息
# <翻译结束>


<原文开始>
// Initialize client channel
<原文结束>

# <翻译开始>
// 初始化客户端通道
# <翻译结束>


<原文开始>
// Send new connection to event server
<原文结束>

# <翻译开始>
// 将新的连接发送至事件服务器
# <翻译结束>


<原文开始>
// Send closed connection to event server
<原文结束>

# <翻译开始>
// 将关闭的连接发送到事件服务器
# <翻译结束>

