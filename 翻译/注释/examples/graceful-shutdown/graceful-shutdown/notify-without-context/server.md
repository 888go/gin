
<原文开始>
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
<原文结束>

# <翻译开始>
// 在运行例程中初始化服务器，使其不会阻塞下面的正常关机处理
# <翻译结束>


<原文开始>
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
<原文结束>

# <翻译开始>
// 等待中断信号，以5秒的超时时间正常关闭服务器
# <翻译结束>


<原文开始>
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
<原文结束>

# <翻译开始>
// Kill(无参数)默认发送系统调用
// SIGTERM kill -2是系统调用
// SIGINT kill -9是系统调用
// 但是不能被捕获，所以不需要添加它
# <翻译结束>


<原文开始>
	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
<原文结束>

# <翻译开始>
// 上下文用于通知服务器，它有5秒的时间来完成当前正在处理的请求
# <翻译结束>

