package i18n

import (
	"github.com/888go/gin"
)

// defaultGetLngHandler ...（默认获取经度处理器）
func defaultGetLngHandler(context *gin.Context, defaultLng string) string {
	if context == nil || context.Request == nil {
		return defaultLng
	}

	lng := context.GetHeader("Accept-Language")
	if lng != "" {
		return lng
	}

	lng = context.Query("lng")
	if lng == "" {
		return defaultLng
	}

	return lng
}
