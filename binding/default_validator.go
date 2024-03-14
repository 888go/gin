// 版权所有2017马努·马丁内斯-阿尔梅达
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到

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

// Error将SliceValidationError中的所有错误元素连接成一个以\n分隔的字符串
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

// ValidateStruct接受任何类型，但只能执行结构或指向结构类型的指针
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

// validateStruct接收结构类型
func (v *defaultValidator) validateStruct(obj any) error {
	v.lazyinit()
	return v.validate.Struct(obj)
}

// Engine返回为默认validator实例提供动力的底层验证器引擎
// 如果您想注册自定义验证或结构层验证，这将非常有用
// 请参阅验证器GoDoc获取更多信息- https://pkg.go.dev/github.com/go-playground/validator/v10
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
