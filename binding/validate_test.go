// Manu Martinez-Almeida版权所有
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到

package binding
import (
	"bytes"
	"testing"
	"time"
	
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	)
type testInterface interface {
	String() string
}

type substructNoValidation struct {
	IString string
	IInt    int
}

type mapNoValidationSub map[string]substructNoValidation

type structNoValidationValues struct {
	substructNoValidation

	Boolean bool

	Uinteger   uint
	Integer    int
	Integer8   int8
	Integer16  int16
	Integer32  int32
	Integer64  int64
	Uinteger8  uint8
	Uinteger16 uint16
	Uinteger32 uint32
	Uinteger64 uint64

	Float32 float32
	Float64 float64

	String string

	Date time.Time

	Struct        substructNoValidation
	InlinedStruct struct {
		String  []string
		Integer int
	}

	IntSlice           []int
	IntPointerSlice    []*int
	StructPointerSlice []*substructNoValidation
	StructSlice        []substructNoValidation
	InterfaceSlice     []testInterface

	UniversalInterface any
	CustomInterface    testInterface

	FloatMap  map[string]float32
	StructMap mapNoValidationSub
}

func createNoValidationValues() structNoValidationValues {
	integer := 1
	s := structNoValidationValues{
		Boolean:            true,
		Uinteger:           1 << 29,
		Integer:            -10000,
		Integer8:           120,
		Integer16:          -20000,
		Integer32:          1 << 29,
		Integer64:          1 << 61,
		Uinteger8:          250,
		Uinteger16:         50000,
		Uinteger32:         1 << 31,
		Uinteger64:         1 << 62,
		Float32:            123.456,
		Float64:            123.456789,
		String:             "text",
		Date:               time.Time{},
		CustomInterface:    &bytes.Buffer{},
		Struct:             substructNoValidation{},
		IntSlice:           []int{-3, -2, 1, 0, 1, 2, 3},
		IntPointerSlice:    []*int{&integer},
		StructSlice:        []substructNoValidation{},
		UniversalInterface: 1.2,
		FloatMap: map[string]float32{
			"foo": 1.23,
			"bar": 232.323,
		},
		StructMap: mapNoValidationSub{
			"foo": substructNoValidation{},
			"bar": substructNoValidation{},
		},
// StructPointerSlice []noValidationSub interfacesslice []testInterface
	}
	s.InlinedStruct.Integer = 1000
	s.InlinedStruct.String = []string{"first", "second"}
	s.IString = "substring"
	s.IInt = 987654
	return s
}

func TestValidateNoValidationValues(t *testing.T) {
	origin := createNoValidationValues()
	test := createNoValidationValues()
	empty := structNoValidationValues{}

	assert.Nil(t, validate(test))
	assert.Nil(t, validate(&test))
	assert.Nil(t, validate(empty))
	assert.Nil(t, validate(&empty))

	assert.Equal(t, origin, test)
}

type structNoValidationPointer struct {
	substructNoValidation

	Boolean bool

	Uinteger   *uint
	Integer    *int
	Integer8   *int8
	Integer16  *int16
	Integer32  *int32
	Integer64  *int64
	Uinteger8  *uint8
	Uinteger16 *uint16
	Uinteger32 *uint32
	Uinteger64 *uint64

	Float32 *float32
	Float64 *float64

	String *string

	Date *time.Time

	Struct *substructNoValidation

	IntSlice           *[]int
	IntPointerSlice    *[]*int
	StructPointerSlice *[]*substructNoValidation
	StructSlice        *[]substructNoValidation
	InterfaceSlice     *[]testInterface

	FloatMap  *map[string]float32
	StructMap *mapNoValidationSub
}

func TestValidateNoValidationPointers(t *testing.T) {
// origin:= createNoValidation_values() test:= createNoValidation_values()
	empty := structNoValidationPointer{}

// ,
// Nil(t, validate(test))断言
// 尼罗河(t,执行极为&test))
	assert.Nil(t, validate(empty))
	assert.Nil(t, validate(&empty))

// 断言
// 等于(t，原点，检验)
}

type Object map[string]any

func TestValidatePrimitives(t *testing.T) {
	obj := Object{"foo": "bar", "bar": 1}
	assert.NoError(t, validate(obj))
	assert.NoError(t, validate(&obj))
	assert.Equal(t, Object{"foo": "bar", "bar": 1}, obj)

	obj2 := []Object{{"foo": "bar", "bar": 1}, {"foo": "bar", "bar": 1}}
	assert.NoError(t, validate(obj2))
	assert.NoError(t, validate(&obj2))

	nu := 10
	assert.NoError(t, validate(nu))
	assert.NoError(t, validate(&nu))
	assert.Equal(t, 10, nu)

	str := "value"
	assert.NoError(t, validate(str))
	assert.NoError(t, validate(&str))
	assert.Equal(t, "value", str)
}

// structCustomValidation是一个辅助结构体，我们使用它来检查是否可以在其上注册自定义验证
// ' notone '绑定指令用于自定义验证并在以后注册
type structCustomValidation struct {
	Integer int `binding:"notone"`
}

func notOne(f1 validator.FieldLevel) bool {
	if val, ok := f1.Field().Interface().(int); ok {
		return val != 1
	}
	return false
}

func TestValidatorEngine(t *testing.T) {
// 这将验证函数' notOne '是否与' defaultValidator '和验证器库所期望的函数签名匹配
	engine, ok := Validator.Engine().(*validator.Validate)
	assert.True(t, ok)

	err := engine.RegisterValidation("notone", notOne)
// 检查我们是否可以注册自定义验证而不会出错
	assert.Nil(t, err)

// 创建一个验证失败的实例
	withOne := structCustomValidation{Integer: 1}
	errs := validate(withOne)

// 检查我们是否得到非nil错误
	assert.NotNil(t, errs)
// 检查错误是否与预期相符
	assert.Error(t, errs, "", "", "notone")
}
