// 版权所有 ? 2014 Manu Martinez-Almeida。保留所有权利。
// 本源代码的使用受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package sse

import (
	"bytes"
	"io"
	"io/ioutil"
)

type decoder struct {
	events []Event
}

func Decode(r io.Reader) ([]Event, error) {
	var dec decoder
	return dec.decode(r)
}

func (d *decoder) dispatchEvent(event Event, data string) {
	dataLength := len(data)
	if dataLength > 0 {
		// 如果数据缓冲区的最后一个字符是U+000A换行符（LF），则从数据缓冲区中移除最后一个字符。
		data = data[:dataLength-1]
		dataLength--
	}
	if dataLength == 0 && event.Event == "" {
		return
	}
	if event.Event == "" {
		event.Event = "message"
	}
	event.Data = data
	d.events = append(d.events, event)
}

func (d *decoder) decode(r io.Reader) ([]Event, error) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var currentEvent Event
	var dataBuffer *bytes.Buffer = new(bytes.Buffer)
// TODO (并编写单元测试)
// 行与行之间必须由以下任意一种字符对或单个字符分隔：
// U+000D 回车（CARRIAGE RETURN）和 U+000A 换行（LINE FEED，即 CRLF 组合）
// 或者
// 单个 U+000A 换行（LINE FEED，LF）字符
// 或者
// 单个 U+000D 回车（CARRIAGE RETURN，CR）字符。
	lines := bytes.Split(buf, []byte{'\n'})
	for _, line := range lines {
		if len(line) == 0 {
			// 如果该行为空（即空白行），则分派该事件。
			d.dispatchEvent(currentEvent, dataBuffer.String())

			// 重置当前事件和数据缓冲区
			currentEvent = Event{}
			dataBuffer.Reset()
			continue
		}
		if line[0] == byte(':') {
			// 如果行以 U+003A 字符（冒号）开始，则忽略该行。
			continue
		}

		var field, value []byte
		colonIndex := bytes.IndexRune(line, ':')
		if colonIndex != -1 {
// 如果该行包含一个 U+003A 字符（冒号）：
// 收集该行第一个 U+003A 字符（冒号）之前的所有字符，
// 并将这些字符组成的字符串赋值给 field。
			field = line[:colonIndex]
// 收集行中第一个 U+003A COLON 字符 (:) 之后的字符，并将该字符串赋值给 value。
			value = line[colonIndex+1:]
			// 如果value以单个U+0020（空格）字符开头，则从value中移除它。
			if len(value) > 0 && value[0] == ' ' {
				value = value[1:]
			}
		} else {
// 否则，字符串不为空，但不包含 U+003A 字符（冒号）：
// 将整行用作字段名，并使用空字符串作为字段值。
			field = line
			value = []byte{}
		}
// 根据给定的字段名和字段值来处理该字段的步骤取决于字段名，具体如下所示。字段名必须进行逐字比较，且不应进行大小写折叠处理。
		switch string(field) {
		case "event":
			// 将事件名称缓冲区设置为字段值。
			currentEvent.Event = string(value)
		case "id":
			// 将事件流的最后一个事件ID设置为字段值。
			currentEvent.Id = string(value)
		case "retry":
// 如果字段值仅包含从 U+0030（数字零，0）到 U+0039（数字九，9）范围内的字符，则将该字段值解释为十进制整数，并将事件流的重连时间设置为该整数。否则，忽略该字段。
			currentEvent.Id = string(value)
		case "data":
			// 将字段值追加到数据缓冲区，
// 
// ensuring there's a space separator if the buffer already has content.
// 确保如果缓冲区已有内容，则在追加前添加一个空格分隔符。
			dataBuffer.Write(value)
			// 然后向数据缓冲区追加一个 U+000A 换行符（LF）字符。
			dataBuffer.WriteString("\n")
		default:
			// 否则。该字段将被忽略。
			continue
		}
	}
	// 当到达文件末尾时，用户代理必须最后一次派发该事件。
	d.dispatchEvent(currentEvent, dataBuffer.String())

	return d.events, nil
}
