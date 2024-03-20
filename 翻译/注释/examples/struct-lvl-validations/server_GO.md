
<原文开始>
// User contains user information.
<原文结束>

# <翻译开始>
// User 包含用户信息。
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
// UserStructLevelValidation 包含一些自定义的结构级别验证，这些验证在字段级别上并不总是适用。例如，此函数验证 FirstName 或 LastName 至少有一个存在；虽然也可以通过自定义字段验证来实现，但那样就需要在两个字段上都添加该验证逻辑，导致代码重复和额外开销，而这种方式只需验证一次。
// 
// 注意：你可能会问为什么不直接在 validator 之外进行这种验证。采用这种方式将验证过程直接融入到 validator 中，可以与验证标签结合使用，并且仍然保持统一的错误输出格式。
# <翻译结束>


<原文开始>
// user := structLevel.CurrentStruct.Interface().(User)
<原文结束>

# <翻译开始>
// 获取当前结构体的接口表示，并将其转换为 User 类型，赋值给 user 变量
# <翻译结束>


<原文开始>
// plus can to more, even with different tag than "fnameorlname"
<原文结束>

# <翻译开始>
// plus 可以做更多事情，即使标签不同于 "fnameorlname"
# <翻译结束>

