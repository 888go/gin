
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
// Content-Type MIME of the most common data formats.
<原文结束>

# <翻译开始>
// 内容类型MIME最常用的数据格式
# <翻译结束>


<原文开始>
// Binding describes the interface which needs to be implemented for binding the
// data present in the request such as JSON request body, query parameters or
// the form POST.
<原文结束>

# <翻译开始>
// 绑定描述了需要实现的接口，用于绑定请求中的数据，如JSON请求体、查询参数或表单POST
# <翻译结束>


<原文开始>
// BindingBody adds BindBody method to Binding. BindBody is similar with Bind,
// but it reads the body from supplied bytes instead of req.Body.
<原文结束>

# <翻译开始>
// BindingBody在Binding中添加BindBody方法
// BindBody与Bind类似，但它从提供的字节中读取主体，而不是从req.Body中读取主体
# <翻译结束>


<原文开始>
// BindingUri adds BindUri method to Binding. BindUri is similar with Bind,
// but it reads the Params.
<原文结束>

# <翻译开始>
// BindingUri将BindUri方法添加到Binding中
// BindUri与Bind类似，但它读取Params
# <翻译结束>


<原文开始>
// StructValidator is the minimal interface which needs to be implemented in
// order for it to be used as the validator engine for ensuring the correctness
// of the request. Gin provides a default implementation for this using
// https://github.com/go-playground/validator/tree/v10.6.1.
<原文结束>

# <翻译开始>
// StructValidator是需要实现的最小接口，以便将其用作确保请求正确性的验证器引擎
// Gin为此提供了一个默认实现，使用https://github.com/go-playground/validator/tree/v10.6.1
# <翻译结束>


<原文开始>
	// ValidateStruct can receive any kind of type and it should never panic, even if the configuration is not right.
	// If the received type is a slice|array, the validation should be performed travel on every element.
	// If the received type is not a struct or slice|array, any validation should be skipped and nil must be returned.
	// If the received type is a struct or pointer to a struct, the validation should be performed.
	// If the struct is not valid or the validation itself fails, a descriptive error should be returned.
	// Otherwise nil must be returned.
<原文结束>

# <翻译开始>
// ValidateStruct可以接收任何类型，即使配置不正确，它也不会panic
// 如果接收到的类型是slice数组，则应该对每个元素执行验证
// 如果接收到的类型不是struct或slice|array，则应该跳过任何验证，并且必须返回nil
// 如果接收到的类型是结构体或指向结构体的指针，则应该执行验证
// 如果结构无效或验证本身失败，则应返回描述性错误
// 否则必须返回nil
# <翻译结束>


<原文开始>
	// Engine returns the underlying validator engine which powers the
	// StructValidator implementation.
<原文结束>

# <翻译开始>
// Engine返回为StructValidator实现提供动力的底层验证器引擎
# <翻译结束>


<原文开始>
// Validator is the default validator which implements the StructValidator
// interface. It uses https://github.com/go-playground/validator/tree/v10.6.1
// under the hood.
<原文结束>

# <翻译开始>
// Validator是默认的验证器，它实现了StructValidator接口
// 它在引擎盖下使用https://github.com/go-playground/validator/tree/v10.6.1
# <翻译结束>


<原文开始>
// These implement the Binding interface and can be used to bind the data
// present in the request to struct instances.
<原文结束>

# <翻译开始>
// 它们实现了Binding接口，可用于将请求中的数据绑定到struct实例
# <翻译结束>


<原文开始>
// Default returns the appropriate Binding instance based on the HTTP method
// and the content type.
<原文结束>

# <翻译开始>
// Default根据HTTP方法和内容类型返回适当的Binding实例
# <翻译结束>


<原文开始>
// case MIMEPOSTForm:
<原文结束>

# <翻译开始>
// 案例MIMEPOSTForm:
# <翻译结束>

