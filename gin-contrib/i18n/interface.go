package i18n

import (
	"github.com/888go/gin"
)

// GinI18n ...
type GinI18n interface {
	getMessage(context *gin类.Context, param interface{}) (string, error)
	mustGetMessage(context *gin类.Context, param interface{}) string
	setBundle(cfg *BundleCfg)
	setGetLngHandler(handler GetLngHandler)
}
