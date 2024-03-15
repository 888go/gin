
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
// ErrorType is an unsigned 64-bit error code as defined in the gin spec.
<原文结束>

# <翻译开始>
// ErrorType是在gin规范中定义的无符号64位错误代码
# <翻译结束>


<原文开始>
	// ErrorTypeBind is used when Context.Bind() fails.
<原文结束>

# <翻译开始>
// 当Context.Bind()失败时使用ErrorTypeBind
# <翻译结束>


<原文开始>
	// ErrorTypeRender is used when Context.Render() fails.
<原文结束>

# <翻译开始>
// 当Context.Render()失败时使用ErrorTypeRender
# <翻译结束>


<原文开始>
	// ErrorTypePrivate indicates a private error.
<原文结束>

# <翻译开始>
// ErrorTypePrivate私有错误
# <翻译结束>


<原文开始>
	// ErrorTypePublic indicates a public error.
<原文结束>

# <翻译开始>
// ErrorTypePublic表示公共错误
# <翻译结束>


<原文开始>
	// ErrorTypeAny indicates any other error.
<原文结束>

# <翻译开始>
// ErrorTypeAny表示任何其他错误
# <翻译结束>


<原文开始>
	// ErrorTypeNu indicates any other error.
<原文结束>

# <翻译开始>
// ErrorTypeNu表示任何其他错误
# <翻译结束>


<原文开始>
// Error represents a error's specification.
<原文结束>

# <翻译开始>
// Error表示错误的说明
# <翻译结束>


<原文开始>
// SetType sets the error's type.
<原文结束>

# <翻译开始>
// SetType设置错误的类型
# <翻译结束>


<原文开始>
// SetMeta sets the error's meta data.
<原文结束>

# <翻译开始>
// SetMeta设置错误的元数据
# <翻译结束>


<原文开始>
// JSON creates a properly formatted JSON
<原文结束>

# <翻译开始>
// JSON创建一个格式正确的JSON
# <翻译结束>


<原文开始>
// MarshalJSON implements the json.Marshaller interface.
<原文结束>

# <翻译开始>
// MarshalJSON实现json
// Marshaller接口
# <翻译结束>


<原文开始>
// Error implements the error interface.
<原文结束>

# <翻译开始>
// Error实现错误接口
# <翻译结束>


<原文开始>
// IsType judges one error.
<原文结束>

# <翻译开始>
// IsType判断一个错误
# <翻译结束>


<原文开始>
// Unwrap returns the wrapped error, to allow interoperability with errors.Is(), errors.As() and errors.Unwrap()
<原文结束>

# <翻译开始>
// Unwrap返回包装后的错误，以允许与errors.Is()、errors.As()和errors.Unwrap()互操作
# <翻译结束>


<原文开始>
// ByType returns a readonly copy filtered the byte.
// ie ByType(gin.ErrorTypePublic) returns a slice of errors with type=ErrorTypePublic.
<原文结束>

# <翻译开始>
// ByType返回经过字节过滤的只读副本
// 即ByType(gin.ErrorTypePublic)返回一个类型=ErrorTypePublic的错误切片
# <翻译结束>


<原文开始>
// Last returns the last error in the slice. It returns nil if the array is empty.
// Shortcut for errors[len(errors)-1].
<原文结束>

# <翻译开始>
// Last返回切片中的最后一个错误
// 如果数组为空，则返回nil
// 错误的快捷方式[len(errors)-1]
# <翻译结束>


<原文开始>
// Errors returns an array with all the error messages.
// Example:
//
//	c.Error(errors.New("first"))
//	c.Error(errors.New("second"))
//	c.Error(errors.New("third"))
//	c.Errors.Errors() // == []string{"first", "second", "third"}
<原文结束>

# <翻译开始>
// Errors返回一个包含所有错误消息的数组
// 示例:c.Error(errors.New("first")) c.Error(errors.New("second")) c. errors (errors.New("third")) c. errors () == []string{"first"， "second"， "third"}
# <翻译结束>

