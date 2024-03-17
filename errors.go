// Manu Martinez-Almeida版权所有
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到

package gin

import (
	"fmt"
	"reflect"
	"strings"
	
	"github.com/888go/gin/internal/json"
)

// ErrorType是在gin规范中定义的无符号64位错误代码
type ErrorType uint64

const (
// 当Context.Bind()失败时使用ErrorTypeBind
	ErrorTypeBind ErrorType = 1 << 63
// 当Context.Render()失败时使用ErrorTypeRender
	ErrorTypeRender ErrorType = 1 << 62
// ErrorTypePrivate私有错误
	ErrorTypePrivate ErrorType = 1 << 0
// ErrorTypePublic表示公共错误
	ErrorTypePublic ErrorType = 1 << 1
// ErrorTypeAny表示任何其他错误
	ErrorTypeAny ErrorType = 1<<64 - 1
// ErrorTypeNu表示任何其他错误
	ErrorTypeNu = 2
)

// Error表示错误的说明
type Error struct {
	Err  error
	Type ErrorType
	Meta any
}

type errorMsgs []*Error

var _ error = (*Error)(nil)

// SetType设置错误的类型

// ff:
// flags:
func (msg *Error) SetType(flags ErrorType) *Error {
	msg.Type = flags
	return msg
}

// SetMeta设置错误的元数据

// ff:
// data:
func (msg *Error) SetMeta(data any) *Error {
	msg.Meta = data
	return msg
}

// JSON创建一个格式正确的JSON

// ff:
func (msg *Error) JSON() any {
	jsonData := H{}
	if msg.Meta != nil {
		value := reflect.ValueOf(msg.Meta)
		switch value.Kind() {
		case reflect.Struct:
			return msg.Meta
		case reflect.Map:
			for _, key := range value.MapKeys() {
				jsonData[key.String()] = value.MapIndex(key).Interface()
			}
		default:
			jsonData["meta"] = msg.Meta
		}
	}
	if _, ok := jsonData["error"]; !ok {
		jsonData["error"] = msg.Error()
	}
	return jsonData
}

// MarshalJSON实现json
// Marshaller接口

// ff:
func (msg *Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(msg.JSON())
}

// Error实现错误接口

// ff:
func (msg Error) Error() string {
	return msg.Err.Error()
}

// IsType判断一个错误

// ff:
// flags:
func (msg *Error) IsType(flags ErrorType) bool {
	return (msg.Type & flags) > 0
}

// Unwrap返回包装后的错误，以允许与errors.Is()、errors.As()和errors.Unwrap()互操作

// ff:
func (msg *Error) Unwrap() error {
	return msg.Err
}

// ByType返回经过字节过滤的只读副本
// 即ByType(gin.ErrorTypePublic)返回一个类型=ErrorTypePublic的错误切片

// ff:
// typ:
func (a errorMsgs) ByType(typ ErrorType) errorMsgs {
	if len(a) == 0 {
		return nil
	}
	if typ == ErrorTypeAny {
		return a
	}
	var result errorMsgs
	for _, msg := range a {
		if msg.IsType(typ) {
			result = append(result, msg)
		}
	}
	return result
}

// Last返回切片中的最后一个错误
// 如果数组为空，则返回nil
// 错误的快捷方式[len(errors)-1]

// ff:
func (a errorMsgs) Last() *Error {
	if length := len(a); length > 0 {
		return a[length-1]
	}
	return nil
}

// Errors返回一个包含所有错误消息的数组
// 示例:c.Error(errors.New("first")) c.Error(errors.New("second")) c. errors (errors.New("third")) c. errors () == []string{"first"， "second"， "third"}

// ff:
func (a errorMsgs) Errors() []string {
	if len(a) == 0 {
		return nil
	}
	errorStrings := make([]string, len(a))
	for i, err := range a {
		errorStrings[i] = err.Error()
	}
	return errorStrings
}


// ff:
func (a errorMsgs) JSON() any {
	switch length := len(a); length {
	case 0:
		return nil
	case 1:
		return a.Last().JSON()
	default:
		jsonData := make([]any, length)
		for i, err := range a {
			jsonData[i] = err.JSON()
		}
		return jsonData
	}
}

// MarshalJSON实现json
// Marshaller接口

// ff:
func (a errorMsgs) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.JSON())
}


// ff:
func (a errorMsgs) String() string {
	if len(a) == 0 {
		return ""
	}
	var buffer strings.Builder
	for i, msg := range a {
		fmt.Fprintf(&buffer, "Error #%02d: %s\n", i+1, msg.Err)
		if msg.Meta != nil {
			fmt.Fprintf(&buffer, "     Meta: %v\n", msg.Meta)
		}
	}
	return buffer.String()
}
