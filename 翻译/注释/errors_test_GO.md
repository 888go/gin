
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
// TestErrorUnwrap tests the behavior of gin.Error with "errors.Is()" and "errors.As()".
// "errors.Is()" and "errors.As()" have been added to the standard library in go 1.13.
<原文结束>

# <翻译开始>
// TestErrorUnwrap 测试 gin.Error 与 "errors.Is()" 和 "errors.As()" 的交互行为。
// "errors.Is()" 和 "errors.As()" 在 Go 1.13 版本中被添加到标准库中。
# <翻译结束>


<原文开始>
// 2 layers of wrapping : use 'fmt.Errorf("%w")' to wrap a gin.Error{}, which itself wraps innerErr
<原文结束>

# <翻译开始>
// 两层包装：使用 'fmt.Errorf("%w")' 来包装一个 gin.Error{}，该 gin.Error{} 自身又封装了内部错误 innerErr。
# <翻译结束>


<原文开始>
// check that 'errors.Is()' and 'errors.As()' behave as expected :
<原文结束>

# <翻译开始>
// 检查 'errors.Is()' 和 'errors.As()' 是否按预期工作：
# <翻译结束>

