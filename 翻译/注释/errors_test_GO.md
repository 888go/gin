
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
//nolint: errcheck
<原文结束>

# <翻译开始>
// nolint: errcheck
// 翻译：// 不进行errcheck检查
# <翻译结束>


<原文开始>
// TestErrorUnwrap tests the behavior of gin.Error with "errors.Is()" and "errors.As()".
// "errors.Is()" and "errors.As()" have been added to the standard library in go 1.13.
<原文结束>

# <翻译开始>
// testrorunwrap测试gin的行为
// Error . is ()"“和“误差()
// “errors.Is()“;和“误差()“;已经被添加到go 1.13的标准库中
# <翻译结束>


<原文开始>
	// 2 layers of wrapping : use 'fmt.Errorf("%w")' to wrap a gin.Error{}, which itself wraps innerErr
<原文结束>

# <翻译开始>
// 2层包装:使用'fmt. error ("%w")'来包装杜松子酒
// Error{}，它本身包装了innerErr
# <翻译结束>


<原文开始>
	// check that 'errors.Is()' and 'errors.As()' behave as expected :
<原文结束>

# <翻译开始>
// 检查'errors.Is()'和'errors.As()'的行为是否符合预期:
# <翻译结束>

