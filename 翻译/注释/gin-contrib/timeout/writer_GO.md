
<原文开始>
// Writer is a writer with memory buffer
<原文结束>

# <翻译开始>
// Writer is a writer with memory buffer
# <翻译结束>


<原文开始>
// NewWriter will return a timeout.Writer pointer
<原文结束>

# <翻译开始>
// NewWriter will return a timeout.Writer pointer
# <翻译结束>


<原文开始>
// Write will write data to response body
<原文结束>

# <翻译开始>
// Write will write data to response body
# <翻译结束>


<原文开始>
// WriteHeader sends an HTTP response header with the provided status code.
// If the response writer has already written headers or if a timeout has occurred,
// this method does nothing.
<原文结束>

# <翻译开始>
// WriteHeader sends an HTTP response header with the provided status code.
// If the response writer has already written headers or if a timeout has occurred,
// this method does nothing.
# <翻译结束>


<原文开始>
	// gin is using -1 to skip writing the status code
	// see https://github.com/gin-gonic/gin/blob/a0acf1df2814fcd828cb2d7128f2f4e2136d3fac/response_writer.go#L61
<原文结束>

# <翻译开始>
	// gin is using -1 to skip writing the status code
	// see https://github.com/gin-gonic/gin/blob/a0acf1df2814fcd828cb2d7128f2f4e2136d3fac/response_writer.go#L61
# <翻译结束>


<原文开始>
// Header will get response headers
<原文结束>

# <翻译开始>
// Header will get response headers
# <翻译结束>


<原文开始>
// WriteString will write string to response body
<原文结束>

# <翻译开始>
// WriteString will write string to response body
# <翻译结束>


<原文开始>
// FreeBuffer will release buffer pointer
<原文结束>

# <翻译开始>
// FreeBuffer will release buffer pointer
# <翻译结束>


<原文开始>
	// if not reset body,old bytes will put in bufPool
<原文结束>

# <翻译开始>
	// if not reset body,old bytes will put in bufPool
# <翻译结束>


<原文开始>
// Status we must override Status func here,
// or the http status code returned by gin.Context.Writer.Status()
// will always be 200 in other custom gin middlewares.
<原文结束>

# <翻译开始>
// Status we must override Status func here,
// or the http status code returned by gin.Context.Writer.Status()
// will always be 200 in other custom gin middlewares.
# <翻译结束>

