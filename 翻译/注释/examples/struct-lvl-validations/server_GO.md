
<原文开始>
// User contains user information.
<原文结束>

# <翻译开始>
// User包含用户信息
# <翻译结束>


<原文开始>
// UserStructLevelValidation contains custom struct level validations that don't always
// make sense at the field validation level. For example, this function validates that either
// FirstName or LastName exist; could have done that with a custom field validation but then
// would have had to add it to both fields duplicating the logic + overhead, this way it's
// only validated once.
//
// NOTE: you may ask why wouldn't not just do this outside of validator. Doing this way
// hooks right into validator and you can combine with validation tags and still have a
// common error output format.
<原文结束>

# <翻译开始>
// UserStructLevelValidation包含自定义结构级验证，这些验证在字段验证级别上并不总是有意义的
// 例如，这个函数验证FirstName或LastName是否存在;本可以使用自定义字段验证来完成此操作，但随后必须将其添加到复制逻辑+开销的两个字段中，这样只验证一次
// 注意:你可能会问为什么不在验证器之外做这个
// 这样做可以直接与验证器挂钩，并且可以与验证标记结合使用，并且仍然具有常见的错误输出格式
# <翻译结束>


<原文开始>
	// user := structLevel.CurrentStruct.Interface().(User)
<原文结束>

# <翻译开始>
// 获取当前结构体的接口表示，然后强制类型转换为 user 类型，并赋值给变量 user
// ```go
// user := structLevel.CurrentStruct.Interface().(user)
// 这里的代码是 Go 语言中的类型断言，将 interface 类型转换为已知的具体类型（这里是 user 类型）。`structLevel.CurrentStruct.Interface()` 表示获取一个包含当前结构体实例的 interface 值，后面的 `.(`user`)` 是用于断言这个 interface 实例实际上是 user 类型。
# <翻译结束>


<原文开始>
	// plus can to more, even with different tag than "fnameorlname"
<原文结束>

# <翻译开始>
// Plus可以添加更多，即使标签与“fnameorlname”不同
# <翻译结束>

