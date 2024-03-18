package i18n

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// BundleCfg ...
type BundleCfg struct {
	DefaultLanguage  language.Tag
	FormatBundleFile string
	AcceptLanguage   []language.Tag
	RootPath         string
	UnmarshalFunc    i18n.UnmarshalFunc
	Loader           Loader
}

type Loader interface {
	LoadMessage(path string) ([]byte, error)
}

type LoaderFunc func(path string) ([]byte, error)


// ff:
// path:
// path:

// ff:
// path:
// path:

// ff:
// path:
// path:

// ff:
// path:
// path:

// ff:
// path:
// path:
func (f LoaderFunc) LoadMessage(path string) ([]byte, error) { return f(path) }

// WithBundle ... // 使用Bundle

// ff:
// config:

// ff:
// config:

// ff:
// config:

// ff:
// config:

// ff:
// config:
func WithBundle(config *BundleCfg) Option {
	return func(g GinI18n) {
		if config.Loader == nil {
			config.Loader = defaultLoader
		}
		g.setBundle(config)
	}
}

// WithGetLngHandle ...
// ...
// （由于上下文信息不足，无法准确翻译该注释含义，请提供更多代码或上下文信息。）
// 一般情况下，根据函数命名习惯，这个注释可能是在描述一个带有获取经度处理功能的方法或选项，"WithGetLngHandle"可以理解为“带有获取经度处理器”，通常这种形式的函数用于在初始化结构体或者设置配置时，注入获取经度的操作句柄或回调函数。

// ff:
// handler:

// ff:
// handler:

// ff:
// handler:

// ff:
// handler:

// ff:
// handler:
func WithGetLngHandle(handler GetLngHandler) Option {
	return func(g GinI18n) {
		g.setGetLngHandler(handler)
	}
}
