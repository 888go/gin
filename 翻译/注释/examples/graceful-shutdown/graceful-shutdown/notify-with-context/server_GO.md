
<原文开始>
// build go1.16
<原文结束>

# <翻译开始>
// 构建go1.16
# <翻译结束>


<原文开始>
	// Create context that listens for the interrupt signal from the OS.
<原文结束>

# <翻译开始>
// 创建上下文，监听来自操作系统的中断信号
# <翻译结束>


<原文开始>
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
<原文结束>

# <翻译开始>
// 在运行例程中初始化服务器，使其不会阻塞下面的正常关机处理
# <翻译结束>


<原文开始>
	// Listen for the interrupt signal.
<原文结束>

# <翻译开始>
// 监听中断信号
# <翻译结束>


<原文开始>
	// Restore default behavior on the interrupt signal and notify user of shutdown.
<原文结束>

# <翻译开始>
// 恢复中断信号的默认行为，并通知用户关机
# <翻译结束>


<原文开始>
	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
<原文结束>

# <翻译开始>
// 上下文用于通知服务器，它有5秒的时间来完成当前正在处理的请求
# <翻译结束>

