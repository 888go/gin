// 版权所有 ? 2017 Manu Martinez-Almeida。保留所有权利。
// 本源代码的使用受 MIT 风格许可证协议约束，
// 该协议可在 LICENSE 文件中查阅。

package binding

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
	
	"github.com/go-playground/validator/v10"
)

type defaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

type SliceValidationError []error

// Error 将 SliceValidationError 中的所有错误元素连接成一个由 \n 分隔的单个字符串。

// ff:
func (err SliceValidationError) Error() string {
	n := len(err)
	switch n {
	case 0:
		return ""
	default:
		var b strings.Builder
		if err[0] != nil {
			fmt.Fprintf(&b, "[%d]: %s", 0, err[0].Error())
		}
		if n > 1 {
			for i := 1; i < n; i++ {
				if err[i] != nil {
					b.WriteString("\n")
					fmt.Fprintf(&b, "[%d]: %s", i, err[i].Error())
				}
			}
		}
		return b.String()
	}
}

var _ StructValidator = (*defaultValidator)(nil)

// ValidateStruct 接收任何类型的参数，但仅对结构体或指向结构体的指针类型执行验证操作。

// ff:
// obj:
func (v *defaultValidator) ValidateStruct(obj any) error {
	if obj == nil {
		return nil
	}

	value := reflect.ValueOf(obj)
	switch value.Kind() {
	case reflect.Ptr:
		return v.ValidateStruct(value.Elem().Interface())
	case reflect.Struct:
		return v.validateStruct(obj)
	case reflect.Slice, reflect.Array:
		count := value.Len()
		validateRet := make(SliceValidationError, 0)
		for i := 0; i < count; i++ {
			if err := v.ValidateStruct(value.Index(i).Interface()); err != nil {
				validateRet = append(validateRet, err)
			}
		}
		if len(validateRet) == 0 {
			return nil
		}
		return validateRet
	default:
		return nil
	}
}

// validateStruct 接收结构体类型
func (v *defaultValidator) validateStruct(obj any) error {
	v.lazyinit()
	return v.validate.Struct(obj)
}

// Engine 返回驱动默认验证器的底层验证引擎。如果要注册自定义验证或结构级别验证，这将非常有用。更多详情请参阅 validator 的 GoDoc 文档：
// https://pkg.go.dev/github.com/go-playground/validator/v10

// ff:
func (v *defaultValidator) Engine() any {
	v.lazyinit()
	return v.validate
}

func (v *defaultValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("binding")
	})
}
