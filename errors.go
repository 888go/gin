// 版权所有 2014 Manu Martinez-Almeida。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package gin类

import (
	"fmt"
	"reflect"
	"strings"
	
	"github.com/888go/gin/internal/json"
)

// ErrorType 是一个无符号的64位错误代码，遵循gin规范定义。
type ErrorType uint64

const (
	// ErrorTypeBind 用于当 Context.Bind() 失败时。
	ErrorTypeBind ErrorType = 1 << 63
	// ErrorTypeRender 用于当 Context.Render() 失败时。
	ErrorTypeRender ErrorType = 1 << 62
	// ErrorTypePrivate 表示一个私有错误。
	ErrorTypePrivate ErrorType = 1 << 0
	// ErrorTypePublic 表示一个公开的错误。
	ErrorTypePublic ErrorType = 1 << 1
	// ErrorTypeAny 表示任何其他错误。
	ErrorTypeAny ErrorType = 1<<64 - 1
	// ErrorTypeNu 表示任何其他错误。
	ErrorTypeNu = 2
)

// Error代表了一个错误的规格说明。
type Error struct {
	Err  error
	Type ErrorType
	Meta any
}

type errorMsgs []*Error

var _ error = (*Error)(nil)

// SetType 设置错误的类型。
func (msg *Error) SetType(flags ErrorType) *Error {
	msg.Type = flags
	return msg
}

// SetMeta 设置错误的元数据。
func (msg *Error) SetMeta(data any) *Error {
	msg.Meta = data
	return msg
}

// JSON 创建一个格式正确的 JSON
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

// MarshalJSON 实现了 json.Marshaller 接口。
func (msg *Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(msg.JSON())
}

// Error 实现了 error 接口。
func (msg Error) Error() string {
	return msg.Err.Error()
}

// IsType 判断一个错误。
func (msg *Error) IsType(flags ErrorType) bool {
	return (msg.Type & flags) > 0
}

// Unwrap 返回封装的错误，以便与 errors.Is()、errors.As() 和 errors.Unwrap() 之间进行互操作性
func (msg *Error) Unwrap() error {
	return msg.Err
}

// ByType 返回一个只读副本，其中包含了经过过滤的错误信息。具体来说，ByType(gin.ErrorTypePublic) 将返回一个类型为 ErrorTypePublic 的错误信息切片。
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

// Last 函数返回切片中的最后一个错误。如果该数组为空，则返回 nil。
// 这是 errors[len(errors)-1] 的快捷方式。
func (a errorMsgs) Last() *Error {
	if length := len(a); length > 0 {
		return a[length-1]
	}
	return nil
}

// Errors 返回包含所有错误消息的数组。
// 示例：
//
//	c.Error(errors.New("第一个错误"))
//	c.Error(errors.New("第二个错误"))
//	c.Error(errors.New("第三个错误"))
//	c.Errors.Errors() // == []string{"第一个", "第二个", "第三个"}
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

// MarshalJSON 实现了 json.Marshaller 接口。
func (a errorMsgs) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.JSON())
}

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
