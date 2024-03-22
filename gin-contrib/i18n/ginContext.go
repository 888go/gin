package i18n

import (
	"github.com/888go/gin"
)

// defaultGetLngHandler ...
func defaultGetLngHandler(context *gin类.Context, defaultLng string) string {
	if context == nil || context.X请求 == nil {
		return defaultLng
	}

	lng := context.X取请求协议头值("Accept-Language")
	if lng != "" {
		return lng
	}

	lng = context.X取URL参数值("lng")
	if lng == "" {
		return defaultLng
	}

	return lng
}
