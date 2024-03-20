
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
// EnableDecoderUseNumber is used to call the UseNumber method on the JSON
// Decoder instance. UseNumber causes the Decoder to unmarshal a number into an
// any as a Number instead of as a float64.
<原文结束>

# <翻译开始>
// EnableDecoderUseNumber 用于调用 JSON 解码器实例上的 UseNumber 方法。启用 UseNumber 后，解码器将在反序列化数字时将其解析为 Number 类型而不是 float64 类型并存储到 any 类型变量中。
# <翻译结束>


<原文开始>
// EnableDecoderDisallowUnknownFields is used to call the DisallowUnknownFields method
// on the JSON Decoder instance. DisallowUnknownFields causes the Decoder to
// return an error when the destination is a struct and the input contains object
// keys which do not match any non-ignored, exported fields in the destination.
<原文结束>

# <翻译开始>
// EnableDecoderDisallowUnknownFields 用于调用 JSON 解码器实例上的 DisallowUnknownFields 方法。该方法启用后，当目标是一个结构体且输入包含与目标中非忽略的导出字段不匹配的对象键时，解码器会返回一个错误。
# <翻译结束>

