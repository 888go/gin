
<原文开始>
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
<原文结束>

# <翻译开始>
// 在一个goroutine中初始化服务器，以便于
// 不会阻塞下面的优雅关闭处理流程
# <翻译结束>


<原文开始>
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
<原文结束>

# <翻译开始>
// 等待中断信号以优雅地关闭服务器，超时时间为5秒。
# <翻译结束>


<原文开始>
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
<原文结束>

# <翻译开始>
// (无参数) kill 默认发送 syscall.SIGTERM
// kill -2 等同于 syscall.SIGINT
// kill -9 等同于 syscall.SIGKILL，但无法被捕获，因此不需要添加它
# <翻译结束>


<原文开始>
	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
<原文结束>

# <翻译开始>
// 上下文用于通知服务器，它有5秒钟的时间来完成当前正在处理的请求
# <翻译结束>

