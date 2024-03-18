
<原文开始>
		// The underlying io.Reader may not return (0, io.EOF)
		// at EOF if the requested size is 0, so read 1 byte
		// instead. The io.Reader docs are a bit ambiguous
		// about the return value of Read when 0 bytes are
		// requested, and {bytes,strings}.Reader gets it wrong
		// too (it returns (0, nil) even at EOF).
<原文结束>

# <翻译开始>
// 当请求的大小为0时，底层的io.Reader在遇到EOF时可能不会返回（0, io.EOF），因此改为读取1个字节。关于在请求0字节时Read方法的返回值，io.Reader的文档有些模糊不清，而且{bytes,strings}.Reader也处理得不正确（即使在EOF时，它也会返回(0, nil)）。
# <翻译结束>


<原文开始>
		// If we had zero bytes to read remaining (but hadn't seen EOF)
		// and we get a byte here, that means we went over our limit.
<原文结束>

# <翻译开始>
// 如果我们之前剩余0字节可读（但尚未遇到EOF）
// 然后在这里获取到一个字节，这意味着我们超过了限制。
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
// RequestSizeLimiter 返回一个中间件，用于限制请求的大小
// 当请求超过限制时，将会发生以下情况：
// * 将错误添加到上下文中
// * 设置 "Connection: close" 头部信息
// * 向客户端发送 413 错误（http.StatusRequestEntityTooLarge，表示请求实体过大）
// * 中断当前上下文
# <翻译结束>


<原文开始>
// ff:
// err:
// n:
// p:
<原文结束>

# <翻译开始>
// ff:
// err:
// n:
// p:
# <翻译结束>

