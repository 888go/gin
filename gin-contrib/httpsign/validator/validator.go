package validator

import (
	"net/http"
)

// Validator 接口，用于检查请求是否有效
type Validator interface {
	Validate(*http.Request) error
}
