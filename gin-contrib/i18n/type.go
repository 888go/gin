package i18n

import (
	"github.com/888go/gin"
)

type (
	// GetLngHandler ...
	GetLngHandler = func(context *gin.Context, defaultLng string) string

	// Option ...
	Option func(GinI18n)
)
