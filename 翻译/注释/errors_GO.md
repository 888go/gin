
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
// ErrorType is an unsigned 64-bit error code as defined in the gin spec.
<原文结束>

# <翻译开始>
// ErrorType 是一个无符号的64位错误代码，遵循gin规范定义。
# <翻译结束>


<原文开始>
// ErrorTypeBind is used when Context.Bind() fails.
<原文结束>

# <翻译开始>
// ErrorTypeBind 用于当 Context.Bind() 失败时。
# <翻译结束>


<原文开始>
// ErrorTypeRender is used when Context.Render() fails.
<原文结束>

# <翻译开始>
// ErrorTypeRender 用于当 Context.Render() 失败时。
# <翻译结束>


<原文开始>
// ErrorTypePrivate indicates a private error.
<原文结束>

# <翻译开始>
// ErrorTypePrivate 表示一个私有错误。
# <翻译结束>


<原文开始>
// ErrorTypePublic indicates a public error.
<原文结束>

# <翻译开始>
// ErrorTypePublic 表示一个公开的错误。
# <翻译结束>


<原文开始>
// ErrorTypeAny indicates any other error.
<原文结束>

# <翻译开始>
// ErrorTypeAny 表示任何其他错误。
# <翻译结束>


<原文开始>
// ErrorTypeNu indicates any other error.
<原文结束>

# <翻译开始>
// ErrorTypeNu 表示任何其他错误。
# <翻译结束>


<原文开始>
// Error represents a error's specification.
<原文结束>

# <翻译开始>
// Error代表了一个错误的规格说明。
# <翻译结束>


<原文开始>
// SetType sets the error's type.
<原文结束>

# <翻译开始>
// SetType 设置错误的类型。
# <翻译结束>


<原文开始>
// SetMeta sets the error's meta data.
<原文结束>

# <翻译开始>
// SetMeta 设置错误的元数据。
# <翻译结束>


<原文开始>
// JSON creates a properly formatted JSON
<原文结束>

# <翻译开始>
// JSON 创建一个格式正确的 JSON
# <翻译结束>


<原文开始>
// MarshalJSON implements the json.Marshaller interface.
<原文结束>

# <翻译开始>
// MarshalJSON 实现了 json.Marshaller 接口。
# <翻译结束>


<原文开始>
// Error implements the error interface.
<原文结束>

# <翻译开始>
// Error 实现了 error 接口。
# <翻译结束>


<原文开始>
// IsType judges one error.
<原文结束>

# <翻译开始>
// IsType 判断一个错误。
# <翻译结束>


<原文开始>
// Unwrap returns the wrapped error, to allow interoperability with errors.Is(), errors.As() and errors.Unwrap()
<原文结束>

# <翻译开始>
// Unwrap 返回封装的错误，以便与 errors.Is()、errors.As() 和 errors.Unwrap() 之间进行互操作性
# <翻译结束>


<原文开始>
// ByType returns a readonly copy filtered the byte.
// ie ByType(gin.ErrorTypePublic) returns a slice of errors with type=ErrorTypePublic.
<原文结束>

# <翻译开始>
// ByType 返回一个只读副本，其中包含了经过过滤的错误信息。具体来说，ByType(gin.ErrorTypePublic) 将返回一个类型为 ErrorTypePublic 的错误信息切片。
# <翻译结束>


<原文开始>
// Last returns the last error in the slice. It returns nil if the array is empty.
// Shortcut for errors[len(errors)-1].
<原文结束>

# <翻译开始>
// Last 函数返回切片中的最后一个错误。如果该数组为空，则返回 nil。
// 这是 errors[len(errors)-1] 的快捷方式。
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
// Errors 返回包含所有错误消息的数组。
// 示例：
//
//	c.Error(errors.New("第一个错误"))
//	c.Error(errors.New("第二个错误"))
//	c.Error(errors.New("第三个错误"))
//	c.Errors.Errors() // == []string{"第一个", "第二个", "第三个"}
# <翻译结束>

