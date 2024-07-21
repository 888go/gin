
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
		// StructPointerSlice []noValidationSub
		// InterfaceSlice     []testInterface
<原文结束>

# <翻译开始>
		// StructPointerSlice []noValidationSub 		// 结构体指针切片
		// InterfaceSlice     []testInterface   		// 接口类型切片
# <翻译结束>


<原文开始>
	//origin := createNoValidation_values()
	//test := createNoValidation_values()
<原文结束>

# <翻译开始>
	// origin := 创建无验证值()
	// test := 创建无验证值()
# <翻译结束>


<原文开始>
	//assert.Nil(t, validate(test))
	//assert.Nil(t, validate(&test))
<原文结束>

# <翻译开始>
	// 断言validate(test)的结果为nil
	// 断言validate(&test)的结果为nil
	// 
	// 这里是对Go语言中测试断言库（如 testify/assert）的注释翻译，这两行代码在进行单元测试时使用。它们的作用是分别检查函数`validate(test)`和`validate(&test)`的返回值是否为`nil`，如果实际结果确实是`nil`，则测试通过；否则，测试失败。
# <翻译结束>


<原文开始>
//assert.Equal(t, origin, test)
<原文结束>

# <翻译开始>
// 断言：在测试用例t中，origin（原始值）应等于test（测试值）
# <翻译结束>


<原文开始>
// structCustomValidation is a helper struct we use to check that
// custom validation can be registered on it.
// The `notone` binding directive is for custom validation and registered later.
<原文结束>

# <翻译开始>
// structCustomValidation 是一个辅助结构体，我们使用它来检查是否能够在其上注册自定义验证。
// `notone` 绑定指令用于自定义验证，并将在后续进行注册。
# <翻译结束>


<原文开始>
	// This validates that the function `notOne` matches
	// the expected function signature by `defaultValidator`
	// and by extension the validator library.
<原文结束>

# <翻译开始>
	// 这验证了函数 `notOne` 与 `defaultValidator` 预期的函数签名相匹配，
	// 从而也就验证了该函数与 validator 库的兼容性。
# <翻译结束>


<原文开始>
// Check that we can register custom validation without error
<原文结束>

# <翻译开始>
// 检查我们可以无错误地注册自定义验证
# <翻译结束>


<原文开始>
// Create an instance which will fail validation
<原文结束>

# <翻译开始>
// 创建一个在验证时会失败的实例
# <翻译结束>


<原文开始>
// Check that we got back non-nil errs
<原文结束>

# <翻译开始>
// 检查返回的 errs 是否非空
# <翻译结束>


<原文开始>
// Check that the error matches expectation
<原文结束>

# <翻译开始>
// 检查错误是否符合预期
# <翻译结束>

