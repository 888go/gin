
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
		// StructPointerSlice []noValidationSub
		// InterfaceSlice     []testInterface
<原文结束>

# <翻译开始>
// StructPointerSlice 是一个指向 noValidationSub 结构体的指针切片
// interfacesslice 是一个 testInterface 接口类型的切片
// var StructPointerSlice []*noValidationSub
// var interfacesslice []testInterface
// 这里的注释翻译为：
// StructPointerSlice 表示一个包含多个指向 noValidationSub 结构体指针的切片。
// interfacesslice 表示一个包含多个 testInterface 类型值的切片。
# <翻译结束>


<原文开始>
	//origin := createNoValidation_values()
	//test := createNoValidation_values()
<原文结束>

# <翻译开始>
// origin := createNoValidation_values() // 创建并初始化不进行验证的值，将结果赋给变量origin
// test := createNoValidation_values() // 创建并初始化另一个不进行验证的值，将结果赋给变量test
// 上述代码中，`createNoValidation_values()` 是一个假设存在的函数，用于生成不需要进行验证的值。这里对两行代码分别进行了中文注释说明。
# <翻译结束>


<原文开始>
	//assert.Nil(t, validate(test))
	//assert.Nil(t, validate(&test))
<原文结束>

# <翻译开始>
// ,
// Nil(t, validate(test))断言
// 尼罗河(t,执行极为&test))
# <翻译结束>


<原文开始>
	//assert.Equal(t, origin, test)
<原文结束>

# <翻译开始>
// 断言
// 等于(t，原点，检验)
# <翻译结束>


<原文开始>
// structCustomValidation is a helper struct we use to check that
// custom validation can be registered on it.
// The `notone` binding directive is for custom validation and registered later.
<原文结束>

# <翻译开始>
// structCustomValidation是一个辅助结构体，我们使用它来检查是否可以在其上注册自定义验证
// ' notone '绑定指令用于自定义验证并在以后注册
# <翻译结束>


<原文开始>
	// This validates that the function `notOne` matches
	// the expected function signature by `defaultValidator`
	// and by extension the validator library.
<原文结束>

# <翻译开始>
// 这将验证函数' notOne '是否与' defaultValidator '和验证器库所期望的函数签名匹配
# <翻译结束>


<原文开始>
	// Check that we can register custom validation without error
<原文结束>

# <翻译开始>
// 检查我们是否可以注册自定义验证而不会出错
# <翻译结束>


<原文开始>
	// Create an instance which will fail validation
<原文结束>

# <翻译开始>
// 创建一个验证失败的实例
# <翻译结束>


<原文开始>
	// Check that we got back non-nil errs
<原文结束>

# <翻译开始>
// 检查我们是否得到非nil错误
# <翻译结束>


<原文开始>
	// Check that the error matches expectation
<原文结束>

# <翻译开始>
// 检查错误是否与预期相符
# <翻译结束>

