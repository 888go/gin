
<原文开始>
	// only send crash reporting
	// r.Use(sentry.Recovery(raven.DefaultClient, true))
<原文结束>

# <翻译开始>
// 仅发送崩溃报告
// r.Use(sentry.Recovery(raven.DefaultClient, true)) 
// 
// （翻译为：）
// 
// 只启用崩溃报告发送功能
// r.Use(sentry提供的恢复中间件，使用raven的默认客户端，并开启true参数以捕获所有panic信息）
# <翻译结束>

