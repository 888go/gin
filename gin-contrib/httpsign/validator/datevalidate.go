package validator

import (
	"errors"
	"fmt"
	"net/http"
	"time"
	
	"github.com/888go/gin"
)

const maxTimeGap = 30 * time.Second // 30 secs

func newPublicError(msg string) *gin.Error {
	return &gin.Error{
		Err:  errors.New(msg),
		Type: gin.ErrorTypePublic,
	}
}

// ErrDateNotInRange：当日期不在可接受的范围内时返回错误
var ErrDateNotInRange = newPublicError("Date submit is not in acceptable range")

// DateValidator 校验器，通过时间范围进行校验
type DateValidator struct {
// TimeGap 是客户端提交时间戳与服务器时间之间允许的最大时间差，以毫秒为精度，被认为是有效的。
	TimeGap time.Duration
}

// NewDateValidator 返回一个具有默认值（30秒）的 DateValidator
func NewDateValidator() *DateValidator {
	return &DateValidator{
		TimeGap: maxTimeGap,
	}
}

// Validate在检查头部日期是否有效时返回错误
func (v *DateValidator) Validate(r *http.Request) error {
	t, err := http.ParseTime(r.Header.Get("date"))
	if err != nil {
		return newPublicError(fmt.Sprintf("Could not parse date header. Error: %s", err.Error()))
	}
	serverTime := time.Now()
	start := serverTime.Add(-v.TimeGap)
	stop := serverTime.Add(v.TimeGap)
	if t.Before(start) || t.After(stop) {
		return ErrDateNotInRange
	}
	return nil
}
