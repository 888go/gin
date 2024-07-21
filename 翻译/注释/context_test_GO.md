
<原文开始>
// Copyright 2014 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2014 Manu Martinez-Almeida。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// Unit tests TODO
// func (c *Context) File(filepath string) {
// func (c *Context) Negotiate(code int, config Negotiate) {
// BAD case: func (c *Context) Render(code int, render render.Render, obj ...any) {
// test that information is not leaked when reusing Contexts (using the Pool)
<原文结束>

# <翻译开始>
// 单元测试 TODO
// func (c *Context) File(filepath string) { //（待办事项：编写此函数的单元测试）
// func (c *Context) Negotiate(code int, config Negotiate) { //（待办事项：编写此函数的单元测试）
// 不良案例：func (c *Context) Render(code int, render render.Render, obj ...any) { //（这个函数设计可能存在问题）
// 测试在重用 Contexts（利用 Pool）时，确保不会泄露信息
# <翻译结束>


<原文开始>
// TestContextSetGet tests that a parameter is set correctly on the
// current context and can be retrieved using Get.
<原文结束>

# <翻译开始>
// TestContextSetGet 测试当前上下文中参数设置正确，
// 并且可以使用 Get 方法成功获取。
# <翻译结束>


<原文开始>
// postform should not mess
<原文结束>

# <翻译开始>
// postform 不应弄乱
# <翻译结束>


<原文开始>
// Tests that the response is serialized as JSON
// and Content-Type is set to application/json
// and special HTML characters are escaped
<原文结束>

# <翻译开始>
// 测试响应是否已序列化为JSON格式
// 并且Content-Type设置为application/json
// 特殊HTML字符已转义
# <翻译结束>


<原文开始>
// Tests that the response is serialized as JSONP
// and Content-Type is set to application/javascript
<原文结束>

# <翻译开始>
// 测试响应是否被序列化为 JSONP
// 并且 Content-Type 被设置为 application/javascript
# <翻译结束>


<原文开始>
// Tests that the response is serialized as JSONP
// and Content-Type is set to application/json
<原文结束>

# <翻译开始>
// 测试响应是否被序列化为JSONP格式
// 并且Content-Type头被设置为application/json
# <翻译结束>


<原文开始>
// Tests that no JSON is rendered if code is 204
<原文结束>

# <翻译开始>
// 测试当状态码为204时，不渲染任何JSON内容
# <翻译结束>


<原文开始>
// Tests that the response is serialized as JSON
// we change the content-type before
<原文结束>

# <翻译开始>
// 测试响应是否被序列化为JSON
// 我们在之前更改了content-type
# <翻译结束>


<原文开始>
// Tests that no Custom JSON is rendered if code is 204
<原文结束>

# <翻译开始>
// 测试当状态码为204时，不渲染任何自定义JSON
# <翻译结束>


<原文开始>
// Tests that the response is serialized as Secure JSON
// and Content-Type is set to application/json
<原文结束>

# <翻译开始>
// 测试响应被序列化为 Secure JSON
// 并且 Content-Type 被设置为 application/json
# <翻译结束>


<原文开始>
// Tests that the response is serialized as JSON
// and Content-Type is set to application/json
// and special HTML characters are preserved
<原文结束>

# <翻译开始>
// 测试响应是否被序列化为JSON格式
// 并且Content-Type设置为application/json
// 同时保留特殊的HTML字符
# <翻译结束>


<原文开始>
// Tests that the response executes the templates
// and responds with Content-Type set to text/html
<原文结束>

# <翻译开始>
// 测试响应会执行模板并以 text/html 设置 Content-Type 进行响应
# <翻译结束>


<原文开始>
// print debug warning log when Engine.trees > 0
<原文结束>

# <翻译开始>
// 当 Engine.trees > 0 时，打印调试警告日志
# <翻译结束>


<原文开始>
// Tests that no HTML is rendered if code is 204
<原文结束>

# <翻译开始>
// 测试当代码为204时不会渲染任何HTML内容
# <翻译结束>


<原文开始>
// TestContextXML tests that the response is serialized as XML
// and Content-Type is set to application/xml
<原文结束>

# <翻译开始>
// TestContextXML测试响应被序列化为XML格式
// 并且Content-Type设置为application/xml
# <翻译结束>


<原文开始>
// Tests that no XML is rendered if code is 204
<原文结束>

# <翻译开始>
// 测试当代码为204时，不会输出任何XML内容
# <翻译结束>


<原文开始>
// TestContextString tests that the response is returned
// with Content-Type set to text/plain
<原文结束>

# <翻译开始>
// TestContextString 测试响应返回时
// 将 Content-Type 设置为 text/plain
# <翻译结束>


<原文开始>
// Tests that no String is rendered if code is 204
<原文结束>

# <翻译开始>
// 测试当状态码为204时，不渲染任何字符串
# <翻译结束>


<原文开始>
// TestContextString tests that the response is returned
// with Content-Type set to text/html
<原文结束>

# <翻译开始>
// TestContextString 测试响应返回时
// 其Content-Type被设置为text/html
# <翻译结束>


<原文开始>
// Tests that no HTML String is rendered if code is 204
<原文结束>

# <翻译开始>
// 测试当状态码为204时，不渲染任何HTML字符串
# <翻译结束>


<原文开始>
// TestContextData tests that the response can be written from `bytestring`
// with specified MIME type
<原文结束>

# <翻译开始>
// TestContextData 测试响应能够通过 `bytestring` 指定的 MIME 类型写入
# <翻译结束>


<原文开始>
// Tests that no Custom Data is rendered if code is 204
<原文结束>

# <翻译开始>
// 测试当代码为204时，不会渲染任何自定义数据
# <翻译结束>


<原文开始>
	// Content-Type='text/plain; charset=utf-8' when go version <= 1.16,
	// else, Content-Type='text/x-go; charset=utf-8'
<原文结束>

# <翻译开始>
	// 当Go版本小于等于1.16时，Content-Type='text/plain; charset=utf-8'，
	// 否则，Content-Type='text/x-go; charset=utf-8'
# <翻译结束>


<原文开始>
// TestContextRenderYAML tests that the response is serialized as YAML
// and Content-Type is set to application/x-yaml
<原文结束>

# <翻译开始>
// TestContextRenderYAML 测试响应被序列化为 YAML 格式
// 并且 Content-Type 被设置为 application/x-yaml
# <翻译结束>


<原文开始>
// TestContextRenderTOML tests that the response is serialized as TOML
// and Content-Type is set to application/toml
<原文结束>

# <翻译开始>
// TestContextRenderTOML 测试响应是否已序列化为 TOML 格式
// 并且 Content-Type 已设置为 application/toml
# <翻译结束>


<原文开始>
// TestContextRenderProtoBuf tests that the response is serialized as ProtoBuf
// and Content-Type is set to application/x-protobuf
// and we just use the example protobuf to check if the response is correct
<原文结束>

# <翻译开始>
// TestContextRenderProtoBuf 测试响应是否被序列化为 ProtoBuf
// 并且 Content-Type 被设置为 application/x-protobuf
// 我们仅使用示例 protobuf 来检查响应是否正确
# <翻译结束>


<原文开始>
	// Legacy tests (validating that the defaults don't break the
	// (insecure!) old behaviour)
<原文结束>

# <翻译开始>
	// 向后兼容测试（验证默认设置不会破坏
	// （不安全！）的旧版行为）
# <翻译结束>


<原文开始>
// Tests exercising the TrustedProxies functionality
<原文结束>

# <翻译开始>
// 测试检验 TrustedProxies 功能的实现
# <翻译结束>


<原文开始>
// Disabled TrustedProxies feature
<原文结束>

# <翻译开始>
// 禁用 TrustedProxies 功能
# <翻译结束>


<原文开始>
// Last proxy is trusted, but the RemoteAddr is not
<原文结束>

# <翻译开始>
// 最后的代理是可信的，但 RemoteAddr 不是
# <翻译结束>


<原文开始>
// Use hostname that resolves to all the proxies
<原文结束>

# <翻译开始>
// 使用解析到所有代理的主机名
# <翻译结束>


<原文开始>
// Use hostname that returns an error
<原文结束>

# <翻译开始>
// 使用返回错误的主机名
# <翻译结束>


<原文开始>
// X-Forwarded-For has a non-IP element
<原文结束>

# <翻译开始>
// X-Forwarded-For 包含非 IP 元素
# <翻译结束>


<原文开始>
	// Result from LookupHost has non-IP element. This should never
	// happen, but we should test it to make sure we handle it
	// gracefully.
<原文结束>

# <翻译开始>
	// LookupHost 返回的结果包含非 IP 元素。这种情况本不应该发生，但我们应当对其进行测试以确保我们能够优雅地处理此类异常情况。
# <翻译结束>


<原文开始>
// Use custom TrustedPlatform header
<原文结束>

# <翻译开始>
// 使用自定义 TrustedPlatform 头部
# <翻译结束>


<原文开始>
// TrustedPlatform is empty
<原文结束>

# <翻译开始>
// TrustedPlatform为空
# <翻译结束>


<原文开始>
// bodyA to typeA and typeB
<原文结束>

# <翻译开始>
// 将bodyA转换为typeA和typeB
# <翻译结束>


<原文开始>
			// When it binds to typeA and typeB, it finds the body is
			// not typeB but typeA.
<原文结束>

# <翻译开始>
			// 当它绑定到 typeA 和 typeB 时，它发现主体不是 typeB，而是 typeA。
# <翻译结束>


<原文开始>
// bodyB to typeA and typeB
<原文结束>

# <翻译开始>
// 将bodyB转换为typeA和typeB
# <翻译结束>


<原文开始>
			// When it binds to typeA and typeB, it finds the body is
			// not typeA but typeB.
<原文结束>

# <翻译开始>
			// 当它绑定到 typeA 和 typeB 时，会发现其实体不是 typeA，而是 typeB。
# <翻译结束>


<原文开始>
// Example request from spec: https://tools.ietf.org/html/rfc6455#section-1.2
<原文结束>

# <翻译开始>
// 示例请求来自规范：https://tools.ietf.org/html/rfc6455#section-1.2
# <翻译结束>


<原文开始>
// Normal request, no websocket required.
<原文结束>

# <翻译开始>
// 正常请求，无需使用websocket。
# <翻译结束>


<原文开始>
// First assert must be executed after the second request
<原文结束>

# <翻译开始>
// 第一个断言必须在第二个请求执行后执行
# <翻译结束>


<原文开始>
// enable ContextWithFallback feature flag
<原文结束>

# <翻译开始>
// 启用 ContextWithFallback 功能标志
# <翻译结束>


<原文开始>
// start async goroutine for calling srv
<原文结束>

# <翻译开始>
// 启动异步goroutine以调用srv
# <翻译结束>


<原文开始>
	// Result() has headers frozen when WriteHeaderNow() has been called
	// Compared to this time, this is when the response headers will be flushed
	// As response is flushed on c.String, the Header cannot be set by the first
	// middleware. Assert this
<原文结束>

# <翻译开始>
	// 当 WriteHeaderNow() 被调用时，Result() 会冻结头部信息
	// 相对于此时，这是响应头将被刷新的时间点
	// 由于在 c.String 上进行响应刷新，因此第一个中间件无法设置 Header。请确认这一点
# <翻译结束>

