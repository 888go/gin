
<原文开始>
// It keeps a list of clients those are currently attached
// and broadcasting events to those clients.
<原文结束>

# <翻译开始>
// 它保留当前附加的客户机列表，并向这些客户机广播事件
# <翻译结束>


<原文开始>
	// Events are pushed to this channel by the main events-gathering routine
<原文结束>

# <翻译开始>
// 事件由主事件收集例程推送到此通道
# <翻译结束>


<原文开始>
	// New client connections
<原文结束>

# <翻译开始>
// 新的客户端连接
# <翻译结束>


<原文开始>
	// Closed client connections
<原文结束>

# <翻译开始>
// 关闭的客户端连接
# <翻译结束>


<原文开始>
	// Total client connections
<原文结束>

# <翻译开始>
// 客户端总连接数
# <翻译结束>


<原文开始>
// New event messages are broadcast to all registered client connection channels
<原文结束>

# <翻译开始>
// 将新事件消息广播到所有已注册的客户端连接通道
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
// 我们以10秒的间隔将当前时间流式传输给客户端
# <翻译结束>


<原文开始>
			// Send current time to clients message channel
<原文结束>

# <翻译开始>
// 发送当前时间到客户端消息通道
# <翻译结束>


<原文开始>
	// Basic Authentication
<原文结束>

# <翻译开始>
// 基本身份验证
# <翻译结束>


<原文开始>
// username : admin, password : admin123
<原文结束>

# <翻译开始>
// 用户名:admin，密码:admin123
# <翻译结束>


<原文开始>
	// Authorized client can stream the event
	// Add event-streaming headers
<原文结束>

# <翻译开始>
// 授权客户端可以流式传输事件
# <翻译结束>


<原文开始>
			// Stream message to client from message channel
<原文结束>

# <翻译开始>
// 从消息通道流消息到客户端
# <翻译结束>


<原文开始>
	// Parse Static files
<原文结束>

# <翻译开始>
// 解析静态文件
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
// 它监听来自客户端的所有传入请求
// 处理添加和删除客户端以及向客户端广播消息
# <翻译结束>


<原文开始>
		// Add new available client
<原文结束>

# <翻译开始>
// 添加新的可用客户端
# <翻译结束>


<原文开始>
		// Remove closed client
<原文结束>

# <翻译开始>
// 删除关闭的客户端
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
// 向事件服务器发送新连接
# <翻译结束>


<原文开始>
			// Send closed connection to event server
<原文结束>

# <翻译开始>
// 向事件服务器发送关闭的连接
# <翻译结束>

