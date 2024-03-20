
<原文开始>
// Writer is a writer with memory buffer
<原文结束>

# <翻译开始>
// Writer 是一个带有内存缓冲区的写入器
# <翻译结束>


<原文开始>
// NewWriter will return a timeout.Writer pointer
<原文结束>

# <翻译开始>
// NewWriter 将返回一个 timeout.Writer 指针
# <翻译结束>


<原文开始>
// Write will write data to response body
<原文结束>

# <翻译开始>
// Write 将数据写入响应体
# <翻译结束>


<原文开始>
// WriteHeader sends an HTTP response header with the provided status code.
// If the response writer has already written headers or if a timeout has occurred,
// this method does nothing.
<原文结束>

# <翻译开始>
// WriteHeader 向客户端发送带有指定状态码的 HTTP 响应头。
// 如果响应写入器已写入了头信息，或者发生超时，
// 此方法将不做任何操作。
# <翻译结束>


<原文开始>
	// gin is using -1 to skip writing the status code
	// see https://github.com/gin-gonic/gin/blob/a0acf1df2814fcd828cb2d7128f2f4e2136d3fac/response_writer.go#L61
<原文结束>

# <翻译开始>
// gin 使用 -1 来跳过写入状态码
// 详情见 https://github.com/gin-gonic/gin/blob/a0acf1df2814fcd828cb2d7128f2f4e2136d3fac/response_writer.go#L61
# <翻译结束>


<原文开始>
// Header will get response headers
<原文结束>

# <翻译开始>
// Header 将获取响应头
# <翻译结束>


<原文开始>
// WriteString will write string to response body
<原文结束>

# <翻译开始>
// WriteString 将字符串写入响应体
# <翻译结束>


<原文开始>
// FreeBuffer will release buffer pointer
<原文结束>

# <翻译开始>
// FreeBuffer 会释放缓冲区指针
# <翻译结束>


<原文开始>
// if not reset body,old bytes will put in bufPool
<原文结束>

# <翻译开始>
// 如果不重置body，旧的字节数据将会被放入bufPool中
# <翻译结束>


<原文开始>
// Status we must override Status func here,
// or the http status code returned by gin.Context.Writer.Status()
// will always be 200 in other custom gin middlewares.
<原文结束>

# <翻译开始>
// 我们必须在这里覆盖 Status 函数，
// 否则在其他自定义 gin 中间件中，gin.Context.Writer.Status() 返回的 HTTP 状态码将始终为 200。
# <翻译结束>

