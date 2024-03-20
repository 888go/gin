
<原文开始>
// Copyright 2020 Gin Core Team. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2020 Gin 核心团队。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// Content-Type MIME of the most common data formats.
<原文结束>

# <翻译开始>
// Content-Type MIME 是最常见的数据格式的 MIME 类型。
# <翻译结束>


<原文开始>
// Binding describes the interface which needs to be implemented for binding the
// data present in the request such as JSON request body, query parameters or
// the form POST.
<原文结束>

# <翻译开始>
// Binding描述了需要实现的接口，目的是为了将请求中携带的数据（如JSON请求体、查询参数或表单POST数据）进行绑定。
# <翻译结束>


<原文开始>
// BindingBody adds BindBody method to Binding. BindBody is similar with Bind,
// but it reads the body from supplied bytes instead of req.Body.
<原文结束>

# <翻译开始>
// BindingBody 为 Binding 添加了 BindBody 方法。BindBody 与 Bind 类似，
// 但是它从提供的字节中读取请求体，而不是从 req.Body 中读取。
# <翻译结束>


<原文开始>
// BindingUri adds BindUri method to Binding. BindUri is similar with Bind,
// but it reads the Params.
<原文结束>

# <翻译开始>
// BindingUri 向 Binding 结构体添加 BindUri 方法。BindUri 与 Bind 类似，
// 但它读取的是 Params 参数。
# <翻译结束>


<原文开始>
// StructValidator is the minimal interface which needs to be implemented in
// order for it to be used as the validator engine for ensuring the correctness
// of the request. Gin provides a default implementation for this using
// https://github.com/go-playground/validator/tree/v10.6.1.
<原文结束>

# <翻译开始>
// StructValidator 是一个最小接口，为了能够用作验证请求正确性的验证引擎，需要实现这个接口。
// Gin 提供了一个默认实现，使用了 https://github.com/go-playground/validator/tree/v10.6.1。
# <翻译结束>


<原文开始>
	// ValidateStruct can receive any kind of type and it should never panic, even if the configuration is not right.
	// If the received type is not a struct, any validation should be skipped and nil must be returned.
	// If the received type is a struct or pointer to a struct, the validation should be performed.
	// If the struct is not valid or the validation itself fails, a descriptive error should be returned.
	// Otherwise nil must be returned.
<原文结束>

# <翻译开始>
// ValidateStruct 可以接收任何类型的值，并且即使配置不正确，也绝不应该引发 panic。如果接收到的类型不是结构体，则应跳过所有验证并返回 nil。
// 如果接收到的类型是结构体或指向结构体的指针，则应执行验证操作。
// 若结构体无效或验证过程本身失败，则应返回一个描述性错误信息。否则必须返回 nil。
# <翻译结束>


<原文开始>
	// Engine returns the underlying validator engine which powers the
	// StructValidator implementation.
<原文结束>

# <翻译开始>
// Engine 方法返回底层驱动验证器引擎，该引擎为 StructValidator 实现提供支持。
# <翻译结束>


<原文开始>
// Validator is the default validator which implements the StructValidator
// interface. It uses https://github.com/go-playground/validator/tree/v10.6.1
// under the hood.
<原文结束>

# <翻译开始>
// Validator 是默认的验证器，实现了 StructValidator 接口。在底层，它使用了 https://github.com/go-playground/validator/tree/v10.6.1 。
# <翻译结束>


<原文开始>
// These implement the Binding interface and can be used to bind the data
// present in the request to struct instances.
<原文结束>

# <翻译开始>
// 这些实现了Binding接口，可用于将请求中呈现的数据绑定到结构体实例。
# <翻译结束>


<原文开始>
// Default returns the appropriate Binding instance based on the HTTP method
// and the content type.
<原文结束>

# <翻译开始>
// Default 根据 HTTP 方法和内容类型返回相应的 Binding 实例。
# <翻译结束>

