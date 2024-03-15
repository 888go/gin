package i18n

import (
	"github.com/888go/gin"
)

// GinI18n ...
// GinI18n 是 Gin 框架的一个国际化（Internationalization，简写为 I18n）扩展包，
// 用于在基于Gin框架开发的Web应用中实现多语言支持功能。
type GinI18n interface {
	getMessage(context *gin.Context, param interface{}) (string, error)
	mustGetMessage(context *gin.Context, param interface{}) string
	setBundle(cfg *BundleCfg)
	setGetLngHandler(handler GetLngHandler)
}
