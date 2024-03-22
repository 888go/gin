package validator

import (
	"net/http"
)

// 验证器接口，用于检查请求是否有效
type Validator interface {
	Validate(*http.Request) error
}
