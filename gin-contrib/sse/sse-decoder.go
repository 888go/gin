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


// ff:
// []Event:
// r:

// ff:
// []Event:
// r:

// ff:
// []Event:
// r:

// ff:
// []Event:
// r:

// ff:
// []Event:
// r:
func Decode(r io.Reader) ([]Event, error) {
	var dec decoder
	return dec.decode(r)
}

func (d *decoder) dispatchEvent(event Event, data string) {
	dataLength := len(data)
	if dataLength > 0 {
// 如果数据缓冲区的最后一个字符是 U+000A 换行符（LF），则从数据缓冲区中移除最后一个字符。
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
// TODO (以及单元测试)
// 行必须由以下字符分隔：
// - U+000D 回车（CARRIAGE RETURN，CR）和 U+000A 换行（LINE FEED，LF）字符对
// - 单个 U+000A 换行（LINE FEED，LF）字符
// - 或者单个 U+000D 回车（CARRIAGE RETURN，CR）字符。
	lines := bytes.Split(buf, []byte{'\n'})
	for _, line := range lines {
		if len(line) == 0 {
// 如果该行为空（即空白行），则分发事件。
			d.dispatchEvent(currentEvent, dataBuffer.String())

// 重置当前事件和数据缓冲区
			currentEvent = Event{}
			dataBuffer.Reset()
			continue
		}
		if line[0] == byte(':') {
// 如果行以 U+003A（冒号）字符开始，则忽略该行。
			continue
		}

		var field, value []byte
		colonIndex := bytes.IndexRune(line, ':')
		if colonIndex != -1 {
// 如果该行包含 U+003A 字符（冒号）：
// 收集该行第一个 U+003A 字符（冒号）之前的所有字符，
// 并将这个字符串赋值给 field 变量。
			field = line[:colonIndex]
// 获取第一个 U+003A（冒号）字符之后的行上字符，并将这些字符组成的字符串赋值给 value。
			value = line[colonIndex+1:]
// 如果value以一个U+0020（空格）字符开头，则从value中移除它。
			if len(value) > 0 && value[0] == ' ' {
				value = value[1:]
			}
		} else {
// 否则，字符串非空但不包含 U+003A（冒号）字符
// 将整行作为字段名，并使用空字符串作为字段值。
			field = line
			value = []byte{}
		}
// 根据给定的字段名和字段值处理该字段的步骤取决于字段名，
// 以下列表给出了具体规则。字段名必须进行逐字比较，
// 不进行大小写折叠处理。
		switch string(field) {
		case "event":
// 将事件名称缓冲区设置为字段值。
			currentEvent.Event = string(value)
		case "id":
// 将事件流的最后事件ID设置为字段值。
			currentEvent.Id = string(value)
		case "retry":
// 如果字段值仅包含范围从 U+0030（数字零 0）到 U+0039（数字九 9）之间的字符，
// 则将字段值解释为十进制整数，并将事件流的重连时间设置为该整数。
// 否则，忽略该字段。
			currentEvent.Id = string(value)
		case "data":
// 将字段值追加到数据缓冲区，
			dataBuffer.Write(value)
// 然后向数据缓冲区追加一个 U+000A 换行符（LF）字符。
			dataBuffer.WriteString("\n")
		default:
// 否则，该字段将被忽略。
			continue
		}
	}
// 当到达文件末尾时，用户代理必须最后一次派发该事件。
	d.dispatchEvent(currentEvent, dataBuffer.String())

	return d.events, nil
}
