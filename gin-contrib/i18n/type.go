package i18n
import (
	"e.coding.net/gogit/go/gin"
	)
type (
	// GetLngHandler ...
	GetLngHandler = func(context *gin.Context, defaultLng string) string

	// Option ...
	Option func(GinI18n)
)
