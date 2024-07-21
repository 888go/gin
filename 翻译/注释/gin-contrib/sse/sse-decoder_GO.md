
<原文开始>
// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2014 Manu Martinez-Almeida。保留所有权利。
// 本源代码的使用受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
//If the data buffer's last character is a U+000A LINE FEED (LF) character, then remove the last character from the data buffer.
<原文结束>

# <翻译开始>
// 如果数据缓冲区的最后一个字符是U+000A换行符（LF），则从数据缓冲区中移除最后一个字符。
# <翻译结束>


<原文开始>
	// TODO (and unit tests)
	// Lines must be separated by either a U+000D CARRIAGE RETURN U+000A LINE FEED (CRLF) character pair,
	// a single U+000A LINE FEED (LF) character,
	// or a single U+000D CARRIAGE RETURN (CR) character.
<原文结束>

# <翻译开始>
	// TODO (并编写单元测试)
	// 行与行之间必须由以下任意一种字符对或单个字符分隔：
	// U+000D 回车（CARRIAGE RETURN）和 U+000A 换行（LINE FEED，即 CRLF 组合）
	// 或者
	// 单个 U+000A 换行（LINE FEED，LF）字符
	// 或者
	// 单个 U+000D 回车（CARRIAGE RETURN，CR）字符。
# <翻译结束>


<原文开始>
// If the line is empty (a blank line). Dispatch the event.
<原文结束>

# <翻译开始>
// 如果该行为空（即空白行），则分派该事件。
# <翻译结束>


<原文开始>
// reset current event and data buffer
<原文结束>

# <翻译开始>
// 重置当前事件和数据缓冲区
# <翻译结束>


<原文开始>
// If the line starts with a U+003A COLON character (:), ignore the line.
<原文结束>

# <翻译开始>
// 如果行以 U+003A 字符（冒号）开始，则忽略该行。
# <翻译结束>


<原文开始>
			// If the line contains a U+003A COLON character character (:)
			// Collect the characters on the line before the first U+003A COLON character (:),
			// and let field be that string.
<原文结束>

# <翻译开始>
			// 如果该行包含一个 U+003A 字符（冒号）：
			// 收集该行第一个 U+003A 字符（冒号）之前的所有字符，
			// 并将这些字符组成的字符串赋值给 field。
# <翻译结束>


<原文开始>
			// Collect the characters on the line after the first U+003A COLON character (:),
			// and let value be that string.
<原文结束>

# <翻译开始>
			// 收集行中第一个 U+003A COLON 字符 (:) 之后的字符，并将该字符串赋值给 value。
# <翻译结束>


<原文开始>
// If value starts with a single U+0020 SPACE character, remove it from value.
<原文结束>

# <翻译开始>
// 如果value以单个U+0020（空格）字符开头，则从value中移除它。
# <翻译结束>


<原文开始>
			// Otherwise, the string is not empty but does not contain a U+003A COLON character character (:)
			// Use the whole line as the field name, and the empty string as the field value.
<原文结束>

# <翻译开始>
			// 否则，字符串不为空，但不包含 U+003A 字符（冒号）：
			// 将整行用作字段名，并使用空字符串作为字段值。
# <翻译结束>


<原文开始>
		// The steps to process the field given a field name and a field value depend on the field name,
		// as given in the following list. Field names must be compared literally,
		// with no case folding performed.
<原文结束>

# <翻译开始>
		// 根据给定的字段名和字段值来处理该字段的步骤取决于字段名，具体如下所示。字段名必须进行逐字比较，且不应进行大小写折叠处理。
# <翻译结束>


<原文开始>
// Set the event name buffer to field value.
<原文结束>

# <翻译开始>
// 将事件名称缓冲区设置为字段值。
# <翻译结束>


<原文开始>
// Set the event stream's last event ID to the field value.
<原文结束>

# <翻译开始>
// 将事件流的最后一个事件ID设置为字段值。
# <翻译结束>


<原文开始>
			// If the field value consists of only characters in the range U+0030 DIGIT ZERO (0) to U+0039 DIGIT NINE (9),
			// then interpret the field value as an integer in base ten, and set the event stream's reconnection time to that integer.
			// Otherwise, ignore the field.
<原文结束>

# <翻译开始>
			// 如果字段值仅包含从 U+0030（数字零，0）到 U+0039（数字九，9）范围内的字符，则将该字段值解释为十进制整数，并将事件流的重连时间设置为该整数。否则，忽略该字段。
# <翻译结束>


<原文开始>
// Append the field value to the data buffer,
<原文结束>

# <翻译开始>
// 将字段值追加到数据缓冲区，
// 
// ensuring there's a space separator if the buffer already has content.
// 确保如果缓冲区已有内容，则在追加前添加一个空格分隔符。
# <翻译结束>


<原文开始>
// then append a single U+000A LINE FEED (LF) character to the data buffer.
<原文结束>

# <翻译开始>
// 然后向数据缓冲区追加一个 U+000A 换行符（LF）字符。
# <翻译结束>


<原文开始>
//Otherwise. The field is ignored.
<原文结束>

# <翻译开始>
// 否则。该字段将被忽略。
# <翻译结束>


<原文开始>
// Once the end of the file is reached, the user agent must dispatch the event one final time.
<原文结束>

# <翻译开始>
// 当到达文件末尾时，用户代理必须最后一次派发该事件。
# <翻译结束>

