
<原文开始>
// Copyright 2014 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Manu Martinez-Almeida版权所有
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到
# <翻译结束>


<原文开始>
// Unit tests TODO
// func (c *Context) File(filepath string) {
// func (c *Context) Negotiate(code int, config Negotiate) {
// BAD case: func (c *Context) Render(code int, render render.Render, obj ...any) {
// test that information is not leaked when reusing Contexts (using the Pool)
<原文结束>

# <翻译开始>
// 单元测试TODO func (c *Context) File(filepath string) {func (c *Context) Negotiate(code int, config Negotiate) {BAD case: func (c *Context) Render(code int, Render Render)
// 渲染，obj…any){测试重用上下文(使用池)时信息是否泄漏
# <翻译结束>


<原文开始>
//nolint: errcheck
<原文结束>

# <翻译开始>
// nolint: errcheck
// 翻译：// 不进行errcheck检查
# <翻译结束>


<原文开始>
// TestContextSetGet tests that a parameter is set correctly on the
// current context and can be retrieved using Get.
<原文结束>

# <翻译开始>
// TestContextSetGet测试在当前上下文中是否正确设置了参数，是否可以使用Get检索参数
# <翻译结束>


<原文开始>
	// postform should not mess
<原文结束>

# <翻译开始>
// Postform不应该乱
# <翻译结束>


<原文开始>
// here c.Request == nil
<原文结束>

# <翻译开始>
// c.Request == nil 表示c的Request属性为空(nil)
# <翻译结束>


<原文开始>
// Tests that the response is serialized as JSON
// and Content-Type is set to application/json
// and special HTML characters are escaped
<原文结束>

# <翻译开始>
// 测试响应是否序列化为JSON, Content-Type是否设置为application/ JSON，是否转义了特殊的HTML字符
# <翻译结束>


<原文开始>
// Tests that the response is serialized as JSONP
// and Content-Type is set to application/javascript
<原文结束>

# <翻译开始>
// 测试响应是否序列化为JSONP，并且Content-Type设置为application/javascript
# <翻译结束>


<原文开始>
// Tests that the response is serialized as JSONP
// and Content-Type is set to application/json
<原文结束>

# <翻译开始>
// 测试响应是否序列化为JSONP，并且Content-Type设置为application/json
# <翻译结束>


<原文开始>
// Tests that no JSON is rendered if code is 204
<原文结束>

# <翻译开始>
// 测试如果code为204，是否不呈现JSON
# <翻译结束>


<原文开始>
// Tests that the response is serialized as JSON
// we change the content-type before
<原文结束>

# <翻译开始>
// 测试响应是否序列化为JSON(之前更改了内容类型)
# <翻译结束>


<原文开始>
// Tests that no Custom JSON is rendered if code is 204
<原文结束>

# <翻译开始>
// 测试如果code为204，是否不呈现自定义JSON
# <翻译结束>


<原文开始>
// Tests that the response is serialized as Secure JSON
// and Content-Type is set to application/json
<原文结束>

# <翻译开始>
// 测试响应是否序列化为安全JSON，并且Content-Type设置为application/ JSON
# <翻译结束>


<原文开始>
// Tests that the response is serialized as JSON
// and Content-Type is set to application/json
// and special HTML characters are preserved
<原文结束>

# <翻译开始>
// 测试响应是否序列化为JSON, Content-Type是否设置为application/ JSON，是否保留特殊的HTML字符
# <翻译结束>


<原文开始>
// Tests that the response executes the templates
// and responds with Content-Type set to text/html
<原文结束>

# <翻译开始>
// 测试响应是否执行模板，并将Content-Type设置为text/html
# <翻译结束>


<原文开始>
	// print debug warning log when Engine.trees > 0
<原文结束>

# <翻译开始>
// 输出调试警告日志
// 树比;0
# <翻译结束>


<原文开始>
// Tests that no HTML is rendered if code is 204
<原文结束>

# <翻译开始>
// 测试如果code为204，是否不呈现HTML
# <翻译结束>


<原文开始>
// TestContextXML tests that the response is serialized as XML
// and Content-Type is set to application/xml
<原文结束>

# <翻译开始>
// TestContextXML测试响应是否序列化为XML，并且Content-Type是否设置为application/ XML
# <翻译结束>


<原文开始>
// Tests that no XML is rendered if code is 204
<原文结束>

# <翻译开始>
// 测试如果代码为204，是否不呈现XML
# <翻译结束>


<原文开始>
// TestContextString tests that the response is returned
// with Content-Type set to text/plain
<原文结束>

# <翻译开始>
// TestContextString测试返回的响应内容类型是否设置为text/plain
# <翻译结束>


<原文开始>
// Tests that no String is rendered if code is 204
<原文结束>

# <翻译开始>
// 测试如果code为204，是否不呈现String
# <翻译结束>


<原文开始>
// TestContextString tests that the response is returned
// with Content-Type set to text/html
<原文结束>

# <翻译开始>
// TestContextString测试返回的响应是否将Content-Type设置为text/html
# <翻译结束>


<原文开始>
// Tests that no HTML String is rendered if code is 204
<原文结束>

# <翻译开始>
// 测试如果code为204，是否不呈现HTML字符串
# <翻译结束>


<原文开始>
// TestContextData tests that the response can be written from `bytestring`
// with specified MIME type
<原文结束>

# <翻译开始>
// TestContextData测试响应是否可以用指定的MIME类型从' bytestring '写入
# <翻译结束>


<原文开始>
// Tests that no Custom Data is rendered if code is 204
<原文结束>

# <翻译开始>
// 测试如果code为204，是否不呈现自定义数据
# <翻译结束>


<原文开始>
	// Content-Type='text/plain; charset=utf-8' when go version <= 1.16,
	// else, Content-Type='text/x-go; charset=utf-8'
<原文结束>

# <翻译开始>
// - type = '文本/平原;当go版本<= 1.16时，charset=utf-8'，否则，Content-Type='text/x-go;charset = utf - 8 '
# <翻译结束>


<原文开始>
// TestContextRenderYAML tests that the response is serialized as YAML
// and Content-Type is set to application/x-yaml
<原文结束>

# <翻译开始>
// TestContextRenderYAML测试响应是否被序列化为YAML，并且Content-Type设置为application/x-yaml
# <翻译结束>


<原文开始>
// TestContextRenderTOML tests that the response is serialized as TOML
// and Content-Type is set to application/toml
<原文结束>

# <翻译开始>
// TestContextRenderTOML测试响应是否序列化为TOML，并且Content-Type设置为application/ TOML
# <翻译结束>


<原文开始>
// TestContextRenderProtoBuf tests that the response is serialized as ProtoBuf
// and Content-Type is set to application/x-protobuf
// and we just use the example protobuf to check if the response is correct
<原文结束>

# <翻译开始>
// TestContextRenderProtoBuf测试响应是否被序列化为ProtoBuf，并且Content-Type设置为application/x-protobuf，我们只使用示例ProtoBuf来检查响应是否正确
# <翻译结束>


<原文开始>
// TODO
<原文结束>

# <翻译开始>
// 一切
# <翻译结束>


<原文开始>
	// Legacy tests (validating that the defaults don't break the
	// (insecure!) old behaviour)
<原文结束>

# <翻译开始>
// 遗留测试(验证默认值不会破坏(不安全!)旧行为)
# <翻译结束>


<原文开始>
	// no port
<原文结束>

# <翻译开始>
// 没有港口
# <翻译结束>


<原文开始>
	// Tests exercising the TrustedProxies functionality
<原文结束>

# <翻译开始>
// 测试执行TrustedProxies功能
# <翻译结束>


<原文开始>
	// IPv6 support
<原文结束>

# <翻译开始>
// IPv6支架
# <翻译结束>


<原文开始>
	// No trusted proxies
<原文结束>

# <翻译开始>
// 没有可信代理
# <翻译结束>


<原文开始>
	// Disabled TrustedProxies feature
<原文结束>

# <翻译开始>
// 禁用TrustedProxies特性
# <翻译结束>


<原文开始>
	// Last proxy is trusted, but the RemoteAddr is not
<原文结束>

# <翻译开始>
// 最后一个代理是可信的，但RemoteAddr不可信
# <翻译结束>


<原文开始>
	// Only trust RemoteAddr
<原文结束>

# <翻译开始>
// 只信任RemoteAddr
# <翻译结束>


<原文开始>
	// All steps are trusted
<原文结束>

# <翻译开始>
// 所有步骤都是可信的
# <翻译结束>


<原文开始>
	// Use CIDR
<原文结束>

# <翻译开始>
// 使用CIDR
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
// x - forward - for有一个非ip元素
# <翻译结束>


<原文开始>
	// Result from LookupHost has non-IP element. This should never
	// happen, but we should test it to make sure we handle it
	// gracefully.
<原文结束>

# <翻译开始>
// LookupHost的结果有非ip元素
// 这种情况不应该发生，但我们应该对其进行测试，以确保我们能够优雅地处理它
# <翻译结束>


<原文开始>
	// Use custom TrustedPlatform header
<原文结束>

# <翻译开始>
// 使用自定义TrustedPlatform头
# <翻译结束>


<原文开始>
	// wrong header
<原文结束>

# <翻译开始>
// 错误的标题
# <翻译结束>


<原文开始>
	// TrustedPlatform is empty
<原文结束>

# <翻译开始>
// TrustedPlatform为空
# <翻译结束>


<原文开始>
	// Test the legacy flag
<原文结束>

# <翻译开始>
// 测试遗留标志
# <翻译结束>


<原文开始>
// set fake content-type
<原文结束>

# <翻译开始>
// 设置虚假内容类型
# <翻译结束>


<原文开始>
		// bodyA to typeA and typeB
<原文结束>

# <翻译开始>
// 身体a到类型a和类型b
# <翻译结束>


<原文开始>
			// When it binds to typeA and typeB, it finds the body is
			// not typeB but typeA.
<原文结束>

# <翻译开始>
// 当它绑定到类型a和类型b时，它发现主体不是类型b而是类型a
# <翻译结束>


<原文开始>
		// bodyB to typeA and typeB
<原文结束>

# <翻译开始>
// 身体b到类型a和类型b
# <翻译结束>


<原文开始>
			// When it binds to typeA and typeB, it finds the body is
			// not typeA but typeB.
<原文结束>

# <翻译开始>
// 当它绑定到类型a和类型b时，它发现主体不是类型a而是类型b
# <翻译结束>


<原文开始>
	// Example request from spec: https://tools.ietf.org/html/rfc6455#section-1.2
<原文结束>

# <翻译开始>
// 来自spec: https://tools.ietf.org/html/rfc6455#section-1.2的示例请求
# <翻译结束>


<原文开始>
	// Normal request, no websocket required.
<原文结束>

# <翻译开始>
// 正常请求，不需要websocket
# <翻译结束>


<原文开始>
				// First assert must be executed after the second request
<原文结束>

# <翻译开始>
// 第一个断言必须在第二个请求之后执行
# <翻译结束>


<原文开始>
//nolint:staticcheck
<原文结束>

# <翻译开始>
// nolint: staticcheck
// （翻译）：忽略静态检查工具对本行代码的检查。
# <翻译结束>


<原文开始>
	// enable ContextWithFallback feature flag
<原文结束>

# <翻译开始>
// 启用ContextWithFallback特性标志
# <翻译结束>


<原文开始>
				// enable ContextWithFallback feature flag
<原文结束>

# <翻译开始>
// 启用ContextWithFallback特性标志
# <翻译结束>


<原文开始>
		// start async goroutine for calling srv
<原文结束>

# <翻译开始>
// 启动调用SRV的异步例程
# <翻译结束>


<原文开始>
// ensure request is done
<原文结束>

# <翻译开始>
// 确保完成请求
# <翻译结束>


<原文开始>
	// Result() has headers frozen when WriteHeaderNow() has been called
	// Compared to this time, this is when the response headers will be flushed
	// As response is flushed on c.String, the Header cannot be set by the first
	// middleware. Assert this
<原文结束>

# <翻译开始>
// 当WriteHeaderNow()被调用时，Result()已经冻结了报头
// 与此相比，这是响应报头将被刷新的时间
// 维护这
# <翻译结束>

