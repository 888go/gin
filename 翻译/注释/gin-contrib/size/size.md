
<原文开始>
		// The underlying io.Reader may not return (0, io.EOF)
		// at EOF if the requested size is 0, so read 1 byte
		// instead. The io.Reader docs are a bit ambiguous
		// about the return value of Read when 0 bytes are
		// requested, and {bytes,strings}.Reader gets it wrong
		// too (it returns (0, nil) even at EOF).
<原文结束>

# <翻译开始>
		// The underlying io.Reader may not return (0, io.EOF)
		// at EOF if the requested size is 0, so read 1 byte
		// instead. The io.Reader docs are a bit ambiguous
		// about the return value of Read when 0 bytes are
		// requested, and {bytes,strings}.Reader gets it wrong
		// too (it returns (0, nil) even at EOF).
# <翻译结束>


<原文开始>
		// If we had zero bytes to read remaining (but hadn't seen EOF)
		// and we get a byte here, that means we went over our limit.
<原文结束>

# <翻译开始>
		// If we had zero bytes to read remaining (but hadn't seen EOF)
		// and we get a byte here, that means we went over our limit.
# <翻译结束>


<原文开始>
// RequestSizeLimiter returns a middleware that limits the size of request
// When a request is over the limit, the following will happen:
// * Error will be added to the context
// * Connection: close header will be set
// * Error 413 will be sent to the client (http.StatusRequestEntityTooLarge)
// * Current context will be aborted
<原文结束>

# <翻译开始>
// RequestSizeLimiter returns a middleware that limits the size of request
// When a request is over the limit, the following will happen:
// * Error will be added to the context
// * Connection: close header will be set
// * Error 413 will be sent to the client (http.StatusRequestEntityTooLarge)
// * Current context will be aborted
# <翻译结束>

