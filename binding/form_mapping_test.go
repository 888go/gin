// 版权所有2019 Gin Core Team
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到

package binding

import (
	"reflect"
	"testing"
	"time"
	
	"github.com/stretchr/testify/assert"
)


// ff:
// t:
func TestMappingBaseTypes(t *testing.T) {
	intPtr := func(i int) *int {
		return &i
	}
	for _, tt := range []struct {
		name   string
		value  any
		form   string
		expect any
	}{
		{"base type", struct{ F int }{}, "9", int(9)},
		{"base type", struct{ F int8 }{}, "9", int8(9)},
		{"base type", struct{ F int16 }{}, "9", int16(9)},
		{"base type", struct{ F int32 }{}, "9", int32(9)},
		{"base type", struct{ F int64 }{}, "9", int64(9)},
		{"base type", struct{ F uint }{}, "9", uint(9)},
		{"base type", struct{ F uint8 }{}, "9", uint8(9)},
		{"base type", struct{ F uint16 }{}, "9", uint16(9)},
		{"base type", struct{ F uint32 }{}, "9", uint32(9)},
		{"base type", struct{ F uint64 }{}, "9", uint64(9)},
		{"base type", struct{ F bool }{}, "True", true},
		{"base type", struct{ F float32 }{}, "9.1", float32(9.1)},
		{"base type", struct{ F float64 }{}, "9.1", float64(9.1)},
		{"base type", struct{ F string }{}, "test", string("test")},
		{"base type", struct{ F *int }{}, "9", intPtr(9)},

// 零值
		{"zero value", struct{ F int }{}, "", int(0)},
		{"zero value", struct{ F uint }{}, "", uint(0)},
		{"zero value", struct{ F bool }{}, "", false},
		{"zero value", struct{ F float32 }{}, "", float32(0)},
	} {
		tp := reflect.TypeOf(tt.value)
		testName := tt.name + ":" + tp.Field(0).Type.String()

		val := reflect.New(reflect.TypeOf(tt.value))
		val.Elem().Set(reflect.ValueOf(tt.value))

		field := val.Elem().Type().Field(0)

		_, err := mapping(val, emptyField, formSource{field.Name: {tt.form}}, "form")
		assert.NoError(t, err, testName)

		actual := val.Elem().Field(0).Interface()
		assert.Equal(t, tt.expect, actual, testName)
	}
}


// ff:
// t:
func TestMappingDefault(t *testing.T) {
	var s struct {
		Int   int    `form:",default=9"`
		Slice []int  `form:",default=9"`
		Array [1]int `form:",default=9"`
	}
	err := mappingByPtr(&s, formSource{}, "form")
	assert.NoError(t, err)

	assert.Equal(t, 9, s.Int)
	assert.Equal(t, []int{9}, s.Slice)
	assert.Equal(t, [1]int{9}, s.Array)
}


// ff:
// t:
func TestMappingSkipField(t *testing.T) {
	var s struct {
		A int
	}
	err := mappingByPtr(&s, formSource{}, "form")
	assert.NoError(t, err)

	assert.Equal(t, 0, s.A)
}


// ff:
// t:
func TestMappingIgnoreField(t *testing.T) {
	var s struct {
		A int `form:"A"`
		B int `form:"-"`
	}
	err := mappingByPtr(&s, formSource{"A": {"9"}, "B": {"9"}}, "form")
	assert.NoError(t, err)

	assert.Equal(t, 9, s.A)
	assert.Equal(t, 0, s.B)
}


// ff:
// t:
func TestMappingUnexportedField(t *testing.T) {
	var s struct {
		A int `form:"a"`
		b int `form:"b"`
	}
	err := mappingByPtr(&s, formSource{"a": {"9"}, "b": {"9"}}, "form")
	assert.NoError(t, err)

	assert.Equal(t, 9, s.A)
	assert.Equal(t, 0, s.b)
}


// ff:
// t:
func TestMappingPrivateField(t *testing.T) {
	var s struct {
		f int `form:"field"`
	}
	err := mappingByPtr(&s, formSource{"field": {"6"}}, "form")
	assert.NoError(t, err)
	assert.Equal(t, 0, s.f)
}


// ff:
// t:
func TestMappingUnknownFieldType(t *testing.T) {
	var s struct {
		U uintptr
	}

	err := mappingByPtr(&s, formSource{"U": {"unknown"}}, "form")
	assert.Error(t, err)
	assert.Equal(t, errUnknownType, err)
}


// ff:
// t:
func TestMappingURI(t *testing.T) {
	var s struct {
		F int `uri:"field"`
	}
	err := mapURI(&s, map[string][]string{"field": {"6"}})
	assert.NoError(t, err)
	assert.Equal(t, 6, s.F)
}


// ff:
// t:
func TestMappingForm(t *testing.T) {
	var s struct {
		F int `form:"field"`
	}
	err := mapForm(&s, map[string][]string{"field": {"6"}})
	assert.NoError(t, err)
	assert.Equal(t, 6, s.F)
}


// ff:
// t:
func TestMapFormWithTag(t *testing.T) {
	var s struct {
		F int `externalTag:"field"`
	}
	err := MapFormWithTag(&s, map[string][]string{"field": {"6"}}, "externalTag")
	assert.NoError(t, err)
	assert.Equal(t, 6, s.F)
}


// ff:
// t:
func TestMappingTime(t *testing.T) {
	var s struct {
		Time      time.Time
		LocalTime time.Time `time_format:"2006-01-02"`
		ZeroValue time.Time
		CSTTime   time.Time `time_format:"2006-01-02" time_location:"Asia/Shanghai"`
		UTCTime   time.Time `time_format:"2006-01-02" time_utc:"1"`
	}

	var err error
	time.Local, err = time.LoadLocation("Europe/Berlin")
	assert.NoError(t, err)

	err = mapForm(&s, map[string][]string{
		"Time":      {"2019-01-20T16:02:58Z"},
		"LocalTime": {"2019-01-20"},
		"ZeroValue": {},
		"CSTTime":   {"2019-01-20"},
		"UTCTime":   {"2019-01-20"},
	})
	assert.NoError(t, err)

	assert.Equal(t, "2019-01-20 16:02:58 +0000 UTC", s.Time.String())
	assert.Equal(t, "2019-01-20 00:00:00 +0100 CET", s.LocalTime.String())
	assert.Equal(t, "2019-01-19 23:00:00 +0000 UTC", s.LocalTime.UTC().String())
	assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", s.ZeroValue.String())
	assert.Equal(t, "2019-01-20 00:00:00 +0800 CST", s.CSTTime.String())
	assert.Equal(t, "2019-01-19 16:00:00 +0000 UTC", s.CSTTime.UTC().String())
	assert.Equal(t, "2019-01-20 00:00:00 +0000 UTC", s.UTCTime.String())

// 错误的位置
	var wrongLoc struct {
		Time time.Time `time_location:"wrong"`
	}
	err = mapForm(&wrongLoc, map[string][]string{"Time": {"2019-01-20T16:02:58Z"}})
	assert.Error(t, err)

// 时间值错误
	var wrongTime struct {
		Time time.Time
	}
	err = mapForm(&wrongTime, map[string][]string{"Time": {"wrong"}})
	assert.Error(t, err)
}


// ff:
// t:
func TestMappingTimeDuration(t *testing.T) {
	var s struct {
		D time.Duration
	}

// 好吧
	err := mappingByPtr(&s, formSource{"D": {"5s"}}, "form")
	assert.NoError(t, err)
	assert.Equal(t, 5*time.Second, s.D)

	// error
	err = mappingByPtr(&s, formSource{"D": {"wrong"}}, "form")
	assert.Error(t, err)
}


// ff:
// t:
func TestMappingSlice(t *testing.T) {
	var s struct {
		Slice []int `form:"slice,default=9"`
	}

// 默认值
	err := mappingByPtr(&s, formSource{}, "form")
	assert.NoError(t, err)
	assert.Equal(t, []int{9}, s.Slice)

// 好吧
	err = mappingByPtr(&s, formSource{"slice": {"3", "4"}}, "form")
	assert.NoError(t, err)
	assert.Equal(t, []int{3, 4}, s.Slice)

	// error
	err = mappingByPtr(&s, formSource{"slice": {"wrong"}}, "form")
	assert.Error(t, err)
}


// ff:
// t:
func TestMappingArray(t *testing.T) {
	var s struct {
		Array [2]int `form:"array,default=9"`
	}

// 错误的默认
	err := mappingByPtr(&s, formSource{}, "form")
	assert.Error(t, err)

// 好吧
	err = mappingByPtr(&s, formSource{"array": {"3", "4"}}, "form")
	assert.NoError(t, err)
	assert.Equal(t, [2]int{3, 4}, s.Array)

// 错误-没有足够的值
	err = mappingByPtr(&s, formSource{"array": {"3"}}, "form")
	assert.Error(t, err)

// 错误-错误的值
	err = mappingByPtr(&s, formSource{"array": {"wrong"}}, "form")
	assert.Error(t, err)
}


// ff:
// t:
func TestMappingStructField(t *testing.T) {
	var s struct {
		J struct {
			I int
		}
	}

	err := mappingByPtr(&s, formSource{"J": {`{"I": 9}`}}, "form")
	assert.NoError(t, err)
	assert.Equal(t, 9, s.J.I)
}


// ff:
// t:
func TestMappingPtrField(t *testing.T) {
	type ptrStruct struct {
		Key int64 `json:"key"`
	}

	type ptrRequest struct {
		Items []*ptrStruct `json:"items" form:"items"`
	}

	var err error

// 0项
	var req0 ptrRequest
	err = mappingByPtr(&req0, formSource{}, "form")
	assert.NoError(t, err)
	assert.Empty(t, req0.Items)

// 只有一件物品
	var req1 ptrRequest
	err = mappingByPtr(&req1, formSource{"items": {`{"key": 1}`}}, "form")
	assert.NoError(t, err)
	assert.Len(t, req1.Items, 1)
	assert.EqualValues(t, 1, req1.Items[0].Key)

// 有2个项目
	var req2 ptrRequest
	err = mappingByPtr(&req2, formSource{"items": {`{"key": 1}`, `{"key": 2}`}}, "form")
	assert.NoError(t, err)
	assert.Len(t, req2.Items, 2)
	assert.EqualValues(t, 1, req2.Items[0].Key)
	assert.EqualValues(t, 2, req2.Items[1].Key)
}


// ff:
// t:
func TestMappingMapField(t *testing.T) {
	var s struct {
		M map[string]int
	}

	err := mappingByPtr(&s, formSource{"M": {`{"one": 1}`}}, "form")
	assert.NoError(t, err)
	assert.Equal(t, map[string]int{"one": 1}, s.M)
}


// ff:
// t:
func TestMappingIgnoredCircularRef(t *testing.T) {
	type S struct {
		S *S `form:"-"`
	}
	var s S

	err := mappingByPtr(&s, formSource{}, "form")
	assert.NoError(t, err)
}
