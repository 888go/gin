
<原文开始>
		// Always set Vary headers
		// see https://github.com/rs/cors/issues/10,
		// https://github.com/rs/cors/commit/dbdca4d95feaa7511a46e6f1efb3b3aa505bc43f#commitcomment-12352001
<原文结束>

# <翻译开始>
// 总是设置 Vary 头信息
// 参见 https://github.com/rs/cors/issues/10，
// 及 https://github.com/rs/cors/commit/dbdca4d95feaa7511a46e6f1efb3b3aa505bc43f#commitcomment-12352001
// 此处的注释指出，为了正确处理跨域资源共享（CORS），应始终设置“Vary”HTTP头。这有助于确保在处理预检请求（OPTIONS）以及其他CORS相关场景时，服务器能够提供正确的响应头信息。参考给出的GitHub链接以获取更多上下文和详细信息。
# <翻译结束>

