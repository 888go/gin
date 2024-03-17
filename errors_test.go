// Manu Martinez-Almeida版权所有
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到

package gin

import (
	"errors"
	"fmt"
	"testing"
	
	"github.com/888go/gin/internal/json"
	"github.com/stretchr/testify/assert"
)


// ff:
// t:
func TestError(t *testing.T) {
	baseError := errors.New("test error")
	err := &Error{
		Err:  baseError,
		Type: ErrorTypePrivate,
	}
	assert.Equal(t, err.Error(), baseError.Error())
	assert.Equal(t, H{"error": baseError.Error()}, err.JSON())

	assert.Equal(t, err.SetType(ErrorTypePublic), err)
	assert.Equal(t, ErrorTypePublic, err.Type)

	assert.Equal(t, err.SetMeta("some data"), err)
	assert.Equal(t, "some data", err.Meta)
	assert.Equal(t, H{
		"error": baseError.Error(),
		"meta":  "some data",
	}, err.JSON())

	jsonBytes, _ := json.Marshal(err)
	assert.Equal(t, "{\"error\":\"test error\",\"meta\":\"some data\"}", string(jsonBytes))

	err.SetMeta(H{ // nolint: errcheck
// 翻译：// 不进行errcheck检查
		"status": "200",
		"data":   "some data",
	})
	assert.Equal(t, H{
		"error":  baseError.Error(),
		"status": "200",
		"data":   "some data",
	}, err.JSON())

	err.SetMeta(H{ // nolint: errcheck
// 翻译：// 不进行errcheck检查
		"error":  "custom error",
		"status": "200",
		"data":   "some data",
	})
	assert.Equal(t, H{
		"error":  "custom error",
		"status": "200",
		"data":   "some data",
	}, err.JSON())

	type customError struct {
		status string
		data   string
	}
	err.SetMeta(customError{status: "200", data: "other data"}) // nolint: errcheck
// 翻译：// 不进行errcheck检查
	assert.Equal(t, customError{status: "200", data: "other data"}, err.JSON())
}


// ff:
// t:
func TestErrorSlice(t *testing.T) {
	errs := errorMsgs{
		{Err: errors.New("first"), Type: ErrorTypePrivate},
		{Err: errors.New("second"), Type: ErrorTypePrivate, Meta: "some data"},
		{Err: errors.New("third"), Type: ErrorTypePublic, Meta: H{"status": "400"}},
	}

	assert.Equal(t, errs, errs.ByType(ErrorTypeAny))
	assert.Equal(t, "third", errs.Last().Error())
	assert.Equal(t, []string{"first", "second", "third"}, errs.Errors())
	assert.Equal(t, []string{"third"}, errs.ByType(ErrorTypePublic).Errors())
	assert.Equal(t, []string{"first", "second"}, errs.ByType(ErrorTypePrivate).Errors())
	assert.Equal(t, []string{"first", "second", "third"}, errs.ByType(ErrorTypePublic|ErrorTypePrivate).Errors())
	assert.Empty(t, errs.ByType(ErrorTypeBind))
	assert.Empty(t, errs.ByType(ErrorTypeBind).String())

	assert.Equal(t, `Error #01: first
Error #02: second
     Meta: some data
Error #03: third
     Meta: map[status:400]
`, errs.String())
	assert.Equal(t, []any{
		H{"error": "first"},
		H{"error": "second", "meta": "some data"},
		H{"error": "third", "status": "400"},
	}, errs.JSON())
	jsonBytes, _ := json.Marshal(errs)
	assert.Equal(t, "[{\"error\":\"first\"},{\"error\":\"second\",\"meta\":\"some data\"},{\"error\":\"third\",\"status\":\"400\"}]", string(jsonBytes))
	errs = errorMsgs{
		{Err: errors.New("first"), Type: ErrorTypePrivate},
	}
	assert.Equal(t, H{"error": "first"}, errs.JSON())
	jsonBytes, _ = json.Marshal(errs)
	assert.Equal(t, "{\"error\":\"first\"}", string(jsonBytes))

	errs = errorMsgs{}
	assert.Nil(t, errs.Last())
	assert.Nil(t, errs.JSON())
	assert.Empty(t, errs.String())
}

type TestErr string


// ff:
// e:
func (e TestErr) Error() string { return string(e) }

// testrorunwrap测试gin的行为
// Error . is ()"“和“误差()
// “errors.Is()“;和“误差()“;已经被添加到go 1.13的标准库中

// ff:
// t:
func TestErrorUnwrap(t *testing.T) {
	innerErr := TestErr("some error")

// 2层包装:使用'fmt. error ("%w")'来包装杜松子酒
// Error{}，它本身包装了innerErr
	err := fmt.Errorf("wrapped: %w", &Error{
		Err:  innerErr,
		Type: ErrorTypeAny,
	})

// 检查'errors.Is()'和'errors.As()'的行为是否符合预期:
	assert.True(t, errors.Is(err, innerErr))
	var testErr TestErr
	assert.True(t, errors.As(err, &testErr))
}
