
<原文开始>
// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
		//If the data buffer's last character is a U+000A LINE FEED (LF) character, then remove the last character from the data buffer.
<原文结束>

# <翻译开始>
		//If the data buffer's last character is a U+000A LINE FEED (LF) character, then remove the last character from the data buffer.
# <翻译结束>


<原文开始>
	// TODO (and unit tests)
	// Lines must be separated by either a U+000D CARRIAGE RETURN U+000A LINE FEED (CRLF) character pair,
	// a single U+000A LINE FEED (LF) character,
	// or a single U+000D CARRIAGE RETURN (CR) character.
<原文结束>

# <翻译开始>
	// TODO (and unit tests)
	// Lines must be separated by either a U+000D CARRIAGE RETURN U+000A LINE FEED (CRLF) character pair,
	// a single U+000A LINE FEED (LF) character,
	// or a single U+000D CARRIAGE RETURN (CR) character.
# <翻译结束>


<原文开始>
			// If the line is empty (a blank line). Dispatch the event.
<原文结束>

# <翻译开始>
			// If the line is empty (a blank line). Dispatch the event.
# <翻译结束>


<原文开始>
			// reset current event and data buffer
<原文结束>

# <翻译开始>
			// reset current event and data buffer
# <翻译结束>


<原文开始>
			// If the line starts with a U+003A COLON character (:), ignore the line.
<原文结束>

# <翻译开始>
			// If the line starts with a U+003A COLON character (:), ignore the line.
# <翻译结束>


<原文开始>
			// If the line contains a U+003A COLON character character (:)
			// Collect the characters on the line before the first U+003A COLON character (:),
			// and let field be that string.
<原文结束>

# <翻译开始>
			// If the line contains a U+003A COLON character character (:)
			// Collect the characters on the line before the first U+003A COLON character (:),
			// and let field be that string.
# <翻译结束>


<原文开始>
			// Collect the characters on the line after the first U+003A COLON character (:),
			// and let value be that string.
<原文结束>

# <翻译开始>
			// Collect the characters on the line after the first U+003A COLON character (:),
			// and let value be that string.
# <翻译结束>


<原文开始>
			// If value starts with a single U+0020 SPACE character, remove it from value.
<原文结束>

# <翻译开始>
			// If value starts with a single U+0020 SPACE character, remove it from value.
# <翻译结束>


<原文开始>
			// Otherwise, the string is not empty but does not contain a U+003A COLON character character (:)
			// Use the whole line as the field name, and the empty string as the field value.
<原文结束>

# <翻译开始>
			// Otherwise, the string is not empty but does not contain a U+003A COLON character character (:)
			// Use the whole line as the field name, and the empty string as the field value.
# <翻译结束>


<原文开始>
		// The steps to process the field given a field name and a field value depend on the field name,
		// as given in the following list. Field names must be compared literally,
		// with no case folding performed.
<原文结束>

# <翻译开始>
		// The steps to process the field given a field name and a field value depend on the field name,
		// as given in the following list. Field names must be compared literally,
		// with no case folding performed.
# <翻译结束>


<原文开始>
			// Set the event name buffer to field value.
<原文结束>

# <翻译开始>
			// Set the event name buffer to field value.
# <翻译结束>


<原文开始>
			// Set the event stream's last event ID to the field value.
<原文结束>

# <翻译开始>
			// Set the event stream's last event ID to the field value.
# <翻译结束>


<原文开始>
			// If the field value consists of only characters in the range U+0030 DIGIT ZERO (0) to U+0039 DIGIT NINE (9),
			// then interpret the field value as an integer in base ten, and set the event stream's reconnection time to that integer.
			// Otherwise, ignore the field.
<原文结束>

# <翻译开始>
			// If the field value consists of only characters in the range U+0030 DIGIT ZERO (0) to U+0039 DIGIT NINE (9),
			// then interpret the field value as an integer in base ten, and set the event stream's reconnection time to that integer.
			// Otherwise, ignore the field.
# <翻译结束>


<原文开始>
			// Append the field value to the data buffer,
<原文结束>

# <翻译开始>
			// Append the field value to the data buffer,
# <翻译结束>


<原文开始>
			// then append a single U+000A LINE FEED (LF) character to the data buffer.
<原文结束>

# <翻译开始>
			// then append a single U+000A LINE FEED (LF) character to the data buffer.
# <翻译结束>


<原文开始>
			//Otherwise. The field is ignored.
<原文结束>

# <翻译开始>
			//Otherwise. The field is ignored.
# <翻译结束>


<原文开始>
	// Once the end of the file is reached, the user agent must dispatch the event one final time.
<原文结束>

# <翻译开始>
	// Once the end of the file is reached, the user agent must dispatch the event one final time.
# <翻译结束>

