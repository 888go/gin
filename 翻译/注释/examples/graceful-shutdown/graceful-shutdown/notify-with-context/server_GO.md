
<原文开始>
// Create context that listens for the interrupt signal from the OS.
<原文结束>

# <翻译开始>
// 创建一个上下文，用于监听来自操作系统的中断信号。
# <翻译结束>


<原文开始>
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
<原文结束>

# <翻译开始>
	// 在一个goroutine中初始化服务器，以便于
	// 不会阻塞下面的优雅关闭处理流程
# <翻译结束>


<原文开始>
// Listen for the interrupt signal.
<原文结束>

# <翻译开始>
// 监听中断信号。
# <翻译结束>


<原文开始>
// Restore default behavior on the interrupt signal and notify user of shutdown.
<原文结束>

# <翻译开始>
// 恢复对中断信号的默认处理行为，并通知用户系统即将关闭。
# <翻译结束>


<原文开始>
	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
<原文结束>

# <翻译开始>
	// 上下文用于通知服务器，它有5秒钟的时间来完成当前正在处理的请求
# <翻译结束>

